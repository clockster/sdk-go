package clockster

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// What a call is missing before it is worth sending.
var (
	errNoBody = errors.New("clockster: this call needs a body")
	errNoFile = errors.New("clockster: this call needs a file to upload")
)

// request is what a generated method hands the transport. One place decides how a parameter is
// written, where the token goes and what a refusal becomes, so the generated half is the
// operations and nothing else.
type request struct {
	method string
	path   string
	query  url.Values
	body   any
	form   *formBody
}

// formBody is the one operation that carries bytes rather than JSON.
type formBody struct {
	file     io.Reader
	filename string
	fields   url.Values
}

// RequestOption changes one call rather than the client.
type RequestOption func(*http.Header)

// WithIdempotencyKey makes a retry of a write answer the first result rather than performing it
// again. Four writes create something with nothing to match a second attempt against — a rota, a
// webhook endpoint, a rotated secret, a delivery sent again — and the header is what makes those
// safe to repeat. A key is remembered for 24 hours and belongs to the request it was first used
// for.
func WithIdempotencyKey(key string) RequestOption {
	return WithHeader("Idempotency-Key", key)
}

// WithHeader sends a header with this call.
func WithHeader(name, value string) RequestOption {
	return func(headers *http.Header) {
		headers.Set(name, value)
	}
}

func (t *transport) do(ctx context.Context, spec request, out any, opts []RequestOption) error {
	body, contentType, err := payload(spec)
	if err != nil {
		return err
	}

	target := t.baseURL + spec.path

	if encoded := spec.query.Encode(); encoded != "" {
		target += "?" + encoded
	}

	call, err := http.NewRequestWithContext(ctx, spec.method, target, body)
	if err != nil {
		return fmt.Errorf("clockster: %s %s: %w", spec.method, spec.path, err)
	}

	// Read per call rather than held from construction, so a rotated key does not need a new
	// client.
	call.Header.Set("Authorization", "Bearer "+t.token)
	call.Header.Set("Accept", "application/json")
	call.Header.Set("User-Agent", t.userAgent)

	if contentType != "" {
		call.Header.Set("Content-Type", contentType)
	}

	for _, opt := range opts {
		opt(&call.Header)
	}

	response, err := t.httpClient.Do(call)
	if err != nil {
		return fmt.Errorf("clockster: %s %s: %w", spec.method, spec.path, err)
	}

	defer response.Body.Close()

	answer, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("clockster: reading the answer to %s %s: %w", spec.method, spec.path, err)
	}

	if response.StatusCode >= 400 {
		return refusal(response, answer)
	}

	// A 204 has no body, and neither has a call that answers nothing. Everything else is JSON.
	if out == nil || len(answer) == 0 {
		return nil
	}

	if err := json.Unmarshal(answer, out); err != nil {
		return fmt.Errorf("clockster: the answer to %s %s is not the JSON this client expects: %w", spec.method, spec.path, err)
	}

	return nil
}

func payload(spec request) (io.Reader, string, error) {
	if spec.form != nil {
		return multipartBody(spec.form)
	}

	if spec.body == nil {
		return nil, "", nil
	}

	encoded, err := json.Marshal(spec.body)
	if err != nil {
		return nil, "", fmt.Errorf("clockster: writing the body of %s %s: %w", spec.method, spec.path, err)
	}

	return bytes.NewReader(encoded), "application/json", nil
}

func multipartBody(form *formBody) (io.Reader, string, error) {
	// Buffered rather than streamed: the API answers the stored file, so the request is not
	// finished until the bytes are, and a pipe would only move where they are held.
	var buffer bytes.Buffer

	writer := multipart.NewWriter(&buffer)

	for name, values := range form.fields {
		for _, value := range values {
			if err := writer.WriteField(name, value); err != nil {
				return nil, "", fmt.Errorf("clockster: writing the form field %s: %w", name, err)
			}
		}
	}

	part, err := writer.CreateFormFile("file", form.filename)
	if err != nil {
		return nil, "", fmt.Errorf("clockster: writing the file part: %w", err)
	}

	if _, err := io.Copy(part, form.file); err != nil {
		return nil, "", fmt.Errorf("clockster: reading the file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("clockster: writing the form: %w", err)
	}

	return &buffer, writer.FormDataContentType(), nil
}

func refusal(response *http.Response, answer []byte) *Error {
	// An edge refusal can answer before the application does, and need not be JSON.
	var envelope struct {
		Error struct {
			Code      string              `json:"code"`
			Message   string              `json:"message"`
			RequestID string              `json:"request_id"`
			Errors    map[string][]string `json:"errors"`
		} `json:"error"`
	}

	_ = json.Unmarshal(answer, &envelope)

	refused := &Error{
		Status:    response.StatusCode,
		Code:      envelope.Error.Code,
		Message:   envelope.Error.Message,
		RequestID: envelope.Error.RequestID,
		Errors:    envelope.Error.Errors,
		Body:      answer,
	}

	if refused.Code == "" {
		refused.Code = "unknown"
	}

	if refused.Message == "" {
		refused.Message = fmt.Sprintf("The API answered %d.", response.StatusCode)
	}

	if seconds, err := strconv.Atoi(response.Header.Get("Retry-After")); err == nil {
		refused.RetryAfter = seconds
	}

	return refused
}

// Query parameters as this API reads them. A list travels comma-separated rather than as a
// repeated `field[]` — both are accepted, and this is the one the document describes. A parameter
// nobody set is not sent: an omitted filter and an empty one mean different things.

func queryOpt[T any](query url.Values, name string, value Opt[T]) {
	if held, ok := value.Value(); ok {
		query.Set(name, scalar(held))
	}
}

func queryList[T any](query url.Values, name string, values []T) {
	if len(values) == 0 {
		return
	}

	written := make([]string, len(values))

	for index, value := range values {
		written[index] = scalar(value)
	}

	query.Set(name, strings.Join(written, ","))
}

func scalar(value any) string {
	switch held := value.(type) {
	case string:
		return held
	case bool:
		return strconv.FormatBool(held)
	case int:
		return strconv.Itoa(held)
	case int64:
		return strconv.FormatInt(held, 10)
	case float64:
		return strconv.FormatFloat(held, 'f', -1, 64)
	default:
		return fmt.Sprint(held)
	}
}

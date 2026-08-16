// What the transport does with a call, against a stub rather than the API.

package clockster

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// call is one request as the stub received it.
type call struct {
	method  string
	path    string
	query   url.Values
	headers http.Header
	body    []byte
}

func serve(t *testing.T, status int, answer string, headers map[string]string) (*Client, *[]call) {
	t.Helper()

	seen := &[]call{}

	return serveFunc(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*seen = append(*seen, call{r.Method, r.URL.Path, r.URL.Query(), r.Header.Clone(), body})

		for name, value := range headers {
			w.Header().Set(name, value)
		}

		w.WriteHeader(status)
		io.WriteString(w, answer)
	}), seen
}

func serveFunc(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	client, err := New("token", WithBaseURL(server.URL))
	if err != nil {
		t.Fatalf("the client was refused: %v", err)
	}

	return client
}

func TestAnswersTheParsedBody(t *testing.T) {
	client, _ := serve(t, 200, `{"data":[{"id":7,"first_name":"Aisulu"}]}`, nil)

	answer, err := client.Users.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	if len(answer.Data) != 1 || answer.Data[0].ID != 7 || answer.Data[0].FirstName != "Aisulu" {
		t.Fatalf("the rows are not what the API answered: %+v", answer.Data)
	}
}

func TestCarriesTheTokenOnEveryCall(t *testing.T) {
	client, seen := serve(t, 200, `{"data":[]}`, nil)

	if _, err := client.Users.List(context.Background(), nil); err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	if got := (*seen)[0].headers.Get("Authorization"); got != "Bearer token" {
		t.Fatalf("the key did not travel: %q", got)
	}

	if got := (*seen)[0].headers.Get("Accept"); got != "application/json" {
		t.Fatalf("the answer was not asked for as JSON: %q", got)
	}

	// So the request log says which client made a call rather than which HTTP library did.
	if got := (*seen)[0].headers.Get("User-Agent"); got != "clockster-go/"+Version {
		t.Fatalf("the client named itself %q", got)
	}
}

func TestAnIntegrationCanNameItselfInstead(t *testing.T) {
	seen := &[]call{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*seen = append(*seen, call{headers: r.Header.Clone()})
		io.WriteString(w, `{"data":[]}`)
	}))

	t.Cleanup(server.Close)

	client, err := New("token", WithBaseURL(server.URL), WithUserAgent("acme-hr/1.4"))
	if err != nil {
		t.Fatalf("the client was refused: %v", err)
	}

	if _, err := client.Users.List(context.Background(), nil); err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	if got := (*seen)[0].headers.Get("User-Agent"); got != "acme-hr/1.4" {
		t.Fatalf("the client named itself %q", got)
	}
}

func TestWritesAListParameterCommaSeparated(t *testing.T) {
	client, seen := serve(t, 200, `{"data":[]}`, nil)

	_, err := client.Users.List(context.Background(), &UsersListParams{
		Include: []string{"location", "department"},
		IDs:     []int64{1, 2},
	})
	if err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	// Both forms are accepted by the API; this is the one the document describes, and the one that
	// keeps the parameter called `include` rather than `include[]`.
	if got := (*seen)[0].query.Get("include"); got != "location,department" {
		t.Fatalf("the relations asked for are %q", got)
	}

	if got := (*seen)[0].query.Get("ids"); got != "1,2" {
		t.Fatalf("the ids asked for are %q", got)
	}
}

func TestLeavesOutWhatWasNotAskedFor(t *testing.T) {
	client, seen := serve(t, 200, `{"data":[]}`, nil)

	_, err := client.Users.List(context.Background(), &UsersListParams{PerPage: Set(int64(50))})
	if err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	// An omitted filter and an empty one mean different things, and a field nobody set is the
	// first.
	if len((*seen)[0].query) != 1 || (*seen)[0].query.Get("per_page") != "50" {
		t.Fatalf("the query is %v", (*seen)[0].query)
	}
}

func TestWritesAParameterTheOperationRequires(t *testing.T) {
	client, seen := serve(t, 200, `{"data":[]}`, nil)

	_, err := client.Attendance.List(context.Background(), &AttendanceListParams{
		DateFrom: "2026-08-01",
		DateTo:   "2026-08-31",
	})
	if err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	if got := (*seen)[0].query.Get("date_from"); got != "2026-08-01" {
		t.Fatalf("the window starts at %q", got)
	}
}

func TestPathParametersLandInThePath(t *testing.T) {
	client, seen := serve(t, 200, `{"data":{"id":42}}`, nil)

	if _, err := client.Users.Get(context.Background(), 42, nil); err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	if (*seen)[0].path != "/company/v3/users/42" {
		t.Fatalf("the call went to %q", (*seen)[0].path)
	}
}

func TestSendsAnIdempotencyKeyWhenAsked(t *testing.T) {
	client, seen := serve(t, 201, `{"data":{"id":1}}`, nil)

	_, err := client.Webhooks.Create(context.Background(), &WebhooksCreateBody{
		URL:    "https://example.test/hook",
		Events: []string{"user.updated"},
	}, WithIdempotencyKey("2f8a"))
	if err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	if got := (*seen)[0].headers.Get("Idempotency-Key"); got != "2f8a" {
		t.Fatalf("the key travelled as %q", got)
	}
}

func TestAFieldNobodySetIsNotWritten(t *testing.T) {
	client, seen := serve(t, 200, `{"data":[]}`, nil)

	_, err := client.Users.Upsert(context.Background(), &UsersUpsertBody{
		Users: []UsersUpsertUser{{
			ExternalID: Set("HR-1"),
			FirstName:  "Aisulu",
			Role:       "employee",
			LocationID: 3,
			PositionID: Null[int64](),
		}},
	})
	if err != nil {
		t.Fatalf("the call was refused: %v", err)
	}

	var written struct {
		Users []map[string]any `json:"users"`
	}

	if err := json.Unmarshal((*seen)[0].body, &written); err != nil {
		t.Fatalf("what was sent is not JSON: %v", err)
	}

	person := written.Users[0]

	// The whole point of Opt: nothing was said about the department, so the stored one stays,
	// where the position was cleared on purpose.
	if _, ok := person["department_id"]; ok {
		t.Fatalf("a field nobody set was written: %v", person)
	}

	if held, ok := person["position_id"]; !ok || held != nil {
		t.Fatalf("the cleared field is %v", held)
	}

	if person["external_id"] != "HR-1" {
		t.Fatalf("the key is %v", person["external_id"])
	}
}

func TestARefusalIsAnError(t *testing.T) {
	client, _ := serve(t, 422, `{"error":{"code":"validation_failed","message":"The given data was invalid.","request_id":"9f2b","errors":{"users.0.role":["The role is required."]}}}`, nil)

	_, err := client.Users.Upsert(context.Background(), &UsersUpsertBody{Users: []UsersUpsertUser{}})
	if err == nil {
		t.Fatal("a refused call answered no error")
	}

	if !errors.Is(err, ErrValidation) {
		t.Fatalf("the refusal does not match ErrValidation: %v", err)
	}

	var refused *Error

	if !errors.As(err, &refused) {
		t.Fatalf("the refusal is not an *Error: %v", err)
	}

	if refused.Code != "validation_failed" || refused.RequestID != "9f2b" {
		t.Fatalf("the envelope was not read: %+v", refused)
	}

	if len(refused.Errors["users.0.role"]) != 1 {
		t.Fatalf("the named fields are %v", refused.Errors)
	}
}

func TestRateLimitCarriesRetryAfter(t *testing.T) {
	client, _ := serve(t, 429, `{"error":{"code":"too_many_requests","message":"Slow down.","request_id":"1a"}}`,
		map[string]string{"Retry-After": "30"})

	_, err := client.Users.List(context.Background(), nil)

	var refused *Error

	if !errors.As(err, &refused) || !errors.Is(err, ErrRateLimit) {
		t.Fatalf("the refusal is %v", err)
	}

	if refused.RetryAfter != 30 {
		t.Fatalf("the wait is %d seconds", refused.RetryAfter)
	}
}

func TestAnotherCompanysRowIsNotFound(t *testing.T) {
	client, _ := serve(t, 404, `{"error":{"code":"not_found","message":"No such user.","request_id":"3c"}}`, nil)

	_, err := client.Users.Get(context.Background(), 999, nil)

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("the refusal is %v", err)
	}
}

func TestAServerErrorIsOurs(t *testing.T) {
	client, _ := serve(t, 503, `{"error":{"code":"unavailable","message":"Down.","request_id":"7d"}}`, nil)

	_, err := client.Users.List(context.Background(), nil)

	if !errors.Is(err, ErrServer) {
		t.Fatalf("the refusal is %v", err)
	}
}

func TestARefusalNeedNotBeJSON(t *testing.T) {
	// An edge refusal can answer before the application does.
	client, _ := serve(t, 502, "<html>bad gateway</html>", nil)

	_, err := client.Users.List(context.Background(), nil)

	var refused *Error

	if !errors.As(err, &refused) {
		t.Fatalf("the refusal is %v", err)
	}

	if refused.Code != "unknown" || !strings.Contains(string(refused.Body), "bad gateway") {
		t.Fatalf("what came back was not kept: %+v", refused)
	}
}

func TestAnEmptyAnswerIsNotAnError(t *testing.T) {
	client, _ := serve(t, 204, "", nil)

	answer, err := client.Users.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("an empty answer was refused: %v", err)
	}

	if len(answer.Data) != 0 {
		t.Fatalf("rows appeared from nowhere: %v", answer.Data)
	}
}

func TestUploadSendsTheFileAsMultipart(t *testing.T) {
	client, seen := serve(t, 201, `{"data":{"id":1,"name":"agreement","format":"pdf","url":"https://x","created_at":"2026-03-02T09:00:00+00:00"}}`, nil)

	_, err := client.Files.Upload(context.Background(), &FilesUploadForm{
		File:     strings.NewReader("%PDF-1.7"),
		Filename: "agreement.pdf",
		Name:     Set("agreement"),
	})
	if err != nil {
		t.Fatalf("the upload was refused: %v", err)
	}

	sent := (*seen)[0]

	_, params, err := mime.ParseMediaType(sent.headers.Get("Content-Type"))
	if err != nil {
		t.Fatalf("the request is not multipart: %v", err)
	}

	reader := multipart.NewReader(strings.NewReader(string(sent.body)), params["boundary"])
	parts := map[string]string{}
	names := map[string]string{}

	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}

		body, _ := io.ReadAll(part)
		parts[part.FormName()] = string(body)
		names[part.FormName()] = part.FileName()
	}

	if parts["file"] != "%PDF-1.7" || names["file"] != "agreement.pdf" {
		t.Fatalf("the bytes travelled as %q named %q", parts["file"], names["file"])
	}

	if parts["name"] != "agreement" {
		t.Fatalf("the field beside the bytes is %q", parts["name"])
	}
}

func TestUploadNeedsAFile(t *testing.T) {
	client, _ := serve(t, 201, `{}`, nil)

	if _, err := client.Files.Upload(context.Background(), nil); !errors.Is(err, errNoFile) {
		t.Fatalf("an upload with nothing to upload answered %v", err)
	}
}

func TestNewNeedsAKey(t *testing.T) {
	if _, err := New(""); !errors.Is(err, ErrNoToken) {
		t.Fatalf("a client without a key was built: %v", err)
	}
}

func TestTheBaseURLLosesItsTrailingSlash(t *testing.T) {
	client, err := New("token", WithBaseURL("https://demo.clockster.com/"))
	if err != nil {
		t.Fatalf("the client was refused: %v", err)
	}

	if client.baseURL != "https://demo.clockster.com" {
		t.Fatalf("the calls would go to %q", client.baseURL)
	}
}

func TestATimeoutIsNotWrittenOntoSomebodyElsesClient(t *testing.T) {
	held := &http.Client{Timeout: time.Second}

	if _, err := New("token", WithHTTPClient(held), WithTimeout(time.Minute)); err != nil {
		t.Fatalf("the client was refused: %v", err)
	}

	if held.Timeout != time.Second {
		t.Fatalf("the supplied client now times out after %s", held.Timeout)
	}
}

package clockster

import (
	"errors"
	"fmt"
)

// The refusals worth telling apart, as targets for errors.Is. The API answers every refusal with
// one envelope, so there is one error type carrying it and a sentinel per status.
var (
	// ErrAuthentication is 401 — no token, or one this surface does not accept.
	ErrAuthentication = errors.New("clockster: unauthenticated")
	// ErrForbidden is 403 — a token without the ability this call needs.
	ErrForbidden = errors.New("clockster: forbidden")
	// ErrNotFound is 404 — no such row in the calling company. Another company's id answers this
	// rather than 403.
	ErrNotFound = errors.New("clockster: not found")
	// ErrConflict is 409 — the row is there and cannot be changed the way you asked.
	ErrConflict = errors.New("clockster: conflict")
	// ErrValidation is 422 — the request was understood and refused. Error.Errors names the fields.
	ErrValidation = errors.New("clockster: validation failed")
	// ErrRateLimit is 429 — over the limit. Error.RetryAfter is the seconds to wait.
	ErrRateLimit = errors.New("clockster: rate limited")
	// ErrServer is 5xx — ours to fix. Retrying is safe on a read and on anything carrying an
	// idempotency key.
	ErrServer = errors.New("clockster: server error")
)

var statusErrors = map[int]error{
	401: ErrAuthentication,
	403: ErrForbidden,
	404: ErrNotFound,
	409: ErrConflict,
	422: ErrValidation,
	429: ErrRateLimit,
}

// Error is a call the API refused.
//
// Code is what to branch on: it names the reason and does not change, where Message is prose and
// may. RequestID identifies this exact call in our logs and is worth quoting when asking us about
// it.
//
//	var refused *clockster.Error
//
//	if errors.As(err, &refused) {
//		log.Printf("%s %s", refused.Code, refused.RequestID)
//	}
//
//	if errors.Is(err, clockster.ErrNotFound) {
//		// no such row in this company
//	}
type Error struct {
	// Status is the HTTP status the API answered.
	Status int
	// Code names the reason, e.g. "validation_failed".
	Code string
	// Message is prose, for a log rather than a branch.
	Message string
	// RequestID identifies the call in our logs.
	RequestID string
	// Errors names the fields a 422 refused, and is empty otherwise.
	Errors map[string][]string
	// RetryAfter is the seconds to wait, from the header of a 429.
	RetryAfter int
	// Body is the response as it arrived, for a refusal this package does not know the shape of.
	Body []byte
}

func (e *Error) Error() string {
	if e.RequestID == "" {
		return fmt.Sprintf("clockster: [%d %s] %s", e.Status, e.Code, e.Message)
	}

	return fmt.Sprintf("clockster: [%d %s] %s (request_id %s)", e.Status, e.Code, e.Message, e.RequestID)
}

// Is matches the sentinel for this refusal's status, so errors.Is answers without unwrapping the
// status by hand.
func (e *Error) Is(target error) bool {
	if sentinel, ok := statusErrors[e.Status]; ok && target == sentinel {
		return true
	}

	return target == ErrServer && e.Status >= 500
}

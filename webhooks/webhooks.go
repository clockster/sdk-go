// Package webhooks verifies a delivery and reads the event out of it.
//
// The check is the only way to the event: Verify takes the body as it arrived and answers the
// parsed event, so there is no path that acts on one that was not verified.
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		event, err := webhooks.VerifyRequest(r, os.Getenv("CLOCKSTER_WEBHOOK_SECRET"))
//		if err != nil {
//			http.Error(w, "refused", http.StatusBadRequest)
//
//			return
//		}
//
//		w.WriteHeader(http.StatusOK)
//
//		go handle(event)
//	}
//
// Answer 2xx quickly and do the work afterwards — a timeout is retried — and deduplicate on the
// event's ID, since the same event may arrive twice.
package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// The headers a delivery carries. Event names what happened and Delivery identifies the attempt,
// so a repeat can be recognised without parsing the body.
const (
	HeaderSignature = "X-Clockster-Signature"
	HeaderTimestamp = "X-Clockster-Timestamp"
	HeaderEvent     = "X-Clockster-Event"
	HeaderDelivery  = "X-Clockster-Delivery"
)

const scheme = "sha256="

// DefaultTolerance is the oldest delivery accepted. Refusing an old one is what stops a replay.
const DefaultTolerance = 5 * time.Minute

// Reason is why a delivery was refused, for a log or a metric.
type Reason string

const (
	ReasonMissingSignature  Reason = "missing_signature"
	ReasonMissingTimestamp  Reason = "missing_timestamp"
	ReasonUnknownScheme     Reason = "unknown_scheme"
	ReasonSignatureMismatch Reason = "signature_mismatch"
	ReasonTimestampUnusable Reason = "timestamp_unreadable"
	ReasonOutsideTolerance  Reason = "timestamp_outside_tolerance"
	ReasonBodyUnparseable   Reason = "body_unparseable"
	ReasonBodyUnreadable    Reason = "body_unreadable"
)

// Error is anything that makes a delivery unsafe to act on.
type Error struct {
	Reason  Reason
	Message string
}

func (e *Error) Error() string {
	return "clockster/webhooks: " + e.Message
}

// Event is what a delivery carries.
type Event struct {
	// ID is null on a trial delivery, which stands for no recorded event. Deduplicate on it: the
	// same event may arrive twice.
	ID *int64 `json:"id"`
	// Event is what happened, e.g. "user.updated".
	Event string `json:"event"`
	// OccurredAt is when it happened, which may be long past on an event recorded late.
	OccurredAt string `json:"occurred_at"`
	// Data is the event's own body, left as it arrived: what it holds depends on Event.
	Data json.RawMessage `json:"data"`
}

// Delivery is a request as it arrived.
type Delivery struct {
	// Body is the bytes as received. Re-serialising a parsed object does not reproduce the signed
	// bytes and the check will fail — read the raw request body.
	Body []byte
	// Signature is the X-Clockster-Signature header.
	Signature string
	// Timestamp is the X-Clockster-Timestamp header, an ISO 8601 instant rather than a Unix time.
	Timestamp string
	// Secret is the signing secret of the endpoint.
	Secret string
}

// Option changes how strict the check is.
type Option func(*settings)

type settings struct {
	tolerance time.Duration
	now       func() time.Time
}

// WithTolerance replaces DefaultTolerance. Zero or less accepts a delivery of any age, which is
// worth having in a test and not in production.
func WithTolerance(tolerance time.Duration) Option {
	return func(s *settings) {
		s.tolerance = tolerance
	}
}

// VerifyRequest verifies a delivery as it arrived over HTTP and answers the event it carries.
//
// The body is read here, so a handler that wants the raw bytes as well should read them itself and
// call Verify.
func VerifyRequest(r *http.Request, secret string, opts ...Option) (*Event, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, &Error{Reason: ReasonBodyUnreadable, Message: "The request body cannot be read: " + err.Error()}
	}

	return Verify(Delivery{
		Body:      body,
		Signature: r.Header.Get(HeaderSignature),
		Timestamp: r.Header.Get(HeaderTimestamp),
		Secret:    secret,
	}, opts...)
}

// Verify verifies a delivery and answers the event it carries. Everything that would make one
// unsafe to act on is an *Error naming its Reason.
func Verify(delivery Delivery, opts ...Option) (*Event, error) {
	held := &settings{tolerance: DefaultTolerance, now: time.Now}

	for _, opt := range opts {
		opt(held)
	}

	if err := check(delivery, held); err != nil {
		return nil, err
	}

	var event Event

	if err := json.Unmarshal(delivery.Body, &event); err != nil {
		return nil, &Error{Reason: ReasonBodyUnparseable, Message: "The body is signed but is not the event this package reads: " + err.Error()}
	}

	return &event, nil
}

func check(delivery Delivery, held *settings) error {
	switch {
	case delivery.Signature == "":
		return &Error{Reason: ReasonMissingSignature, Message: "No " + HeaderSignature + " header."}
	case delivery.Timestamp == "":
		return &Error{Reason: ReasonMissingTimestamp, Message: "No " + HeaderTimestamp + " header."}
	case !strings.HasPrefix(delivery.Signature, scheme):
		return &Error{Reason: ReasonUnknownScheme, Message: "The signature is not " + scheme + "<hex>."}
	}

	// The timestamp is inside what is signed, so it cannot be edited to widen the check below.
	signed := hmac.New(sha256.New, []byte(delivery.Secret))
	signed.Write([]byte(delivery.Timestamp + "."))
	signed.Write(delivery.Body)

	expected := hex.EncodeToString(signed.Sum(nil))

	if !hmac.Equal([]byte(strings.TrimPrefix(delivery.Signature, scheme)), []byte(expected)) {
		return &Error{
			Reason:  ReasonSignatureMismatch,
			Message: "The signature does not match the body. Verify the bytes as received, before parsing them.",
		}
	}

	return fresh(delivery.Timestamp, held)
}

func fresh(timestamp string, held *settings) error {
	if held.tolerance <= 0 {
		return nil
	}

	sent, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return &Error{
			Reason:  ReasonTimestampUnusable,
			Message: fmt.Sprintf("The timestamp %q is not an ISO 8601 instant.", timestamp),
		}
	}

	age := held.now().Sub(sent)

	// Absolute, so a receiver whose clock runs behind refuses rather than accepting indefinitely.
	if age < 0 {
		age = -age
	}

	if age > held.tolerance {
		return &Error{
			Reason:  ReasonOutsideTolerance,
			Message: fmt.Sprintf("The delivery is outside the %s tolerance.", held.tolerance),
		}
	}

	return nil
}

// Verifying a delivery, against one this test signed rather than one that arrived.

package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const secret = "whsec_test"

func sign(timestamp string, body []byte) string {
	signed := hmac.New(sha256.New, []byte(secret))
	signed.Write([]byte(timestamp + "."))
	signed.Write(body)

	return scheme + hex.EncodeToString(signed.Sum(nil))
}

func delivered(body string) Delivery {
	timestamp := time.Now().UTC().Format(time.RFC3339)

	return Delivery{
		Body:      []byte(body),
		Signature: sign(timestamp, []byte(body)),
		Timestamp: timestamp,
		Secret:    secret,
	}
}

func refusal(t *testing.T, err error) *Error {
	t.Helper()

	var refused *Error

	if !errors.As(err, &refused) {
		t.Fatalf("the delivery answered %v rather than a refusal", err)
	}

	return refused
}

func TestVerifiesADeliveryAndAnswersTheEvent(t *testing.T) {
	event, err := Verify(delivered(`{"id":123,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+05:00","data":{"id":7}}`))
	if err != nil {
		t.Fatalf("a delivery we signed was refused: %v", err)
	}

	if event.Event != "user.updated" || event.ID == nil || *event.ID != 123 {
		t.Fatalf("the event read as %+v", event)
	}

	if !strings.Contains(string(event.Data), `"id":7`) {
		t.Fatalf("the event's own body is %s", event.Data)
	}
}

func TestATrialDeliveryCarriesNoEventID(t *testing.T) {
	event, err := Verify(delivered(`{"id":null,"event":"ping","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`))
	if err != nil {
		t.Fatalf("a trial delivery was refused: %v", err)
	}

	// Null stands for no recorded event, which is what a trial delivery is.
	if event.ID != nil {
		t.Fatalf("the trial delivery carries id %d", *event.ID)
	}
}

func TestRefusesABodyThatWasChanged(t *testing.T) {
	delivery := delivered(`{"id":1,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`)
	delivery.Body = []byte(strings.Replace(string(delivery.Body), `"id":1`, `"id":2`, 1))

	_, err := Verify(delivery)

	if reason := refusal(t, err).Reason; reason != ReasonSignatureMismatch {
		t.Fatalf("the changed body was refused as %s", reason)
	}
}

func TestRefusesADeliveryUnderAnotherSecret(t *testing.T) {
	delivery := delivered(`{"id":1,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`)
	delivery.Secret = "whsec_somebody_else"

	_, err := Verify(delivery)

	if reason := refusal(t, err).Reason; reason != ReasonSignatureMismatch {
		t.Fatalf("the wrong secret was refused as %s", reason)
	}
}

func TestTheTimestampIsInsideWhatIsSigned(t *testing.T) {
	delivery := delivered(`{"id":1,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`)
	delivery.Timestamp = time.Now().UTC().Add(-time.Minute).Format(time.RFC3339)

	_, err := Verify(delivery)

	// Moving the timestamp to widen the tolerance breaks the signature instead.
	if reason := refusal(t, err).Reason; reason != ReasonSignatureMismatch {
		t.Fatalf("the moved timestamp was refused as %s", reason)
	}
}

func TestRefusesWhatIsMissing(t *testing.T) {
	body := `{"id":1,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`

	cases := []struct {
		name     string
		change   func(*Delivery)
		expected Reason
	}{
		{"no signature", func(d *Delivery) { d.Signature = "" }, ReasonMissingSignature},
		{"no timestamp", func(d *Delivery) { d.Timestamp = "" }, ReasonMissingTimestamp},
		{"another scheme", func(d *Delivery) { d.Signature = "md5=abc" }, ReasonUnknownScheme},
	}

	for _, held := range cases {
		t.Run(held.name, func(t *testing.T) {
			delivery := delivered(body)
			held.change(&delivery)

			_, err := Verify(delivery)

			if reason := refusal(t, err).Reason; reason != held.expected {
				t.Fatalf("it was refused as %s", reason)
			}
		})
	}
}

func TestRefusesADeliveryOutsideTheTolerance(t *testing.T) {
	timestamp := time.Now().UTC().Add(-time.Hour).Format(time.RFC3339)
	body := []byte(`{"id":1,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`)

	delivery := Delivery{Body: body, Signature: sign(timestamp, body), Timestamp: timestamp, Secret: secret}

	// Signed by us and an hour old: refusing it is what stops a replay.
	if _, err := Verify(delivery); refusal(t, err).Reason != ReasonOutsideTolerance {
		t.Fatalf("an hour-old delivery was refused as %v", err)
	}

	if _, err := Verify(delivery, WithTolerance(0)); err != nil {
		t.Fatalf("with the check off it was still refused: %v", err)
	}
}

func TestRefusesATimestampThatIsNotAnInstant(t *testing.T) {
	body := []byte(`{"id":1,"event":"user.updated","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`)
	timestamp := "the second tuesday"

	_, err := Verify(Delivery{Body: body, Signature: sign(timestamp, body), Timestamp: timestamp, Secret: secret})

	if reason := refusal(t, err).Reason; reason != ReasonTimestampUnusable {
		t.Fatalf("it was refused as %s", reason)
	}
}

func TestRefusesABodyThatIsSignedButNotAnEvent(t *testing.T) {
	_, err := Verify(delivered("not json at all"))

	if reason := refusal(t, err).Reason; reason != ReasonBodyUnparseable {
		t.Fatalf("it was refused as %s", reason)
	}
}

func TestVerifyRequestReadsTheHeaders(t *testing.T) {
	delivery := delivered(`{"id":9,"event":"attendance.recorded","occurred_at":"2026-03-02T09:00:00+00:00","data":{}}`)

	request := httptest.NewRequest("POST", "/hook", strings.NewReader(string(delivery.Body)))
	request.Header.Set(HeaderSignature, delivery.Signature)
	request.Header.Set(HeaderTimestamp, delivery.Timestamp)

	event, err := VerifyRequest(request, secret)
	if err != nil {
		t.Fatalf("the delivery was refused: %v", err)
	}

	if event.Event != "attendance.recorded" {
		t.Fatalf("the event read as %+v", event)
	}
}

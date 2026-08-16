// What a field left out, a field set and a field cleared come to on the wire.

package clockster

import (
	"encoding/json"
	"testing"
)

type fields struct {
	Untouched Opt[string] `json:"untouched,omitzero"`
	Cleared   Opt[string] `json:"cleared,omitzero"`
	Written   Opt[int64]  `json:"written,omitzero"`
}

func TestOnlyWhatWasSetIsWritten(t *testing.T) {
	written, err := json.Marshal(fields{Cleared: Null[string](), Written: Set(int64(7))})
	if err != nil {
		t.Fatalf("the fields cannot be written: %v", err)
	}

	if string(written) != `{"cleared":null,"written":7}` {
		t.Fatalf("what would be sent is %s", written)
	}
}

func TestReadsBackWhatWasWritten(t *testing.T) {
	var held fields

	if err := json.Unmarshal([]byte(`{"cleared":null,"written":7}`), &held); err != nil {
		t.Fatalf("the fields cannot be read: %v", err)
	}

	if held.Untouched.Sent() {
		t.Fatal("a field that was not there reads as sent")
	}

	if !held.Cleared.Sent() || !held.Cleared.IsNull() {
		t.Fatalf("the cleared field reads as %+v", held.Cleared)
	}

	value, ok := held.Written.Value()

	if !ok || value != 7 {
		t.Fatalf("the written field reads as %v, %v", value, ok)
	}
}

func TestValueAnswersNothingForANull(t *testing.T) {
	if _, ok := Null[string]().Value(); ok {
		t.Fatal("a cleared field answered a value")
	}

	if held := Null[string]().Or("fallback"); held != "fallback" {
		t.Fatalf("the fallback is %q", held)
	}

	if held := Set("held").Or("fallback"); held != "held" {
		t.Fatalf("the value is %q", held)
	}
}

func TestDerefAnswersTheZeroValueForNothing(t *testing.T) {
	if held := Deref[string](nil); held != "" {
		t.Fatalf("nothing dereferenced to %q", held)
	}

	if held := Deref(Ptr("held")); held != "held" {
		t.Fatalf("the value is %q", held)
	}
}

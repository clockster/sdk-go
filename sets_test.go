// The sets of values, and the two things worth asserting about a generated constant.

package clockster

import (
	"slices"
	"testing"
)

func TestSetConstantIsTheStringItStandsFor(t *testing.T) {
	// Nothing to unwrap and nothing to convert: it goes wherever the string goes, which is what
	// makes an untyped constant the right shape for a field that stayed a string.
	if UsersRoleEmployee != "employee" {
		t.Fatalf("UsersRoleEmployee is %q, wanted %q", UsersRoleEmployee, "employee")
	}

	if WebhooksEventUserCreated != "user.created" {
		t.Fatalf("WebhooksEventUserCreated is %q, wanted %q", WebhooksEventUserCreated, "user.created")
	}
}

func TestSetValuesAreEveryOneInOrder(t *testing.T) {
	if got := UsersStatusValues(); !slices.Equal(got, []string{"active", "dismissed", "all"}) {
		t.Fatalf("UsersStatusValues is %v", got)
	}
}

func TestSetValuesCannotBeReachedTwice(t *testing.T) {
	// A function rather than a package-level slice, so one caller sorting it in place does not
	// reorder it for the next.
	first := UsersStatusValues()
	first[0] = "meddled"

	if UsersStatusValues()[0] != "active" {
		t.Fatal("the values are shared between callers")
	}
}

func TestOneResourceInsideAnotherSharesTheSet(t *testing.T) {
	// `webhooks` and `webhooks/deliveries` name the same events, and the document says so once.
	if !slices.Contains(WebhooksEventValues(), WebhooksEventTaskApproved) {
		t.Fatal("the event names of a subscription and of a delivery filter are not one set")
	}
}

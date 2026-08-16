// Package clockster is the official Go SDK for the Clockster Company API.
//
// A server-to-server client for a company's employees, structure, schedules, attendance, tasks and
// documents, generated from the API's own OpenAPI document. No dependencies.
//
//	clockster, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
//	if err != nil {
//		return err
//	}
//
//	for user, err := range clockster.Users.ListAll(ctx, &clockster.UsersListParams{
//		Include: []string{"location"},
//	}) {
//		if err != nil {
//			return err
//		}
//
//		fmt.Println(user.ID, user.FirstName)
//	}
//
// # What a method answers
//
// The parsed body, so rows are answer.Data, and an error on anything the API refused. Nothing is
// validated on the way in: a field we add tomorrow reaches your code today rather than being
// refused by this package.
//
// # Absent, null and set
//
// A key that was not asked for is absent, never null: null means the value is known to be empty,
// and absent means you did not ask. In an answer both are a nil pointer. In a request they are
// told apart, because they mean different things there — a field left out keeps whatever is
// stored, and one written as null clears it. That is what [Opt] is for: [Set] carries a value,
// [Null] clears, and a field nobody touched is not written at all.
//
// # Refusals
//
// A refusal is an [Error], carrying the code to branch on, the message, and the request id to
// quote when asking us about a call. The statuses worth telling apart have a sentinel each, so
// errors.Is answers without unwrapping anything:
//
//	if errors.Is(err, clockster.ErrRateLimit) {
//		// over the limit; Error.RetryAfter says how long to wait
//	}
//
// # Dates, times and numbers
//
// Instants, dates and clock times are strings, in the shapes the document states, rather than
// time.Time: a date carries no zone and a clock time is read in the timezone stated beside it, and
// converting either would decide something this package does not know. Decimal amounts are JSON
// numbers rounded to two places; do not accumulate them in binary floating point.
//
// Deliveries are verified with the [github.com/clockster/sdk-go/webhooks] package.
package clockster

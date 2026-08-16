# sdk-go

[![Go Reference](https://pkg.go.dev/badge/github.com/clockster/sdk-go.svg)](https://pkg.go.dev/github.com/clockster/sdk-go)

Official Go SDK for the [Clockster Company API](https://api.clockster.com/openapi/v3.json).

Server-to-server client for a company's employees, structure, schedules, attendance, tasks and
documents. Generated from the API's OpenAPI document. No dependencies.

```bash
go get github.com/clockster/sdk-go
```

Requires Go 1.24 or newer.

## Quickstart

One token authenticates one company. Create it under Settings → API in the web application.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/clockster/sdk-go"
)

func main() {
	client, err := clockster.New(os.Getenv("CLOCKSTER_TOKEN"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	me, err := client.Me(ctx)
	if err != nil {
		panic(err)
	}

	locations, err := client.Locations.Upsert(ctx, &clockster.LocationsUpsertBody{
		Items: []clockster.LocationsUpsertItem{{ExternalID: "HQ", Title: "Head office"}},
	})
	if err != nil {
		panic(err)
	}

	_, err = client.Users.Upsert(ctx, &clockster.UsersUpsertBody{
		Users: []clockster.UsersUpsertUser{{
			ExternalID: clockster.Set("HR-1"),
			FirstName:  "Aisulu",
			Role:       "employee",
			LocationID: locations.Data[0].ID,
		}},
	})
	if err != nil {
		panic(err)
	}

	timesheets, err := client.Timesheets.List(ctx, &clockster.TimesheetsListParams{
		DateFrom: "2026-08-01",
		DateTo:   "2026-08-31",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(me.Data.Title, len(timesheets.Data))
}
```

A method answers the parsed body, so rows are `answer.Data`. Nothing is validated on the way in:
the answer is the JSON as it arrived, and a field we add tomorrow reaches your code today.

The methods are named after the operations, so the API documentation is the reference for both:
`GET /users` is `client.Users.List(...)`, `POST /users/upsert` is `client.Users.Upsert(...)`. The
TypeScript and Python clients use the same names.

## Absent, null and set

A key that was not asked for is absent, never null. `null` means the value is known to be empty; an
absent key means you did not ask. In an answer both are a nil pointer:

```go
if user.Department != nil {
	fmt.Println(user.Department.Title)
}
```

In a request the two are told apart, because there they mean different things — a field left out
keeps whatever is stored, and one written as null clears it. That is what `Opt` is for:

```go
clockster.UsersUpsertUser{
	ExternalID:   clockster.Set("HR-1"),
	FirstName:    "Aisulu",
	PositionID:   clockster.Null[int64](), // clear the position
	// DepartmentID is not set: not written, and the stored department stays.
}
```

`clockster.Deref` reads a pointer that may be nil, and `clockster.Ptr` writes a required field that
accepts null.

## Paging

Thirteen listings page on a cursor, and each has a `ListAll` beside its `List` that walks them:

```go
for user, err := range client.Users.ListAll(ctx, &clockster.UsersListParams{
	PerPage: clockster.Set(int64(100)),
	Include: []string{"location"},
}) {
	if err != nil {
		return err
	}

	fmt.Println(clockster.Deref(user.ExternalID))
}
```

A refused page is answered where it was refused, so half a listing is never mistaken for the whole
of one. A cursor belongs to the filters it was issued under: change them and walk again. `List`
itself is there when you want to hold the cursor yourself — `answer.Meta.NextCursor`.

## Refusals

A refusal is an error carrying the whole envelope. `Code` is what to branch on: it names the reason
and does not change, where `Message` is prose and may. `RequestID` identifies the call in our logs.

```go
_, err := client.Users.Upsert(ctx, body)

var refused *clockster.Error

if errors.As(err, &refused) {
	log.Printf("%s %v %s", refused.Code, refused.Errors, refused.RequestID)
}

if errors.Is(err, clockster.ErrRateLimit) {
	time.Sleep(time.Duration(refused.RetryAfter) * time.Second)
}
```

The statuses worth telling apart have a sentinel each: `ErrAuthentication` (401), `ErrForbidden`
(403), `ErrNotFound` (404), `ErrConflict` (409), `ErrValidation` (422), `ErrRateLimit` (429) and
`ErrServer` (5xx). Another company's id answers `ErrNotFound` rather than `ErrForbidden` — you
cannot learn that it exists.

## Retries and idempotency

Retry a 429 and a 5xx; do not retry a 4xx. A keyed write converges on what you meant rather than
doubling anything, so a timed-out upsert is safe to send again. Four writes have no key of your own
to match a second attempt against — a rota, a webhook endpoint, a rotated secret, a delivery sent
again — and those take a header instead:

```go
client.Schedules.Create(ctx, body, clockster.WithIdempotencyKey(attempt))
```

## Uploading a file

One operation carries bytes rather than JSON:

```go
file, err := os.Open("agreement.pdf")
if err != nil {
	return err
}

defer file.Close()

stored, err := client.Files.Upload(ctx, &clockster.FilesUploadForm{
	File:     file,
	Filename: "agreement.pdf",
	Name:     clockster.Set("agreement"),
})
```

## Webhooks

Verifying a delivery is the only way to the event it carries, so there is no path that acts on one
that was not verified:

```go
import "github.com/clockster/sdk-go/webhooks"

func handler(w http.ResponseWriter, r *http.Request) {
	event, err := webhooks.VerifyRequest(r, os.Getenv("CLOCKSTER_WEBHOOK_SECRET"))
	if err != nil {
		http.Error(w, "refused", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)

	go handle(event)
}
```

Verify the bytes as received: re-serialising a parsed object does not reproduce what was signed.
Answer 2xx quickly and do the work afterwards — a timeout is retried — and deduplicate on
`event.ID`, since the same event may arrive twice.

## Dates, times and numbers

Instants, dates and clock times are strings, in the shapes the document states, rather than
`time.Time`: a date carries no zone and a clock time is read in the timezone stated beside it, and
converting either would decide something this package does not know. Durations are seconds.
Decimal amounts are JSON numbers rounded to two places; do not accumulate them in binary floating
point.

## Configuration

```go
client, err := clockster.New(token,
	clockster.WithBaseURL("https://demo.clockster.com"), // a demo stand instead of production
	clockster.WithTimeout(60*time.Second),               // the default is 30 seconds
	clockster.WithUserAgent("acme-hr/1.4"),              // names your integration in our request log
	clockster.WithHTTPClient(retrying),                  // a proxy, a retrying wrapper, a recording one
)
```

The key is read per request, so rotating it does not need a new client.

## Generated from the document

`models.gen.go` and `api.gen.go` are written by `internal/cmd/generate` from
`openapi/company-v3.json`, and committed — an API change appears in review as the lines of the
client it moves. To refresh:

```bash
make spec generate check
```

A nightly job compares the committed document with the published one, so drift is noticed here
rather than by you.

## Examples

Two whole integrations live in [examples](examples): a roster sync in, a timesheet export out.

## Versions

This package follows its own semver, unrelated to the version of the API and to the other SDKs. A
new API version would be a major release of this package rather than a second package.

## License

MIT. See [LICENSE](LICENSE).

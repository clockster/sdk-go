# Examples

Two integrations of the shape most of them have: one writes a roster in, one reads a month out.
Both are single files and use the package as published.

```bash
export CLOCKSTER_TOKEN=...   # Settings → API in the web application
```

## `rostersync`

Sync employees from a CSV, and dismiss whoever is no longer in it.

```bash
go run ./examples/rostersync people.csv
```

```csv
external_id,first_name,last_name,email,phone,location_code,location_title
HR-1,Aisulu,Serik,aisulu@example.com,+77010000001,WH-01,Warehouse
HR-2,Bolat,Nurlan,bolat@example.com,+77010000002,WH-01,Warehouse
```

What it shows: writing locations and reading their ids back out of the answer, `external_id` as the
key that makes a second run an update rather than a duplicate, batching at the hundred the endpoint
takes, and dismissing by difference — everybody active here who is not in the file. It also shows
what `Opt` is for: a column the file leaves empty is not written, so the stored value stays.

Run it twice. The second run writes the same people and dismisses nobody, which is the property a
nightly sync needs.

## `timesheetexport`

Export a month of timesheets as CSV, a row per person per day.

```bash
go run ./examples/timesheetexport 2026-08 > august.csv
```

What it shows: walking a listing with `ListAll`, asking for the facts with `Include`, and the two
things that catch people out — times are seconds, and a day nobody was scheduled for answers a null
`Planned` rather than being left out.

## Writing your own

The methods are named after the operations, so the
[API documentation](https://api.clockster.com/openapi/v3.json) reads as the reference for both:
`GET /users` is `clockster.Users.List(...)`, `POST /users/upsert` is `clockster.Users.Upsert(...)`.
Every method answers the parsed body and an error on anything the API refused.

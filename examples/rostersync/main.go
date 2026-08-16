// Sync a roster from a CSV into Clockster, and dismiss whoever is no longer in it.
//
// The shape a real HR integration has: your system knows people by a key of its own, and every
// sync is "here is everybody, make it so". ExternalID is what makes that idempotent — send it on
// every person and the second run updates rather than duplicates.
//
//	export CLOCKSTER_TOKEN=...
//	go run ./examples/rostersync people.csv
//
// The file wants a header row:
//
//	external_id,first_name,last_name,email,phone,location_code,location_title
//
// Nothing is deleted. Somebody missing from the file is dismissed, which frees their seat and
// leaves their attendance, payroll and documents readable.
package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/clockster/sdk-go"
)

// The roster endpoint takes 100 people a call, and says so; the rest is one round trip each.
const batch = 100

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(argv []string) error {
	if len(argv) != 2 {
		return fmt.Errorf("usage: %s people.csv", argv[0])
	}

	token := os.Getenv("CLOCKSTER_TOKEN")
	if token == "" {
		return errors.New("set CLOCKSTER_TOKEN to a company API key (Settings, API)")
	}

	source, err := rows(argv[1])
	if err != nil {
		return err
	}

	var opts []clockster.Option

	// A demo stand answers the same API; production is the default.
	if held := os.Getenv("CLOCKSTER_BASE_URL"); held != "" {
		opts = append(opts, clockster.WithBaseURL(held))
	}

	client, err := clockster.New(token, opts...)
	if err != nil {
		return err
	}

	ctx := context.Background()

	company, err := client.Me(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Syncing %d people into %s.\n", len(source), company.Data.Title)

	locations, err := locationsByCode(ctx, client, source)
	if err != nil {
		return err
	}

	written, err := upsert(ctx, client, source, locations)
	if err != nil {
		return err
	}

	leaving, err := dismissed(ctx, client, keys(source))
	if err != nil {
		return err
	}

	if len(leaving) > 0 {
		if _, err := client.Users.Dismiss(ctx, &clockster.UsersDismissBody{Users: leaving}); err != nil {
			return err
		}
	}

	fmt.Printf("%d written, %d dismissed.\n", written, len(leaving))

	return nil
}

func rows(path string) ([]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%s cannot be read: %w", path, err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("%s has a header row and nothing under it", path)
	}

	header := records[0]
	people := make([]map[string]string, 0, len(records)-1)

	for _, record := range records[1:] {
		person := map[string]string{}

		for index, column := range header {
			if index < len(record) {
				person[strings.TrimSpace(column)] = strings.TrimSpace(record[index])
			}
		}

		people = append(people, person)
	}

	return people, nil
}

// locationsByCode writes the locations the file names, and reads their ids back out of the answer.
//
// An employee is filed against a location by id, and the file knows only its own code — so the
// codes go up as ExternalID and come back beside the id they were given.
func locationsByCode(ctx context.Context, client *clockster.Client, source []map[string]string) (map[string]int64, error) {
	named := map[string]string{}

	for _, person := range source {
		if code := person["location_code"]; code != "" {
			named[code] = person["location_title"]
		}
	}

	if len(named) == 0 {
		return map[string]int64{}, nil
	}

	codes := make([]string, 0, len(named))

	for code := range named {
		codes = append(codes, code)
	}

	sort.Strings(codes)

	items := make([]clockster.LocationsUpsertItem, 0, len(codes))

	for _, code := range codes {
		items = append(items, clockster.LocationsUpsertItem{ExternalID: code, Title: named[code]})
	}

	answer, err := client.Locations.Upsert(ctx, &clockster.LocationsUpsertBody{Items: items})
	if err != nil {
		return nil, err
	}

	found := map[string]int64{}

	for _, outcome := range answer.Data {
		if outcome.ExternalID != nil {
			found[*outcome.ExternalID] = outcome.ID
		}
	}

	return found, nil
}

func person(row map[string]string, locations map[string]int64) clockster.UsersUpsertUser {
	held := clockster.UsersUpsertUser{
		ExternalID: clockster.Set(row["external_id"]),
		FirstName:  row["first_name"],
		Role:       "employee",
		LocationID: locations[row["location_code"]],
	}

	// Set only where the file has something to say: an empty string is not an absent value, and
	// writing one would blank a field somebody filled in the web application. A field left unset
	// is not written at all, so the stored value stays.
	if value := row["last_name"]; value != "" {
		held.LastName = clockster.Set(value)
	}

	if value := row["email"]; value != "" {
		held.Email = clockster.Set(value)
	}

	if value := row["phone"]; value != "" {
		held.Phone = clockster.Set(value)
	}

	return held
}

func upsert(ctx context.Context, client *clockster.Client, source []map[string]string, locations map[string]int64) (int, error) {
	people := make([]clockster.UsersUpsertUser, 0, len(source))

	for _, row := range source {
		people = append(people, person(row, locations))
	}

	written := 0

	for start := 0; start < len(people); start += batch {
		end := min(start+batch, len(people))

		answer, err := client.Users.Upsert(ctx, &clockster.UsersUpsertBody{Users: people[start:end]})
		if err != nil {
			// None of the batch landed: the write is all or nothing, so the file is fixable and the
			// whole run can be repeated.
			var refused *clockster.Error

			if errors.As(err, &refused) && errors.Is(err, clockster.ErrValidation) {
				return written, fmt.Errorf("refused: %s %v", refused.Code, refused.Errors)
			}

			return written, err
		}

		written += len(answer.Data)

		fmt.Printf("  %d/%d\n", written, len(people))
	}

	return written, nil
}

// dismissed is everyone on file here, not in the roster, and not already gone.
func dismissed(ctx context.Context, client *clockster.Client, held map[string]bool) ([]clockster.UsersDismissUser, error) {
	var leaving []clockster.UsersDismissUser

	for employee, err := range client.Users.ListAll(ctx, &clockster.UsersListParams{
		PerPage: clockster.Set(int64(batch)),
		Status:  clockster.Set("active"),
	}) {
		if err != nil {
			return nil, err
		}

		if employee.ExternalID != nil && !held[*employee.ExternalID] {
			leaving = append(leaving, clockster.UsersDismissUser{ExternalID: clockster.Set(*employee.ExternalID)})
		}
	}

	return leaving, nil
}

func keys(source []map[string]string) map[string]bool {
	held := map[string]bool{}

	for _, row := range source {
		held[row["external_id"]] = true
	}

	return held
}

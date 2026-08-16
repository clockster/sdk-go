// Export a month of timesheets as CSV, a row per person per day.
//
//	export CLOCKSTER_TOKEN=...
//	go run ./examples/timesheetexport 2026-08 > august.csv
//
// What it shows: walking a listing with ListAll, asking for the facts with Include, and the two
// things that catch people out — times are seconds, and a day nobody was scheduled for answers a
// null Planned rather than being left out.
package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/clockster/sdk-go"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(argv []string) error {
	if len(argv) != 2 {
		return fmt.Errorf("usage: %s 2026-08", argv[0])
	}

	token := os.Getenv("CLOCKSTER_TOKEN")
	if token == "" {
		return errors.New("set CLOCKSTER_TOKEN to a company API key (Settings, API)")
	}

	from, to, err := month(argv[1])
	if err != nil {
		return err
	}

	var opts []clockster.Option

	if held := os.Getenv("CLOCKSTER_BASE_URL"); held != "" {
		opts = append(opts, clockster.WithBaseURL(held))
	}

	client, err := clockster.New(token, opts...)
	if err != nil {
		return err
	}

	out := csv.NewWriter(os.Stdout)
	defer out.Flush()

	if err := out.Write([]string{
		"date", "external_id", "name", "planned_start", "planned_end",
		"planned_seconds", "worked_seconds", "late_seconds", "overworked_seconds",
	}); err != nil {
		return err
	}

	// The whole month in one walk: the cursor is the iterator's business, the filters are ours.
	for row, err := range client.Timesheets.ListAll(context.Background(), &clockster.TimesheetsListParams{
		DateFrom: from,
		DateTo:   to,
		Include:  []string{"actual", "variance"},
	}) {
		if err != nil {
			return err
		}

		if err := out.Write(record(row)); err != nil {
			return err
		}
	}

	return out.Error()
}

func month(held string) (string, string, error) {
	first, err := time.Parse("2006-01", held)
	if err != nil {
		return "", "", fmt.Errorf("%q is not a month, which is written 2026-08", held)
	}

	return first.Format("2006-01-02"), first.AddDate(0, 1, -1).Format("2006-01-02"), nil
}

func record(row clockster.TimesheetsListRow) []string {
	held := []string{row.Date, clockster.Deref(row.User.ExternalID), name(row.User)}

	// A day nobody was scheduled for answers null here rather than being left out of the listing,
	// and a day nobody clocked in on answers null for the actual.
	if row.Planned == nil {
		held = append(held, "", "", "")
	} else {
		held = append(held, row.Planned.Start, row.Planned.End, strconv.FormatInt(row.Planned.TimePlanned, 10))
	}

	worked := ""

	if row.Actual != nil {
		worked = strconv.FormatInt(row.Actual.TimeWorked, 10)
	}

	late, overworked := "", ""

	if row.Variance != nil {
		late = strconv.FormatInt(row.Variance.TimeLate, 10)
		overworked = strconv.FormatInt(row.Variance.TimeOverworked, 10)
	}

	return append(held, worked, late, overworked)
}

func name(user clockster.TimesheetsListRowUser) string {
	parts := []string{clockster.Deref(user.FirstName), clockster.Deref(user.LastName)}

	return strings.TrimSpace(strings.Join(parts, " "))
}

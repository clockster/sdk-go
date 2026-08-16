// Walking a listing, against a stub listing rather than the API.

package clockster

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// listing answers `pages` pages of rows, recording the cursor it was asked for each time.
func listing(t *testing.T, pages [][]int64, asked *[]string, queries *[]url.Values) *Client {
	t.Helper()

	return serveFunc(t, func(w http.ResponseWriter, r *http.Request) {
		cursor := r.URL.Query().Get("cursor")

		if asked != nil {
			*asked = append(*asked, cursor)
		}

		if queries != nil {
			*queries = append(*queries, r.URL.Query())
		}

		index := 0

		if cursor != "" {
			index, _ = strconv.Atoi(cursor)
		}

		next := "null"

		if index+1 < len(pages) {
			next = strconv.Quote(strconv.Itoa(index + 1))
		}

		rows := make([]string, 0, len(pages[index]))

		for _, id := range pages[index] {
			rows = append(rows, fmt.Sprintf(`{"id":%d,"first_name":"row"}`, id))
		}

		fmt.Fprintf(w, `{"data":[%s],"links":{"first":null,"last":null,"prev":null,"next":null},`+
			`"meta":{"path":"/users","per_page":50,"next_cursor":%s,"prev_cursor":null}}`,
			strings.Join(rows, ","), next)
	})
}

func TestWalksEveryRowAcrossEveryPage(t *testing.T) {
	var asked []string

	client := listing(t, [][]int64{{1, 2}, {3}, {4, 5}}, &asked, nil)

	var found []int64

	for row, err := range client.Users.ListAll(context.Background(), nil) {
		if err != nil {
			t.Fatalf("the walk was refused: %v", err)
		}

		found = append(found, row.ID)
	}

	if fmt.Sprint(found) != "[1 2 3 4 5]" {
		t.Fatalf("the rows walked are %v", found)
	}

	// The first page carries no cursor, and every page after it carries the one it was given.
	if fmt.Sprint(asked) != "[ 1 2]" {
		t.Fatalf("the cursors asked for are %q", asked)
	}
}

func TestPassesTheFiltersThroughUnchanged(t *testing.T) {
	var queries []url.Values

	client := listing(t, [][]int64{{1}, {2}}, nil, &queries)

	for _, err := range client.Users.ListAll(context.Background(), &UsersListParams{
		PerPage: Set(int64(100)),
		Include: []string{"location"},
	}) {
		if err != nil {
			t.Fatalf("the walk was refused: %v", err)
		}
	}

	for index, query := range queries {
		if query.Get("per_page") != "100" || query.Get("include") != "location" {
			t.Fatalf("page %d asked %v", index, query)
		}
	}
}

func TestStopsOnACursorThatRepeats(t *testing.T) {
	calls := 0

	client := serveFunc(t, func(w http.ResponseWriter, r *http.Request) {
		calls++

		// A cursor that repeats would page until the process is killed.
		fmt.Fprint(w, `{"data":[{"id":1,"first_name":"row"}],`+
			`"links":{"first":null,"last":null,"prev":null,"next":null},`+
			`"meta":{"path":"/users","per_page":50,"next_cursor":"same","prev_cursor":null}}`)
	})

	rows := 0

	for _, err := range client.Users.ListAll(context.Background(), nil) {
		if err != nil {
			t.Fatalf("the walk was refused: %v", err)
		}

		rows++
	}

	if calls != 2 || rows != 2 {
		t.Fatalf("%d calls and %d rows before it stopped", calls, rows)
	}
}

func TestARefusedPageIsAnsweredWhereItWasRefused(t *testing.T) {
	calls := 0

	client := serveFunc(t, func(w http.ResponseWriter, r *http.Request) {
		calls++

		if calls > 1 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `{"error":{"code":"unavailable","message":"Down.","request_id":"1"}}`)

			return
		}

		fmt.Fprint(w, `{"data":[{"id":1,"first_name":"row"}],`+
			`"links":{"first":null,"last":null,"prev":null,"next":null},`+
			`"meta":{"path":"/users","per_page":50,"next_cursor":"1","prev_cursor":null}}`)
	})

	var (
		rows    int
		refused error
	)

	for _, err := range client.Users.ListAll(context.Background(), nil) {
		if err != nil {
			refused = err

			break
		}

		rows++
	}

	// Half a listing is never mistaken for the whole of one.
	if rows != 1 || refused == nil {
		t.Fatalf("%d rows and %v", rows, refused)
	}
}

func TestBreakingOutStopsTheWalk(t *testing.T) {
	calls := 0

	client := serveFunc(t, func(w http.ResponseWriter, r *http.Request) {
		calls++

		fmt.Fprint(w, `{"data":[{"id":1,"first_name":"row"},{"id":2,"first_name":"row"}],`+
			`"links":{"first":null,"last":null,"prev":null,"next":null},`+
			`"meta":{"path":"/users","per_page":50,"next_cursor":"1","prev_cursor":null}}`)
	})

	for range client.Users.ListAll(context.Background(), nil) {
		break
	}

	if calls != 1 {
		t.Fatalf("the walk asked for %d pages after it was left", calls)
	}
}

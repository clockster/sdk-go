// Refresh openapi/company-v3.json from the deployment that implements it.
//
// `-check` writes nothing and exits 1 on a difference, which is what CI runs nightly.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"time"
)

const target = "openapi/company-v3.json"

const published = "https://api.clockster.com/openapi/v3.json"

func main() {
	check := flag.Bool("check", false, "write nothing, and answer 1 on a difference")
	flag.Parse()

	source := published

	if held := os.Getenv("CLOCKSTER_SPEC_URL"); held != "" {
		source = held
	}

	current, err := fetch(source)
	if err != nil {
		fail("%v", err)
	}

	committed, err := os.ReadFile(target)
	if err != nil && !os.IsNotExist(err) {
		fail("%s cannot be read: %v", target, err)
	}

	// Compared as documents rather than as text: the builder writes four-space indentation and
	// this writes two, and that difference is not drift.
	if same(committed, current) {
		fmt.Println("Specification is current.")

		return
	}

	if *check {
		fail("The specification has drifted from %s. Run `make spec generate`.", source)
	}

	var written bytes.Buffer

	if err := json.Indent(&written, current, "", "  "); err != nil {
		fail("%s did not answer JSON: %v", source, err)
	}

	written.WriteString("\n")

	if err := os.WriteFile(target, written.Bytes(), 0o644); err != nil {
		fail("%s cannot be written: %v", target, err)
	}

	fmt.Println("Specification updated. Run `make generate`.")
}

func fetch(source string) ([]byte, error) {
	client := &http.Client{Timeout: 60 * time.Second}

	response, err := client.Get(source)
	if err != nil {
		return nil, fmt.Errorf("%s cannot be read: %w", source, err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s answered %d", source, response.StatusCode)
	}

	return io.ReadAll(response.Body)
}

func same(left, right []byte) bool {
	var first, second any

	if err := json.Unmarshal(left, &first); err != nil {
		return false
	}

	if err := json.Unmarshal(right, &second); err != nil {
		return false
	}

	return reflect.DeepEqual(first, second)
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

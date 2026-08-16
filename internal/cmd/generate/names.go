package main

import (
	"regexp"
	"strings"
	"unicode"
)

// Resource-controller verbs of the operation ids, mapped to the vocabulary of the SDK. The same
// table the TypeScript and Python clients use, so one operation is called one thing in all three.
var verbs = map[string]string{
	"index":   "list",
	"show":    "get",
	"store":   "create",
	"destroy": "delete",
}

// Operations whose summary names a verb the map cannot reach, and two listings without an `index`.
var overrides = map[string][]string{
	"company.v3.attendance.store": {"attendance", "record"},
	"company.v3.files.store":      {"files", "upload"},
	"company.v3.webhooks.secret":  {"webhooks", "rotate_secret"},
	"company.v3.payroll.payslips": {"payroll", "payslips", "list"},
	"company.v3.webhooks.events":  {"webhooks", "events", "list"},
}

// Words Go writes in capitals. A field is `ExternalID` rather than `ExternalId`, which is what a
// reader of any other Go package expects.
var initialisms = map[string]string{
	"id":    "ID",
	"ids":   "IDs",
	"url":   "URL",
	"urls":  "URLs",
	"uri":   "URI",
	"api":   "API",
	"http":  "HTTP",
	"https": "HTTPS",
	"uuid":  "UUID",
	"ip":    "IP",
	"sms":   "SMS",
	"json":  "JSON",
	"utc":   "UTC",
	"iso":   "ISO",
	"bin":   "BIN",
	"iin":   "IIN",
}

var keywords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true, "interface": true,
	"map": true, "package": true, "range": true, "return": true, "select": true,
	"struct": true, "switch": true, "type": true, "var": true,
}

var separators = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func words(name string) []string {
	var out []string

	for _, part := range separators.Split(name, -1) {
		if part != "" {
			out = append(out, part)
		}
	}

	return out
}

// exported is a Go name for a snake_case or dotted one: `external_id` is `ExternalID`.
func exported(name string) string {
	var out strings.Builder

	for _, part := range words(name) {
		lower := strings.ToLower(part)

		if capitalised, ok := initialisms[lower]; ok {
			out.WriteString(capitalised)

			continue
		}

		out.WriteString(strings.ToUpper(part[:1]) + strings.ToLower(part[1:]))
	}

	written := out.String()

	// A name cannot start with a digit; nothing in this document does, and a later one would
	// otherwise emit source that does not compile.
	if written != "" && unicode.IsDigit(rune(written[0])) {
		return "N" + written
	}

	return written
}

// unexported is a Go name for an argument: `user_id` is `userID`.
func unexported(name string) string {
	parts := words(name)

	if len(parts) == 0 {
		return "value"
	}

	out := strings.ToLower(parts[0])

	for _, part := range parts[1:] {
		out += exported(part)
	}

	if keywords[out] {
		return out + "Value"
	}

	return out
}

// singular is what one row of a listing is called: `Shifts` is a list of `Shift`, `Data` is a list
// of `Row`.
func singular(name string) string {
	switch {
	case strings.HasSuffix(name, "Data"):
		return strings.TrimSuffix(name, "Data") + "Row"
	case strings.HasSuffix(name, "ies"):
		return strings.TrimSuffix(name, "ies") + "y"
	case strings.HasSuffix(name, "ss"), !strings.HasSuffix(name, "s"):
		return name + "Item"
	default:
		return strings.TrimSuffix(name, "s")
	}
}

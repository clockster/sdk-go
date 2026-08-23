package main

import (
	"encoding/json"
	"sort"
	"strings"
)

// The closed sets of values, as constants a caller can reach for.
//
// The name of each is read rather than worked out: the document carries it as `x-clockster-set`, so
// this client and the other three call one set one thing without four generators each needing the
// same rule. What is decided here is only what Go does with a name it is given.
//
// Constants rather than a named type. A field stays a `string`, so nothing existing changes shape
// and a value the API starts answering with tomorrow still reaches the caller as itself — which is
// safe precisely because only what you send is ever a closed set: every one of these is on a query
// parameter or in a request body, and none is in an answer.

// rememberSet records a set met while walking, under the name the document gives it.
func (m *models) rememberSet(sch *schema) {
	if sch.Set == "" {
		return
	}

	values := make([]string, 0, len(sch.Enum))

	for _, value := range sch.Enum {
		held, ok := value.(string)
		if !ok {
			return
		}

		values = append(values, held)
	}

	held, seen := m.sets[sch.Set]
	if !seen {
		m.sets[sch.Set] = values

		return
	}

	// One name over two different sets would publish constants for one of them and a note about
	// the other. The document has a test against this; so does this.
	if strings.Join(held, "\x00") != strings.Join(values, "\x00") {
		fail("%s names two different sets of values.", sch.Set)
	}
}

func setsSource(m *models) string {
	head := "// Code generated from " + specPath + `; DO NOT EDIT.

// The closed sets of values the Company API accepts, named by the document.
//
// Constants rather than a named type of their own, so one goes wherever the string goes:
// ` + "`clockster.UsersRoleEmployee`" + ` is ` + "`\"employee\"`" + ` and drops into the field that takes it. The
// fields stay strings, so a value the API starts answering with tomorrow reaches you as itself.
//
// Only what you send is ever closed. Every set here is on a query parameter or in a request body,
// and none of them is in an answer.

package clockster

`

	names := make([]string, 0, len(m.sets))

	for name := range m.sets {
		names = append(names, name)
	}

	sort.Strings(names)

	var body strings.Builder

	for index, name := range names {
		if index > 0 {
			body.WriteString("\n")
		}

		body.WriteString(setSource(m, name, m.sets[name]))
	}

	return head + body.String()
}

func setSource(m *models, name string, values []string) string {
	written := make([]string, 0, len(values))
	constants := make([]string, 0, len(values))

	for _, value := range values {
		encoded, err := json.Marshal(value)
		if err != nil {
			fail("a value of %s cannot be written: %v", name, value)
		}

		written = append(written, string(encoded))

		constant := name + exported(value)

		if m.taken[constant] {
			fail("%s is a name this package already holds. Rename the set in the document's overlay.", constant)
		}

		constants = append(constants, constant)
	}

	var out strings.Builder

	out.WriteString(docComment(name+" is what this field is allowed to be. One of "+strings.Join(written, ", ")+".", ""))
	out.WriteString("const (\n")

	for index, constant := range constants {
		out.WriteString("\t" + constant + " = " + written[index] + "\n")
	}

	out.WriteString(")\n\n")
	out.WriteString(docComment(name+"Values is every value of "+name+", in the order the document names them.", ""))
	out.WriteString("func " + name + "Values() []string {\n\treturn []string{\n")

	for _, constant := range constants {
		out.WriteString("\t\t" + constant + ",\n")
	}

	out.WriteString("\t}\n}\n")

	return out.String()
}

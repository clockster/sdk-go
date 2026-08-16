// Generate models.gen.go and api.gen.go from openapi/company-v3.json.
//
// The output is committed, so an API change appears in review as the lines of the client it moves.
//
// Two files come out of it. models.gen.go holds a struct per shape the document describes: the
// components it names, and one per request body, query and answer besides. api.gen.go holds the
// namespaces — clockster.Users.List(ctx, …) — the client fields that reach them, and an iterator
// over every listing that pages on a cursor.
//
// Nothing here validates. A response is the JSON as it arrived, decoded into the shape the
// document describes; a field the API adds tomorrow reaches the caller today rather than being
// refused on the way in.
package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

// What the comments this writes are wrapped to, which is what the rest of the repository is
// written to.
const commentWidth = 96

const (
	specPath   = "openapi/company-v3.json"
	modelsPath = "models.gen.go"
	apiPath    = "api.gen.go"
)

// Exported names the hand-written half of the package already holds. A generated shape that wants
// one of these is renamed rather than shadowing it.
var handWritten = []string{
	"Client", "Error", "Opt", "Option", "RequestOption", "Set", "Null", "Ptr", "Deref",
	"Version", "DefaultBaseURL", "DefaultTimeout", "ErrNoToken", "ErrAuthentication",
	"ErrForbidden", "ErrNotFound", "ErrConflict", "ErrValidation", "ErrRateLimit", "ErrServer",
	"New", "WithBaseURL", "WithHTTPClient", "WithTimeout", "WithUserAgent", "WithHeader",
	"WithIdempotencyKey", "transport", "request", "formBody",
}

func main() {
	raw, err := os.ReadFile(specPath)
	if err != nil {
		fail("%s cannot be read: %v", specPath, err)
	}

	doc := &document{}

	if err := json.Unmarshal(raw, doc); err != nil {
		fail("%s is not the document this generator reads: %v", specPath, err)
	}

	endpoints := parse(doc)
	tree := namespaceTree(endpoints)

	reserved := append([]string{}, handWritten...)

	for _, node := range tree.all() {
		reserved = append(reserved, node.typeName)
	}

	m := newModels(doc, reserved)

	for _, e := range endpoints {
		e.types(m)
	}

	write(modelsPath, modelsSource(m))
	write(apiPath, apiSource(endpoints, tree))

	fmt.Printf("%d operations, %d shapes.\n", len(endpoints), len(m.order))
}

func parse(doc *document) []*endpoint {
	var endpoints []*endpoint

	for path, item := range doc.Paths {
		for method, spec := range item {
			endpoints = append(endpoints, parseEndpoint(path, method, spec))
		}
	}

	// Sorted by where the method lands rather than by path, so the file reads as the client does
	// and two runs write the same bytes.
	sort.Slice(endpoints, func(i, j int) bool {
		left, right := endpoints[i], endpoints[j]

		if left.namespace != right.namespace {
			return left.namespace < right.namespace
		}

		return left.name < right.name
	})

	return endpoints
}

// A group of operations, and the groups under it.
type namespace struct {
	typeName string
	field    string
	// How a caller reaches it: `Webhooks.Deliveries`.
	path     string
	children []*namespace
}

func (n *namespace) child(field, typeName string) *namespace {
	for _, held := range n.children {
		if held.field == field {
			return held
		}
	}

	held := &namespace{typeName: typeName, field: field, path: strings.TrimPrefix(n.path+"."+field, ".")}
	n.children = append(n.children, held)

	sort.Slice(n.children, func(i, j int) bool { return n.children[i].field < n.children[j].field })

	return held
}

func (n *namespace) all() []*namespace {
	var found []*namespace

	for _, held := range n.children {
		found = append(found, held)
		found = append(found, held.all()...)
	}

	return found
}

func namespaceTree(endpoints []*endpoint) *namespace {
	root := &namespace{}

	for _, e := range endpoints {
		node := root

		for depth := 1; depth < len(e.segments); depth++ {
			node = node.child(exported(e.segments[depth-1]), exported(strings.Join(e.segments[:depth], "_")))
		}
	}

	return root
}

func modelsSource(m *models) string {
	head := "// Code generated from " + specPath + `; DO NOT EDIT.

// Every shape the Company API answers with or accepts.
//
// A field the document marks optional is a pointer in an answer and an Opt in a request: a
// relation is absent from an answer unless ` + "`include`" + ` asked for it, and a field left out of a
// write keeps whatever is stored where a null one clears it.

package clockster

`

	body := m.body()

	return head + imports(body) + "\n" + body
}

func apiSource(endpoints []*endpoint, tree *namespace) string {
	var body strings.Builder

	body.WriteString(clientSource(tree))

	for _, node := range tree.all() {
		body.WriteString(namespaceSource(node))
	}

	for _, e := range endpoints {
		body.WriteString("\n" + e.source())

		if walk := e.walkSource(); walk != "" {
			body.WriteString("\n" + walk)
		}
	}

	head := `// Code generated from openapi/company-v3.json; DO NOT EDIT.

// The operations of the Company API, as they are called.
//
// Every method answers the parsed body of the response, and a refusal is an error rather than a
// status to read — so there is no branch between the call and the rows.

package clockster

`

	return head + imports(body.String()) + "\n" + body.String()
}

// imports is what the generated file uses, worked out from what it wrote rather than declared
// ahead of it.
func imports(body string) string {
	needed := []struct {
		token string
		path  string
	}{
		{"context.Context", "context"},
		{"fmt.Sprintf", "fmt"},
		{"io.Reader", "io"},
		{"iter.Seq2", "iter"},
		{"http.Method", "net/http"},
		{"url.Values", "net/url"},
	}

	var used []string

	for _, held := range needed {
		if strings.Contains(body, held.token) {
			used = append(used, "\t\""+held.path+"\"")
		}
	}

	if len(used) == 0 {
		return ""
	}

	return "import (\n" + strings.Join(used, "\n") + "\n)\n"
}

// clientSource is the type a caller holds. Written here rather than beside New, so the groups of
// operations are fields of it and read as such in the documentation.
func clientSource(tree *namespace) string {
	lines := []string{
		"// Client is the Company API, one token to one company.",
		"//",
		"// Build one with [New]. Every operation hangs off a field below, named after the section of",
		"// the documentation it belongs to: `clockster.Users.List(ctx, …)`. A Client is safe for",
		"// concurrent use.",
		"type Client struct {",
		"\t*transport",
	}

	for _, node := range tree.children {
		lines = append(lines, "", fmt.Sprintf("\t// %s is `clockster.%s`.", node.field, node.path))
		lines = append(lines, fmt.Sprintf("\t%s *%s", node.field, node.typeName))
	}

	lines = append(lines, "}", "", "func newClient(t *transport) *Client {", "\treturn &Client{", "\t\ttransport: t,")

	for _, node := range tree.children {
		lines = append(lines, fmt.Sprintf("\t\t%s: new%s(t),", node.field, node.typeName))
	}

	return strings.Join(append(lines, "\t}", "}", ""), "\n") + "\n"
}

func namespaceSource(node *namespace) string {
	lines := []string{
		fmt.Sprintf("// %s holds the operations of `clockster.%s`.", node.typeName, node.path),
		fmt.Sprintf("type %s struct {", node.typeName),
		"\t*transport",
	}

	for _, child := range node.children {
		lines = append(lines, "", fmt.Sprintf("\t// %s is `clockster.%s`.", child.field, child.path))
		lines = append(lines, fmt.Sprintf("\t%s *%s", child.field, child.typeName))
	}

	lines = append(lines, "}", "", fmt.Sprintf("func new%s(t *transport) *%s {", node.typeName, node.typeName))
	lines = append(lines, fmt.Sprintf("\treturn &%s{", node.typeName), "\t\ttransport: t,")

	for _, child := range node.children {
		lines = append(lines, fmt.Sprintf("\t\t%s: new%s(t),", child.field, child.typeName))
	}

	return strings.Join(append(lines, "\t}", "}", ""), "\n") + "\n"
}

func write(path, source string) {
	formatted, err := format.Source([]byte(source))
	if err != nil {
		broken := path + ".broken"

		_ = os.WriteFile(broken, []byte(source), 0o644)

		fail("what was written for %s does not parse: %v. The source is in %s.", path, err, broken)
	}

	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		fail("%s cannot be written: %v", path, err)
	}
}

// docComment writes prose from the document as a Go comment. A fenced block becomes an indented
// one, which is how godoc reads code.
func docComment(text, indent string) string {
	var out strings.Builder

	fenced := false

	for _, line := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			fenced = !fenced

			continue
		}

		switch {
		case fenced:
			out.WriteString(indent + "//\t" + line + "\n")
		case strings.TrimSpace(line) == "":
			out.WriteString(indent + "//\n")
		default:
			for _, wrapped := range wrap(line, commentWidth-utf8.RuneCountInString(indent)) {
				out.WriteString(indent + "// " + wrapped + "\n")
			}
		}
	}

	return out.String()
}

// wrap breaks a line of prose to the width the rest of this repository is written to. The
// document wraps its own paragraphs; a parameter's description is one long line.
func wrap(line string, width int) []string {
	// A table row is not prose, and breaking one destroys it.
	if utf8.RuneCountInString(line) <= width || strings.Contains(line, "|") {
		return []string{line}
	}

	margin := line[:len(line)-len(strings.TrimLeft(line, " \t"))]

	var (
		wrapped []string
		current string
	)

	for _, word := range strings.Fields(line) {
		switch {
		case current == "":
			current = margin + word
		case utf8.RuneCountInString(current+" "+word) <= width:
			current += " " + word
		default:
			wrapped = append(wrapped, current)
			current = margin + word
		}
	}

	return append(wrapped, current)
}

// commentLines is prose written above a field of a struct, wrapped and one line to an entry.
func commentLines(text string) []string {
	return strings.Split(strings.TrimSuffix(docComment(text, "\t"), "\n"), "\n")
}

func lowerFirst(text string) string {
	if text == "" {
		return text
	}

	return strings.ToLower(text[:1]) + text[1:]
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

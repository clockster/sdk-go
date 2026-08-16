package main

import (
	"fmt"
	"net/http"
	"strings"
)

const prefix = "company.v3."

const base = "/company/v3"

// endpoint is one route, and everything the client needs to call it.
type endpoint struct {
	path   string
	method string
	spec   *operation

	// The segments of the operation id after `company.v3`, with the verb mapped, and the Go names
	// they turn into.
	segments  []string
	namespace string
	receiver  string
	name      string
	stem      string

	pathParams []*parameter
	query      []*parameter
	idempotent bool

	paramsType string
	bodyType   string
	formType   string
	returns    string

	// Set on a listing that pages on a cursor, which is what ListAll walks.
	rowType     string
	metaField   string
	metaPointer bool
	cursorField string
}

func parseEndpoint(path, method string, spec *operation) *endpoint {
	segments, ok := overrides[spec.ID]

	if !ok {
		segments = strings.Split(strings.TrimPrefix(spec.ID, prefix), ".")
		last := len(segments) - 1

		if mapped, found := verbs[segments[last]]; found {
			segments[last] = mapped
		}
	}

	e := &endpoint{
		path:     path,
		method:   strings.ToUpper(method),
		spec:     spec,
		segments: segments,
		name:     exported(segments[len(segments)-1]),
		stem:     exported(strings.Join(segments, "_")),
		receiver: "c",
	}

	if group := segments[:len(segments)-1]; len(group) > 0 {
		e.namespace = exported(strings.Join(group, "_"))
		e.receiver = "n"
	}

	for _, held := range spec.Parameters {
		switch held.In {
		case "path":
			e.pathParams = append(e.pathParams, held)
		case "query":
			e.query = append(e.query, held)
		case "header":
			e.idempotent = e.idempotent || held.Name == "Idempotency-Key"
		}
	}

	return e
}

// types names everything this endpoint takes and answers, and registers the shapes.
func (e *endpoint) types(m *models) {
	m.scope = e.stem
	m.origin = e.method + " " + e.path

	e.declareParams(m)
	e.declareBody(m)
	e.declareAnswer(m)
}

func (e *endpoint) declareParams(m *models) {
	if len(e.query) == 0 {
		return
	}

	name := m.unique(e.stem + "Params")
	m.reserve(name)

	lines := []string{
		fmt.Sprintf("// %s is the query of %s %s.", name, e.method, e.path),
		fmt.Sprintf("type %s struct {", name),
	}

	for _, held := range e.query {
		lines = append(lines, fields(lines, m.fieldSource(name, held.Name, held.Schema, held.Required, sent, held.Description))...)

		if held.Name == "cursor" {
			e.cursorField = exported(held.Name)
		}
	}

	m.blocks[name] = strings.Join(append(lines, "}"), "\n") + "\n"
	e.paramsType = name
}

func (e *endpoint) declareBody(m *models) {
	if e.spec.RequestBody == nil {
		return
	}

	if held := e.spec.RequestBody.schema("application/json"); held != nil {
		e.bodyType = m.typeOf(held, e.stem+"Body", sent).name

		return
	}

	held := e.spec.RequestBody.schema("multipart/form-data")

	if held == nil {
		fail("%s takes a body in a format this generator does not know.", e.spec.ID)
	}

	name := m.unique(e.stem + "Form")
	m.reserve(name)

	lines := []string{
		fmt.Sprintf("// %s is what %s %s takes. It carries bytes rather than JSON,", name, e.method, e.path),
		"// which is what makes this the one operation sent as multipart.",
		fmt.Sprintf("type %s struct {", name),
		"\t// File is the bytes to store, read to the end before the request is sent.",
		"\tFile io.Reader",
		"",
		"\t// Filename is the name the bytes travel under. Empty is sent as \"upload\".",
		"\tFilename string",
	}

	for _, key := range held.Properties.keys {
		if key == "file" {
			continue
		}

		field := held.Properties.get(key)

		lines = append(lines, "")
		lines = append(lines, m.fieldSource(name, key, field, held.requires(key), sent, field.Description)...)
	}

	m.blocks[name] = strings.Join(append(lines, "}"), "\n") + "\n"
	e.formType = name
}

func (e *endpoint) declareAnswer(m *models) {
	status := "200"

	if _, ok := e.spec.Responses["201"]; ok {
		status = "201"
	}

	held := e.spec.Responses[status].schema("application/json")

	if held == nil {
		return
	}

	answered := m.typeOf(held, e.stem+"Response", answer)

	if answered.kind != kindStruct {
		fail("%s answers %s, which is not a shape this client can name.", e.spec.ID, answered.name)
	}

	e.returns = answered.name
	e.findCursor(m, held, answered.name)
}

// findCursor works out whether this listing pages on a cursor, and what one row of it is.
func (e *endpoint) findCursor(m *models, held *schema, name string) {
	if e.cursorField == "" {
		return
	}

	data, meta := held.Properties.get("data"), held.Properties.get("meta")

	if data == nil || meta == nil {
		return
	}

	rows := m.typeOf(data, childHint(name, "data"), answer)

	if rows.kind != kindSlice {
		return
	}

	cursor := meta

	if meta.Ref != "" {
		cursor = m.doc.Components.Schemas.get(meta.Ref[strings.LastIndex(meta.Ref, "/")+1:])
	}

	if cursor.Properties.get("next_cursor") == nil {
		return
	}

	e.rowType = strings.TrimPrefix(rows.name, "[]")
	e.metaField = exported("meta")
	e.metaPointer = strings.HasPrefix(render(m.typeOf(meta, childHint(name, "meta"), answer), held.requires("meta"), answer), "*")
}

func (e *endpoint) receiverSource() string {
	if e.namespace == "" {
		return "(c *Client)"
	}

	return fmt.Sprintf("(n *%s)", e.namespace)
}

func (e *endpoint) arguments() []string {
	args := []string{"ctx context.Context"}

	for _, held := range e.pathParams {
		args = append(args, fmt.Sprintf("%s %s", unexported(held.Name), pathType(held)))
	}

	switch {
	case e.bodyType != "":
		args = append(args, "body *"+e.bodyType)
	case e.formType != "":
		args = append(args, "form *"+e.formType)
	}

	if e.paramsType != "" {
		args = append(args, "params *"+e.paramsType)
	}

	return append(args, "opts ...RequestOption")
}

func pathType(held *parameter) string {
	if held.Schema != nil && len(held.Schema.concrete()) > 0 && held.Schema.concrete()[0] == "string" {
		return "string"
	}

	return "int64"
}

func (e *endpoint) pathSource() string {
	if len(e.pathParams) == 0 {
		return fmt.Sprintf("%q", base+strings.TrimPrefix(e.path, base))
	}

	written := e.path
	args := make([]string, 0, len(e.pathParams))

	for _, held := range e.pathParams {
		verb := "%d"

		if pathType(held) == "string" {
			verb = "%s"
		}

		written = strings.Replace(written, "{"+held.Name+"}", verb, 1)
		args = append(args, unexported(held.Name))
	}

	return fmt.Sprintf("fmt.Sprintf(%q, %s)", written, strings.Join(args, ", "))
}

func (e *endpoint) doc() string {
	lines := []string{fmt.Sprintf("// %s is %s %s.", e.name, e.method, e.path)}

	if summary := e.spec.Summary; summary != "" {
		// Stopped rather than left bare: gofmt reads a lone unpunctuated line as a heading and
		// rewrites it as one.
		if !strings.ContainsAny(summary[len(summary)-1:], ".!?") {
			summary += "."
		}

		lines = append(lines, "//", "// "+summary)
	}

	if e.spec.Description != "" {
		lines = append(lines, "//", strings.TrimSuffix(docComment(e.spec.Description, ""), "\n"))
	}

	if e.idempotent {
		lines = append(lines,
			"//",
			"// This write has nothing of your own to match a second attempt against, so a retry of it is",
			"// safe only with WithIdempotencyKey.",
		)
	}

	return strings.Join(lines, "\n")
}

func (e *endpoint) source() string {
	var out strings.Builder

	answers := "error"

	if e.returns != "" {
		answers = fmt.Sprintf("(*%s, error)", e.returns)
	}

	out.WriteString(e.doc() + "\n")
	out.WriteString(fmt.Sprintf("func %s %s(%s) %s {\n", e.receiverSource(), e.name, strings.Join(e.arguments(), ", "), answers))
	out.WriteString(e.guards())
	out.WriteString(e.querySource())
	out.WriteString(e.callSource())
	out.WriteString("}\n")

	return out.String()
}

func (e *endpoint) fails(answer string) string {
	if e.returns == "" {
		return "return " + answer
	}

	return "return nil, " + answer
}

func (e *endpoint) guards() string {
	switch {
	case e.bodyType != "":
		return fmt.Sprintf("if body == nil {\n%s\n}\n\n", e.fails("errNoBody"))
	case e.formType != "":
		return fmt.Sprintf("if form == nil || form.File == nil {\n%s\n}\n\n", e.fails("errNoFile"))
	default:
		return ""
	}
}

func (e *endpoint) querySource() string {
	if e.paramsType == "" {
		return ""
	}

	var out strings.Builder

	out.WriteString("query := url.Values{}\n\n")
	out.WriteString("if params != nil {\n")

	for _, held := range e.query {
		out.WriteString(parameterSource(held))
	}

	out.WriteString("}\n\n")

	return out.String()
}

func parameterSource(held *parameter) string {
	field := "params." + exported(held.Name)

	if len(held.Schema.concrete()) > 0 && held.Schema.concrete()[0] == "array" {
		return fmt.Sprintf("queryList(query, %q, %s)\n", held.Name, field)
	}

	if !held.Required {
		return fmt.Sprintf("queryOpt(query, %q, %s)\n", held.Name, field)
	}

	if held.Schema.nullable() {
		return fmt.Sprintf("if %s != nil {\nquery.Set(%q, scalar(*%s))\n}\n", field, held.Name, field)
	}

	return fmt.Sprintf("query.Set(%q, scalar(%s))\n", held.Name, field)
}

func (e *endpoint) callSource() string {
	var out strings.Builder

	if e.formType != "" {
		out.WriteString(formSource(e))
	}

	target := "nil"

	if e.returns != "" {
		out.WriteString(fmt.Sprintf("var out %s\n\n", e.returns))

		target = "&out"
	}

	parts := []string{
		fmt.Sprintf("method: %s,", methodConstant(e.method)),
		fmt.Sprintf("path:   %s,", e.pathSource()),
	}

	if e.paramsType != "" {
		parts = append(parts, "query:  query,")
	}

	if e.bodyType != "" {
		parts = append(parts, "body:   body,")
	}

	if e.formType != "" {
		parts = append(parts, "form:   &formBody{file: form.File, filename: filename, fields: fields},")
	}

	// The transport is embedded rather than held, so the namespace calls it as its own.
	out.WriteString(fmt.Sprintf("if err := %s.do(ctx, request{\n%s\n}, %s, opts); err != nil {\n%s\n}\n\n",
		e.receiver, strings.Join(parts, "\n"), target, e.fails("err")))

	if e.returns == "" {
		out.WriteString("return nil\n")
	} else {
		out.WriteString("return &out, nil\n")
	}

	return out.String()
}

func formSource(e *endpoint) string {
	var out strings.Builder

	out.WriteString("fields := url.Values{}\n\n")

	for _, key := range formFields(e) {
		out.WriteString(fmt.Sprintf("queryOpt(fields, %q, form.%s)\n", key, exported(key)))
	}

	out.WriteString("\nfilename := form.Filename\n\n")
	out.WriteString("if filename == \"\" {\nfilename = \"upload\"\n}\n\n")

	return out.String()
}

func formFields(e *endpoint) []string {
	held := e.spec.RequestBody.schema("multipart/form-data")

	var keys []string

	for _, key := range held.Properties.keys {
		if key != "file" {
			keys = append(keys, key)
		}
	}

	return keys
}

func methodConstant(method string) string {
	switch method {
	case http.MethodGet:
		return "http.MethodGet"
	case http.MethodPost:
		return "http.MethodPost"
	case http.MethodPut:
		return "http.MethodPut"
	case http.MethodPatch:
		return "http.MethodPatch"
	case http.MethodDelete:
		return "http.MethodDelete"
	default:
		fail("%s is a method this generator does not write.", method)

		return ""
	}
}

// walkSource is the iterator over a listing that pages on a cursor.
func (e *endpoint) walkSource() string {
	if e.rowType == "" {
		return ""
	}

	guard := ""

	if e.metaPointer {
		guard = fmt.Sprintf("if page.%s == nil {\nreturn\n}\n\n", e.metaField)
	}

	return fmt.Sprintf(`// %[1]sAll walks every page of %[1]s and answers a row at a time.
//
//	for row, err := range %[7]s.%[1]sAll(ctx, params) {
//		if err != nil {
//			return err
//		}
//	}
//
// A refused page is answered where it was refused, so half a listing is never mistaken for the
// whole of one. A cursor belongs to the filters it was issued under: change them and walk again.
func %[2]s %[1]sAll(ctx context.Context, params *%[3]s, opts ...RequestOption) iter.Seq2[%[4]s, error] {
	return func(yield func(%[4]s, error) bool) {
		walked := %[3]s{}

		if params != nil {
			walked = *params
		}

		seen := map[string]bool{}

		for {
			page, err := %[8]s.%[1]s(ctx, &walked, opts...)
			if err != nil {
				var none %[4]s

				yield(none, err)

				return
			}

			for _, row := range page.Data {
				if !yield(row, nil) {
					return
				}
			}

			%[9]scursor := page.%[5]s.NextCursor

			// A cursor that repeats would page until the process is killed, which is worse than
			// stopping.
			if cursor == nil || *cursor == "" || seen[*cursor] {
				return
			}

			seen[*cursor] = true
			walked.%[6]s = Set(*cursor)
		}
	}
}
`, e.name, e.receiverSource(), e.paramsType, e.rowType, e.metaField, e.cursorField, exampleReceiver(e), e.receiver, guard)
}

// exampleReceiver is how the call reads in the comment above it: `clockster.Webhooks.Deliveries`.
func exampleReceiver(e *endpoint) string {
	parts := []string{"clockster"}

	for _, segment := range e.segments[:len(e.segments)-1] {
		parts = append(parts, exported(segment))
	}

	return strings.Join(parts, ".")
}

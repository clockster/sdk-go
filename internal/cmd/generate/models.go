package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Which way a shape travels. A field that may be left out is `Opt` on the way out and a pointer on
// the way back: a request has to tell an omitted field from a null one, and a response is read
// rather than written.
type side uint8

const (
	answer side = iota
	sent
)

type kind uint8

const (
	kindScalar kind = iota
	kindStruct
	kindSlice
	kindMap
	kindAny
	kindIface
)

type gotype struct {
	name     string
	kind     kind
	nullable bool
	// What the values are, which the document says and Go cannot.
	note string
	// The shape of the string, said only where the field has nothing else written above it.
	format string
}

// models is every type the document implies, named once and emitted in the order they were needed.
type models struct {
	doc        *document
	order      []string
	blocks     map[string]string
	taken      map[string]bool
	shapes     map[string]string
	components map[string]side

	// One operation's shapes, so the same object under two keys of one answer is named once, and
	// where the shape came from, for the comment above it.
	scope  string
	origin string
}

func newModels(doc *document, reserved []string) *models {
	m := &models{
		doc:        doc,
		blocks:     map[string]string{},
		taken:      map[string]bool{},
		shapes:     map[string]string{},
		components: map[string]side{},
	}

	for _, name := range reserved {
		m.taken[name] = true
	}

	return m
}

func (m *models) reserve(name string) {
	m.taken[name] = true
	m.order = append(m.order, name)
}

func (m *models) unique(hint string) string {
	name := hint

	for attempt := 2; m.taken[name]; attempt++ {
		name = fmt.Sprintf("%s%d", hint, attempt)
	}

	return name
}

// componentType is a shape the document names, which keeps that name.
func (m *models) componentType(ref string, s side) gotype {
	name := ref[strings.LastIndex(ref, "/")+1:]

	if reached, ok := m.components[name]; ok {
		if reached != s {
			fail("%s is used both in a request and in an answer, and the two want different Go "+
				"types. Split it in the document, or emit a second shape here.", name)
		}

		return gotype{name: name, kind: kindStruct}
	}

	m.components[name] = s
	// Reserved before the body is built: a component that reaches itself would otherwise recur
	// forever.
	m.reserve(name)

	// A component is shared by the operations that use it, so it is not written up as belonging to
	// whichever of them reached it first.
	origin := m.origin
	m.origin = ""

	m.blocks[name] = m.structSource(name, m.doc.Components.Schemas.get(name), s)
	m.origin = origin

	return gotype{name: name, kind: kindStruct}
}

func (m *models) typeOf(sch *schema, hint string, s side) gotype {
	if sch == nil {
		return gotype{name: "any", kind: kindAny}
	}

	if sch.Ref != "" {
		return m.componentType(sch.Ref, s)
	}

	if len(sch.OneOf) > 0 {
		return m.union(hint, sch, s)
	}

	rest := sch.concrete()
	nullable := sch.nullable()

	if len(rest) == 0 {
		return gotype{name: "any", kind: kindAny, nullable: nullable}
	}

	if len(rest) > 1 {
		// Neither type is the value's, so the caller reads the JSON as it arrived.
		return gotype{
			name:     "any",
			kind:     kindAny,
			nullable: nullable,
			note:     "One of " + strings.Join(rest, ", ") + ".",
		}
	}

	held := m.oneType(sch, rest[0], hint, s)
	held.nullable = nullable

	return held
}

func (m *models) oneType(sch *schema, declared, hint string, s side) gotype {
	switch declared {
	case "array":
		return m.arrayType(sch, hint, s)
	case "object":
		return m.objectType(sch, hint, s)
	case "string":
		return m.stringType(sch)
	case "integer":
		return gotype{name: "int64", kind: kindScalar}
	case "number":
		return gotype{name: "float64", kind: kindScalar}
	case "boolean":
		return gotype{name: "bool", kind: kindScalar}
	default:
		fail("%s is a type this generator does not know.", declared)

		return gotype{}
	}
}

func (m *models) arrayType(sch *schema, hint string, s side) gotype {
	if sch.Items == nil {
		return gotype{name: "[]any", kind: kindSlice}
	}

	item := m.typeOf(sch.Items, singular(hint), s)
	held := item.name

	// A row that may be null is a pointer; a list of lists or of maps is neither.
	if item.nullable && (item.kind == kindScalar || item.kind == kindStruct) {
		held = "*" + held
	}

	return gotype{name: "[]" + held, kind: kindSlice, note: item.note}
}

func (m *models) objectType(sch *schema, hint string, s side) gotype {
	if sch.Properties.len() > 0 {
		return gotype{name: m.register(hint, sch, s), kind: kindStruct}
	}

	// A map: its keys are values rather than field names, and the document says so.
	if len(sch.AdditionalProperties) > 0 {
		var extra schema

		if err := json.Unmarshal(sch.AdditionalProperties, &extra); err == nil && len(extra.types)+len(extra.Ref) > 0 {
			value := m.typeOf(&extra, hint+"Value", s)

			return gotype{name: "map[string]" + value.name, kind: kindMap}
		}
	}

	return gotype{name: "map[string]any", kind: kindMap}
}

func (m *models) stringType(sch *schema) gotype {
	held := gotype{name: "string", kind: kindScalar}

	if len(sch.Enum) > 0 {
		written := make([]string, 0, len(sch.Enum))

		for _, value := range sch.Enum {
			encoded, err := json.Marshal(value)
			if err != nil {
				fail("an enum value cannot be written: %v", value)
			}

			written = append(written, string(encoded))
		}

		// A string rather than a named type of its own: a value we add tomorrow reaches the caller
		// as itself rather than as something this package refuses to hold.
		held.note = "One of " + strings.Join(written, ", ") + "."

		return held
	}

	switch sch.Format {
	case "date":
		held.format = "A plain date, YYYY-MM-DD."
	case "date-time":
		held.format = "An instant, ISO 8601 with an offset."
	case "time":
		held.format = "A clock time, HH:MM:SS, read in the timezone beside it."
	}

	return held
}

// register names an inline shape, and answers the name of an identical one already seen.
//
// Deduplicated by the name it wants rather than by its keys alone: a listing's `location` and its
// `locations` are one shape under one name, where a `position` shaped exactly like a `department`
// is its own type. Two operations that answer the same keys keep their own names too, which is why
// the hint carries the operation. Shapes the whole surface shares are components in the document,
// and keep the names it gives them.
func (m *models) register(hint string, sch *schema, s side) string {
	key := fmt.Sprintf("%s\x00%d\x00%s", hint, s, sch.signature)

	if name, ok := m.shapes[key]; ok {
		return name
	}

	name := m.unique(hint)
	m.shapes[key] = name

	m.reserve(name)
	m.blocks[name] = m.structSource(name, sch, s)

	return name
}

func (m *models) union(hint string, sch *schema, s side) gotype {
	name := m.unique(hint)
	m.reserve(name)

	variants := make([]string, 0, len(sch.OneOf))

	for _, one := range sch.OneOf {
		variants = append(variants, m.typeOf(one, hint+"Variant", s).name)
	}

	lines := []string{
		fmt.Sprintf("// %s is one of %s.", name, strings.Join(variants, ", ")),
	}

	if sch.Discriminator != nil {
		lines = append(lines,
			"//",
			fmt.Sprintf("// Which one is read off %q, so that field carries the value named beside each type",
				sch.Discriminator.PropertyName),
			"// below rather than any of the values it accepts.",
		)
	}

	lines = append(lines,
		fmt.Sprintf("type %s interface {", name),
		fmt.Sprintf("\tis%s()", name),
		"}",
		"",
	)

	for _, variant := range variants {
		lines = append(lines, fmt.Sprintf("func (%s) is%s() {}", variant, name))
	}

	m.blocks[name] = strings.Join(lines, "\n") + "\n"

	return gotype{name: name, kind: kindIface}
}

func (m *models) structSource(name string, sch *schema, s side) string {
	if sch == nil || sch.Properties.len() == 0 {
		return fmt.Sprintf("%stype %s map[string]any\n", m.comment(name, sch, s), name)
	}

	lines := []string{m.comment(name, sch, s) + fmt.Sprintf("type %s struct {", name)}

	for _, key := range sch.Properties.keys {
		held := sch.Properties.get(key)
		lines = append(lines, fields(lines, m.fieldSource(name, key, held, sch.requires(key), s, held.Description))...)
	}

	lines = append(lines, "}")

	return strings.Join(lines, "\n") + "\n"
}

// fields keeps a field with something written above it apart from the one before, and leaves a
// run of plain ones together.
func fields(written, field []string) []string {
	opening := strings.HasSuffix(written[len(written)-1], "{")

	if len(field) > 1 && !opening {
		return append([]string{""}, field...)
	}

	return field
}

func (m *models) comment(name string, sch *schema, s side) string {
	if sch != nil && sch.Description != "" {
		return docComment(name+" "+lowerFirst(sch.Description), "")
	}

	if m.origin == "" {
		return ""
	}

	if s == sent {
		return fmt.Sprintf("// %s is part of what %s takes.\n", name, m.origin)
	}

	return fmt.Sprintf("// %s is part of what %s answers.\n", name, m.origin)
}

func (m *models) fieldSource(parent, key string, sch *schema, required bool, s side, description string) []string {
	held := m.typeOf(sch, childHint(parent, key), s)
	written := render(held, required, s)
	tag := key

	if s == sent && !required {
		tag += ",omitzero"
	}

	var lines []string

	if description != "" {
		lines = append(lines, commentLines(description)...)
	}

	if held.note != "" {
		lines = append(lines, commentLines(held.note)...)
	}

	if held.format != "" && description == "" {
		lines = append(lines, commentLines(held.format)...)
	}

	return append(lines, fmt.Sprintf("\t%s %s `json:%q`", exported(key), written, tag))
}

// render is how a field of that type is written, which is the whole of the difference between the
// two directions.
func render(held gotype, required bool, s side) string {
	switch held.kind {
	case kindSlice, kindMap, kindIface:
		return held.name
	case kindAny:
		return "any"
	}

	if s == answer {
		if required && !held.nullable {
			return held.name
		}

		return "*" + held.name
	}

	if !required {
		if held.kind == kindStruct {
			return "*" + held.name
		}

		return "Opt[" + held.name + "]"
	}

	if held.nullable {
		return "*" + held.name
	}

	return held.name
}

// childHint is what a nested shape is called: `UsersListResponse` plus `data` is `UsersListData`,
// minus the noise, and a row of that is `UsersListRow`.
func childHint(parent, key string) string {
	for _, suffix := range []string{"Response", "Body", "Params", "Form"} {
		if strings.HasSuffix(parent, suffix) {
			parent = strings.TrimSuffix(parent, suffix)

			break
		}
	}

	return parent + exported(key)
}

// body is every shape, in the order it was first needed.
func (m *models) body() string {
	blocks := make([]string, 0, len(m.order))

	for _, name := range m.order {
		blocks = append(blocks, strings.TrimSuffix(m.blocks[name], "\n"))
	}

	return strings.Join(blocks, "\n\n") + "\n"
}

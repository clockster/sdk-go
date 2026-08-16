package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// The part of OpenAPI this document uses. Anything it does not describe is not decoded, so a
// keyword arriving later is ignored rather than half-read.

type document struct {
	Paths      map[string]map[string]*operation `json:"paths"`
	Components struct {
		Schemas *props `json:"schemas"`
	} `json:"components"`
}

type operation struct {
	ID          string             `json:"operationId"`
	Summary     string             `json:"summary"`
	Description string             `json:"description"`
	Parameters  []*parameter       `json:"parameters"`
	RequestBody *content           `json:"requestBody"`
	Responses   map[string]content `json:"responses"`
}

type content struct {
	Content map[string]struct {
		Schema *schema `json:"schema"`
	} `json:"content"`
}

func (c content) schema(media string) *schema {
	held, ok := c.Content[media]
	if !ok {
		return nil
	}

	return held.Schema
}

type parameter struct {
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Description string  `json:"description"`
	Required    bool    `json:"required"`
	Schema      *schema `json:"schema"`
}

type discriminator struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping"`
}

type schema struct {
	Ref                  string          `json:"$ref"`
	Type                 json.RawMessage `json:"type"`
	Items                *schema         `json:"items"`
	Properties           *props          `json:"properties"`
	Required             []string        `json:"required"`
	Enum                 []any           `json:"enum"`
	AdditionalProperties json.RawMessage `json:"additionalProperties"`
	OneOf                []*schema       `json:"oneOf"`
	Discriminator        *discriminator  `json:"discriminator"`
	Format               string          `json:"format"`
	Description          string          `json:"description"`

	// The declared types with `null` still among them, and a canonical form of the whole schema.
	// Two identical shapes are named once, and identical here means these bytes.
	types     []string
	signature string
}

func (s *schema) UnmarshalJSON(data []byte) error {
	type plain schema

	var held plain

	if err := json.Unmarshal(data, &held); err != nil {
		return err
	}

	*s = schema(held)

	if len(s.Type) > 0 {
		var single string

		if err := json.Unmarshal(s.Type, &single); err == nil {
			s.types = []string{single}
		} else if err := json.Unmarshal(s.Type, &s.types); err != nil {
			return fmt.Errorf("type is neither a name nor a list: %s", s.Type)
		}
	}

	// Re-marshalled rather than kept as it arrived: the same shape written at two depths differs
	// in indentation alone, and that is not a difference.
	var canonical any

	if err := json.Unmarshal(data, &canonical); err != nil {
		return err
	}

	rewritten, err := json.Marshal(canonical)
	if err != nil {
		return err
	}

	s.signature = string(rewritten)

	return nil
}

// nullable reports whether the document says the value may be null.
func (s *schema) nullable() bool {
	for _, declared := range s.types {
		if declared == "null" {
			return true
		}
	}

	return false
}

// concrete is the declared types with `null` taken out.
func (s *schema) concrete() []string {
	var rest []string

	for _, declared := range s.types {
		if declared != "null" {
			rest = append(rest, declared)
		}
	}

	return rest
}

func (s *schema) requires(name string) bool {
	for _, key := range s.Required {
		if key == name {
			return true
		}
	}

	return false
}

// props is a JSON object whose keys are read in the order the document writes them, so a struct
// reads the way the documentation does rather than alphabetically.
type props struct {
	keys   []string
	values map[string]*schema
}

func (p *props) UnmarshalJSON(data []byte) error {
	p.values = map[string]*schema{}
	decoder := json.NewDecoder(bytes.NewReader(data))

	if _, err := decoder.Token(); err != nil {
		return err
	}

	for decoder.More() {
		key, err := decoder.Token()
		if err != nil {
			return err
		}

		name, ok := key.(string)
		if !ok {
			return fmt.Errorf("object key %v is not a name", key)
		}

		held := &schema{}

		if err := decoder.Decode(held); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}

		p.keys = append(p.keys, name)
		p.values[name] = held
	}

	_, err := decoder.Token()

	return err
}

func (p *props) len() int {
	if p == nil {
		return 0
	}

	return len(p.keys)
}

func (p *props) get(name string) *schema {
	if p == nil {
		return nil
	}

	return p.values[name]
}

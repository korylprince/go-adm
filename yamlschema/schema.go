// Package yamlschema implements a minimal [JSON Schema] parser that reads the schema as YAML.
//
// It parses a YAML-encoded JSON Schema into a [Schema] AST and can generate Go
// structs from it. The primary use case is parsing Apple's
// [device management root schema], but the parser is generic enough for other
// YAML-encoded JSON Schemas.
//
// Use [New] or [NewFromFile] to parse a schema, and [GenerateFromFile] or
// [GenerateFromGit] to generate Go code directly.
//
// [JSON Schema]: https://json-schema.org/
// [device management root schema]: https://github.com/apple/device-management/blob/release/docs/schema.yaml
package schema

import (
	"fmt"
	"io"
	"os"

	"github.com/korylprince/go-yaml"
)

type Type string

const (
	TypeArray   Type = "array"
	TypeBool    Type = "boolean"
	TypeInteger Type = "integer"
	TypeNumber  Type = "number"
	TypeObject  Type = "object"
	TypeString  Type = "string"
)

// TypeSlice can represent a single string Type or a slice of Types
type TypeSlice []Type

func (t *TypeSlice) UnmarshalYAML(f func(interface{}) error) error {
	if err := f((*[]Type)(t)); err == nil {
		return nil
	}

	*t = []Type{""}
	if err := f(&(*t)[0]); err != nil {
		return err
	}
	return nil
}

// OrderedMap is an ordered map[string]*Schema
type OrderedMap struct {
	Map   map[string]*Schema
	Order []string
}

// Iter iterates through keys and values in order
func (m *OrderedMap) Iter(f func(key string, schm *Schema)) {
	if m == nil {
		return
	}
	for _, k := range m.Order {
		f(k, m.Map[k])
	}
}

func (m *OrderedMap) UnmarshalYAML(f func(interface{}) error) error {
	mp := map[string]*Schema{}
	if err := f(&mp); err != nil {
		return fmt.Errorf("could not unmarshal map: %w", err)
	}
	m.Map = mp

	var ms yaml.MapSlice
	if err := f(&ms); err != nil {
		return fmt.Errorf("could not unmarshal mapslice: %w", err)
	}

	for _, item := range ms {
		m.Order = append(m.Order, item.Key.(string))
	}

	return nil
}

// Schema represents a JSON Schema (encoded in YAML)
type Schema struct {
	Title                string      `yaml:"title,omitempty"`
	Type                 TypeSlice   `yaml:"type"`
	Description          string      `yaml:"description,omitempty"`
	AdditionalProperties bool        `yaml:"additionalProperties,omitempty"`
	Required             []string    `yaml:"required,omitempty"`
	Properties           *OrderedMap `yaml:"properties,omitempty"`
	Items                *Schema     `yaml:"items,omitempty"`
	Enum                 []string    `yaml:"enum,omitempty"`
	Default              any         `yaml:"default,omitempty"`
}

// New parses the yaml schema from r
func New(r io.Reader) (*Schema, error) {
	cmd := new(Schema)
	if err := yaml.NewDecoder(r).Decode(cmd); err != nil {
		return nil, fmt.Errorf("could not unmarshal schema: %w", err)
	}

	return cmd, nil
}

// NewFromFile parses the yaml schema at path
func NewFromFile(path string) (*Schema, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open: %w", err)
	}
	defer f.Close()

	return New(f)
}

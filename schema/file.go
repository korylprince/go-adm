package schema

import (
	"strings"
)

type Type interface {
	PayloadKey() *PayloadKey
}

type Enum struct {
	*Schema
	Parents []*PayloadKey
	Key     *PayloadKey
}

func (e *Enum) PayloadKey() *PayloadKey {
	return e.Key
}

type Source string

const (
	SourcePayloadKeys  Source = "PayloadKeys"
	SourceResponseKeys Source = "ResponseKeys"
	SourceSubKeys      Source = "SubKeys"
)

type StructField struct {
	*Schema
	Parents        []*PayloadKey
	Key            *PayloadKey
	AdditionalTags map[string]string
}

type Struct struct {
	*Schema
	Source  Source
	Parents []*PayloadKey
	Key     *PayloadKey
	Fields  []*StructField
}

func (s *Struct) PayloadKey() *PayloadKey {
	return s.Key
}

// Map is a top level map, e.g. a schema that maps to a map, not a struct
type Map struct {
	*Schema
	Source  Source
	Parents []*PayloadKey
	Key     *PayloadKey
}

func (m *Map) PayloadKey() *PayloadKey {
	return m.Key
}

// File is a set of Enums and Structs generated from a set of one or more Schemas
type File struct {
	Schemas []*Schema
	Types   []Type
}

// NewFile returns a File by traversing one or more Schemas
func NewFile(schemas ...*Schema) *File {
	var types []Type

	// add enums, structs, and maps in order of appearance
	for _, s := range schemas {
		var structs []*Struct

		// build struct/map for top level schema payload keys
		if len(s.PayloadKeys) > 0 {
			var content []string
			if strings.TrimSpace(s.Description) != "" {
				content = append(content, strings.TrimSpace(s.Description))
			}
			if s.Payload != nil && strings.TrimSpace(s.Payload.Content) != "" {
				content = append(content, strings.TrimSpace(s.Payload.Content))
			}

			key := &PayloadKey{
				Title:   s.Title,
				Key:     s.Title,
				SubKeys: s.PayloadKeys,
				Type:    PayloadKeyTypeDictionary,
				Content: strings.Join(content, "\n"),
			}

			if key.IsStruct() {
				strct := &Struct{
					Schema: s,
					Source: SourcePayloadKeys,
					Key:    key,
				}

				structs = append(structs, strct)
				types = append(types, strct)
			}

			if key.IsMap() {
				m := &Map{
					Schema: s,
					Source: SourcePayloadKeys,
					Key:    key,
				}

				types = append(types, m)
			}
		}

		// build struct/map for top level schema response keys
		if len(s.ResponseKeys) > 0 {
			// TODO: set a description/content field for response?
			key := &PayloadKey{
				Title:   s.Title + "Response",
				Key:     s.Title + "Response",
				SubKeys: s.ResponseKeys,
				Type:    PayloadKeyTypeDictionary,
			}

			if key.IsStruct() {
				strct := &Struct{
					Schema: s,
					Source: SourceResponseKeys,
					Key:    key,
				}

				structs = append(structs, strct)
				types = append(types, strct)
			}

			if key.IsMap() {
				m := &Map{
					Schema: s,
					Source: SourceResponseKeys,
					Key:    key,
				}

				types = append(types, m)
			}
		}

		seen := make(map[*PayloadKey]struct{})
		s.Iter(func(parents []*PayloadKey, key *PayloadKey) {
			// don't add type more than once
			if _, ok := seen[key]; ok {
				return
			}
			seen[key] = struct{}{}

			if key.IsEnum() {
				// don't add enum more than once
				if len(parents) > 0 && parents[len(parents)-1].Type == PayloadKeyTypeArray {
					return
				}

				types = append(types, &Enum{
					Schema:  s,
					Parents: parents,
					Key:     key,
				})
			}

			// build struct
			if key.IsStruct() {
				strct := &Struct{
					Schema:  s,
					Source:  SourceSubKeys,
					Parents: parents,
					Key:     key,
				}
				structs = append(structs, strct)
				types = append(types, strct)
			}
		})

		// build struct fields
		for _, strct := range structs {
			var subParents []*PayloadKey
			copy(subParents, strct.Parents)
			subParents = append(subParents, strct.Key)
			for _, subkey := range strct.Key.SubKeys {
				strct.Fields = append(strct.Fields, &StructField{
					Schema:  s,
					Parents: subParents,
					Key:     subkey,
				})
			}
		}
	}

	return &File{
		Schemas: schemas,
		Types:   types,
	}
}

package schema

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/git"
	"github.com/korylprince/go-adm/replace"
	"github.com/korylprince/go-adm/text"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// EncodeOption allows configuring an Encoder
type EncodeOption func(*Encoder)

func WithReplacements(reps replace.Replacements) EncodeOption {
	return func(e *Encoder) {
		e.reps = reps
	}
}

// Encoder encodes Schemas to Go types
type Encoder struct {
	f    *jen.File
	reps replace.Replacements
}

func NewEncoder(f *jen.File, opts ...EncodeOption) *Encoder {
	e := &Encoder{f: f}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// findCommonName finds a common suffix between strings, e.g. [JohnSmith, JaneSmith, JoeAmith] -> mith
func findCommonSuffix(names []string) string {
	l := 1
	suffix := names[0][len(names[0])-1-l:]
	for l < len(names[0]) {
		for _, n := range names[1:] {
			if l > len(n) || strings.HasSuffix(n, suffix) {
				return suffix
			}
		}
		l += 1
		suffix = names[0][len(names[0])-1-l:]
	}
	return suffix
}

// findStructs recurses through the root schema, finding all type: object schemas that should be marshaled as Go structs
// An ordered name -> schema mapping is returned where name is the best guess at what the struct should be called
func findStructs(root *Schema) *OrderedMap {
	seen := make(map[*Schema]struct{})
	found := map[*Schema]map[string]struct{}{root: {"schema": {}}}
	var order []*Schema

	// dfs through schema to preserve order
	var dfs func(s *Schema)
	dfs = func(s *Schema) {
		if _, ok := seen[s]; ok {
			return
		}
		seen[s] = struct{}{}
		if !slices.Contains(order, s) {
			order = append(order, s)
		}

		s.Properties.Iter(func(k string, v *Schema) {
			if v.Title == "" {
				v.Title = k
			}
			if len(v.Type) == 1 && v.Type[0] == TypeObject {
				key := v.Title
				if key == "" {
					key = k
				}
				if _, ok := found[v]; !ok {
					found[v] = map[string]struct{}{key: {}}
				} else {
					found[v][key] = struct{}{}
				}
			}
			if v.Items != nil && len(v.Items.Type) == 1 && v.Items.Type[0] == TypeObject {
				key := v.Items.Title
				if key == "" {
					key = v.Title
				}
				if _, ok := found[v.Items]; !ok {
					found[v.Items] = map[string]struct{}{key: {}}
				} else {
					found[v.Items][key] = struct{}{}
				}
			}
			dfs(v)
		})
		if s.Items != nil {
			dfs(s.Items)
		}
	}
	dfs(root)

	// find best name and convert to OrderedMap
	structs := &OrderedMap{Map: make(map[string]*Schema, len(found))}
	for _, s := range order {
		nameSet, ok := found[s]
		if !ok {
			continue
		}
		names := maps.Keys(nameSet)
		if len(names) == 1 {
			structs.Map[names[0]] = s
			structs.Order = append(structs.Order, names[0])
			continue
		}
		name := findCommonSuffix(names)
		if name == "" && len(names) > 0 {
			name = names[0]
		}
		structs.Map[name] = s
		structs.Order = append(structs.Order, name)
	}

	return structs
}

func (e *Encoder) normalizeName(name, typ string) string {
	return e.reps.Replace(text.NormalizeName(name), typ)
}

func schemaType(s *Schema) jen.Code {
	if len(s.Type) == 0 {
		return jen.Any()
	} else if len(s.Type) > 1 {
		// multi-type
		typs := slices.Clone(s.Type)
		slices.Sort(typs)
		typstrs := make([]string, len(typs))
		for idx, t := range typs {
			typstrs[idx] = cases.Title(language.AmericanEnglish).String(string(t))
		}
		return jen.Qual("github.com/korylprince/go-adm/yamlschema", strings.Join(typstrs, ""))
	}
	switch s.Type[0] {
	case TypeBool:
		return jen.Bool()
	case TypeInteger:
		return jen.Int64()
	case TypeNumber:
		return jen.Float64()
	case TypeObject:
		return jen.Any()
	case TypeString:
		return jen.String()
	case TypeArray:
		if s.Items == nil {
			return jen.Index().Any()
		}
		return jen.Index().Add(schemaType(s.Items))
	}
	return jen.Any()
}

func (e *Encoder) Encode(s *Schema) {
	// find structs
	structs := findStructs(s)
	structMap := map[*Schema]string{}
	structs.Iter(func(name string, schm *Schema) {
		structMap[schm] = name
	})

	// marshal enums
	enumMap := map[*Schema]string{}
	structs.Iter(func(parentName string, strct *Schema) {
		strct.Properties.Iter(func(name string, enum *Schema) {
			if len(enum.Enum) == 0 {
				return
			}
			enumName := e.normalizeName(parentName, "struct") + e.normalizeName(name, "field")
			enumMap[enum] = enumName

			// render type definition
			if enum.Description != "" {
				e.f.Comment(text.DocComment(enum.Description))
			}
			e.f.Type().Id(enumName).String()

			// render enum values
			consts := make([]jen.Code, len(enum.Enum))
			for idx, v := range enum.Enum {
				consts[idx] = jen.Id(enumName + e.normalizeName(v, "const")).Id(enumName).Op("=").Lit(v)
			}
			e.f.Const().Defs(consts...)
		})
	})

	// marshal structs
	structs.Iter(func(structName string, strct *Schema) {
		if strct.Description != "" {
			e.f.Comment(text.DocComment(strct.Description))
		}
		var fields []jen.Code

		strct.Properties.Iter(func(propName string, prop *Schema) {
			if prop.Description != "" {
				fields = append(fields, jen.Comment(text.DocComment(prop.Description)))
			}

			// render field tag
			tag := map[string]string{"yaml": propName}
			if !slices.Contains(strct.Required, propName) {
				tag["yaml"] = propName + ",omitempty"
			}

			// select field type
			typ := schemaType(prop)
			if prop.Items != nil {
				if t, ok := structMap[prop.Items]; ok {
					typ = jen.Index().Op("*").Id(e.normalizeName(t, "struct"))
				}
			}
			if t, ok := structMap[prop]; ok {
				typ = jen.Op("*").Id(e.normalizeName(t, "struct"))
			}
			if t, ok := enumMap[prop]; ok {
				typ = jen.Id(t)
			}

			fields = append(fields, jen.Id(e.normalizeName(propName, "field")).Add(typ).Tag(tag))
		})

		e.f.Type().Id(e.normalizeName(structName, "struct")).Struct(fields...)
	})
}

// GenerateFromFile generates Go types from the schema at path using the given pkgName and optional replacements
func GenerateFromFile(path, pkgName string, reps replace.Replacements) ([]byte, error) {
	schema, err := NewFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not parse file: %w", err)
	}

	f := jen.NewFile(pkgName)
	f.HeaderComment("DO NOT EDIT")
	f.HeaderComment(fmt.Sprintf("generated from %s", path))

	NewEncoder(f, WithReplacements(reps)).Encode(schema)

	buf := new(bytes.Buffer)
	if err := f.Render(buf); err != nil {
		return nil, fmt.Errorf("could not render code: %w", err)
	}

	return buf.Bytes(), nil
}

// GenerateFromGit generates Go types from the schema at the git repo/commit/path using the given pkgName and optional replacements
func GenerateFromGit(repoURL, commit, path, pkgName string, reps replace.Replacements) ([]byte, error) {
	repo, err := git.New(repoURL, commit)
	if err != nil {
		return nil, fmt.Errorf("could not check out repository: %w", err)
	}

	file, err := repo.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	schema, err := New(file)
	if err != nil {
		return nil, fmt.Errorf("could not parse file: %w", err)
	}

	hash, err := repo.Hash()
	if err != nil {
		return nil, fmt.Errorf("could not get hash: %w", err)
	}

	f := jen.NewFile(pkgName)
	f.HeaderComment("DO NOT EDIT")
	f.HeaderComment(fmt.Sprintf("generated from %s:%s/%s", repoURL, hash, path))

	f.Const().Id("DeviceManagementGenerateHash").Op("=").Lit(hash)

	NewEncoder(f, WithReplacements(reps)).Encode(schema)

	buf := new(bytes.Buffer)
	if err := f.Render(buf); err != nil {
		return nil, fmt.Errorf("could not render code: %w", err)
	}

	return buf.Bytes(), nil
}

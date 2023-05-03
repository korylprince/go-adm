package schema

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"regexp"
	"strings"

	"github.com/korylprince/go-adm/git"
	"github.com/korylprince/go-adm/replace"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// EncodeOption allows configuring an Encoder
type EncodeOption func(*Encoder) error

func WithReplacements(reps replace.Replacements) EncodeOption {
	return func(e *Encoder) error {
		e.reps = reps
		return nil
	}
}

// Encoder encodes Schemas to Go types
type Encoder struct {
	w    io.Writer
	opts []EncodeOption
	reps replace.Replacements
}

func NewEncoder(w io.Writer, opts ...EncodeOption) *Encoder {
	return &Encoder{w: w, opts: opts}
}

func (e *Encoder) normalizeName(s string) string {
	stripped := regexp.MustCompile("[^a-zA-Z0-9\\-]").ReplaceAllLiteralString(s, "")
	split := strings.Split(stripped, "-")
	name := ""
	for _, n := range split {
		name += cases.Title(language.AmericanEnglish).String(n)
	}
	return e.reps.Replace(name)
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

		s.Properties.Iter(func(k string, v *Schema) error {
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
			return nil
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

func multiType(s *Schema) string {
	typs := slices.Clone(s.Type)
	slices.Sort(typs)
	typstrs := make([]string, len(typs))
	for idx, t := range typs {
		typstrs[idx] = cases.Title(language.AmericanEnglish).String(string(t))
	}
	// TODO: resolve this import
	return "yamlschema." + strings.Join(typstrs, "")
}

func schemaType(s *Schema) string {
	if len(s.Type) == 0 {
		return "any"
	} else if len(s.Type) > 1 {
		return multiType(s)
	}
	switch s.Type[0] {
	case TypeBool:
		return "bool"
	case TypeInteger:
		return "int64"
	case TypeNumber:
		return "float64"
	case TypeObject:
		return "any"
	case TypeString:
		return "string"
	case TypeArray:
		if s.Items == nil {
			return "[]any"
		}
		return "[]" + schemaType(s.Items)
	}
	return "any"
}

func docComment(doc string) string {
	comment := ""
	for _, line := range strings.Split(strings.TrimSpace(doc), "\n") {
		if c := strings.TrimSpace(line); c != "" {
			comment += fmt.Sprintf("// %s\n", c)
		}
	}
	return comment
}

func (e *Encoder) Encode(s *Schema) error {
	for _, opt := range e.opts {
		if err := opt(e); err != nil {
			return fmt.Errorf("option failed: %w", err)
		}
	}

	// find structs
	structs := findStructs(s)
	structMap := map[*Schema]string{}
	structs.Iter(func(name string, schm *Schema) error {
		structMap[schm] = name
		return nil
	})

	buf := new(bytes.Buffer)

	// marshal multitypes
	multitypes := map[string]struct{}{}
	if err := structs.Iter(func(_ string, strct *Schema) error {
		return strct.Properties.Iter(func(name string, prop *Schema) error {
			if len(prop.Type) < 2 {
				return nil
			}
			mtype := multiType(prop)
			if _, ok := multitypes[mtype]; ok {
				return nil
			}
			multitypes[mtype] = struct{}{}

			return nil
		})
	}); err != nil {
		return err
	}

	// marshal enums
	enumMap := map[*Schema]string{}
	if err := structs.Iter(func(parentName string, strct *Schema) error {
		return strct.Properties.Iter(func(name string, enum *Schema) error {
			if len(enum.Enum) == 0 {
				return nil
			}
			enumName := e.normalizeName(parentName) + e.normalizeName(name)
			enumMap[enum] = enumName

			if enum.Description != "" {
				buf.WriteString(docComment(enum.Description))
			}
			buf.WriteString(fmt.Sprintf("type %s string\nconst (\n", enumName))

			for _, v := range enum.Enum {
				buf.WriteString(fmt.Sprintf("\t%s%s %s = \"%s\"\n", enumName, e.normalizeName(v), enumName, v))
			}
			buf.WriteString(")\n\n")

			if _, err := buf.WriteTo(e.w); err != nil {
				return fmt.Errorf("could not write enum %s: %w", enumName, err)
			}
			return nil
		})
	}); err != nil {
		return err
	}

	// marshal structs
	if err := structs.Iter(func(structName string, strct *Schema) error {
		if strct.Description != "" {
			buf.WriteString(docComment(strct.Description))
		}
		buf.WriteString(fmt.Sprintf("type %s struct {\n", e.normalizeName(structName)))

		strct.Properties.Iter(func(propName string, prop *Schema) error {
			if prop.Description != "" {
				buf.WriteString(docComment(prop.Description))
			}

			typ := schemaType(prop)
			if prop.Items != nil {
				if t, ok := structMap[prop.Items]; ok {
					typ = "[]*" + e.normalizeName(t)
				}
			}
			if t, ok := structMap[prop]; ok {
				typ = "*" + e.normalizeName(t)
			}
			if t, ok := enumMap[prop]; ok {
				typ = t
			}

			buf.WriteString(fmt.Sprintf("\t%s %s `yaml:\"%s", e.normalizeName(propName), typ, propName))
			if !slices.Contains(strct.Required, propName) {
				buf.WriteString(",omitempty")
			}
			buf.WriteString("\"`\n")

			return nil
		})
		buf.WriteString("}\n\n")

		if _, err := buf.WriteTo(e.w); err != nil {
			return fmt.Errorf("could not write struct %s: %w", structName, err)
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GenerateFromFile generates Go types from the schema at path using the given pkgName and optional replacements
func GenerateFromFile(path, pkgName string, reps replace.Replacements) ([]byte, error) {
	schema, err := NewFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not parse file: %w", err)
	}

	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("// package %s is generated from %s\n// DO NOT EDIT\npackage %s\n\n",
		pkgName, path, pkgName))

	if err := NewEncoder(buf, WithReplacements(reps)).Encode(schema); err != nil {
		return nil, fmt.Errorf("could not write schema: %w", err)
	}

	fmtd, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("could not format generated code: %w", err)
	}

	return bytes.TrimSpace(fmtd), nil
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

	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("// DO NOT EDIT\n// generated from %s:%s/%s\n\npackage %s\n\n",
		repoURL, hash, path, pkgName))

	// TODO: resolve this correctly
	buf.WriteString("import yamlschema \"github.com/korylprince/go-adm/yamlschema\"\n")

	buf.WriteString(fmt.Sprintf("const DeviceManagementGenerateHash = \"%s\"\n\n", hash))

	if err := NewEncoder(buf, WithReplacements(reps)).Encode(schema); err != nil {
		return nil, fmt.Errorf("could not write schema: %w", err)
	}

	fmtd, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("could not format generated code: %w", err)
	}

	return bytes.TrimSpace(fmtd), nil
}

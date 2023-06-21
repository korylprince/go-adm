package declarations

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-git/go-billy/v5/util"
	"github.com/korylprince/go-adm/git"
	"github.com/korylprince/go-adm/replace"
	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/text"
)

type EncodeOption func(*Encoder)

func WithReplacements(reps replace.Replacements) EncodeOption {
	return func(e *Encoder) {
		e.reps = reps
	}
}

type Encoder struct {
	f    *jen.File
	reps replace.Replacements
	gn   *schema.GlobalNamer

	enums   []jen.Code
	structs []jen.Code

	structMap map[string]string
}

func NewEncoder(f *jen.File, opts ...EncodeOption) *Encoder {
	e := &Encoder{f: f, gn: schema.NewGlobalNamer(), structMap: make(map[string]string)}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Encoder) normalizeName(name string) string {
	return e.reps.Replace(name)
}

func (e *Encoder) schemaType(key *schema.PayloadKey) jen.Code {
	var typ jen.Code = jen.Any()
	switch key.Type {
	case schema.PayloadKeyTypeString:
		typ = jen.String()
	case schema.PayloadKeyTypeInteger:
		typ = jen.Int64()
	case schema.PayloadKeyTypeReal:
		typ = jen.Float64()
	case schema.PayloadKeyTypeBoolean:
		typ = jen.Bool()
	case schema.PayloadKeyTypeDate:
		typ = jen.Qual("time", "Time")
	case schema.PayloadKeyTypeData:
		typ = jen.Index().Byte()
	case schema.PayloadKeyTypeArray:
		if key.SubKeys[0].Type == schema.PayloadKeyTypeDictionary {
			typ = jen.Index().Add(e.schemaType(key.SubKeys[0]))
		} else {
			// we want *[]string, not *[]*string, so we monkey patch the sub key's presence attribute
			orig := key.SubKeys[0].Presence
			key.SubKeys[0].Presence = schema.PayloadKeyPresenceRequired
			typ = jen.Index().Add(e.schemaType(key.SubKeys[0]))
			key.SubKeys[0].Presence = orig
		}
	case schema.PayloadKeyTypeDictionary:
		return jen.Op("*").Id(e.normalizeName(e.gn.KeyName(key)))
	case schema.PayloadKeyTypeAny:
		typ = jen.Any()
	}

	// make optional fields pointers
	if key.Presence != schema.PayloadKeyPresenceRequired {
		typ = jen.Op("*").Add(typ)
	}

	return typ
}

func (e *Encoder) renderEnum(parentName string, key *schema.PayloadKey) {
	// render comment
	if doc := strings.TrimSpace(key.Content); doc != "" {
		e.enums = append(e.enums, jen.Comment(text.DocComment(key.Content)))
	}

	enumName := e.normalizeName(parentName + text.NormalizeName(key.Key))

	// array with enum subkey
	if key.Type == schema.PayloadKeyTypeArray && len(key.SubKeys) == 1 && len(key.SubKeys[0].Rangelist) > 0 {
		key = key.SubKeys[0]
		// render subkey comment
		if doc := strings.TrimSpace(key.Content); doc != "" {
			e.enums = append(e.enums, jen.Comment(text.DocComment(key.Content)))
		}
	}

	// render type
	switch key.Type {
	case schema.PayloadKeyTypeString:
		e.enums = append(e.enums, jen.Type().Id(enumName).String())
	case schema.PayloadKeyTypeInteger:
		e.enums = append(e.enums, jen.Type().Id(enumName).Int64())
	default:
		panic(fmt.Errorf("unknown key type: %s", key.Type))
	}

	// render consts
	e.enums = append(e.enums, jen.Const().DefsFunc(func(g *jen.Group) {
		for _, v := range key.Rangelist {
			switch key.Type {
			case schema.PayloadKeyTypeString:
				g.Id(e.normalizeName(enumName + text.NormalizeName(v.String()))).Id(enumName).Op("=").Lit(v.String())
			case schema.PayloadKeyTypeInteger:
				g.Id(e.normalizeName(enumName + strconv.Itoa(int(v.Int64())))).Id(enumName).Op("=").Lit(int(v.Int64()))
			}
		}
	}))
}

func (e *Encoder) renderSchema(s *schema.Schema) {
	seen := make(map[*schema.PayloadKey]struct{})
	var dfs func(parentName string, keys []*schema.PayloadKey) (fields []jen.Code)
	dfs = func(parentName string, keys []*schema.PayloadKey) (fields []jen.Code) {
		for _, key := range keys {
			// render field comment
			if doc := strings.TrimSpace(key.Content); doc != "" {
				fields = append(fields, jen.Comment(text.DocComment(doc)))
			}

			// render json tag
			tags := map[string]string{"json": key.Key}
			if key.Presence == schema.PayloadKeyPresenceOptional {
				tags["json"] = key.Key + ",omitempty"
			}

			// render default tag
			if key.Default.Value() != nil {
				switch {
				case key.Default.IsInt64():
					tags["default"] = strconv.FormatInt(key.Default.Int64(), 10)
				case key.Default.IsFloat64():
					tags["default"] = strconv.FormatFloat(key.Default.Float64(), 'f', -1, 64)
				case key.Default.IsString():
					tags["default"] = key.Default.String()
				}
			}

			// render field
			fields = append(fields, jen.Id(e.normalizeName(text.NormalizeName(key.Key))).Add(e.schemaType(key)).Tag(tags))

			// prevent infinite recursion
			if _, ok := seen[key]; ok {
				return
			}
			seen[key] = struct{}{}

			if key.Type == schema.PayloadKeyTypeDictionary || (key.Type == schema.PayloadKeyTypeArray && len(key.SubKeys) == 1 && key.SubKeys[0].Type == schema.PayloadKeyTypeDictionary) {
				keyName := e.normalizeName(e.gn.KeyName(key))
				if keyName == "" {
					keyName = e.normalizeName(text.NormalizeName(key.Key))
				}

				// recurse into fields
				subfields := dfs(keyName, key.SubKeys)

				// render struct for dictionary keys
				if key.Type == schema.PayloadKeyTypeDictionary {
					if doc := strings.TrimSpace(key.Content); doc != "" {
						e.structs = append(e.structs, jen.Comment(text.DocComment(doc)))
					}
					e.structs = append(e.structs, jen.Type().Id(keyName).Struct(subfields...))
				}
			}
		}

		return fields
	}

	schemaName := e.normalizeName(e.gn.SchemaName(s))
	// recurse through payload keys, generating types along the way
	fields := dfs(schemaName, s.PayloadKeys)

	// generate code for schema struct itself
	if doc := strings.TrimSpace(s.Description); doc != "" {
		e.structs = append(e.structs, jen.Comment(text.DocComment(doc)))
	}
	if doc := strings.TrimSpace(s.Payload.Content); doc != "" {
		e.structs = append(e.structs, jen.Comment(text.DocComment(doc)))
	}
	e.structs = append(e.structs, jen.Type().Id(schemaName).Struct(fields...))

	// generate code for DeclarativeType method
	if s.Payload.DeclarationType != "" {
		e.structs = append(e.structs, jen.Func().Parens(jen.Id("p").Op("*").Id(schemaName)).Id("DeclarationType").Parens(nil).String().Block(
			jen.Return().Lit(s.Payload.DeclarationType),
		))

		// add struct to id -> type map
		e.structMap[schemaName] = s.Payload.DeclarationType
	}
}

func (e *Encoder) Encode(schemas ...*schema.Schema) []byte {
	e.gn.Register(schemas...)

	// render enums
	for _, s := range schemas {
		seen := make(map[*schema.PayloadKey]struct{})
		s.Iter(func(parent, key *schema.PayloadKey) {
			// prevent seeing array enum twice
			if _, ok := seen[key]; ok {
				return
			}

			parentName := e.gn.SchemaName(s)
			if parent != nil {
				parentName = e.gn.KeyName(parent)
			}

			// normal enum
			if len(key.Rangelist) > 0 {
				e.renderEnum(parentName, key)
			}

			// array with enum subkey
			if key.Type == schema.PayloadKeyTypeArray && len(key.SubKeys) == 1 && len(key.SubKeys[0].Rangelist) > 0 {
				e.renderEnum(parentName, key)
				seen[key.SubKeys[0]] = struct{}{}
			}
		})
	}

	for _, s := range schemas {
		e.renderSchema(s)
	}

	e.f.Var().Id("DeclarationMap").Op("=").Map(jen.String()).Any().Values(jen.DictFunc(func(d jen.Dict) {
		for typ, id := range e.structMap {
			d[jen.Lit(id)] = jen.Id(typ).Values()
		}
	}))

	for _, stmt := range append(e.enums, e.structs...) {
		e.f.Add(stmt)
	}

	return nil
}

// GenerateFromGit generates Go types from the declarative schema at the git repo/commit/path using the optional replacements
// and outputs it at the given directory
func GenerateFromGit(repoURL, commit, path string, reps replace.Replacements, output string) error {
	repo, err := git.New(repoURL, commit)
	if err != nil {
		return fmt.Errorf("could not check out repository: %w", err)
	}

	hash, err := repo.Hash()
	if err != nil {
		return fmt.Errorf("could not get hash: %w", err)
	}

	dirs, err := repo.ReadDir(path)
	if err != nil {
		return fmt.Errorf("could not read %s: %w", path, err)
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		var schemas []*schema.Schema
		if err = util.Walk(repo, repo.Join(path, dir.Name()), func(filepath string, info fs.FileInfo, _ error) error {
			if !strings.HasSuffix(info.Name(), ".yaml") {
				return nil
			}

			buf, err := util.ReadFile(repo, filepath)
			if err != nil {
				return fmt.Errorf("could not read %s: %w", filepath, err)
			}

			s, err := schema.New(buf)
			if err != nil {
				return fmt.Errorf("could not parse %s: %w", filepath, err)
			}
			schemas = append(schemas, s)

			return nil
		}); err != nil {
			return err
		}

		f := jen.NewFile(dir.Name())
		f.HeaderComment("DO NOT EDIT")
		f.HeaderComment(fmt.Sprintf("generated from %s:%s/%s", repoURL, hash, filepath.Join(path, dir.Name())))

		f.Const().Id("DeviceManagementGenerateHash").Op("=").Lit(hash)

		NewEncoder(f, WithReplacements(reps)).Encode(schemas...)

		buf := new(bytes.Buffer)
		if err = f.Render(buf); err != nil {
			return fmt.Errorf("could not render code: %w", err)
		}

		if err = os.MkdirAll(filepath.Join(output, dir.Name()), 0755); err != nil {
			return fmt.Errorf("could not create output directory %s: %w", filepath.Join(".", dir.Name()), err)
		}

		if err = os.WriteFile(filepath.Join(output, dir.Name(), dir.Name()+".go"), buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("could not create write %s: %w", filepath.Join(output, dir.Name(), dir.Name()+".go"), err)
		}
	}

	return nil
}

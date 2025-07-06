package declarative

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-git/go-billy/v5/util"
	"github.com/korylprince/go-adm/git"
	"github.com/korylprince/go-adm/replace"
	"github.com/korylprince/go-adm/schema"
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
	enc  *schema.Encoder
}

func NewEncoder(f *jen.File, opts ...EncodeOption) *Encoder {
	e := &Encoder{f: f}
	for _, opt := range opts {
		opt(e)
	}
	e.enc = schema.NewEncoder(f, schema.WithReplacements(e.reps))
	return e
}

func declarationType(typ schema.Type) string {
	switch t := typ.(type) {
	case *schema.Struct:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.DeclarationType == "" {
			return ""
		}
		return t.Schema.Payload.DeclarationType
	case *schema.Map:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.DeclarationType == "" {
			return ""
		}
		return t.Schema.Payload.DeclarationType
	}
	return ""
}

func credentialType(typ schema.Type) string {
	switch t := typ.(type) {
	case *schema.Struct:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.CredentialType == "" {
			return ""
		}
		return t.Schema.Payload.CredentialType
	case *schema.Map:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.CredentialType == "" {
			return ""
		}
		return t.Schema.Payload.CredentialType
	}
	return ""
}

func statusItemType(typ schema.Type) string {
	switch t := typ.(type) {
	case *schema.Struct:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.StatusItemType == "" {
			return ""
		}
		return t.Schema.Payload.StatusItemType
	case *schema.Map:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.StatusItemType == "" {
			return ""
		}
		return t.Schema.Payload.StatusItemType
	}
	return ""
}

func (e *Encoder) Encode(file *schema.File) {
	e.enc.RegisterFile(file)
	// render DeclarationType -> struct map
	var decls []schema.Type
	for _, typ := range file.Types {
		if dt := declarationType(typ); dt != "" {
			decls = append(decls, typ)
		}
	}
	if len(decls) > 0 {
		e.f.Var().Id("DeclarationMap").Op("=").Map(jen.String()).Any().Values(jen.DictFunc(func(d jen.Dict) {
			for _, typ := range decls {
				dt := declarationType(typ)
				d[jen.Lit(dt)] = jen.Id(e.enc.Name(typ.PayloadKey(), replace.Struct)).Values()
			}
		}))
	}

	// render CredentialType -> struct map
	var creds []schema.Type
	for _, typ := range file.Types {
		if ct := credentialType(typ); ct != "" {
			creds = append(creds, typ)
		}
	}
	if len(creds) > 0 {
		e.f.Var().Id("CredentialMap").Op("=").Map(jen.String()).Any().Values(jen.DictFunc(func(d jen.Dict) {
			for _, typ := range creds {
				ct := credentialType(typ)
				d[jen.Lit(ct)] = jen.Id(e.enc.Name(typ.PayloadKey(), replace.Struct)).Values()
			}
		}))
	}

	// render StatusItemType -> struct map
	var statuses []schema.Type
	for _, typ := range file.Types {
		if st := statusItemType(typ); st != "" {
			statuses = append(statuses, typ)
		}
	}
	if len(statuses) > 0 {
		e.f.Var().Id("StatusItemType").Op("=").Map(jen.String()).Any().Values(jen.DictFunc(func(d jen.Dict) {
			for _, typ := range statuses {
				st := statusItemType(typ)
				d[jen.Lit(st)] = jen.Id(e.enc.Name(typ.PayloadKey(), replace.Struct)).Values()
			}
		}))
	}

	for _, typ := range file.Types {
		switch t := typ.(type) {
		case *schema.Enum:
			e.enc.EncodeEnum(t)
		case *schema.Struct:
			e.enc.EncodeStruct(t)

			if dt := declarationType(typ); dt != "" {
				structName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Op("*").Id(structName)
				e.f.Func().Parens(rcvr).Id("DeclarationType").Parens(nil).String().Block(
					jen.Return().Lit(dt),
				)
			}
			if ct := credentialType(typ); ct != "" {
				structName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Op("*").Id(structName)
				e.f.Func().Parens(rcvr).Id("CredentialType").Parens(nil).String().Block(
					jen.Return().Lit(ct),
				)
			}
			if st := statusItemType(typ); st != "" {
				structName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Op("*").Id(structName)
				e.f.Func().Parens(rcvr).Id("StatusItemType").Parens(nil).String().Block(
					jen.Return().Lit(st),
				)
			}
		case *schema.Map:
			e.enc.EncodeMap(t)

			if dt := declarationType(typ); dt != "" {
				mapName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Id(mapName)
				e.f.Func().Parens(rcvr).Id("DeclarationType").Parens(nil).String().Block(
					jen.Return().Lit(dt),
				)
			}
			if ct := credentialType(typ); ct != "" {
				mapName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Id(mapName)
				e.f.Func().Parens(rcvr).Id("CredentialType").Parens(nil).String().Block(
					jen.Return().Lit(ct),
				)
			}
			if st := statusItemType(typ); st != "" {
				structName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Id(structName)
				e.f.Func().Parens(rcvr).Id("StatusItemType").Parens(nil).String().Block(
					jen.Return().Lit(st),
				)
			}
		}
	}
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

	type file struct {
		path string
		s    *schema.Schema
	}

	var files []*file
	if err = util.Walk(repo, path, func(filePath string, info fs.FileInfo, _ error) error {
		if !strings.HasSuffix(info.Name(), ".yaml") {
			return nil
		}

		buf, err := util.ReadFile(repo, filePath)
		if err != nil {
			return fmt.Errorf("could not read %s: %w", filePath, err)
		}

		s, err := schema.New(buf)
		if err != nil {
			return fmt.Errorf("could not parse %s: %w", filePath, err)
		}

		// get relative directory
		rel, err := filepath.Rel(path, filePath)
		if err != nil {
			return fmt.Errorf("could not get relative path for %s: %w", filePath, err)
		}
		dir, _ := filepath.Split(rel)

		files = append(files, &file{
			path: dir,
			s:    s,
		})

		return nil
	}); err != nil {
		return err
	}

	schemaMap := make(map[string][]*schema.Schema)
	for _, f := range files {
		schemaMap[f.path] = append(schemaMap[f.path], f.s)
	}

	for dir, schemas := range schemaMap {
		fn := filepath.Base(dir)

		f := jen.NewFile(fn)
		f.HeaderComment("DO NOT EDIT")
		f.HeaderComment(fmt.Sprintf("generated from %s:%s/%s", repoURL, hash, filepath.Join(path, dir)))

		f.Const().Id("DeviceManagementGenerateHash").Op("=").Lit(hash)

		file := schema.NewFile(schemas)
		NewEncoder(f, WithReplacements(reps)).Encode(file)

		buf := new(bytes.Buffer)
		if err = f.Render(buf); err != nil {
			return fmt.Errorf("could not render code: %w", err)
		}

		if err = os.MkdirAll(filepath.Join(output, dir), 0755); err != nil {
			return fmt.Errorf("could not create output directory %s: %w", filepath.Join(".", dir), err)
		}

		if err = os.WriteFile(filepath.Join(output, dir, fn+".go"), buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("could not create write %s: %w", filepath.Join(output, dir, fn+".go"), err)
		}
	}

	return nil
}

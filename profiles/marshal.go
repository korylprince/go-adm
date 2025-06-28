package profiles

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

func payloadType(typ schema.Type) string {
	switch t := typ.(type) {
	case *schema.Struct:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.PayloadType == "" {
			return ""
		}
		return t.Schema.Payload.PayloadType
	case *schema.Map:
		if t.Source != schema.SourcePayloadKeys || t.Schema.Payload == nil || t.Schema.Payload.PayloadType == "" {
			return ""
		}
		return t.Schema.Payload.PayloadType
	}
	return ""
}

func (e *Encoder) Encode(file *schema.File) {
	e.enc.RegisterFile(file)
	// FIXME: there are multiple schemas with the same PayloadType (e.g. MCX)
	// render PayloadType -> struct map
	// e.f.Var().Id("PayloadMap").Op("=").Map(jen.String()).Any().Values(jen.DictFunc(func(d jen.Dict) {
	// 	for _, typ := range file.Types {
	// 		if pt := payloadType(typ); pt != "" {
	// 			d[jen.Lit(pt)] = jen.Id(e.enc.Name(typ.PayloadKey(), replace.Struct)).Values()
	// 		}
	// 	}
	// }))

	for _, typ := range file.Types {
		switch t := typ.(type) {
		case *schema.Enum:
			e.enc.EncodeEnum(t)
		case *schema.Struct:
			e.enc.EncodeStruct(t)

			if pt := payloadType(typ); pt != "" {
				structName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Op("*").Id(structName)
				e.f.Func().Parens(rcvr).Id("PayloadType").Parens(nil).String().Block(
					jen.Return().Lit(pt),
				)
			}
		case *schema.Map:
			e.enc.EncodeMap(t)

			if pt := payloadType(typ); pt != "" {
				structName := e.enc.Name(t.Key, replace.Struct)
				rcvr := jen.Id("p").Op("*").Id(structName)
				e.f.Func().Parens(rcvr).Id("PayloadType").Parens(nil).String().Block(
					jen.Return().Lit(pt),
				)
			}
		}
	}
}

// GenerateFromGit generates Go types from the profile schema at the git repo/commit/path using the optional replacements
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

	fi, err := repo.Stat(path)
	if err != nil {
		return fmt.Errorf("could not read %s: %w", path, err)
	}
	var schemas []*schema.Schema
	if err = util.Walk(repo, path, func(filepath string, info fs.FileInfo, _ error) error {
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

	f := jen.NewFile(fi.Name())
	f.HeaderComment("DO NOT EDIT")
	f.HeaderComment(fmt.Sprintf("generated from %s:%s/%s", repoURL, hash, filepath.Join(path, fi.Name())))

	f.Const().Id("DeviceManagementGenerateHash").Op("=").Lit(hash)

	file := schema.NewFile(schemas,
		schema.WithIncludeEmptyPayloadKeys(true),
	)
	NewEncoder(f, WithReplacements(reps)).Encode(file)

	buf := new(bytes.Buffer)
	if err = f.Render(buf); err != nil {
		return fmt.Errorf("could not render code: %w", err)
	}

	if err = os.MkdirAll(filepath.Join(output, fi.Name()), 0755); err != nil {
		return fmt.Errorf("could not create output directory %s: %w", filepath.Join(".", fi.Name()), err)
	}

	if err = os.WriteFile(filepath.Join(output, fi.Name(), fi.Name()+".go"), buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("could not create write %s: %w", filepath.Join(output, fi.Name(), fi.Name()+".go"), err)
	}

	return nil
}

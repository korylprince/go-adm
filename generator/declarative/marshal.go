package gen

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/utils/replace"
)

type EncodeOption func(*Encoder)

func WithReplacements(reps replace.Replacements) EncodeOption {
	return func(e *Encoder) {
		e.reps = reps
	}
}

func WithTags(tags []string) EncodeOption {
	return func(e *Encoder) {
		e.tags = tags
	}
}

func WithRequiredDefault() EncodeOption {
	return func(e *Encoder) {
		e.reqDefTags = true
	}
}

// WithNamerOptions passes naming options through to the schema encoder.
// See schema.WithNameOverrides and schema.WithFullyQualifiedNames.
func WithNamerOptions(opts ...schema.NamerOption) EncodeOption {
	return func(e *Encoder) {
		e.namerOpts = append(e.namerOpts, opts...)
	}
}

// WithStatusItemTypes renders the given dotted status item identifiers as a
// StatusItemTypes slice.
//
// The nested status tree turns each dotted identifier into a path of Go fields,
// which leaves no generated value holding the identifier itself -- and it is the
// identifier, not the path, that appears on the wire.
func WithStatusItemTypes(itemTypes []string) EncodeOption {
	return func(e *Encoder) {
		e.statusItemTypes = itemTypes
	}
}

type Encoder struct {
	f               *jen.File
	reps            replace.Replacements
	tags            []string
	reqDefTags      bool
	namerOpts       []schema.NamerOption
	statusItemTypes []string
	enc             *schema.Encoder
}

func NewEncoder(f *jen.File, opts ...EncodeOption) *Encoder {
	e := &Encoder{f: f, tags: []string{"json"}}
	for _, opt := range opts {
		opt(e)
	}
	sOpts := []schema.EncodeOption{
		schema.WithReplacements(e.reps),
		schema.WithTags(e.tags),
	}
	if e.reqDefTags {
		sOpts = append(sOpts, schema.WithRequiredDefault())
	}
	if len(e.namerOpts) > 0 {
		sOpts = append(sOpts, schema.WithNamerOptions(e.namerOpts...))
	}
	e.enc = schema.NewEncoder(f, sOpts...)
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

// statusItemTypesDoc explains the generated StatusItemTypes slice in the output
// itself, since the identifiers no longer appear anywhere else as values.
const statusItemTypesDoc = `StatusItemTypes lists every DDM status item these types can represent, by
Apple's dotted status item identifier.

Generated from the statusitemtype of each schema under declarative/status/.
Each identifier corresponds to a field reachable from StatusItems -- the field's
doc comment names the item it carries.

These are the identifiers the protocol uses as string values: the items named by
a status subscription, a status report's Errors[].StatusItem, and the client's
management.client-capabilities.supported-payloads.status-items.`

// docComment prefixes each line of doc with "//".
//
// text.DocComment drops blank lines, which is right for Apple's content but
// wrong here: godoc needs a bare "//" between paragraphs to keep them apart.
func docComment(doc string) string {
	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(doc), "\n") {
		lines = append(lines, strings.TrimRight("// "+strings.TrimSpace(line), " "))
	}
	return strings.Join(lines, "\n")
}

func (e *Encoder) Encode(file *schema.File) {
	e.enc.RegisterFile(file)

	// render the list of known status item identifiers
	if len(e.statusItemTypes) > 0 {
		e.f.Comment(docComment(statusItemTypesDoc))
		e.f.Var().Id("StatusItemTypes").Op("=").Index().String().ValuesFunc(func(g *jen.Group) {
			for _, itemType := range e.statusItemTypes {
				g.Line().Lit(itemType)
			}
			g.Line()
		})
	}
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

package schema

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/replace"
	"github.com/korylprince/go-adm/text"
)

// Encoder encodes a File to Go types
type Encoder struct {
	namer      *GlobalNamer
	f          *jen.File
	reps       replace.Replacements
	tags       []string
	reqDefTags bool
}

type EncodeOption func(*Encoder)

func WithReplacements(reps replace.Replacements) EncodeOption {
	return func(e *Encoder) {
		e.reps = reps
	}
}

// WithTags overrides the default struct tags the encoder generates.
// The default is "json", and "plist."
// These will have ",omitempty" appended for presence-optional keys.
func WithTags(tags []string) EncodeOption {
	return func(e *Encoder) {
		e.tags = tags
	}
}

// WithRequiredDefault turns on generation of the "required" and "default" struct tags.
func WithRequiredDefault() EncodeOption {
	return func(e *Encoder) {
		e.reqDefTags = true
	}
}

func NewEncoder(f *jen.File, opts ...EncodeOption) *Encoder {
	e := &Encoder{f: f, tags: []string{"json", "plist"}}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Encoder) normalizeName(name string, typ replace.ReplacementType) string {
	return e.reps.Replace(text.NormalizeName(name), typ)
}

// Name returns a normalized and globally unique name for the given key and type
func (e *Encoder) Name(key *PayloadKey, typ replace.ReplacementType) string {
	return e.normalizeName(e.namer.KeyName(key), typ)
}

// MapValueType returns the Go type for the value side of a map[string]T
// generated from a dictionary with an ANY subkey.
func (e *Encoder) MapValueType(key *PayloadKey) jen.Code {
	switch key.Type {
	case PayloadKeyTypeAny:
		return jen.Any()
	case PayloadKeyTypeString:
		return jen.String()
	case PayloadKeyTypeArray:
		// Determine the array element type from the subkeys.
		// Temporarily force Required presence to avoid pointer wrapping.
		orig := key.Presence
		key.Presence = PayloadKeyPresenceRequired
		code := e.fieldType(key)
		key.Presence = orig
		return code
	case PayloadKeyTypeDictionary:
		if key.IsStruct() {
			return jen.Id(e.Name(key, replace.Field))
		}
		// nested map (ANY subkey within an ANY subkey)
		return jen.Map(jen.String()).Add(e.MapValueType(key.SubKeys[0]))
	default:
		panic(fmt.Errorf("ANY <dictionary>: unknown value type: %s", key.Type))
	}
}

// get the type of struct field based on the PayloadKey
func (e *Encoder) fieldType(key *PayloadKey) jen.Code {
	var typ jen.Code = jen.Any()
	switch key.Type {
	case PayloadKeyTypeString:
		typ = jen.String()
	case PayloadKeyTypeInteger:
		typ = jen.Int64()
	case PayloadKeyTypeReal:
		typ = jen.Float64()
	case PayloadKeyTypeBoolean:
		typ = jen.Bool()
	case PayloadKeyTypeDate:
		typ = jen.Qual("time", "Time")
	case PayloadKeyTypeData:
		typ = jen.Index().Byte()
	case PayloadKeyTypeArray:
		if len(key.SubKeys) > 0 && key.SubKeys[0].Type == PayloadKeyTypeDictionary {
			typ = jen.Index().Add(e.fieldType(key.SubKeys[0]))
		} else if key.IsEnum() {
			typ = jen.Index().Id(e.Name(key, replace.Const))
		} else if len(key.SubKeys) > 0 {
			// we want *[]string, not *[]*string, so we monkey patch the sub key's presence attribute
			orig := key.SubKeys[0].Presence
			key.SubKeys[0].Presence = PayloadKeyPresenceRequired
			typ = jen.Index().Add(e.fieldType(key.SubKeys[0]))
			key.SubKeys[0].Presence = orig
		} else {
			return jen.Any()
		}
	case PayloadKeyTypeDictionary:
		if key.IsStruct() {
			typ = jen.Id(e.Name(key, replace.Field))
		} else {
			typ = jen.Map(jen.String()).Add(e.MapValueType(key.SubKeys[0]))
		}
	case PayloadKeyTypeAny:
		typ = jen.Any()
	}

	// replace enums
	if key.IsEnum() && key.Type != PayloadKeyTypeArray {
		typ = jen.Id(e.Name(key, replace.Const))
	}

	// make optional fields pointers
	if key.Presence != PayloadKeyPresenceRequired {
		typ = jen.Op("*").Add(typ)
	}

	return typ
}

// get the tags for a struct fields
func (e *Encoder) fieldTags(fld *StructField) map[string]string {
	// render json and plist tag
	tags := make(map[string]string)
	for _, tagName := range e.tags {
		tags[tagName] = fld.Key.Key
	}
	if fld.Key.Presence == PayloadKeyPresenceOptional {
		for k, v := range tags {
			tags[k] = v + ",omitempty"
		}
	}

	// render required tag
	if e.reqDefTags && fld.Key.Presence == PayloadKeyPresenceRequired {
		tags["required"] = "true"
	}

	// render default tag
	if e.reqDefTags && fld.Key.Default.Value() != nil {
		switch {
		case fld.Key.Default.IsInt64():
			tags["default"] = strconv.FormatInt(fld.Key.Default.Int64(), 10)
		case fld.Key.Default.IsFloat64():
			tags["default"] = strconv.FormatFloat(fld.Key.Default.Float64(), 'f', -1, 64)
		case fld.Key.Default.IsString():
			tags["default"] = fld.Key.Default.String()
		}
	}

	// overlay user-supplied tags
	for k, v := range fld.AdditionalTags {
		tags[k] = v
	}

	return tags
}

func (e *Encoder) RegisterFile(f *File) {
	e.namer = NewGlobalNamer(f)
}

func (e *Encoder) Encode(f *File) {
	e.RegisterFile(f)

	for _, typ := range f.Types {
		switch t := typ.(type) {
		case *Enum:
			e.EncodeEnum(t)
		case *Struct:
			e.EncodeStruct(t)
		case *Map:
			e.EncodeMap(t)
		}
	}
}

func (e *Encoder) EncodeEnum(enum *Enum) {
	// render comment
	if doc := strings.TrimSpace(enum.Key.Content); doc != "" {
		e.f.Comment(text.DocComment(enum.Key.Content))
	}

	// get enum name
	enumName := e.Name(enum.Key, replace.Const)

	// if enum is defined by array with enum subkey
	enumKey := enum.Key
	if enum.Key.Type == PayloadKeyTypeArray && len(enum.Key.SubKeys) == 1 && len(enum.Key.SubKeys[0].Rangelist) > 0 {
		enumKey = enum.Key.SubKeys[0]
		// render SubKey comment
		if doc := strings.TrimSpace(enumKey.Content); doc != "" {
			e.f.Comment(text.DocComment(enumKey.Content))
		}
	}

	// render type
	switch enumKey.Type {
	case PayloadKeyTypeString:
		e.f.Type().Id(enumName).String()
	case PayloadKeyTypeInteger:
		e.f.Type().Id(enumName).Int64()
	case PayloadKeyTypeReal:
		e.f.Type().Id(enumName).Float64()
	default:
		panic(fmt.Errorf("unknown enumKey type: %s", enumKey.Type))
	}

	// render consts
	e.f.Const().DefsFunc(func(g *jen.Group) {
		for _, v := range enumKey.Rangelist {
			switch enumKey.Type {
			case PayloadKeyTypeString:
				constName := e.normalizeName(enumName+text.NormalizeName(v.String()), replace.Const)
				g.Id(constName).
					Id(enumName).
					Op("=").
					Lit(v.String())
			case PayloadKeyTypeInteger:
				// replace negative sign with "Neg"
				val := strconv.Itoa(int(v.Int64()))
				val = strings.Replace(val, "-", "Neg", 1)
				constName := e.normalizeName(enumName+val, replace.Const)
				g.Id(constName).
					Id(enumName).
					Op("=").
					Lit(int(v.Int64()))
			case PayloadKeyTypeReal:
				// replace negative sign with "Neg"
				val := strconv.FormatFloat(float64(v.Float64()), 'f', -1, 64)
				val = strings.Replace(val, "-", "Neg", 1)
				constName := e.normalizeName(enumName+val, replace.Const)
				g.Id(constName).
					Id(enumName).
					Op("=").
					Lit(float64(v.Float64()))
			}
		}
	})
}

type structEncodeOptions struct {
	embeds             []jen.Code
	fieldTypeOverrides map[string]jen.Code
}

type StructEncodeOption func(*structEncodeOptions)

// WithStructEmbeds prepends embedded types at the top of the generated struct.
func WithStructEmbeds(embeds ...jen.Code) StructEncodeOption {
	return func(opts *structEncodeOptions) {
		opts.embeds = append(opts.embeds, embeds...)
	}
}

// WithStructFieldTypeOverride replaces the generated field type for the given schema key.
func WithStructFieldTypeOverride(key string, typ jen.Code) StructEncodeOption {
	return func(opts *structEncodeOptions) {
		if opts.fieldTypeOverrides == nil {
			opts.fieldTypeOverrides = make(map[string]jen.Code)
		}
		opts.fieldTypeOverrides[key] = typ
	}
}

// EncodeStruct encodes a struct with optional embedded types and field type overrides.
func (e *Encoder) EncodeStruct(strct *Struct, opts ...StructEncodeOption) {
	var encOpts structEncodeOptions
	for _, opt := range opts {
		opt(&encOpts)
	}

	// render comment
	if doc := strings.TrimSpace(strct.Key.Content); doc != "" {
		e.f.Comment(text.DocComment(strct.Key.Content))
	}

	structName := e.Name(strct.Key, replace.Struct)

	e.f.Type().Id(structName).StructFunc(func(g *jen.Group) {
		for _, embed := range encOpts.embeds {
			g.Add(embed)
		}
		for _, fld := range strct.Fields {
			// render field comment
			if doc := strings.TrimSpace(fld.Key.Content); doc != "" {
				g.Comment(text.DocComment(doc))
			}

			fieldType := e.fieldType(fld.Key)
			if override, ok := encOpts.fieldTypeOverrides[fld.Key.Key]; ok {
				fieldType = override
			}

			g.Id(e.normalizeName(text.NormalizeName(fld.Key.Key), replace.Field)).
				Add(fieldType).
				Tag(e.fieldTags(fld))
		}
	})
}

func (e *Encoder) EncodeMap(m *Map) {
	// render comment
	if doc := strings.TrimSpace(m.Key.Content); doc != "" {
		e.f.Comment(text.DocComment(m.Key.Content))
	}

	mapName := e.Name(m.Key, replace.Struct)

	// render map type
	e.f.Type().
		Id(mapName).
		Map(jen.String()).
		Add(e.MapValueType(m.Key.SubKeys[0]))
}

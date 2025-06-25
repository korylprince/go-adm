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
	namer *GlobalNamer
	f     *jen.File
	reps  replace.Replacements
}

type EncodeOption func(*Encoder)

func WithReplacements(reps replace.Replacements) EncodeOption {
	return func(e *Encoder) {
		e.reps = reps
	}
}

func NewEncoder(f *jen.File, opts ...EncodeOption) *Encoder {
	e := &Encoder{f: f}
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

// determine the value type of a <dictionary> with <any> fields
func anyDictionaryType(key *PayloadKey) jen.Code {
	switch key.Type {
	case PayloadKeyTypeAny:
		return jen.Any()
	case PayloadKeyTypeString:
		return jen.String()
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
		if key.SubKeys[0].Type == PayloadKeyTypeDictionary {
			typ = jen.Index().Add(e.fieldType(key.SubKeys[0]))
		} else if key.IsEnum() {
			typ = jen.Index().Id(e.Name(key, replace.Const))
		} else {
			// we want *[]string, not *[]*string, so we monkey patch the sub key's presence attribute
			orig := key.SubKeys[0].Presence
			key.SubKeys[0].Presence = PayloadKeyPresenceRequired
			typ = jen.Index().Add(e.fieldType(key.SubKeys[0]))
			key.SubKeys[0].Presence = orig
		}
	case PayloadKeyTypeDictionary:
		if key.IsStruct() {
			typ = jen.Id(e.Name(key, replace.Field))
		} else {
			typ = jen.Map(jen.String()).Add(anyDictionaryType(key.SubKeys[0]))
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
	tags := map[string]string{
		"json":  fld.Key.Key,
		"plist": fld.Key.Key,
	}
	if fld.Key.Presence == PayloadKeyPresenceOptional {
		tags["json"] = fld.Key.Key + ",omitempty"
		tags["plist"] = fld.Key.Key + ",omitempty"
	}

	// render required tag
	if fld.Key.Presence == PayloadKeyPresenceRequired {
		tags["required"] = "true"
	}

	// render default tag
	if fld.Key.Default.Value() != nil {
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
				fmt.Println(constName)
				g.Id(constName).
					Id(enumName).
					Op("=").
					Lit(int(v.Int64()))
			}
		}
	})
}

func (e *Encoder) EncodeStruct(strct *Struct) {
	// render comment
	if doc := strings.TrimSpace(strct.Key.Content); doc != "" {
		e.f.Comment(text.DocComment(strct.Key.Content))
	}

	structName := e.Name(strct.Key, replace.Struct)

	e.f.Type().Id(structName).StructFunc(func(g *jen.Group) {
		for _, fld := range strct.Fields {
			// render field comment
			if doc := strings.TrimSpace(fld.Key.Content); doc != "" {
				g.Comment(text.DocComment(doc))
			}

			g.Id(e.normalizeName(text.NormalizeName(fld.Key.Key), replace.Field)).
				Add(e.fieldType(fld.Key)).
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
		Add(anyDictionaryType(m.Key.SubKeys[0]))
}

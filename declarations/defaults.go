package declarations

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var ErrNotStruct = errors.New("not a struct or pointer to a struct")

// StructDefaults sets the default values of a struct (and any nested structs) using the default tag.
// StructDefaults only sets defaults on pointer fields that are unset
func StructDefaults(v any) error {
	val := reflect.ValueOf(v)
	if val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Type().Kind() != reflect.Struct {
		return fmt.Errorf("could not parse %s: %w", val.Type().Name(), ErrNotStruct)
	}

	strct := val.Type()
	for i := 0; i < strct.NumField(); i++ {
		fld := strct.Field(i)

		// nested struct
		if fld.Type.Kind() == reflect.Struct || (fld.Type.Kind() == reflect.Pointer && fld.Type.Elem().Kind() == reflect.Struct) {
			// if struct isn't optional initialize it
			if !strings.Contains(fld.Tag.Get("json"), "omitempty") {
				if fld.Type.Kind() == reflect.Struct {
					val.Field(i).Set(reflect.New(fld.Type))
				} else {
					val.Field(i).Set(reflect.New(fld.Type.Elem()))
				}
			}
			if !val.Field(i).IsNil() {
				if err := StructDefaults(val.Field(i).Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// slice of structs
		if fld.Type.Kind() == reflect.Slice || fld.Type.Kind() == reflect.Pointer && fld.Type.Elem().Kind() == reflect.Slice {
			// check if value is nil
			val := val.Field(i)
			if val.IsNil() {
				continue
			}

			// deference pointer
			typ := fld.Type
			if typ.Kind() == reflect.Pointer {
				typ = typ.Elem()
				val = val.Elem()
			}
			if typ.Elem().Kind() == reflect.Struct || (typ.Elem().Kind() == reflect.Pointer && typ.Elem().Elem().Kind() == reflect.Struct) {
				for j := 0; j < val.Len(); j++ {
					if err := StructDefaults(val.Index(j).Interface()); err != nil {
						return err
					}
				}
				continue
			}
		}

		// get default flag
		tag := fld.Tag.Get("default")
		if tag == "" {
			continue
		}

		// only set defaults on nilable fields
		if fld.Type.Kind() != reflect.Pointer {
			continue
		}

		// only set defaults on fields that are unset (nil)
		if !val.Field(i).IsNil() {
			continue
		}

		typ := fld.Type.Elem()

		// parse default value
		var def reflect.Value
		switch typ.Kind() {
		case reflect.Bool:
			switch tag {
			case "true":
				def = reflect.ValueOf(true)
			case "false":
				def = reflect.ValueOf(false)
			default:
				return fmt.Errorf("could not parse %s.%s default %s as bool", strct.Name(), fld.Name, tag)
			}
		case reflect.Int, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(tag, 10, 64)
			if err != nil {
				return fmt.Errorf("could not parse %s.%s default %s as int64: %w", strct.Name(), fld.Name, tag, err)
			}
			def = reflect.ValueOf(i)
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(tag, 64)
			if err != nil {
				return fmt.Errorf("could not parse %s.%s default %s as float: %w", strct.Name(), fld.Name, tag, err)
			}
			def = reflect.ValueOf(f)
		case reflect.String:
			def = reflect.ValueOf(tag)
		default:
			return fmt.Errorf("could not parse %s.%s default %s: unsupported field type %s", strct.Name(), fld.Name, tag, typ.String())
		}

		// set default value
		fldval := reflect.New(typ)
		fldval.Elem().Set(def.Convert(typ))
		val.Field(i).Set(fldval)
	}

	return nil
}

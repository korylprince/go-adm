package jsonutil

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var ErrNotStruct = errors.New("not a struct or pointer to a struct")

// SetDefaults sets the default values of a struct (and any nested structs) using the default tag.
func SetDefaults(v any) error {
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
		if fld.Type.Kind() == reflect.Struct {
			if err := SetDefaults(val.Field(i).Interface()); err != nil {
				return err
			}
			continue
		}

		// nested *struct
		if fld.Type.Kind() == reflect.Pointer && fld.Type.Elem().Kind() == reflect.Struct {
			if !val.Field(i).IsNil() {
				if err := SetDefaults(val.Field(i).Interface()); err != nil {
					return err
				}
			}
			continue
		}

		// make slices (i.e. end up with [] instead of null in the json)
		if fld.Type.Kind() == reflect.Slice {
			if val.Field(i).IsNil() {
				val.Field(i).Set(reflect.MakeSlice(fld.Type, 0, 0))
			}
		}

		// make maps (i.e. end up with {} instead of null in the json)
		if fld.Type.Kind() == reflect.Map {
			if val.Field(i).IsNil() {
				val.Field(i).Set(reflect.MakeMap(fld.Type))
			}
		}

		// slice of structs
		if fld.Type.Kind() == reflect.Slice || fld.Type.Kind() == reflect.Pointer && fld.Type.Elem().Kind() == reflect.Slice {
			// check if value is nil
			fldval := val.Field(i)
			if fldval.IsNil() {
				continue
			}

			// deference pointer
			typ := fld.Type
			if typ.Kind() == reflect.Pointer {
				typ = typ.Elem()
				fldval = fldval.Elem()
			}

			// check that slice is made of struct or *struct
			if typ.Elem().Kind() == reflect.Struct || (typ.Elem().Kind() == reflect.Pointer && typ.Elem().Elem().Kind() == reflect.Struct) {
				for j := 0; j < fldval.Len(); j++ {
					if err := SetDefaults(fldval.Index(j).Interface()); err != nil {
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

		// only set defaults on fields that are unset (nil) or the zero value
		if !val.Field(i).IsNil() && !val.Field(i).Elem().IsZero() {
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

package jsonutil

import (
	"fmt"
	"reflect"
	"strings"
)

// FullFields uses reflection to fill out types and remove omitempty tags so that all fields will be marshaled to json
func FullFields(v any) any {
	val := reflect.ValueOf(v)
	typ := val.Type()

	switch kind := typ.Kind(); kind {
	case reflect.Pointer:
		// init nil values
		if val.IsNil() {
			val = reflect.New(typ.Elem())
		}
		// dereference pointer
		return FullFields(val.Elem().Interface())
	case reflect.Slice:
		// init slice
		if val.IsNil() {
			val = reflect.MakeSlice(typ, 0, 0)
		}
		// fill slice with at least one item
		if val.Len() == 0 {
			item := reflect.New(typ.Elem()).Elem()
			val = reflect.Append(val, item)
		}
		// rewrite struct items
		if typ.Elem().Kind() == reflect.Struct {
			vals := make([]any, val.Len())
			for idx := 0; idx < val.Len(); idx++ {
				vals[idx] = FullFields(val.Index(idx).Interface())
			}
			val = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(vals[0]).Elem()), len(vals), len(vals))
			for idx := 0; idx < len(vals); idx++ {
				val.Index(idx).Set(reflect.ValueOf(vals[idx]).Elem())
			}
		} else if typ.Elem().Kind() == reflect.Pointer && typ.Elem().Elem().Kind() == reflect.Struct {
			vals := make([]any, val.Len())
			for idx := 0; idx < val.Len(); idx++ {
				vals[idx] = FullFields(val.Index(idx).Interface())
			}
			val = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(vals[0])), len(vals), len(vals))
			for idx := 0; idx < len(vals); idx++ {
				val.Index(idx).Set(reflect.ValueOf(vals[idx]))
			}
		}
		return val.Interface()
	case reflect.Map:
		if val.IsNil() {
			val = reflect.MakeMap(typ)
		}
		return val.Interface()
	case reflect.Struct:
		// pass
	default:
		return val.Interface()
	}

	repl := make(map[int]reflect.Value)

	// rewrite field types
	flds := make([]reflect.StructField, typ.NumField())
	for idx := 0; idx < typ.NumField(); idx++ {
		fld := typ.Field(idx)
		// remove omitempty
		if tag := fld.Tag.Get("json"); len(tag) > 0 {
			split := strings.Split(tag, ",")
			if len(split) > 0 {
				tag = split[0]
			}
			if tag != "-" {
				fld.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, tag))
			} else {
				fld.Tag = reflect.StructTag("")
			}
		}
		// recurse into fields
		if fld.Type.Kind() == reflect.Pointer {
			newv := reflect.ValueOf(FullFields(val.Field(idx).Interface()))
			fld.Type = newv.Type()
			repl[idx] = newv
		} else {
			newv := reflect.ValueOf(FullFields(val.Field(idx).Interface()))
			if newv.Type().Kind() == reflect.Pointer {
				newv = newv.Elem()
			}
			fld.Type = newv.Type()
			repl[idx] = newv
		}
		flds[idx] = fld
	}

	// copy values
	ret := reflect.New(reflect.StructOf(flds))
	for idx := 0; idx < typ.NumField(); idx++ {
		if r, ok := repl[idx]; ok {
			ret.Elem().Field(idx).Set(r)
		} else {
			ret.Elem().Field(idx).Set(val.Field(idx))
		}
	}

	return ret.Interface()
}

package tagutil

import (
	"reflect"

	"github.com/fatih/structtag"
)

// FullFields uses reflection to fill out types and remove omitempty tags so that all fields will be marshaled to json/plist
func FullFields(v any) any {
	return fullFields(v, make(map[reflect.Type]struct{}))
}

func fullFields(v any, seen map[reflect.Type]struct{}) any {
	val := reflect.ValueOf(v)
	typ := val.Type()

	switch kind := typ.Kind(); kind {
	case reflect.Pointer:
		// init nil values
		if val.IsNil() {
			val = reflect.New(typ.Elem())
		}
		// dereference pointer
		return fullFields(val.Elem().Interface(), seen)
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
				vals[idx] = fullFields(val.Index(idx).Interface(), seen)
			}
			val = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(vals[0]).Elem()), len(vals), len(vals))
			for idx := 0; idx < len(vals); idx++ {
				val.Index(idx).Set(reflect.ValueOf(vals[idx]).Elem())
			}
		} else if typ.Elem().Kind() == reflect.Pointer && typ.Elem().Elem().Kind() == reflect.Struct {
			vals := make([]any, val.Len())
			for idx := 0; idx < val.Len(); idx++ {
				vals[idx] = fullFields(val.Index(idx).Interface(), seen)
			}
			val = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(vals[0])), len(vals), len(vals))
			for idx := 0; idx < len(vals); idx++ {
				val.Index(idx).Set(reflect.ValueOf(vals[idx]))
			}
		} else if typ.Elem().Kind() == reflect.Interface {
			// recurse into interface elements (e.g. []any containing structs)
			for idx := 0; idx < val.Len(); idx++ {
				elem := val.Index(idx).Elem()
				expanded := fullFields(elem.Interface(), seen)
				ev := reflect.ValueOf(expanded)
				// dereference pointer results so plist library sees struct values
				if ev.Kind() == reflect.Pointer {
					ev = ev.Elem()
				}
				val.Index(idx).Set(ev)
			}
		}
		return val.Interface()
	case reflect.Map:
		if val.IsNil() {
			val = reflect.MakeMap(typ)
		}
		return val.Interface()
	case reflect.Struct:
		// break recursive type cycles
		if _, ok := seen[typ]; ok {
			return reflect.New(typ).Interface()
		}
		seen[typ] = struct{}{}
		defer delete(seen, typ)

		// skip types with no exported fields (e.g. time.Time)
		hasExported := false
		for i := 0; i < typ.NumField(); i++ {
			if typ.Field(i).IsExported() {
				hasExported = true
				break
			}
		}
		if !hasExported {
			return val.Interface()
		}
	default:
		return val.Interface()
	}

	repl := make(map[int]reflect.Value)

	// rewrite field types
	flds := make([]reflect.StructField, typ.NumField())
	for idx := 0; idx < typ.NumField(); idx++ {
		fld := typ.Field(idx)
		// skip unexported fields (e.g. time.Time internals)
		if !fld.IsExported() {
			flds[idx] = fld
			continue
		}
		// remove omitempty
		tags, err := structtag.Parse(string(fld.Tag))
		if err != nil {
			flds[idx] = fld
			continue
		}
		tags.DeleteOptions("json", "omitempty")
		tags.DeleteOptions("plist", "omitempty")
		fld.Tag = reflect.StructTag(tags.String())
		// recurse into fields
		if fld.Type.Kind() == reflect.Pointer {
			newv := reflect.ValueOf(fullFields(val.Field(idx).Interface(), seen))
			fld.Type = newv.Type()
			if fld.Type.Kind() != reflect.Pointer {
				ptrNewV := reflect.New(newv.Type())
				ptrNewV.Elem().Set(newv)
				newv = ptrNewV
				fld.Type = newv.Type()
			}
			repl[idx] = newv
		} else {
			newv := reflect.ValueOf(fullFields(val.Field(idx).Interface(), seen))
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
		} else if typ.Field(idx).IsExported() {
			ret.Elem().Field(idx).Set(val.Field(idx))
		}
	}

	return ret.Interface()
}

//go:generate /bin/bash -c "profilegen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../../GENERATE_SHA)\" -repl ./repls.yaml -reqdef"

package mdm

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/korylprince/go-adm/generated/mdm/profiles"
	"github.com/korylprince/go-adm/utils/tagutil"
	"github.com/micromdm/plist"
)

var ErrUnknownPayloadType = errors.New("unknown payload type")

// Profile is a configuration profile containing one or more payloads.
type Profile struct {
	profiles.TopLevel
}

func (p *Profile) Type() string {
	return string(p.TopLevel.PayloadType)
}

// initCommonPayloadKeys ensures the embedded *CommonPayloadKeys is initialized
// and its PayloadType and PayloadVersion fields are set.
func initCommonPayloadKeys(payload profiles.ProfilePayload) {
	v := reflect.ValueOf(payload)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	cpk := v.FieldByName("CommonPayloadKeys")
	if !cpk.IsValid() || cpk.Kind() != reflect.Pointer {
		return
	}
	if cpk.IsNil() {
		cpk.Set(reflect.New(cpk.Type().Elem()))
	}
	pt := cpk.Elem().FieldByName("PayloadType")
	if pt.IsValid() && pt.String() == "" {
		pt.SetString(payload.PayloadType())
	}
	pv := cpk.Elem().FieldByName("PayloadVersion")
	if pv.IsValid() && pv.Int() == 0 {
		pv.SetInt(int64(profiles.CommonPayloadKeysPayloadVersion1))
	}
}

// flattenWrapperStruct builds a flat map[string]any from a wrapper struct's
// CommonPayloadKeys fields and Properties map entries.
func flattenWrapperStruct(v reflect.Value) map[string]any {
	result := make(map[string]any)
	if cpk := v.FieldByName("CommonPayloadKeys"); cpk.IsValid() && cpk.Kind() == reflect.Pointer && !cpk.IsNil() {
		cpkVal := cpk.Elem()
		cpkType := cpkVal.Type()
		for j := 0; j < cpkType.NumField(); j++ {
			field := cpkType.Field(j)
			tag := field.Tag.Get("plist")
			if tag == "" || tag == "-" {
				continue
			}
			key, _, _ := strings.Cut(tag, ",")
			fv := cpkVal.Field(j)
			if strings.Contains(tag, "omitempty") && fv.IsZero() {
				continue
			}
			if fv.Kind() == reflect.Pointer {
				if fv.IsNil() {
					continue
				}
				result[key] = fv.Elem().Interface()
			} else {
				result[key] = fv.Interface()
			}
		}
	}
	if props := v.FieldByName("Properties"); props.IsValid() && props.Kind() == reflect.Map && !props.IsNil() {
		for _, mk := range props.MapKeys() {
			result[mk.String()] = props.MapIndex(mk).Interface()
		}
	}
	return result
}

// PlistValue returns a composite type suitable for marshalling to plist.
func (p *Profile) PlistValue() (any, error) {
	if p == nil {
		return nil, errors.New("profile is nil")
	}

	if err := tagutil.SetDefaults(&p.TopLevel); err != nil {
		return nil, fmt.Errorf("could not set top-level defaults: %w", err)
	}

	payloads := make([]any, len(p.PayloadContent))
	for i, payload := range p.PayloadContent {
		if payload == nil {
			return nil, fmt.Errorf("payload %d is nil", i)
		}

		typ := reflect.TypeOf(payload)
		if typ.Kind() == reflect.Pointer {
			typ = typ.Elem()
		}
		if typ.Kind() != reflect.Struct {
			return nil, fmt.Errorf("payload %d (%s): map-based payloads are not supported in PlistValue; marshal manually", i, payload.PayloadType())
		}

		initCommonPayloadKeys(payload)

		if err := tagutil.SetDefaults(payload); err != nil {
			return nil, fmt.Errorf("could not set defaults for payload %d: %w", i, err)
		}

		// Check if this is a wrapper struct (has a Properties map field).
		// If so, flatten CommonPayloadKeys + Properties into a single map.
		v := reflect.ValueOf(payload).Elem()
		if props := v.FieldByName("Properties"); props.IsValid() && props.Kind() == reflect.Map {
			payloads[i] = flattenWrapperStruct(v)
		} else {
			// dereference pointer so plist library sees a struct value
			payloads[i] = v.Interface()
		}
	}

	// Build the top-level struct dynamically, replacing PayloadContent's
	// interface slice with []any of concrete struct values for plist marshalling
	topLevelType := reflect.TypeOf(p.TopLevel)
	topLevelValue := reflect.ValueOf(p.TopLevel)

	flds := make([]reflect.StructField, topLevelType.NumField())
	for i := 0; i < topLevelType.NumField(); i++ {
		fld := topLevelType.Field(i)
		if fld.Name == "PayloadContent" {
			fld.Type = reflect.TypeOf([]any{})
		}
		flds[i] = fld
	}

	result := reflect.New(reflect.StructOf(flds))
	for i := 0; i < topLevelType.NumField(); i++ {
		if topLevelType.Field(i).Name == "PayloadContent" {
			result.Elem().Field(i).Set(reflect.ValueOf(payloads))
		} else {
			result.Elem().Field(i).Set(topLevelValue.Field(i))
		}
	}

	return result.Elem().Interface(), nil
}

func (p *Profile) MarshalIndentPlist(indent string) ([]byte, error) {
	v, err := p.PlistValue()
	if err != nil {
		return nil, err
	}
	return plist.MarshalIndent(v, indent)
}

func (p *Profile) MarshalPlist() ([]byte, error) {
	return p.MarshalIndentPlist("")
}

// NewFromType creates a new Profile with a single payload of the given PayloadType.
func NewProfileFromType(typ, profileUUID, payloadUUID string) (*Profile, error) {
	payload, ok := profiles.PayloadMap[typ]
	if !ok {
		return nil, ErrUnknownPayloadType
	}

	pay := reflect.New(reflect.TypeOf(payload)).Interface().(profiles.ProfilePayload)

	// initialize CommonPayloadKeys on the payload with the provided UUID
	v := reflect.ValueOf(pay).Elem()
	if cpk := v.FieldByName("CommonPayloadKeys"); cpk.IsValid() && cpk.Kind() == reflect.Pointer {
		cpk.Set(reflect.New(cpk.Type().Elem()))
		cpk.Elem().FieldByName("PayloadUUID").SetString(payloadUUID)
	}

	// initialize Properties map for wrapper struct payloads
	if props := v.FieldByName("Properties"); props.IsValid() && props.Kind() == reflect.Map && props.IsNil() {
		props.Set(reflect.MakeMap(props.Type()))
	}

	return &Profile{
		TopLevel: profiles.TopLevel{
			PayloadType:    profiles.PayloadTypeConfiguration,
			PayloadVersion: profiles.TopLevelPayloadVersion1,
			PayloadUUID:    profileUUID,
			PayloadContent: []profiles.ProfilePayload{pay},
		},
	}, nil
}

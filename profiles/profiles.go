//go:generate /bin/bash -c "profilegen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../GENERATE_SHA)\" -repl ./repls.yaml -reqdef"

package profiles

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/korylprince/go-adm/profiles/profiles"
	"github.com/korylprince/go-adm/tagutil"
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

		// dereference pointer so plist library sees a struct value
		payloads[i] = reflect.ValueOf(payload).Elem().Interface()
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
func NewFromType(typ, profileUUID, payloadUUID string) (*Profile, error) {
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

	return &Profile{
		TopLevel: profiles.TopLevel{
			PayloadType:    profiles.PayloadTypeConfiguration,
			PayloadVersion: profiles.TopLevelPayloadVersion1,
			PayloadUUID:    profileUUID,
			PayloadContent: []profiles.ProfilePayload{pay},
		},
	}, nil
}

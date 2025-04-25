package schema

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/korylprince/go-yaml"
)

var ErrUnexpectedBool = errors.New("unexpected boolean")

// IntegerNumber holds either an int64 or a float64
type IntegerNumber struct {
	*int64
	*float64
}

func (in *IntegerNumber) IsInt64() bool {
	return in.int64 != nil
}

func (in *IntegerNumber) IsFloat64() bool {
	return in.float64 != nil
}

// Int64 returns the underlying int64. Int64 will panic if the underlying type is not int64
func (in *IntegerNumber) Int64() int64 {
	return *in.int64
}

// Float64 returns the underlying float64. Float64 will panic if the underlying type is not float64
func (in *IntegerNumber) Float64() float64 {
	return *in.float64
}

func (in *IntegerNumber) Value() any {
	if in.IsInt64() {
		return *in.int64
	} else if in.IsFloat64() {
		return *in.float64
	}
	return nil
}
func (in *IntegerNumber) UnmarshalYAML(data []byte) error {
	if bytes.Contains(data, []byte{'.'}) {
		var f float64
		if err := yaml.Unmarshal(data, &f); err != nil {
			return err
		}
		in.float64 = &f
		return nil
	}

	var i int64
	if err := yaml.Unmarshal(data, &i); err != nil {
		return err
	}
	in.int64 = &i
	return nil
}

func (in *IntegerNumber) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(in.Value())
}

func NewIntegerNumber[T int64 | float64](v T) *IntegerNumber {
	switch val := any(v).(type) {
	case int64:
		return &IntegerNumber{int64: &val}
	case float64:
		return &IntegerNumber{float64: &val}
	default:
		panic(fmt.Sprintf("unexpected type %T", val))
	}
}

// IntegerNumberString holds an int64, float64, or a string
type IntegerNumberString struct {
	*int64
	*float64
	*string
}

func (in *IntegerNumberString) IsInt64() bool {
	return in.int64 != nil
}

func (in *IntegerNumberString) IsFloat64() bool {
	return in.float64 != nil
}

func (in *IntegerNumberString) IsString() bool {
	return in.string != nil
}

// Int64 returns the underlying int64. Int64 will panic if the underlying type is not int64
func (in *IntegerNumberString) Int64() int64 {
	return *in.int64
}

// Float64 returns the underlying float64. Float64 will panic if the underlying type is not float64
func (in *IntegerNumberString) Float64() float64 {
	return *in.float64
}

// String returns the underlying string. String will panic if the underlying type is not string
func (in *IntegerNumberString) String() string {
	return *in.string
}

func (in *IntegerNumberString) Value() any {
	if in.IsInt64() {
		return *in.int64
	} else if in.IsFloat64() {
		return *in.float64
	} else if in.IsString() {
		return *in.string
	}
	return nil
}

func (in *IntegerNumberString) UnmarshalYAML(data []byte) error {
	//verify string isn't a bool
	var b bool
	if err := yaml.Unmarshal(data, &b); err == nil {
		return ErrUnexpectedBool
	}

	if bytes.Contains(data, []byte{'.'}) {
		var f float64
		if err := yaml.Unmarshal(data, &f); err != nil {
			var s string
			if err := yaml.Unmarshal(data, &s); err != nil {
				return err
			}
			in.string = &s
			return nil
		}
		in.float64 = &f
		return nil
	}

	var i int64
	if err := yaml.Unmarshal(data, &i); err != nil {
		var s string
		if err := yaml.Unmarshal(data, &s); err != nil {
			return err
		}
		in.string = &s
		return nil
	}
	in.int64 = &i
	return nil
}

func (in *IntegerNumberString) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(in.Value())
}

func NewIntegerNumberString[T int64 | float64 | string](v T) *IntegerNumberString {
	switch val := any(v).(type) {
	case int64:
		return &IntegerNumberString{int64: &val}
	case float64:
		return &IntegerNumberString{float64: &val}
	case string:
		return &IntegerNumberString{string: &val}
	default:
		panic(fmt.Sprintf("unexpected type %T", val))
	}
}

// BooleanIntegerNumberString holds a bool, int64, float64, or a string
type BooleanIntegerNumberString struct {
	*bool
	*int64
	*float64
	*string
}

func (in *BooleanIntegerNumberString) IsBool() bool {
	return in.bool != nil
}

func (in *BooleanIntegerNumberString) IsInt64() bool {
	return in.int64 != nil
}

func (in *BooleanIntegerNumberString) IsFloat64() bool {
	return in.float64 != nil
}

func (in *BooleanIntegerNumberString) IsString() bool {
	return in.string != nil
}

// Bool returns the underlying bool. Bool will panic if the underlying type is not bool
func (in *BooleanIntegerNumberString) Bool() bool {
	return *in.bool
}

// Int64 returns the underlying int64. Int64 will panic if the underlying type is not int64
func (in *BooleanIntegerNumberString) Int64() int64 {
	return *in.int64
}

// Float64 returns the underlying float64. Float64 will panic if the underlying type is not float64
func (in *BooleanIntegerNumberString) Float64() float64 {
	return *in.float64
}

// String returns the underlying string. String will panic if the underlying type is not string
func (in *BooleanIntegerNumberString) String() string {
	return *in.string
}

func (in *BooleanIntegerNumberString) Value() any {
	if in.IsBool() {
		return *in.bool
	} else if in.IsInt64() {
		return *in.int64
	} else if in.IsFloat64() {
		return *in.float64
	} else if in.IsString() {
		return *in.string
	}
	return nil
}

func (in *BooleanIntegerNumberString) UnmarshalYAML(data []byte) error {
	var b bool
	if err := yaml.Unmarshal(data, &b); err == nil {
		in.bool = &b
		return nil
	}

	if bytes.Contains(data, []byte{'.'}) {
		var f float64
		if err := yaml.Unmarshal(data, &f); err != nil {
			var s string
			if err := yaml.Unmarshal(data, &s); err != nil {
				return err
			}
			in.string = &s
			return nil
		}
		in.float64 = &f
		return nil
	}

	var i int64
	if err := yaml.Unmarshal(data, &i); err != nil {
		var s string
		if err := yaml.Unmarshal(data, &s); err != nil {
			return err
		}
		in.string = &s
		return nil
	}
	in.int64 = &i
	return nil
}

func (in *BooleanIntegerNumberString) MarshalYAML() ([]byte, error) {
	return yaml.Marshal(in.Value())
}

func NewBooleanIntegerNumberString[T bool | int64 | float64 | string](v T) *BooleanIntegerNumberString {
	switch val := any(v).(type) {
	case bool:
		return &BooleanIntegerNumberString{bool: &val}
	case int64:
		return &BooleanIntegerNumberString{int64: &val}
	case float64:
		return &BooleanIntegerNumberString{float64: &val}
	case string:
		return &BooleanIntegerNumberString{string: &val}
	default:
		panic(fmt.Sprintf("unexpected type %T", val))
	}
}

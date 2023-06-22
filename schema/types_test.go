package schema_test

import (
	"testing"

	"github.com/korylprince/go-adm/schema"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func TestValues(t *testing.T) {
	if schema.FromValue(schema.NewValue(10)) != 10 {
		t.Error("failed test: int(10)")
	}
	if schema.FromValue(schema.NewValue[int64](10)) != 10 {
		t.Error("failed test: int64(10)")
	}
	if schema.FromValue(schema.NewValue[uint64](10)) != 10 {
		t.Error("failed test: uint64(10)")
	}
	if schema.FromValue(schema.NewValue[float64](10)) != 10 {
		t.Error("failed test: float64(10)")
	}
	if schema.FromValue(schema.NewValue("test")) != "test" {
		t.Error("failed test: string(\"test\")")
	}
	if schema.FromValue(schema.NewValue(true)) != true {
		t.Error("failed test: bool(true)")
	}
	if schema.FromValue(schema.NewValue(false)) != false {
		t.Error("failed test: bool(false)")
	}
	if !slices.Equal(schema.FromValue(schema.NewValue([]string{"test"})), []string{"test"}) {
		t.Error("failed test: []string{\"test\"}")
	}
	if !maps.Equal(schema.FromValue(schema.NewValue(map[string]string{"test": "test"})), map[string]string{"test": "test"}) {
		t.Error("failed test: map[string]string{\"test\": \"test\"}")
	}

	// testing for panics
	schema.FromValue((*int)(nil))
	schema.FromValue((*int64)(nil))
	schema.FromValue((*float64)(nil))
	schema.FromValue((*string)(nil))
	schema.FromValue((*bool)(nil))
	schema.FromValue((*[]string)(nil))
	schema.FromValue((*map[string]string)(nil))
}

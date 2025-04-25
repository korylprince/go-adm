package schema_test

import (
	"bytes"
	"testing"

	schema "github.com/korylprince/go-adm/yamlschema"
	"github.com/korylprince/go-yaml"
)

const testDoc = `i1: 123
f1: 1.23
i2: 321
f2: 3.21
s1: "123"
s2: "1.23"
i3: 123
f3: 1.23
s3: "123"
s4: "1.23"
b1: false
b2: true
`

var (
	badIntNum        = []string{"1.1.1", "str", `"123"`, "[]", "{}", "\"map.key\": []"}
	badIntNumStr     = []string{"[]", "{}", "\"map.key\": []"}
	badBoolIntNumStr = []string{"[]", "{}", "\"map.key\": []"}
)

func TestTypes(t *testing.T) {
	type test struct {
		I1 *schema.IntegerNumber              `yaml:"i1"`
		F1 *schema.IntegerNumber              `yaml:"f1"`
		I2 *schema.IntegerNumberString        `yaml:"i2"`
		F2 *schema.IntegerNumberString        `yaml:"f2"`
		S1 *schema.IntegerNumberString        `yaml:"s1"`
		S2 *schema.IntegerNumberString        `yaml:"s2"`
		I3 *schema.BooleanIntegerNumberString `yaml:"i3"`
		F3 *schema.BooleanIntegerNumberString `yaml:"f3"`
		S3 *schema.BooleanIntegerNumberString `yaml:"s3"`
		S4 *schema.BooleanIntegerNumberString `yaml:"s4"`
		B1 *schema.BooleanIntegerNumberString `yaml:"b1"`
		B2 *schema.BooleanIntegerNumberString `yaml:"b2"`
	}

	testdata := &test{
		I1: schema.NewIntegerNumber[int64](123),
		F1: schema.NewIntegerNumber(1.23),
		I2: schema.NewIntegerNumberString[int64](321),
		F2: schema.NewIntegerNumberString(3.21),
		S1: schema.NewIntegerNumberString("123"),
		S2: schema.NewIntegerNumberString("1.23"),
		I3: schema.NewBooleanIntegerNumberString[int64](123),
		F3: schema.NewBooleanIntegerNumberString(1.23),
		S3: schema.NewBooleanIntegerNumberString("123"),
		S4: schema.NewBooleanIntegerNumberString("1.23"),
		B1: schema.NewBooleanIntegerNumberString(false),
		B2: schema.NewBooleanIntegerNumberString(true),
	}

	testcase := new(test)
	if err := yaml.Unmarshal([]byte(testDoc), testcase); err != nil {
		t.Fatalf("could not unmarshal test data: %v", err)
	}

	if !testcase.I1.IsInt64() || testcase.I1.Int64() != testdata.I1.Int64() {
		t.Errorf("expected I1 to be equal: want %v, have %v (%[2]T)", testdata.I1.Int64(), testcase.I1.Value())
	}

	if !testcase.F1.IsFloat64() || testcase.F1.Float64() != testdata.F1.Float64() {
		t.Errorf("expected F1 to be equal: want %v, have %v (%[2]T)", testdata.F1.Int64(), testcase.F1.Value())
	}

	if !testcase.I2.IsInt64() || testcase.I2.Int64() != testdata.I2.Int64() {
		t.Errorf("expected I2 to be equal: want %v, have %v (%[2]T)", testdata.I2.Int64(), testcase.I2.Value())
	}

	if !testcase.F2.IsFloat64() || testcase.F2.Float64() != testdata.F2.Float64() {
		t.Errorf("expected F2 to be equal: want %v, have %v (%[2]T)", testdata.F2.Int64(), testcase.F2.Value())
	}

	if !testcase.S1.IsString() || testcase.S1.String() != testdata.S1.String() {
		t.Errorf("expected S1 to be equal: want %v, have %v (%[2]T)", testdata.S1.String(), testcase.S1.Value())
	}

	if !testcase.S2.IsString() || testcase.S2.String() != testdata.S2.String() {
		t.Errorf("expected S2 to be equal: want %v, have %v (%[2]T)", testdata.S2.String(), testcase.S2.Value())
	}

	if !testcase.I3.IsInt64() || testcase.I3.Int64() != testdata.I3.Int64() {
		t.Errorf("expected I3 to be equal: want %v, have %v (%[2]T)", testdata.I3.Int64(), testcase.I3.Value())
	}
	if !testcase.F3.IsFloat64() || testcase.F3.Float64() != testdata.F3.Float64() {
		t.Errorf("expected F3 to be equal: want %v, have %v (%[2]T)", testdata.F3.Int64(), testcase.F3.Value())
	}
	if !testcase.S3.IsString() || testcase.S3.String() != testdata.S3.String() {
		t.Errorf("expected S3 to be equal: want %v, have %v (%[2]T)", testdata.S3.String(), testcase.S3.Value())
	}
	if !testcase.S4.IsString() || testcase.S4.String() != testdata.S4.String() {
		t.Errorf("expected S4 to be equal: want %v, have %v (%[2]T)", testdata.S4.String(), testcase.S4.Value())
	}
	if !testcase.B1.IsBool() || testcase.B1.Bool() != testdata.B1.Bool() {
		t.Errorf("expected B2 to be equal: want %v, have %v (%[2]T)", testdata.B1.Bool(), testcase.B1.Bool())
	}
	if !testcase.B2.IsBool() || testcase.B2.Bool() != testdata.B2.Bool() {
		t.Errorf("expected B2 to be equal: want %v, have %v (%[2]T)", testdata.B2.Bool(), testcase.B2.Bool())
	}

	buf, err := yaml.Marshal(testcase)
	if err != nil {
		t.Fatalf("could not marshal test case: %v", err)
	}

	if !bytes.Equal([]byte(testDoc), buf) {
		t.Errorf("expected end to end marshaling to be equal")
	}

	{
		nilcase := new(schema.IntegerNumber)
		if err := yaml.Unmarshal([]byte("null"), nilcase); err != nil {
			t.Errorf("could not unmarshal IntegerNumber null test case: %v", err)
		} else if nilcase.Value() != nil {
			t.Errorf("expected IntegerNumber null to be nil: %#v", nilcase.Value())
		}
	}

	{
		nilcase := new(schema.IntegerNumberString)
		if err := yaml.Unmarshal([]byte("null"), nilcase); err != nil {
			t.Errorf("could not unmarshal IntegerNumberString null test case: %v", err)
		} else if nilcase.Value() != nil {
			t.Errorf("expected IntegerNumberString null to be nil: %#v", nilcase.Value())
		}
	}

	{
		nilcase := new(schema.BooleanIntegerNumberString)
		if err := yaml.Unmarshal([]byte("null"), nilcase); err != nil {
			t.Errorf("could not unmarshal BooleanIntegerNumberString null test case: %v", err)
		} else if nilcase.Value() != nil {
			t.Errorf("expected BooleanIntegerNumberString null to be nil: %#v", nilcase.Value())
		}
	}

	for _, bad := range badIntNum {
		testcase := new(schema.IntegerNumber)
		if err := yaml.Unmarshal([]byte(bad), testcase); err == nil {
			t.Errorf("expected IntegerNumber test case to fail: %s (%#v)", bad, testcase.Value())
		}
	}

	for _, bad := range badIntNumStr {
		testcase := new(schema.IntegerNumberString)
		if err := yaml.Unmarshal([]byte(bad), testcase); err == nil {
			t.Errorf("expected IntegerNumberString test case to fail: %s (%#v)", bad, testcase.Value())
		}
	}

	for _, bad := range badBoolIntNumStr {
		testcase := new(schema.BooleanIntegerNumberString)
		if err := yaml.Unmarshal([]byte(bad), testcase); err == nil {
			t.Errorf("expected IntegerNumberString test case to fail: %s (%#v)", bad, testcase.Value())
		}
	}
}

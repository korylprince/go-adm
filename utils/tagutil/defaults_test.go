package tagutil_test

import (
	"testing"

	"github.com/korylprince/go-adm/generated/declarative"
	"github.com/korylprince/go-adm/utils/tagutil"
)

type EnumInt int64

const (
	EnumIntTest EnumInt = 2
)

type EnumFloat float64

const (
	EnumFloatTest EnumFloat = 1.618
)

type EnumStr string

const (
	EnumStrTest EnumStr = "enum"
)

type T struct {
	B  *bool      `default:"true"`
	B2 *bool      `default:"false"`
	I  *int       `default:"10"`
	F  *float32   `default:"3.14"`
	S  *string    `default:"test"`
	EI *EnumInt   `default:"2"`
	EF *EnumFloat `default:"1.618"`
	ES *EnumStr   `default:"enum"`
	T  *T
	TA []*T
}

func TestStructDefaults(t *testing.T) {
	test := new(T)
	test2 := new(T)
	test3 := new(T)
	test.T = test2
	test.TA = []*T{test3}

	if err := tagutil.SetDefaults(test); err != nil {
		t.Fatalf("could not set struct defaults: %v", err)
	}

	check := func(test *T) {
		if !*test.B {
			t.Errorf("T.B is not default: have: %v, want: true", *test.B)
		}
		if *test.B2 {
			t.Errorf("T.B2 is not default: have: %v, want: false", *test.B2)
		}
		if *test.I != 10 {
			t.Errorf("T.I is not default: have: %v, want: 10", *test.I)
		}
		if *test.F != 3.14 {
			t.Errorf("T.F is not default: have: %v, want: 3.14", *test.F)
		}
		if *test.S != "test" {
			t.Errorf("T.S is not default: have: %v, want: \"test\"", *test.S)
		}
		if *test.EI != EnumIntTest {
			t.Errorf("T.EI is not default: have: %v, want: 3.14", *test.EI)
		}
		if *test.EF != EnumFloatTest {
			t.Errorf("T.EF is not default: have: %v, want: 1.618", *test.EF)
		}
		if *test.ES != EnumStrTest {
			t.Errorf("T.ES is not default: have: %v, want: \"enum\"", *test.ES)
		}
	}
	check(test)
	check(test2)
	check(test3)

	for typ := range declarative.DeclarationMap {
		decl, err := declarative.NewFromType(typ, "id", "tok")
		if err != nil {
			t.Errorf("could not generate declaration %s: %v", typ, err)
			continue
		}

		err = tagutil.SetDefaults(decl)
		if err != nil {
			t.Errorf("could not set defaults for declaration %s: %v", typ, err)
		}
	}
}

func TestValueStructDefaults(t *testing.T) {
	type child struct {
		Mode   *string `default:"RSA"`
		Values []string
	}
	nested := struct{ Child child }{}
	slice := struct{ Items []child }{Items: []child{{}}}
	items := []child{{}}
	pointerSlice := struct{ Items *[]child }{Items: &items}
	for _, test := range []struct {
		name  string
		value any
		child *child
	}{
		{"nested struct", &nested, &nested.Child},
		{"slice of structs", &slice, &slice.Items[0]},
		{"pointer to slice of structs", &pointerSlice, &items[0]},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := tagutil.SetDefaults(test.value); err != nil {
				t.Fatal(err)
			}
			if test.child.Mode == nil || *test.child.Mode != "RSA" {
				t.Errorf("default was not applied: %+v", test.child)
			}
			if test.child.Values == nil {
				t.Error("slice was not initialized")
			}
		})
	}
}

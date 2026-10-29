package schema_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/utils/replace"
	schema "github.com/korylprince/go-adm/yamlschema"
)

// encode renders the YAML schema src with reps and returns the source.
func encode(t *testing.T, src string, reps replace.Replacements) string {
	t.Helper()

	s, err := schema.New(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}

	f := jen.NewFile("test")
	schema.NewEncoder(f, schema.WithReplacements(reps)).Encode(s)

	buf := new(bytes.Buffer)
	if err := f.Render(buf); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

// Enum type names and constants are assembled from multiple parts. Once
// replacements must still land once on the final identifier, not on each part.
func TestEncodeOnceCompositeNames(t *testing.T) {
	src := encode(t, `
title: Root
type: object
properties:
  DeviceMode:
    type: string
    enum:
    - Caldav
    - Carddav
`, replace.Replacements{
		{Match: "^", Replacement: "Bar", Types: []string{"struct", "const"}, Once: true},
		{Match: "$", Replacement: "Info", Types: []string{"struct", "const"}, Once: true},
		{Match: `^(.*)Caldav(.*)$`, Replacement: "${1}CalDAV${2}", Types: []string{"struct", "const", "field"}},
	})

	for _, want := range []string{
		"type BarSchemaInfo struct",
		"type BarSchemaDeviceModeInfo string",
		"BarSchemaDeviceModeCalDAVInfo",
		"BarSchemaDeviceModeCarddavInfo",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("expected %q in output:\n%s", want, src)
		}
	}
	if strings.Contains(src, "BarBar") || strings.Contains(src, "InfoInfo") {
		t.Errorf("expected prefix and suffix once, got output:\n%s", src)
	}
}

// Ordinary replacements on a complete enum constant are scoped to const, even
// when once rules are configured. A struct-only rule must not reach the
// constants and collapse distinct values into one name.
func TestEncodeOnceConstScope(t *testing.T) {
	const yaml = `
title: Root
type: object
properties:
  Mode:
    type: string
    enum:
    - Url
    - URL
`
	ordinary := &replace.Replacement{Match: "Url", Replacement: "URL", Types: []string{"struct"}}
	never := &replace.Replacement{Match: "^NeverMatches$", Replacement: "X", Types: []string{"const"}, Once: true}

	for name, reps := range map[string]replace.Replacements{
		"without once": {ordinary},
		"with once":    {ordinary, never},
	} {
		t.Run(name, func(t *testing.T) {
			src := encode(t, yaml, reps)
			for _, want := range []string{
				"SchemaModeUrl SchemaMode",
				"SchemaModeURL SchemaMode",
			} {
				if !strings.Contains(src, want) {
					t.Errorf("expected %q in output:\n%s", want, src)
				}
			}
		})
	}
}

// An assembled enum name is normalized with the ordinary rules before the
// once pass, so a once rule sees the ordinary rules' output. The ordinary pass
// runs whether or not once rules are configured, so a once rule that never
// matches doesn't change the output.
func TestEncodeOnceAfterCompositeNormalization(t *testing.T) {
	const yaml = `
title: Root
type: object
properties:
  Mode:
    type: string
    enum:
    - Foo
`
	ordinary := &replace.Replacement{Match: "^SchemaMode", Replacement: "Choice", Types: []string{"struct"}}
	never := &replace.Replacement{Match: "^NeverMatches$", Replacement: "X", Types: []string{"struct"}, Once: true}
	once := &replace.Replacement{Match: "^Choice", Replacement: "AcmeChoice", Types: []string{"struct"}, Once: true}

	for _, test := range []struct {
		name  string
		reps  replace.Replacements
		wants []string
	}{
		{"without once", replace.Replacements{ordinary}, []string{"type Choice string", "ChoiceFoo Choice"}},
		{"with nonmatching once", replace.Replacements{ordinary, never}, []string{"type Choice string", "ChoiceFoo Choice"}},
		{"with once", replace.Replacements{ordinary, once}, []string{"type AcmeChoice string", "AcmeChoiceFoo AcmeChoice"}},
	} {
		t.Run(test.name, func(t *testing.T) {
			src := encode(t, yaml, test.reps)
			for _, want := range test.wants {
				if !strings.Contains(src, want) {
					t.Errorf("expected %q in output:\n%s", want, src)
				}
			}
		})
	}
}

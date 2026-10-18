package schema_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/utils/replace"
	schema "github.com/korylprince/go-adm/yamlschema"
)

// Enum type names and constants are assembled from multiple parts. Once
// replacements must still land once on the final identifier, not on each part.
func TestEncodeOnceCompositeNames(t *testing.T) {
	s, err := schema.New(strings.NewReader(`
title: Root
type: object
properties:
  DeviceMode:
    type: string
    enum:
    - Caldav
    - Carddav
`))
	if err != nil {
		t.Fatal(err)
	}

	reps := replace.Replacements{
		{Match: "^", Replacement: "Bar", Types: []string{"struct", "const"}, Once: true},
		{Match: "$", Replacement: "Info", Types: []string{"struct", "const"}, Once: true},
		{Match: `^(.*)Caldav(.*)$`, Replacement: "${1}CalDAV${2}", Types: []string{"struct", "const", "field"}},
	}

	f := jen.NewFile("test")
	schema.NewEncoder(f, schema.WithReplacements(reps)).Encode(s)

	buf := new(bytes.Buffer)
	if err := f.Render(buf); err != nil {
		t.Fatal(err)
	}
	src := buf.String()

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

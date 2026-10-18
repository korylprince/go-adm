package gen_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	structgen "github.com/korylprince/go-adm/generator/structgen"
	"github.com/korylprince/go-adm/utils/replace"
)

// write drops a schema into a temp file for GenerateFromFiles to read.
func write(t *testing.T, name, body string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(body), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func generate(t *testing.T, path string) string {
	t.Helper()

	buf := new(strings.Builder)
	if err := structgen.GenerateFromFiles([]string{path}, "test", nil, []string{"json"}, buf); err != nil {
		t.Fatal(err)
	}
	return buf.String()
}

// A schema whose top level payload keys describe a free-form dictionary becomes
// a schema.Map rather than a schema.Struct. The encoder only switched on Enum
// and Struct, so such a schema generated nothing at all -- and reported no
// error, which is the part that makes it easy to miss.
func TestGenerateTopLevelMap(t *testing.T) {
	path := write(t, "anydict.yaml", `
title: Any Dictionary Schema
description: Arbitrary keys.
payloadkeys:
- key: ANY
  type: <any>
  presence: optional
`)

	src := generate(t, path)

	if want := "type AnyDictionarySchema map[string]any"; !strings.Contains(src, want) {
		t.Errorf("expected %q in output:\n%s", want, src)
	}
}

// Once replacements apply to the final enum type and constant identifiers,
// not once per component, and ordinary rules still process the result.
func TestGenerateOncePrefixOnEnum(t *testing.T) {
	path := write(t, "enum.yaml", `
title: Enum Schema
payloadkeys:
- key: Mode
  type: <string>
  presence: required
  rangelist:
  - Caldav
  - Carddav
`)

	reps := replace.Replacements{
		{Match: "^", Replacement: "Bar", Types: []string{"struct", "const"}, Once: true},
		{Match: "$", Replacement: "Info", Types: []string{"struct", "const"}, Once: true},
		{Match: `^(.*)Caldav(.*)$`, Replacement: "${1}CalDAV${2}", Types: []string{"struct", "const", "field"}},
	}

	buf := new(strings.Builder)
	if err := structgen.GenerateFromFiles([]string{path}, "test", reps, []string{"json"}, buf); err != nil {
		t.Fatal(err)
	}
	src := buf.String()

	for _, want := range []string{
		"type BarEnumSchemaInfo struct",
		"type BarModeInfo string",
		"BarModeCalDAVInfo",
		"BarModeCarddavInfo",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("expected %q in output:\n%s", want, src)
		}
	}
	if strings.Contains(src, "BarBar") || strings.Contains(src, "InfoInfo") {
		t.Errorf("expected prefix and suffix once, got output:\n%s", src)
	}
}

// The ordinary case, to show the map handling didn't displace it.
func TestGenerateTopLevelStruct(t *testing.T) {
	path := write(t, "normal.yaml", `
title: Normal Schema
payloadkeys:
- key: Name
  type: <string>
  presence: required
`)

	src := generate(t, path)

	if want := "type NormalSchema struct"; !strings.Contains(src, want) {
		t.Errorf("expected %q in output:\n%s", want, src)
	}
}

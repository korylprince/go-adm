package schema_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/schema"
)

// Apple spells a free-form dictionary two ways: with an explicit ANY subkey, or
// by simply omitting subkeys. The tests here cover the second spelling, which
// used to render as an empty struct and silently discard the dictionary's
// contents on unmarshal. See declarative/protocol/statusreport.yaml, where
// StatusItems and Details both take that form.

func TestSubKeylessDictionary(t *testing.T) {
	s := anyTestSchema("Report",
		&schema.PayloadKey{
			Key:      "Details",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceOptional,
			Content:  "A dictionary that contains further details.",
		},
	)

	t.Run("is a map, not a struct", func(t *testing.T) {
		key := s.PayloadKeys[0]
		if key.IsStruct() {
			t.Error("IsStruct should be false for a dictionary with no subkeys")
		}
		if !key.IsMap() {
			t.Error("IsMap should be true for a dictionary with no subkeys")
		}
	})

	t.Run("renders as an inline map", func(t *testing.T) {
		src := renderFile(t, []*schema.Schema{s})

		if !strings.Contains(src, "map[string]any") {
			t.Errorf("expected an inline map:\n%s", src)
		}
		if strings.Contains(src, "type Details struct{}") {
			t.Errorf("free-form dictionary rendered as an empty struct:\n%s", src)
		}
	})
}

// A schema with no payload keys at all is a different thing entirely: an empty
// document, like an MDM command that takes no arguments. Those must keep
// rendering as empty structs, so the fix above cannot be a blanket rule.
func TestEmptySchemaStaysAStruct(t *testing.T) {
	s := &schema.Schema{Title: "DeviceConfiguredCommand", Payload: &schema.Payload{}}

	file := schema.NewFile([]*schema.Schema{s},
		schema.WithIncludeEmptyPayloadKeys(true),
		schema.WithIncludeEmptyResponseKeys(true),
	)

	f := jen.NewFile("test")
	schema.NewEncoder(f).Encode(file)

	buf := new(strings.Builder)
	if err := f.Render(buf); err != nil {
		t.Fatal(err)
	}
	src := buf.String()

	for _, want := range []string{
		"type DeviceConfiguredCommand struct{}",
		"type DeviceConfiguredCommandResponse struct{}",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("expected %q in output:\n%s", want, src)
		}
	}
	if strings.Contains(src, "map[string]any") {
		t.Errorf("an empty schema should not become a free-form map:\n%s", src)
	}
}

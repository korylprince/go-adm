package schema_test

import (
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/korylprince/go-adm/schema"
)

// helper: build a Schema from payloadkeys
func anyTestSchema(title string, payloadkeys ...*schema.PayloadKey) *schema.Schema {
	return &schema.Schema{
		Title:       title,
		PayloadKeys: payloadkeys,
		Payload:     &schema.Payload{},
	}
}

// helper: render a File through the Encoder and return the generated Go source
func renderFile(t *testing.T, schemas []*schema.Schema) string {
	t.Helper()
	file := schema.NewFile(schemas)
	f := jen.NewFile("test")
	e := schema.NewEncoder(f)
	e.Encode(file)
	buf := &strings.Builder{}
	if err := f.Render(buf); err != nil {
		t.Fatalf("render error: %v", err)
	}
	return buf.String()
}

// Pattern 1: key: ANY + type: <any> → map[string]any
// Example: managed.application.configuration.yaml Configuration dict
func TestANY_Pattern1_MapStringAny(t *testing.T) {
	s := anyTestSchema("Pattern1",
		&schema.PayloadKey{
			Key:  "Identifier",
			Type: schema.PayloadKeyTypeString,
		},
		&schema.PayloadKey{
			Key:      "Configuration",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceOptional,
			SubKeys: []*schema.PayloadKey{
				{
					Key:      "ANY",
					Type:     schema.PayloadKeyTypeAny,
					Presence: schema.PayloadKeyPresenceOptional,
				},
			},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	// Should generate a map type, not a struct, for Configuration
	if !strings.Contains(src, "map[string]any") {
		t.Errorf("Pattern 1: expected map[string]any in output:\n%s", src)
	}
	// Should NOT have a struct field named "ANY"
	if strings.Contains(src, "ANY") && strings.Contains(src, "struct") {
		// Only fail if "ANY" appears as a field inside a struct for Configuration
		if strings.Contains(src, "ANY any") || strings.Contains(src, "ANY *any") {
			t.Errorf("Pattern 1: should not have struct field ANY:\n%s", src)
		}
	}
}

// Pattern 2: key: ANY + type: <string> → map[string]string
// Example: TopLevel.yaml ConsentTextItem, relay.managed.yaml HTTP Headers
func TestANY_Pattern2_MapStringString(t *testing.T) {
	s := anyTestSchema("Pattern2",
		&schema.PayloadKey{
			Key:  "Name",
			Type: schema.PayloadKeyTypeString,
		},
		&schema.PayloadKey{
			Key:      "Headers",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceOptional,
			SubKeys: []*schema.PayloadKey{
				{
					Key:      "ANY",
					Type:     schema.PayloadKeyTypeString,
					Presence: schema.PayloadKeyPresenceRequired,
				},
			},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	if !strings.Contains(src, "map[string]string") {
		t.Errorf("Pattern 2: expected map[string]string in output:\n%s", src)
	}
}

// Pattern 3: key: ANY + type: <array> → map[string][]string
// Example: system-extension-policy.yaml AllowedSystemExtensionTypes
func TestANY_Pattern3_MapStringSlice(t *testing.T) {
	s := anyTestSchema("Pattern3",
		&schema.PayloadKey{
			Key:  "AllowedTeamIdentifiers",
			Type: schema.PayloadKeyTypeString,
		},
		&schema.PayloadKey{
			Key:      "AllowedExtensionTypes",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceOptional,
			SubKeys: []*schema.PayloadKey{
				{
					Key:      "ANY",
					Type:     schema.PayloadKeyTypeArray,
					Presence: schema.PayloadKeyPresenceOptional,
					SubKeys: []*schema.PayloadKey{
						{
							Key:      "AllowedExtensionTypesItems",
							Type:     schema.PayloadKeyTypeString,
							Presence: schema.PayloadKeyPresenceRequired,
						},
					},
				},
			},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	if !strings.Contains(src, "map[string][]string") {
		t.Errorf("Pattern 3: expected map[string][]string in output:\n%s", src)
	}
}

// Pattern 4: key: ANY + type: <dictionary> (with subkeys) → map[string]Struct
// Example: ManagedClient.preferences.yaml PayloadContent
func TestANY_Pattern4_MapStringStruct(t *testing.T) {
	s := anyTestSchema("Pattern4",
		&schema.PayloadKey{
			Key:      "Domains",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceRequired,
			SubKeys: []*schema.PayloadKey{
				{
					Key:      "ANY",
					Type:     schema.PayloadKeyTypeDictionary,
					Presence: schema.PayloadKeyPresenceRequired,
					SubKeys: []*schema.PayloadKey{
						{
							Key:      "Forced",
							Type:     schema.PayloadKeyTypeBoolean,
							Presence: schema.PayloadKeyPresenceOptional,
						},
						{
							Key:      "Value",
							Type:     schema.PayloadKeyTypeString,
							Presence: schema.PayloadKeyPresenceOptional,
						},
					},
				},
			},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	// Should produce a map type whose value is a generated struct
	if !strings.Contains(src, "map[string]") {
		t.Errorf("Pattern 4: expected map[string]... in output:\n%s", src)
	}
	// The value struct should have the fields Forced and Value
	if !strings.Contains(src, "Forced") || !strings.Contains(src, "Value") {
		t.Errorf("Pattern 4: expected struct with Forced and Value fields:\n%s", src)
	}
	// Should NOT have a struct field named "ANY" with a type annotation
	// (The struct type NAMED "ANY" is fine — it's the value type of the map)
	lines := strings.Split(src, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// A struct field named ANY would look like: ANY SomeType `json:"ANY"`
		if strings.HasPrefix(trimmed, "ANY ") && strings.Contains(trimmed, "`") {
			t.Errorf("Pattern 4: should not have a struct field named ANY:\n%s", src)
		}
	}
}

// Nested ANY-backed maps should render inline in struct fields, not as extra named
// map declarations that go unused in generated code.
func TestANY_Pattern4_NoUnusedNestedMapType(t *testing.T) {
	s := anyTestSchema("NestedMapTypeTest",
		&schema.PayloadKey{
			Key:      "ExtensionConfigs",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceOptional,
			SubKeys: []*schema.PayloadKey{
				{
					Key:        "ANY",
					Type:       schema.PayloadKeyTypeDictionary,
					Presence:   schema.PayloadKeyPresenceOptional,
					SubKeyType: "AppConfigDictionary",
					SubKeys: []*schema.PayloadKey{
						{
							Key:      "DataAssetReference",
							Type:     schema.PayloadKeyTypeString,
							Presence: schema.PayloadKeyPresenceOptional,
						},
					},
				},
			},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	if !strings.Contains(src, "ExtensionConfigs *map[string]AppConfigDictionary") {
		t.Errorf("expected inline map field type for nested ANY dictionary:\n%s", src)
	}
	if strings.Contains(src, "type ExtensionConfigs map[string]AppConfigDictionary") {
		t.Errorf("unexpected unused named nested map type:\n%s", src)
	}
}

// Pattern 5: key: ANY <descriptor> + type: <dictionary> → map[string]Struct
// Example: device.restrictions.list.yaml "ANY restriction name"
func TestANY_Pattern5_DescriptorMapStringStruct(t *testing.T) {
	s := anyTestSchema("Pattern5",
		&schema.PayloadKey{
			Key:      "Restrictions",
			Type:     schema.PayloadKeyTypeDictionary,
			Presence: schema.PayloadKeyPresenceOptional,
			SubKeys: []*schema.PayloadKey{
				{
					Key:      "ANY restriction name",
					Type:     schema.PayloadKeyTypeDictionary,
					Presence: schema.PayloadKeyPresenceOptional,
					SubKeys: []*schema.PayloadKey{
						{
							Key:      "value",
							Type:     schema.PayloadKeyTypeBoolean,
							Presence: schema.PayloadKeyPresenceRequired,
						},
					},
				},
			},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	// Should produce a map type
	if !strings.Contains(src, "map[string]") {
		t.Errorf("Pattern 5: expected map[string]... in output:\n%s", src)
	}
	// The value struct should have the field "value"
	if !strings.Contains(src, "Value") {
		t.Errorf("Pattern 5: expected struct with Value field:\n%s", src)
	}
}

// Pattern 6: Top-level key: ANY → entire payload is a map
// Example: com.apple.firstethernet.managed.yaml, management/properties.yaml
func TestANY_Pattern6_TopLevelMap(t *testing.T) {
	s := anyTestSchema("Pattern6",
		&schema.PayloadKey{
			Key:      "ANY",
			Type:     schema.PayloadKeyTypeAny,
			Presence: schema.PayloadKeyPresenceOptional,
		},
	)

	file := schema.NewFile([]*schema.Schema{s})

	// The top-level type should be a Map, not a Struct
	var hasMap, hasStruct bool
	for _, typ := range file.Types {
		switch t := typ.(type) {
		case *schema.Map:
			if t.Source == schema.SourcePayloadKeys {
				hasMap = true
			}
		case *schema.Struct:
			if t.Source == schema.SourcePayloadKeys {
				hasStruct = true
			}
		}
	}

	if !hasMap {
		t.Error("Pattern 6: expected top-level Map type")
	}
	if hasStruct {
		t.Error("Pattern 6: should not have top-level Struct type")
	}
}

// Test IsANY helper
func TestIsANY(t *testing.T) {
	tests := []struct {
		key  string
		want bool
	}{
		{"ANY", true},
		{"ANY restriction name", true},
		{"ANY app identifier", true},
		{"ANY index", true},
		{"ANYTHING", false},
		{"NotANY", false},
		{"Configuration", false},
		{"", false},
	}

	for _, tt := range tests {
		pk := &schema.PayloadKey{Key: tt.key}
		if got := pk.IsANY(); got != tt.want {
			t.Errorf("IsANY(%q) = %v, want %v", tt.key, got, tt.want)
		}
	}
}

// Test that IsMap and IsStruct work correctly for all patterns
func TestIsMapIsStruct(t *testing.T) {
	tests := []struct {
		name     string
		key      *schema.PayloadKey
		wantMap  bool
		wantStruc bool
	}{
		{
			name: "Pattern1: ANY+any",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeDictionary,
				SubKeys: []*schema.PayloadKey{
					{Key: "ANY", Type: schema.PayloadKeyTypeAny},
				},
			},
			wantMap: true, wantStruc: false,
		},
		{
			name: "Pattern2: ANY+string",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeDictionary,
				SubKeys: []*schema.PayloadKey{
					{Key: "ANY", Type: schema.PayloadKeyTypeString},
				},
			},
			wantMap: true, wantStruc: false,
		},
		{
			name: "Pattern3: ANY+array",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeDictionary,
				SubKeys: []*schema.PayloadKey{
					{Key: "ANY", Type: schema.PayloadKeyTypeArray, SubKeys: []*schema.PayloadKey{
						{Key: "Items", Type: schema.PayloadKeyTypeString},
					}},
				},
			},
			wantMap: true, wantStruc: false,
		},
		{
			name: "Pattern4: ANY+dictionary",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeDictionary,
				SubKeys: []*schema.PayloadKey{
					{Key: "ANY", Type: schema.PayloadKeyTypeDictionary, SubKeys: []*schema.PayloadKey{
						{Key: "Field1", Type: schema.PayloadKeyTypeString},
					}},
				},
			},
			wantMap: true, wantStruc: false,
		},
		{
			name: "Pattern5: ANY descriptor+dictionary",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeDictionary,
				SubKeys: []*schema.PayloadKey{
					{Key: "ANY restriction name", Type: schema.PayloadKeyTypeDictionary, SubKeys: []*schema.PayloadKey{
						{Key: "value", Type: schema.PayloadKeyTypeBoolean},
					}},
				},
			},
			wantMap: true, wantStruc: false,
		},
		{
			name: "Normal struct",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeDictionary,
				SubKeys: []*schema.PayloadKey{
					{Key: "Field1", Type: schema.PayloadKeyTypeString},
					{Key: "Field2", Type: schema.PayloadKeyTypeInteger},
				},
			},
			wantMap: false, wantStruc: true,
		},
		{
			name: "Non-dictionary",
			key: &schema.PayloadKey{
				Type: schema.PayloadKeyTypeString,
			},
			wantMap: false, wantStruc: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.key.IsMap(); got != tt.wantMap {
				t.Errorf("IsMap() = %v, want %v", got, tt.wantMap)
			}
			if got := tt.key.IsStruct(); got != tt.wantStruc {
				t.Errorf("IsStruct() = %v, want %v", got, tt.wantStruc)
			}
		})
	}
}

// Test that ANY keys don't produce "ANY" in generated type names
func TestANY_Naming(t *testing.T) {
	tests := []struct {
		name       string
		schema     *schema.Schema
		wantNot    string // substring that should NOT appear in output
		wantSubstr string // substring that SHOULD appear in output
	}{
		{
			name: "SubKeyType used for naming",
			schema: anyTestSchema("NamingTest1",
				&schema.PayloadKey{
					Key:  "Apps",
					Type: schema.PayloadKeyTypeDictionary,
					SubKeys: []*schema.PayloadKey{
						{
							Key:        "ANY app identifier",
							Type:       schema.PayloadKeyTypeDictionary,
							SubKeyType: "ManagedApp",
							SubKeys: []*schema.PayloadKey{
								{Key: "Status", Type: schema.PayloadKeyTypeString, Presence: schema.PayloadKeyPresenceRequired},
							},
						},
					},
				},
			),
			wantNot:    "ANYappidentifier",
			wantSubstr: "ManagedApp",
		},
		{
			name: "Descriptor used for naming",
			schema: anyTestSchema("NamingTest2",
				&schema.PayloadKey{
					Key:  "Restrictions",
					Type: schema.PayloadKeyTypeDictionary,
					SubKeys: []*schema.PayloadKey{
						{
							Key:  "ANY restriction name",
							Type: schema.PayloadKeyTypeDictionary,
							SubKeys: []*schema.PayloadKey{
								{Key: "value", Type: schema.PayloadKeyTypeBoolean, Presence: schema.PayloadKeyPresenceRequired},
							},
						},
					},
				},
			),
			wantNot:    "ANYrestriction",
			wantSubstr: "RestrictionName",
		},
		{
			name: "Plain ANY stripped from names",
			schema: anyTestSchema("NamingTest3",
				&schema.PayloadKey{
					Key:  "Identifier",
					Type: schema.PayloadKeyTypeString,
				},
				&schema.PayloadKey{
					Key:  "Config",
					Type: schema.PayloadKeyTypeDictionary,
					SubKeys: []*schema.PayloadKey{
						{
							Key:  "ANY",
							Type: schema.PayloadKeyTypeAny,
						},
					},
				},
			),
			wantNot:    "ANYany",
			wantSubstr: "map[string]any",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := renderFile(t, []*schema.Schema{tt.schema})
			if strings.Contains(src, tt.wantNot) {
				t.Errorf("output should not contain %q:\n%s", tt.wantNot, src)
			}
			if !strings.Contains(src, tt.wantSubstr) {
				t.Errorf("output should contain %q:\n%s", tt.wantSubstr, src)
			}
		})
	}
}

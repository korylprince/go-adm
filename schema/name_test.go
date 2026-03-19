package schema_test

import (
	"testing"

	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/utils/replace"
)

var testSchemaA = &schema.Schema{Title: "SchemaA", PayloadKeys: []*schema.PayloadKey{
	{Key: "KeyA", Type: schema.PayloadKeyTypeDictionary},
	{Key: "KeyB", Type: schema.PayloadKeyTypeDictionary, SubKeys: []*schema.PayloadKey{
		{Key: "SubKeyA", Type: schema.PayloadKeyTypeDictionary},
		{Key: "SubKeyB", Type: schema.PayloadKeyTypeDictionary},
	}},
	{Key: "KeyC", Type: schema.PayloadKeyTypeDictionary},
}}

var testSchemaB = &schema.Schema{Title: "SchemaB", PayloadKeys: []*schema.PayloadKey{
	{Key: "KeyA", Type: schema.PayloadKeyTypeDictionary},
	{Key: "KeyB", Type: schema.PayloadKeyTypeDictionary, SubKeys: []*schema.PayloadKey{
		{Key: "SubKeyA", Type: schema.PayloadKeyTypeDictionary},
		{Key: "SubKeyC", Type: schema.PayloadKeyTypeDictionary},
	}},
	{Key: "KeyD", Type: schema.PayloadKeyTypeDictionary},
}}

var testSchemaB2 = &schema.Schema{Title: "SchemaB", PayloadKeys: []*schema.PayloadKey{
	{Key: "KeyE", Type: schema.PayloadKeyTypeDictionary},
}}

var keyTests = map[*schema.PayloadKey]string{
	testSchemaA.PayloadKeys[0]:            "SchemaAKeyA",
	testSchemaA.PayloadKeys[1]:            "SchemaAKeyB",
	testSchemaA.PayloadKeys[2]:            "KeyC",
	testSchemaA.PayloadKeys[1].SubKeys[0]: "SchemaAKeyBSubKeyA",
	testSchemaA.PayloadKeys[1].SubKeys[1]: "SubKeyB",
	testSchemaB.PayloadKeys[0]:            "SchemaBKeyA",
	testSchemaB.PayloadKeys[1]:            "SchemaBKeyB",
	testSchemaB.PayloadKeys[1].SubKeys[0]: "SchemaBKeyBSubKeyA",
	testSchemaB.PayloadKeys[1].SubKeys[1]: "SubKeyC",
	testSchemaB.PayloadKeys[2]:            "KeyD",
	testSchemaB2.PayloadKeys[0]:           "KeyE",
}

var schemaTests = map[*schema.Schema]string{
	testSchemaA:  "SchemaA",
	testSchemaB:  "SchemaB1",
	testSchemaB2: "SchemaB2",
}

func TestGlobalNamer(t *testing.T) {
	file := schema.NewFile([]*schema.Schema{testSchemaA, testSchemaB, testSchemaB2})
	gn := schema.NewGlobalNamer(file, nil)

	for test, want := range keyTests {
		if have := gn.KeyName(test); have != want {
			t.Errorf("key test failed: have: %s, want: %s", have, want)
		}
	}

	// build map of schemas to payload keys
	schemaMap := make(map[*schema.Schema]*schema.PayloadKey)
	for _, t := range file.Types {
		if strct, ok := t.(*schema.Struct); ok {
			if strct.Source == schema.SourcePayloadKeys {
				schemaMap[strct.Schema] = strct.Key
			}
		}
	}

	for test, want := range schemaTests {
		if have := gn.KeyName(schemaMap[test]); have != want {
			t.Errorf("schema test failed: have: %s, want: %s", have, want)
		}
	}
}

// rootStructKey returns the PayloadKey for the root struct of the given schema.
func rootStructKey(file *schema.File, s *schema.Schema) *schema.PayloadKey {
	for _, t := range file.Types {
		if strct, ok := t.(*schema.Struct); ok {
			if strct.Source == schema.SourcePayloadKeys && strct.Schema == s {
				return strct.Key
			}
		}
	}
	return nil
}

func TestGlobalNamerReplacements(t *testing.T) {
	t.Run("post-replacement collision disambiguates with longer prefix", func(t *testing.T) {
		// Replacement: Apns → APNs for struct type.
		// "apns" normalizes to "Apns"; "APNs" normalizes to "APNs".
		// After replacement both resolve to "APNs", creating a collision
		// that requires GlobalNamer to use schema-title prefixes.
		reps := replace.Replacements{
			{Match: `^(.*)Apns(.*)$`, Replacement: "${1}APNs${2}", Types: []string{"struct"}},
		}

		s1 := &schema.Schema{Title: "SchemaX", PayloadKeys: []*schema.PayloadKey{
			{Key: "apns", Type: schema.PayloadKeyTypeDictionary},
		}}
		s2 := &schema.Schema{Title: "SchemaY", PayloadKeys: []*schema.PayloadKey{
			{Key: "APNs", Type: schema.PayloadKeyTypeDictionary},
		}}

		file := schema.NewFile([]*schema.Schema{s1, s2})
		gn := schema.NewGlobalNamer(file, reps)

		name1 := gn.KeyName(s1.PayloadKeys[0]) // "apns" key
		name2 := gn.KeyName(s2.PayloadKeys[0]) // "APNs" key

		// Raw names must include schema prefix to disambiguate
		if name1 != "SchemaXApns" {
			t.Errorf("expected SchemaXApns, got %s", name1)
		}
		if name2 != "SchemaYAPNs" {
			t.Errorf("expected SchemaYAPNs, got %s", name2)
		}

		// Post-replacement names must be unique
		resolved1 := reps.Replace(name1, replace.Struct)
		resolved2 := reps.Replace(name2, replace.Struct)
		if resolved1 == resolved2 {
			t.Errorf("resolved names should differ: both %q (raw: %q, %q)", resolved1, name1, name2)
		}
	})

	t.Run("fully qualified collision uses numbered suffix", func(t *testing.T) {
		// Replacement: Apn (not followed by lowercase) → APN for struct type.
		// Schema titles "Apn" and "APN" both resolve to "APN" even at the
		// fully-qualified level, forcing the numbered-suffix fallback.
		reps := replace.Replacements{
			{Match: `^(.*)Apn([^a-z].*|)$`, Replacement: "${1}APN${2}", Types: []string{"struct"}},
		}

		s1 := &schema.Schema{Title: "Apn", PayloadKeys: []*schema.PayloadKey{
			{Key: "name", Type: schema.PayloadKeyTypeString},
		}}
		s2 := &schema.Schema{Title: "APN", PayloadKeys: []*schema.PayloadKey{
			{Key: "name", Type: schema.PayloadKeyTypeString},
		}}

		file := schema.NewFile([]*schema.Schema{s1, s2})
		gn := schema.NewGlobalNamer(file, reps)

		key1 := rootStructKey(file, s1)
		key2 := rootStructKey(file, s2)
		if key1 == nil || key2 == nil {
			t.Fatal("could not find root struct keys")
		}

		name1 := gn.KeyName(key1)
		name2 := gn.KeyName(key2)

		// Raw names should be numbered suffixes of the schema title
		if name1 != "Apn1" {
			t.Errorf("expected Apn1, got %s", name1)
		}
		if name2 != "APN2" {
			t.Errorf("expected APN2, got %s", name2)
		}

		// Post-replacement names must be unique
		resolved1 := reps.Replace(name1, replace.Struct)
		resolved2 := reps.Replace(name2, replace.Struct)
		if resolved1 == resolved2 {
			t.Errorf("resolved names should differ: both %q (raw: %q, %q)", resolved1, name1, name2)
		}
	})
}

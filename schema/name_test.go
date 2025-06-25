package schema_test

import (
	"testing"

	"github.com/korylprince/go-adm/schema"
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
	file := schema.NewFile(testSchemaA, testSchemaB, testSchemaB2)
	gn := schema.NewGlobalNamer(file)

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

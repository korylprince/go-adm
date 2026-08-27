package schema_test

import (
	"testing"

	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/utils/replace"
)

// structKey builds a dictionary key that renders as a struct. The placeholder
// subkey is load-bearing: a dictionary with no subkeys at all is free-form and
// renders as an inline map[string]any, which never gets a name of its own.
func structKey(key string, subKeys ...*schema.PayloadKey) *schema.PayloadKey {
	if len(subKeys) == 0 {
		subKeys = []*schema.PayloadKey{{Key: "Placeholder", Type: schema.PayloadKeyTypeString}}
	}
	return &schema.PayloadKey{Key: key, Type: schema.PayloadKeyTypeDictionary, SubKeys: subKeys}
}

var testSchemaA = &schema.Schema{Title: "SchemaA", PayloadKeys: []*schema.PayloadKey{
	structKey("KeyA"),
	structKey("KeyB",
		structKey("SubKeyA"),
		structKey("SubKeyB"),
	),
	structKey("KeyC"),
}}

var testSchemaB = &schema.Schema{Title: "SchemaB", PayloadKeys: []*schema.PayloadKey{
	structKey("KeyA"),
	structKey("KeyB",
		structKey("SubKeyA"),
		structKey("SubKeyC"),
	),
	structKey("KeyD"),
}}

var testSchemaB2 = &schema.Schema{Title: "SchemaB", PayloadKeys: []*schema.PayloadKey{
	structKey("KeyE"),
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

func TestGlobalNamerNameOverrides(t *testing.T) {
	// Two schemas each holding a "Model" key. Left alone the namer would give
	// one of them the bare name "Model" and prefix the other, and which one
	// gets which depends on registration order -- the instability that pinned
	// names exist to remove.
	s1 := &schema.Schema{Title: "SchemaP", PayloadKeys: []*schema.PayloadKey{structKey("Model")}}
	s2 := &schema.Schema{Title: "SchemaQ", PayloadKeys: []*schema.PayloadKey{structKey("Model")}}

	file := schema.NewFile([]*schema.Schema{s1, s2})

	pinned := s1.PayloadKeys[0]
	other := s2.PayloadKeys[0]

	gn := schema.NewGlobalNamer(file, nil, schema.WithNameOverrides(
		map[*schema.PayloadKey]string{pinned: "PinnedModel"},
	))

	if have := gn.KeyName(pinned); have != "PinnedModel" {
		t.Errorf("pinned key name: have: %s, want: PinnedModel", have)
	}

	// The unpinned key is free to take the short name, since the pinned one is
	// no longer competing for it.
	if have := gn.KeyName(other); have != "Model" {
		t.Errorf("unpinned key name: have: %s, want: Model", have)
	}
}

func TestGlobalNamerNameOverridesCollide(t *testing.T) {
	// A pinned name is counted as taken, so an unpinned key that would have
	// claimed it falls back to a longer prefix rather than colliding.
	s1 := &schema.Schema{Title: "SchemaR", PayloadKeys: []*schema.PayloadKey{structKey("Widget")}}
	s2 := &schema.Schema{Title: "SchemaS", PayloadKeys: []*schema.PayloadKey{structKey("Gadget")}}

	file := schema.NewFile([]*schema.Schema{s1, s2})

	gn := schema.NewGlobalNamer(file, nil, schema.WithNameOverrides(
		map[*schema.PayloadKey]string{s1.PayloadKeys[0]: "Gadget"},
	))

	if have := gn.KeyName(s1.PayloadKeys[0]); have != "Gadget" {
		t.Errorf("pinned key name: have: %s, want: Gadget", have)
	}
	if have := gn.KeyName(s2.PayloadKeys[0]); have != "SchemaSGadget" {
		t.Errorf("displaced key name: have: %s, want: SchemaSGadget", have)
	}
}

func TestGlobalNamerFullyQualified(t *testing.T) {
	// SchemaT holds a "Model" that would otherwise win the bare name, and a
	// deeper "State" that nothing else competes for.
	s1 := &schema.Schema{Title: "SchemaT", PayloadKeys: []*schema.PayloadKey{
		structKey("Widget", structKey("Model"), structKey("State")),
	}}
	s2 := &schema.Schema{Title: "SchemaU", PayloadKeys: []*schema.PayloadKey{structKey("Model")}}

	file := schema.NewFile([]*schema.Schema{s1, s2})
	gn := schema.NewGlobalNamer(file, nil, schema.WithFullyQualifiedNames(s1))

	widget := s1.PayloadKeys[0]
	for _, test := range []struct {
		key  *schema.PayloadKey
		want string
	}{
		{widget, "SchemaTWidget"},
		{widget.SubKeys[0], "SchemaTWidgetModel"},
		{widget.SubKeys[1], "SchemaTWidgetState"},
	} {
		if have := gn.KeyName(test.key); have != test.want {
			t.Errorf("qualified name: have: %s, want: %s", have, test.want)
		}
	}

	// the unqualified schema is unaffected, and can still take the short name
	if have := gn.KeyName(s2.PayloadKeys[0]); have != "Model" {
		t.Errorf("unqualified schema name: have: %s, want: Model", have)
	}
}

func TestGlobalNamerFullyQualifiedFromPinnedAncestor(t *testing.T) {
	// A pinned name is already fully qualified, so descendants continue from it
	// rather than repeating the path above it.
	//
	// The array here mirrors how Apple declares one: a `reasons` array holding
	// a `_Reasons` element. Only the element becomes a type, and both segments
	// normalize to the same word, so the qualified name would read
	// PinnedReasonsReasons without collapsing the repeat.
	element := structKey("_Reasons")
	outer := structKey("Outer", &schema.PayloadKey{
		Key:     "reasons",
		Type:    schema.PayloadKeyTypeArray,
		SubKeys: []*schema.PayloadKey{element},
	})
	s := &schema.Schema{Title: "SchemaV", PayloadKeys: []*schema.PayloadKey{outer}}

	file := schema.NewFile([]*schema.Schema{s})

	gn := schema.NewGlobalNamer(file, nil,
		schema.WithNameOverrides(map[*schema.PayloadKey]string{outer: "Pinned"}),
		schema.WithFullyQualifiedNames(s),
	)

	if have := gn.KeyName(outer); have != "Pinned" {
		t.Errorf("pinned: have: %s, want: Pinned", have)
	}
	if have := gn.KeyName(element); have != "PinnedReasons" {
		t.Errorf("array element: have: %s, want: PinnedReasons", have)
	}
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
			structKey("apns"),
		}}
		s2 := &schema.Schema{Title: "SchemaY", PayloadKeys: []*schema.PayloadKey{
			structKey("APNs"),
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

func TestGlobalNamerCollisions(t *testing.T) {
	t.Run("a fixed name displaces an unqualified key", func(t *testing.T) {
		// SchemaW is fully qualified, so its Widget takes SchemaWWidget. SchemaX
		// happens to declare a key of that exact name; since fixed names are
		// reserved first, the unqualified key yields and takes a longer prefix.
		s1 := &schema.Schema{Title: "SchemaW", PayloadKeys: []*schema.PayloadKey{structKey("Widget")}}
		s2 := &schema.Schema{Title: "SchemaX", PayloadKeys: []*schema.PayloadKey{structKey("SchemaWWidget")}}

		file := schema.NewFile([]*schema.Schema{s1, s2})
		gn := schema.NewGlobalNamer(file, nil, schema.WithFullyQualifiedNames(s1))

		if have := gn.KeyName(s1.PayloadKeys[0]); have != "SchemaWWidget" {
			t.Errorf("qualified key: have: %s, want: SchemaWWidget", have)
		}
		if have := gn.KeyName(s2.PayloadKeys[0]); have != "SchemaXSchemaWWidget" {
			t.Errorf("displaced key: have: %s, want: SchemaXSchemaWWidget", have)
		}
	})

	t.Run("two fixed names falls back to an index", func(t *testing.T) {
		// Nothing in the status tree should reach this -- names come from unique
		// dotted paths -- but a duplicate must not silently declare one Go type
		// twice, so the second gets an index.
		s := &schema.Schema{Title: "SchemaY", PayloadKeys: []*schema.PayloadKey{
			structKey("Dup"), structKey("Other"),
		}}
		file := schema.NewFile([]*schema.Schema{s})

		gn := schema.NewGlobalNamer(file, nil, schema.WithNameOverrides(map[*schema.PayloadKey]string{
			s.PayloadKeys[0]: "Same",
			s.PayloadKeys[1]: "Same",
		}))

		first, second := gn.KeyName(s.PayloadKeys[0]), gn.KeyName(s.PayloadKeys[1])
		if first == second {
			t.Fatalf("both keys named %q", first)
		}
		if first != "Same" || second != "Same2" {
			t.Errorf("have: %s, %s; want: Same, Same2", first, second)
		}
	})

	t.Run("a collision created by replacements is caught", func(t *testing.T) {
		// Names are compared after -repl is applied, so two names that differ
		// raw but resolve to the same identifier still get disambiguated.
		reps := replace.Replacements{
			{Match: `^(.*)Mdm(.*)$`, Replacement: "${1}MDM${2}", Types: []string{"struct"}},
		}
		s1 := &schema.Schema{Title: "SchemaZ", PayloadKeys: []*schema.PayloadKey{structKey("Mdm")}}
		s2 := &schema.Schema{Title: "SchemaZZ", PayloadKeys: []*schema.PayloadKey{structKey("MDM")}}

		file := schema.NewFile([]*schema.Schema{s1, s2})
		gn := schema.NewGlobalNamer(file, reps)

		name1 := reps.Replace(gn.KeyName(s1.PayloadKeys[0]), replace.Struct)
		name2 := reps.Replace(gn.KeyName(s2.PayloadKeys[0]), replace.Struct)
		if name1 == name2 {
			t.Errorf("post-replacement names collide: both %q", name1)
		}
	})
}

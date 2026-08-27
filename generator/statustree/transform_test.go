package statustree_test

import (
	"strings"
	"testing"

	"github.com/korylprince/go-adm/generator/statustree"
	"github.com/korylprince/go-adm/schema"
	yamlschema "github.com/korylprince/go-adm/yamlschema"
)

// statusItem builds a schema shaped like one of Apple's status item files: a
// statusitemtype, and a single top level payload key named for it.
func statusItem(path string, key *schema.PayloadKey) *schema.Schema {
	key.Key = path
	return &schema.Schema{
		Title:       path,
		Payload:     &schema.Payload{StatusItemType: path},
		PayloadKeys: []*schema.PayloadKey{key},
	}
}

func stringItem(path string) *schema.Schema {
	return statusItem(path, &schema.PayloadKey{
		Type:     schema.PayloadKeyTypeString,
		Presence: schema.PayloadKeyPresenceRequired,
		Content:  "The " + path + " value.",
	})
}

// envelope mirrors declarative/protocol/statusreport.yaml: a StatusItems
// dictionary with no subkeys of its own, alongside the rest of the report.
func envelope() *schema.Schema {
	return &schema.Schema{
		Title:   "Status Report",
		Payload: &schema.Payload{RequestType: "StatusReport"},
		PayloadKeys: []*schema.PayloadKey{
			{Key: "StatusItems", Type: schema.PayloadKeyTypeDictionary, Presence: schema.PayloadKeyPresenceRequired},
			{Key: "FullReport", Type: schema.PayloadKeyTypeBoolean, Presence: schema.PayloadKeyPresenceOptional},
		},
	}
}

// find returns the subkey with the given key name.
func find(key *schema.PayloadKey, name string) *schema.PayloadKey {
	for _, sub := range key.SubKeys {
		if sub.Key == name {
			return sub
		}
	}
	return nil
}

// walk returns the key at the given path of key names, failing the test if any
// segment is missing.
func walk(t *testing.T, keys []*schema.PayloadKey, path ...string) *schema.PayloadKey {
	t.Helper()

	cur := &schema.PayloadKey{SubKeys: keys}
	for i, name := range path {
		if cur = find(cur, name); cur == nil {
			t.Fatalf("no key at %s", strings.Join(path[:i+1], "."))
		}
	}
	return cur
}

func TestTransform(t *testing.T) {
	items := []*schema.Schema{
		stringItem("device.model.family"),
		stringItem("device.operating-system.build-version"),
		stringItem("device.operating-system.supplemental.build-version"),
		stringItem("passcode.is-present"),
	}

	tree, err := statustree.Transform(items, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("dotted paths become nested keys", func(t *testing.T) {
		key := walk(t, tree.Schema.PayloadKeys, "device", "model", "family")
		if key.Type != schema.PayloadKeyTypeString {
			t.Errorf("leaf type = %s, want %s", key.Type, schema.PayloadKeyTypeString)
		}

		// an interior node that is also a leaf's parent under a longer path:
		// build-version and supplemental are siblings
		os := walk(t, tree.Schema.PayloadKeys, "device", "operating-system")
		if find(os, "build-version") == nil || find(os, "supplemental") == nil {
			t.Error("device.operating-system should hold both build-version and supplemental")
		}
		walk(t, tree.Schema.PayloadKeys, "device", "operating-system", "supplemental", "build-version")
	})

	t.Run("interior nodes are optional dictionaries", func(t *testing.T) {
		key := walk(t, tree.Schema.PayloadKeys, "device", "model")
		if key.Type != schema.PayloadKeyTypeDictionary {
			t.Errorf("interior type = %s, want %s", key.Type, schema.PayloadKeyTypeDictionary)
		}
		if key.Presence != schema.PayloadKeyPresenceOptional {
			t.Errorf("interior presence = %s, want %s", key.Presence, schema.PayloadKeyPresenceOptional)
		}
	})

	t.Run("required leaves are relaxed to optional", func(t *testing.T) {
		// Apple marks status items required, meaning "required within this
		// item", but reports are incremental so any item may be absent.
		key := walk(t, tree.Schema.PayloadKeys, "device", "model", "family")
		if key.Presence != schema.PayloadKeyPresenceOptional {
			t.Errorf("leaf presence = %s, want %s", key.Presence, schema.PayloadKeyPresenceOptional)
		}
	})

	t.Run("leaf doc records the dotted identifier", func(t *testing.T) {
		key := walk(t, tree.Schema.PayloadKeys, "device", "model", "family")
		if want := "Status item: `device.model.family`."; !strings.Contains(key.Content, want) {
			t.Errorf("leaf content %q does not mention %q", key.Content, want)
		}
		if want := "The device.model.family value."; !strings.Contains(key.Content, want) {
			t.Errorf("leaf content %q dropped Apple's own description", key.Content)
		}
	})

	t.Run("source schemas are left untouched", func(t *testing.T) {
		// The transform clones leaves precisely so the parsed schemas stay
		// usable elsewhere in the same process.
		key := items[0].PayloadKeys[0]
		if key.Key != "device.model.family" {
			t.Errorf("source key was rewritten to %q", key.Key)
		}
		if key.Presence != schema.PayloadKeyPresenceRequired {
			t.Errorf("source presence was relaxed to %s", key.Presence)
		}
	})

	t.Run("interior names are pinned to the full path", func(t *testing.T) {
		want := map[string]string{
			"device":                               "StatusItemsDevice",
			"device.model":                         "StatusItemsDeviceModel",
			"device.operating-system":              "StatusItemsDeviceOperatingSystem",
			"device.operating-system.supplemental": "StatusItemsDeviceOperatingSystemSupplemental",
			"passcode":                             "StatusItemsPasscode",
		}

		got := make(map[string]string)
		for key, name := range tree.NameOverrides {
			got[key.Key] = name
		}

		for path, name := range want {
			segs := strings.Split(path, ".")
			last := segs[len(segs)-1]
			if got[last] != name {
				t.Errorf("name for %s = %q, want %q", path, got[last], name)
			}
		}

		// a scalar item has no value type of its own, so nothing to pin
		leaf := walk(t, tree.Schema.PayloadKeys, "device", "model", "family")
		if name, ok := tree.NameOverrides[leaf]; ok {
			t.Errorf("scalar leaf should not be pinned, got %q", name)
		}
	})
}

// A status item's value type is pinned from the dotted path too. Left to the
// namer these take the shortest unique suffix -- InstallReason, BatteryHealth --
// which the flat structs never suffered because each item was a schema of its
// own whose title qualified everything under it.
func TestTransformPinsLeafValueTypes(t *testing.T) {
	dict := statusItem("softwareupdate.install-reason", &schema.PayloadKey{
		Type:     schema.PayloadKeyTypeDictionary,
		Presence: schema.PayloadKeyPresenceRequired,
		SubKeys: []*schema.PayloadKey{
			{Key: "reason", Type: schema.PayloadKeyTypeString, Presence: schema.PayloadKeyPresenceRequired},
		},
	})

	array := statusItem("security.certificate.list", &schema.PayloadKey{
		Type:     schema.PayloadKeyTypeArray,
		Presence: schema.PayloadKeyPresenceRequired,
		SubKeys: []*schema.PayloadKey{
			{Key: "_list", Type: schema.PayloadKeyTypeDictionary, SubKeys: []*schema.PayloadKey{
				{Key: "identifier", Type: schema.PayloadKeyTypeString, Presence: schema.PayloadKeyPresenceRequired},
			}},
		},
	})

	// a non-empty rangelist is what makes a key an enum
	enum := statusItem("device.power.battery-health", &schema.PayloadKey{
		Type:     schema.PayloadKeyTypeString,
		Presence: schema.PayloadKeyPresenceRequired,
		Rangelist: []yamlschema.IntegerNumberString{
			*yamlschema.NewIntegerNumberString("normal"),
			*yamlschema.NewIntegerNumberString("unknown"),
		},
	})

	scalar := stringItem("device.model.family")

	tree, err := statustree.Transform([]*schema.Schema{dict, array, enum, scalar}, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("dictionary item pins its own key", func(t *testing.T) {
		key := walk(t, tree.Schema.PayloadKeys, "softwareupdate", "install-reason")
		if got := tree.NameOverrides[key]; got != "SoftwareupdateInstallReason" {
			t.Errorf("name = %q, want SoftwareupdateInstallReason", got)
		}
	})

	t.Run("array item pins its element type", func(t *testing.T) {
		// the array itself generates no type; the element dictionary does
		key := walk(t, tree.Schema.PayloadKeys, "security", "certificate", "list")
		if _, ok := tree.NameOverrides[key]; ok {
			t.Error("the array key should not be pinned")
		}
		if got := tree.NameOverrides[key.SubKeys[0]]; got != "SecurityCertificateList" {
			t.Errorf("element name = %q, want SecurityCertificateList", got)
		}
	})

	t.Run("enum item pins its own key", func(t *testing.T) {
		key := walk(t, tree.Schema.PayloadKeys, "device", "power", "battery-health")
		if got := tree.NameOverrides[key]; got != "DevicePowerBatteryHealth" {
			t.Errorf("name = %q, want DevicePowerBatteryHealth", got)
		}
	})

	t.Run("scalar item generates no type to pin", func(t *testing.T) {
		key := walk(t, tree.Schema.PayloadKeys, "device", "model", "family")
		if got, ok := tree.NameOverrides[key]; ok {
			t.Errorf("scalar should not be pinned, got %q", got)
		}
	})
}

func TestTransformDeterministic(t *testing.T) {
	// Generated output has to be stable run to run, so the tree is sorted
	// rather than left in whatever order the schema files were walked.
	forward := []*schema.Schema{
		stringItem("device.model.family"),
		stringItem("device.model.identifier"),
		stringItem("app.managed.list"),
	}
	reverse := []*schema.Schema{
		stringItem("app.managed.list"),
		stringItem("device.model.identifier"),
		stringItem("device.model.family"),
	}

	keyNames := func(schemas []*schema.Schema) []string {
		tree, err := statustree.Transform(schemas, nil)
		if err != nil {
			t.Fatal(err)
		}
		var names []string
		var visit func(keys []*schema.PayloadKey, prefix string)
		visit = func(keys []*schema.PayloadKey, prefix string) {
			for _, key := range keys {
				names = append(names, prefix+key.Key)
				visit(key.SubKeys, prefix+key.Key+".")
			}
		}
		visit(tree.Schema.PayloadKeys, "")
		return names
	}

	a, b := keyNames(forward), keyNames(reverse)
	if strings.Join(a, ",") != strings.Join(b, ",") {
		t.Errorf("key order depends on input order:\n%v\n%v", a, b)
	}
}

func TestTransformEnvelope(t *testing.T) {
	env := envelope()
	items := []*schema.Schema{stringItem("device.model.family")}

	tree, err := statustree.Transform(items, env)
	if err != nil {
		t.Fatal(err)
	}

	if tree.Schema.Title != env.Title {
		t.Errorf("title = %q, want %q", tree.Schema.Title, env.Title)
	}

	statusItems := find(&schema.PayloadKey{SubKeys: tree.Schema.PayloadKeys}, "StatusItems")
	if statusItems == nil {
		t.Fatal("grafted schema has no StatusItems key")
	}
	walk(t, statusItems.SubKeys, "device", "model", "family")

	if find(&schema.PayloadKey{SubKeys: tree.Schema.PayloadKeys}, "FullReport") == nil {
		t.Error("grafted schema dropped the rest of the envelope")
	}

	t.Run("envelope is deep copied", func(t *testing.T) {
		// declgen still generates protocol.StatusReport from the original, and
		// sharing PayloadKeys across two schemas corrupts the encoder's naming.
		if len(env.PayloadKeys[0].SubKeys) != 0 {
			t.Error("grafting mutated the source envelope")
		}
		for _, key := range tree.Schema.PayloadKeys {
			for _, orig := range env.PayloadKeys {
				if key == orig {
					t.Errorf("grafted schema shares PayloadKey %q with the source envelope", key.Key)
				}
			}
		}
	})
}

func TestTransformErrors(t *testing.T) {
	tests := []struct {
		name  string
		items []*schema.Schema
		want  string
	}{
		{
			name:  "no items",
			items: nil,
			want:  "no status item schemas",
		},
		{
			name:  "duplicate item",
			items: []*schema.Schema{stringItem("device.model.family"), stringItem("device.model.family")},
			want:  "duplicate status item",
		},
		{
			name:  "existing item is a prefix of a new one",
			items: []*schema.Schema{stringItem("device.model"), stringItem("device.model.family")},
			want:  "collides",
		},
		{
			name:  "new item is a prefix of an existing one",
			items: []*schema.Schema{stringItem("device.model.family"), stringItem("device.model")},
			want:  "collides",
		},
		{
			name: "more than one top level payload key",
			items: []*schema.Schema{{
				Title:   "device.model.family",
				Payload: &schema.Payload{StatusItemType: "device.model.family"},
				PayloadKeys: []*schema.PayloadKey{
					{Key: "device.model.family", Type: schema.PayloadKeyTypeString},
					{Key: "extra", Type: schema.PayloadKeyTypeString},
				},
			}},
			want: "expected exactly 1 top level payload key",
		},
		{
			name: "payload key does not match the statusitemtype",
			items: []*schema.Schema{{
				Title:       "device.model.family",
				Payload:     &schema.Payload{StatusItemType: "device.model.family"},
				PayloadKeys: []*schema.PayloadKey{{Key: "family", Type: schema.PayloadKeyTypeString}},
			}},
			want: "expected it to match the statusitemtype",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := statustree.Transform(test.items, nil)
			if err == nil {
				t.Fatalf("expected an error mentioning %q", test.want)
			}
			if !strings.Contains(err.Error(), test.want) {
				t.Errorf("error = %q, want it to mention %q", err, test.want)
			}
		})
	}
}

func TestTransformEnvelopeErrors(t *testing.T) {
	items := []*schema.Schema{stringItem("device.model.family")}

	t.Run("envelope has no StatusItems key", func(t *testing.T) {
		env := envelope()
		env.PayloadKeys = env.PayloadKeys[1:]
		if _, err := statustree.Transform(items, env); err == nil {
			t.Fatal("expected an error")
		}
	})

	t.Run("StatusItems already declares subkeys", func(t *testing.T) {
		// Apple leaving this dictionary shapeless is the only reason there is
		// room to graft into it; if that changes we would be discarding keys.
		env := envelope()
		env.PayloadKeys[0].SubKeys = []*schema.PayloadKey{{Key: "device", Type: schema.PayloadKeyTypeString}}
		if _, err := statustree.Transform(items, env); err == nil {
			t.Fatal("expected an error")
		}
	})
}

func TestTransformItemTypes(t *testing.T) {
	tree, err := statustree.Transform([]*schema.Schema{
		stringItem("passcode.is-present"),
		stringItem("device.model.family"),
		stringItem("device.identifier.udid"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	// sorted, because the trie is walked in sorted order
	want := []string{"device.identifier.udid", "device.model.family", "passcode.is-present"}
	if strings.Join(tree.ItemTypes, ",") != strings.Join(want, ",") {
		t.Errorf("ItemTypes = %v, want %v", tree.ItemTypes, want)
	}
}

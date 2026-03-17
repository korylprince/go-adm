package schema_test

import (
	"strings"
	"testing"

	"github.com/korylprince/go-adm/schema"
)

// Test recursive struct resolution: an array whose SubKeys contains a struct
// that circularly references the same slice (via shared pointer).
func TestRecursiveArrayRef(t *testing.T) {
	// Simulates com.apple.applicationaccess.new.yaml after go-yaml resolves
	// the recursive alias: subApps.SubKeys points back to the same slice.
	whiteListItem := &schema.PayloadKey{
		Key:  "whiteListItem",
		Type: schema.PayloadKeyTypeDictionary,
	}
	subKeys := []*schema.PayloadKey{whiteListItem}
	whiteListItem.SubKeys = []*schema.PayloadKey{
		{Key: "bundleID", Type: schema.PayloadKeyTypeString, Presence: schema.PayloadKeyPresenceRequired},
		{Key: "appID", Type: schema.PayloadKeyTypeData, Presence: schema.PayloadKeyPresenceRequired},
		{
			Key:        "subApps",
			Type:       schema.PayloadKeyTypeArray,
			Presence:   schema.PayloadKeyPresenceOptional,
			SubKeyType: "ApplicationItem",
			SubKeys:    subKeys, // recursive: same slice as parent
		},
	}

	s := anyTestSchema("RecursiveTest",
		&schema.PayloadKey{
			Key:        "whiteList",
			Type:       schema.PayloadKeyTypeArray,
			Presence:   schema.PayloadKeyPresenceOptional,
			SubKeyType: "ApplicationItem",
			SubKeys:    subKeys,
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	if strings.Contains(src, "SubApps any") || strings.Contains(src, "SubApps *any") {
		t.Errorf("recursive array should not generate 'any'; got:\n%s", src)
	}
	if !strings.Contains(src, "*[]*WhiteListItem") || !strings.Contains(src, "SubApps") {
		t.Errorf("expected SubApps field with *[]*WhiteListItem type; got:\n%s", src)
	}
}

// Test nested recursive arrays (array-of-array pattern).
// Simulates com.apple.homescreenlayout.yaml: Dock -> [IconItem{Pages -> [[IconItem]]}]
func TestRecursiveNestedArrayRef(t *testing.T) {
	iconItem := &schema.PayloadKey{
		Key:  "IconItem",
		Type: schema.PayloadKeyTypeDictionary,
	}
	iconSubKeys := []*schema.PayloadKey{iconItem}

	pagesItem := &schema.PayloadKey{
		Key:        "PagesItem",
		Type:       schema.PayloadKeyTypeArray,
		SubKeyType: "IconItem",
		SubKeys:    iconSubKeys, // recursive: resolves back to iconItem
	}

	iconItem.SubKeys = []*schema.PayloadKey{
		{Key: "Type", Type: schema.PayloadKeyTypeString, Presence: schema.PayloadKeyPresenceRequired},
		{Key: "DisplayName", Type: schema.PayloadKeyTypeString, Presence: schema.PayloadKeyPresenceOptional},
		{
			Key:        "Pages",
			Type:       schema.PayloadKeyTypeArray,
			Presence:   schema.PayloadKeyPresenceOptional,
			SubKeyType: "PagesItem",
			SubKeys:    []*schema.PayloadKey{pagesItem},
		},
	}

	s := anyTestSchema("NestedRecursiveTest",
		&schema.PayloadKey{
			Key:        "Dock",
			Type:       schema.PayloadKeyTypeArray,
			Presence:   schema.PayloadKeyPresenceOptional,
			SubKeyType: "IconItem",
			SubKeys:    iconSubKeys,
		},
		&schema.PayloadKey{
			Key:        "TopPages",
			Type:       schema.PayloadKeyTypeArray,
			Presence:   schema.PayloadKeyPresenceRequired,
			SubKeyType: "PagesItem",
			SubKeys:    []*schema.PayloadKey{pagesItem},
		},
	)

	src := renderFile(t, []*schema.Schema{s})

	if strings.Contains(src, "[]any") {
		t.Errorf("nested recursive array should not generate '[]any'; got:\n%s", src)
	}
	if !strings.Contains(src, "TopPages [][]*IconItem") {
		t.Errorf("expected TopPages [][]*IconItem; got:\n%s", src)
	}
}

// TestRecursiveYAMLParse verifies that the go-yaml fork correctly resolves
// recursive slice aliases, producing populated SubKeys instead of nil.
func TestRecursiveYAMLParse(t *testing.T) {
	const yamlSrc = `
title: RecursiveYAMLTest
payloadkeys:
- key: whiteList
  type: <array>
  subkeytype: ApplicationItem
  subkeys: &id001
  - key: whiteListItem
    type: <dictionary>
    subkeys:
    - key: bundleID
      type: <string>
      presence: required
    - key: subApps
      type: <array>
      subkeytype: ApplicationItem
      subkeys: *id001
`

	s, err := schema.New([]byte(yamlSrc))
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	whiteList := s.PayloadKeys[0]
	if len(whiteList.SubKeys) != 1 {
		t.Fatalf("expected 1 subkey for whiteList, got %d", len(whiteList.SubKeys))
	}

	whiteListItem := whiteList.SubKeys[0]
	if len(whiteListItem.SubKeys) != 2 {
		t.Fatalf("expected 2 subkeys for whiteListItem, got %d", len(whiteListItem.SubKeys))
	}

	subApps := whiteListItem.SubKeys[1]
	if subApps.Key != "subApps" {
		t.Fatalf("expected subApps key, got %s", subApps.Key)
	}
	if len(subApps.SubKeys) != 1 {
		t.Fatalf("expected subApps to have 1 recursive subkey, got %d", len(subApps.SubKeys))
	}
	if subApps.SubKeys[0] != whiteListItem {
		t.Error("expected recursive subkey pointer identity")
	}

	// Verify code generation works with the YAML-parsed recursive schema
	src := renderFile(t, []*schema.Schema{s})

	if strings.Contains(src, "SubApps *any") || strings.Contains(src, "SubApps any") {
		t.Errorf("recursive array should not generate 'any'; got:\n%s", src)
	}
	if !strings.Contains(src, "*[]*WhiteListItem") {
		t.Errorf("expected *[]*WhiteListItem type; got:\n%s", src)
	}
}

// Package statustree decomposes Apple's dotted DDM status item namespace into
// the nested schema that status reports actually use on the wire.
//
// Apple declares each status item in its own file under declarative/status/,
// keyed by a dotted statusitemtype:
//
//	payload:
//	  statusitemtype: device.model.family
//	payloadkeys:
//	- key: device.model.family
//	  type: <string>
//
// Taken literally that generates a struct with a `json:"device.model.family"`
// tag, which matches nothing: a device nests the report instead.
//
//	{"StatusItems": {"device": {"model": {"family": "Mac"}}}}
//
// Transform closes that gap purely as a schema AST rewrite -- split the dotted
// names into a trie, synthesize a dictionary key for each interior node, and
// hand the result to the ordinary schema encoder. Nothing about rendering Go is
// reimplemented here: dictionaries, arrays, enums, recursive YAML anchors,
// naming and struct tags all stay with schema.Encoder, so status types come out
// consistent with every other type go-adm generates.
package statustree

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/utils/text"
)

// StatusItemsKey is the key the status report envelope holds status items under.
const StatusItemsKey = "StatusItems"

// statusReportRequestType identifies the status report envelope schema
// (declarative/protocol/statusreport.yaml).
const statusReportRequestType = "StatusReport"

const statusItemsDoc = "The status items for a status report, decomposed from Apple's dotted status item namespace."

// Tree is a synthesized status item schema, ready to hand to schema.NewFile.
type Tree struct {
	// Schema is the synthesized schema. It replaces the status item schemas it
	// was built from, which must not also be registered in the same
	// schema.File -- see the note on leafKey.
	Schema *schema.Schema

	// NameOverrides pins a stable, path-derived name on each interior namespace
	// node and on each status item's value type. Pass it to
	// schema.WithNameOverrides.
	NameOverrides map[*schema.PayloadKey]string

	// ItemTypes is every status item's dotted identifier, sorted.
	ItemTypes []string
}

// IsStatusItem reports whether the schema describes a single DDM status item.
func IsStatusItem(s *schema.Schema) bool {
	return s.Payload != nil && s.Payload.StatusItemType != ""
}

// IsStatusReport reports whether the schema is the status report envelope, the
// document status items are delivered inside of.
func IsStatusReport(s *schema.Schema) bool {
	return s.Payload != nil && s.Payload.RequestType == statusReportRequestType
}

// Transform decomposes the given status item schemas into a single nested
// schema.
//
// If envelope is nil the result is a standalone StatusItems schema. Given the
// status report envelope, the tree is instead grafted into its StatusItems key,
// so the result is a whole typed report -- status items alongside Errors and
// FullReport. The envelope is deep copied, so the caller's schema is left
// untouched and can still be generated separately.
func Transform(items []*schema.Schema, envelope *schema.Schema) (*Tree, error) {
	if len(items) == 0 {
		return nil, errors.New("no status item schemas to transform")
	}

	root, err := buildTrie(items)
	if err != nil {
		return nil, err
	}

	t := &Tree{NameOverrides: make(map[*schema.PayloadKey]string)}

	roots := make([]*schema.PayloadKey, 0, len(root.children))
	for _, seg := range root.childSegs() {
		roots = append(roots, t.convert(root.children[seg], []string{seg}))
	}

	if envelope == nil {
		t.Schema = &schema.Schema{
			Title:       StatusItemsKey,
			Description: statusItemsDoc,
			PayloadKeys: roots,
		}
		return t, nil
	}

	clone := cloneSchema(envelope)

	idx := slices.IndexFunc(clone.PayloadKeys, func(key *schema.PayloadKey) bool {
		return key.Key == StatusItemsKey
	})
	if idx < 0 {
		return nil, fmt.Errorf("status report envelope %q has no %q payload key", envelope.Title, StatusItemsKey)
	}

	key := clone.PayloadKeys[idx]
	if key.Type != schema.PayloadKeyTypeDictionary {
		return nil, fmt.Errorf("status report envelope %q: %q is %s, expected %s", envelope.Title, StatusItemsKey, key.Type, schema.PayloadKeyTypeDictionary)
	}
	// Apple leaves this dictionary shapeless, which is the only reason there is
	// room to graft the tree in. If it ever gains subkeys of its own we would be
	// silently discarding them.
	if len(key.SubKeys) > 0 {
		return nil, fmt.Errorf("status report envelope %q: %q already declares %d subkeys", envelope.Title, StatusItemsKey, len(key.SubKeys))
	}

	key.SubKeys = roots
	t.NameOverrides[key] = StatusItemsKey
	t.Schema = clone

	return t, nil
}

// convert turns one trie node into a synthesized PayloadKey. path is the node's
// dotted status item path, used for doc comments and for pinning names.
func (t *Tree) convert(n *node, path []string) *schema.PayloadKey {
	if n.isLeaf() {
		t.ItemTypes = append(t.ItemTypes, strings.Join(path, "."))

		key := leafKey(n, path)
		t.pinLeafValue(key, path)
		return key
	}

	key := &schema.PayloadKey{
		Key:      n.seg,
		Type:     schema.PayloadKeyTypeDictionary,
		Presence: schema.PayloadKeyPresenceOptional,
		Content:  fmt.Sprintf("Status items under the `%s` namespace.", strings.Join(path, ".")),
	}

	for _, seg := range n.childSegs() {
		key.SubKeys = append(key.SubKeys, t.convert(n.children[seg], append(slices.Clone(path), seg)))
	}

	// Pin interior names to the full path. Left to the GlobalNamer these take
	// the shortest unique suffix, which for namespace nodes means exported
	// types called List, Model, Report -- and worst, Type. Beyond reading badly
	// as public API, that suffix is computed against every other type in the
	// package, so adding a single status item upstream can silently rename an
	// unrelated one. Deriving the name from the path alone makes it stable.
	t.NameOverrides[key] = StatusItemsKey + qualify(path)

	return key
}

// pinLeafValue pins the name of the type generated for a status item's *value*,
// where the item declares one.
//
// This matters as much as pinning the namespace nodes, and for the same reason.
// The value type's qualified path is only as long as the tree above it, so the
// shortest unique suffix strips it back to names like InstallReason,
// InstallState, or BatteryHealth. In a namespace as large and varied as Apple's
// those read poorly and, worse, shift as items are added: the flat structs got
// SoftwareUpdateInstallReason for free, because each item was a schema of its
// own whose title qualified everything beneath it. Naming from the dotted path
// restores that and makes it stable.
//
// Which key becomes the type depends on the item's shape. A dictionary or a
// rangelist enum is generated from the item's own key; an array of dictionaries
// is generated from its element type instead. Scalar items, and free-form
// dictionaries that render as an inline map, generate no named type at all.
func (t *Tree) pinLeafValue(key *schema.PayloadKey, path []string) {
	name := qualify(path)

	switch {
	case key.Type == schema.PayloadKeyTypeArray && len(key.SubKeys) > 0 && key.SubKeys[0].Type == schema.PayloadKeyTypeDictionary:
		// The element carries the generated type, so it wants the dotted
		// identifier in its doc comment as much as the item's own key does.
		// Clone it before writing: the leaf clone is shallow, so its SubKeys
		// still point into Apple's parsed schema.
		element := *key.SubKeys[0]
		element.Content = withItemPath(element.Content, path)
		key.SubKeys = append([]*schema.PayloadKey{&element}, key.SubKeys[1:]...)

		t.NameOverrides[&element] = name
	case key.IsStruct(), key.IsEnum():
		// IsEnum also covers an array with a rangelist subkey, which the
		// encoder names from the array key rather than the element.
		t.NameOverrides[key] = name
	}
}

// withItemPath appends the status item's dotted identifier to a doc comment.
//
// The identifier is worth carrying into the generated code because it is what
// appears on the wire: in a status subscription's item list, in a report's
// Errors[].StatusItem, and in the client's supported-payloads. Nothing else in
// the nested tree records it.
func withItemPath(doc string, path []string) string {
	return strings.TrimSpace(fmt.Sprintf("%s\nStatus item: `%s`.", strings.TrimSpace(doc), strings.Join(path, ".")))
}

// leafKey clones a status item's own top level PayloadKey down to the single
// segment it occupies in the tree.
//
// The clone is deliberately shallow. SubKeys keep pointing into the status
// item's original schema, which is what makes leaf *value* types come out
// exactly as Apple declares them. Two consequences follow, and both matter:
//
//   - The original key must be cloned rather than rewritten in place, or every
//     other consumer of these parsed schemas sees a mangled key.
//   - The original schema must be dropped from the schema.File that receives
//     the synthesized one. schema.NewFile dedupes types with a `seen` set that
//     resets per schema, and GlobalNamer keys its name table by PayloadKey
//     identity, so a key reachable from two schemas in one File is emitted
//     twice and named wrongly.
func leafKey(n *node, path []string) *schema.PayloadKey {
	key := *n.leaf
	key.Key = n.seg

	// Apple marks a status item's own key `required`, which means "required
	// within this item when the item is reported", not "present in every
	// report". Reports are incremental and may carry any subset of items, so
	// the item itself has to be relaxed to optional. Presence *inside* the
	// value -- a certificate's identifier, say -- is left exactly as declared.
	key.Presence = schema.PayloadKeyPresenceOptional

	doc := key.Content
	if strings.TrimSpace(doc) == "" {
		doc = n.leafSchema.Description
	}
	key.Content = withItemPath(doc, path)

	return &key
}

// qualify joins a dotted path into the Go name fragment for it, normalizing
// each segment on its own.
//
// Per-segment is the whole point: text.NormalizeName splits on [_\-:] but not
// on ".", so normalizing "device.model.family" whole yields Devicemodelfamily.
// Splitting first gives DeviceModelFamily, and is why the tree needs far fewer
// hand-written replacement rules than the flat status structs did.
func qualify(path []string) string {
	var name strings.Builder
	for _, seg := range path {
		name.WriteString(text.NormalizeName(seg))
	}
	return name.String()
}

// cloneSchema deep copies a schema's payload keys so the caller can rewrite them
// without disturbing the parsed original.
//
// Needed because the envelope schema stays in play elsewhere: declgen still
// generates protocol.StatusReport from it. Sharing PayloadKey pointers between
// two schemas breaks the encoder, for the reasons spelled out on leafKey.
//
// The seen map does more than stop infinite recursion. Apple's YAML uses
// anchors to reference one subkey object from several places, and expanding
// those into independent copies would generate a separate Go type per use, so
// aliasing has to survive the copy.
func cloneSchema(s *schema.Schema) *schema.Schema {
	clone := *s
	seen := make(map[*schema.PayloadKey]*schema.PayloadKey)
	clone.PayloadKeys = cloneKeys(s.PayloadKeys, seen)
	clone.ResponseKeys = cloneKeys(s.ResponseKeys, seen)
	return &clone
}

func cloneKeys(keys []*schema.PayloadKey, seen map[*schema.PayloadKey]*schema.PayloadKey) []*schema.PayloadKey {
	if keys == nil {
		return nil
	}
	clones := make([]*schema.PayloadKey, 0, len(keys))
	for _, key := range keys {
		clones = append(clones, cloneKey(key, seen))
	}
	return clones
}

func cloneKey(key *schema.PayloadKey, seen map[*schema.PayloadKey]*schema.PayloadKey) *schema.PayloadKey {
	if clone, ok := seen[key]; ok {
		return clone
	}

	clone := new(schema.PayloadKey)
	// register before recursing, so a key reachable from its own subtree
	// resolves to this clone instead of recursing forever
	seen[key] = clone

	*clone = *key
	clone.SubKeys = cloneKeys(key.SubKeys, seen)

	return clone
}

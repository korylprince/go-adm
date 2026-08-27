package statustree

import (
	"fmt"
	"slices"
	"strings"

	"github.com/korylprince/go-adm/schema"
)

// node is one segment of the dotted status item namespace.
//
// Every status item is a leaf: no item path is a prefix of another across
// Apple's status schemas, so a node is either a namespace with children or a
// status item with a value, never both. `device.operating-system.build-version`
// and `device.operating-system.supplemental.build-version` coexist because
// `supplemental` is purely interior.
type node struct {
	seg      string
	children map[string]*node

	// leaf is the status item's own top level PayloadKey and leafSchema the
	// file it was parsed from. Both are nil on interior nodes.
	leaf       *schema.PayloadKey
	leafSchema *schema.Schema
}

func newNode(seg string) *node {
	return &node{seg: seg, children: make(map[string]*node)}
}

func (n *node) isLeaf() bool {
	return n.leaf != nil
}

// child returns the named child, creating it if this is the first time the
// segment has been seen.
func (n *node) child(seg string) *node {
	c, ok := n.children[seg]
	if !ok {
		c = newNode(seg)
		n.children[seg] = c
	}
	return c
}

// childSegs returns the child segments in sorted order. Insertion order follows
// whatever order the schema files were walked in, so everything that renders
// the trie sorts first -- generated output has to be identical run to run.
func (n *node) childSegs() []string {
	segs := make([]string, 0, len(n.children))
	for seg := range n.children {
		segs = append(segs, seg)
	}
	slices.Sort(segs)
	return segs
}

// buildTrie inserts every status item schema into a trie keyed on its dotted
// statusitemtype.
//
// Each schema is expected to carry exactly one top level payload key, named for
// the statusitemtype itself. Every check here is fatal rather than skip-and-
// continue: the whole transform rests on these invariants holding, so if Apple
// ever breaks one, failing loudly beats emitting a tree that quietly misses or
// misplaces status items.
func buildTrie(items []*schema.Schema) (*node, error) {
	root := newNode("")

	for _, s := range items {
		path := s.Payload.StatusItemType

		if len(s.PayloadKeys) != 1 {
			return nil, fmt.Errorf("status item %q: expected exactly 1 top level payload key, got %d", path, len(s.PayloadKeys))
		}
		pk := s.PayloadKeys[0]
		if pk.Key != path {
			return nil, fmt.Errorf("status item %q: top level payload key is %q, expected it to match the statusitemtype", path, pk.Key)
		}

		segs := strings.Split(path, ".")

		cur := root
		for i, seg := range segs[:len(segs)-1] {
			cur = cur.child(seg)
			// Descending through a node that already holds a value means some
			// earlier item's path is a prefix of this one, so one of the two
			// would have to be both a namespace and a value.
			if cur.isLeaf() {
				return nil, fmt.Errorf("status item %q collides with %q, which is a prefix of it", path, strings.Join(segs[:i+1], "."))
			}
		}

		leaf := cur.child(segs[len(segs)-1])
		if leaf.isLeaf() {
			return nil, fmt.Errorf("duplicate status item %q", path)
		}
		// ...and the same collision seen from the other side: this item's path
		// is a prefix of one already inserted.
		if len(leaf.children) > 0 {
			return nil, fmt.Errorf("status item %q collides with %q, which it is a prefix of", path, path+"."+leaf.childSegs()[0])
		}

		leaf.leaf, leaf.leafSchema = pk, s
	}

	return root, nil
}

package schema

import (
	"fmt"

	"github.com/korylprince/go-adm/text"
	"golang.org/x/exp/slices"
)

// KeyANY is a special *PayloadKey.Key to represent a generic map[string]any instead of a struct
const KeyANY = "ANY"

type nameVal struct {
	key    *PayloadKey
	parent *PayloadKey
	schema *Schema
}

// GlobalNamer namespaces schema and key names the minimum amount to achieve globally unique names among the registered schemas
type GlobalNamer struct {
	names   map[string][]*nameVal
	keys    map[*PayloadKey]*nameVal
	schemas map[*Schema]*nameVal
}

func NewGlobalNamer() *GlobalNamer {
	return &GlobalNamer{
		names:   make(map[string][]*nameVal),
		keys:    make(map[*PayloadKey]*nameVal),
		schemas: make(map[*Schema]*nameVal),
	}
}

func (n *GlobalNamer) Register(schemas ...*Schema) {
	for _, s := range schemas {
		// don't register a schema twice
		if _, ok := n.schemas[s]; ok {
			continue
		}
		name := text.NormalizeName(s.Title)
		val := &nameVal{schema: s}
		n.names[name] = append(n.names[name], val)
		n.schemas[s] = val

		// add keys
		s.Iter(func(parent, key *PayloadKey) {
			if key.Type != PayloadKeyTypeDictionary {
				return
			}
			// don't register ANY <dictionary>s
			if len(key.SubKeys) > 0 && key.SubKeys[0].Type == KeyANY {
				return
			}
			name := text.NormalizeName(key.Key)
			val := &nameVal{key: key, parent: parent}
			// if parent is schema, add fake parent with name
			if val.parent == nil {
				val.parent = &PayloadKey{Key: s.Title}
			}
			n.names[name] = append(n.names[name], val)
			n.keys[key] = val
		})
	}
}

// distinct returns all items in strs are unique
func distinct(strs []string) bool {
	set := make(map[string]struct{})
	for _, s := range strs {
		set[s] = struct{}{}
	}

	return len(strs) == len(set)
}

// KeyName returns a unique name for the given key, amongst the registered schemas
func (n *GlobalNamer) KeyName(key *PayloadKey) string {
	name := text.NormalizeName(key.Key)
	vals := n.names[name]

	// not registered
	if len(vals) == 0 {
		return ""
	}

	// name is unique
	if len(vals) == 1 {
		return name
	}

	// if there is a name collision, recurse through parents until a distinct chain is found
	depth := 1
	for ; ; depth++ {
		var parents []string
	outer:
		for _, val := range vals {
			// get parent at depth
			node := val.key
			for p := depth; p > 0; p-- {
				// if we reach the root schema, don't include it in the distinct check
				// TODO: this may not be sufficient in all cases; we may need to be a bit smarter about this
				if node == nil {
					continue outer
				}
				node = n.keys[node].parent
			}
			parents = append(parents, text.NormalizeName(node.Key))
		}

		if distinct(parents) {
			break
		}
	}

	// use depth to form name
	node := key
	for p := depth; p > 0; p-- {
		node = n.keys[node].parent
		name = text.NormalizeName(node.Key) + name
	}

	return name
}

// SchemaName returns a unique name for the given schema, amongst the registered schemas
func (n *GlobalNamer) SchemaName(s *Schema) string {
	name := text.NormalizeName(s.Title)
	vals := n.names[name]

	// not registered
	if len(vals) == 0 {
		return ""
	}

	// name is unique
	if len(vals) == 1 {
		return name
	}

	count := 0
	for _, val := range vals {
		if val.schema != nil {
			count++
		}
	}

	// if this is the only schema with the conflicting name, it can take the name
	if count == 1 {
		return name
	}

	// we somehow ended up with multiple schemas with the same name. Append the index number to make the name unique. Yuck.
	schemas := make([]*Schema, len(vals))
	for idx, val := range vals {
		schemas[idx] = val.schema
	}

	return fmt.Sprintf("%s%d", name, slices.Index(schemas, s)+1)
}

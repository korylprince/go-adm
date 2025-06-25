package schema

import (
	"fmt"
	"strings"

	"github.com/korylprince/go-adm/text"
)

type payloadKey struct {
	schema  *Schema
	parents []*PayloadKey
	key     *PayloadKey
	root    bool
	name    string
}

func (key *payloadKey) fullyQualified() []string {
	var name []string
	if !key.root {
		name = []string{text.NormalizeName(key.schema.Title)}
	}
	for _, parent := range key.parents {
		name = append(name, text.NormalizeName(parent.Key))
	}
	name = append(name, text.NormalizeName(key.key.Key))
	return name
}

// GlobalNamer registers Schemas to determine globally unique names for all generated types
type GlobalNamer struct {
	keyNames map[*PayloadKey]*payloadKey
}

func NewGlobalNamer(file *File) *GlobalNamer {
	namer := &GlobalNamer{
		keyNames: make(map[*PayloadKey]*payloadKey),
	}

	// get all PayloadKeys needing globally unique names
	var keys []*payloadKey
	for _, typ := range file.Types {
		switch t := typ.(type) {
		case *Enum:
			keys = append(keys, &payloadKey{
				schema:  t.Schema,
				parents: t.Parents,
				key:     t.Key,
			})
		case *Struct:
			keys = append(keys, &payloadKey{
				schema:  t.Schema,
				parents: t.Parents,
				key:     t.Key,
				root:    t.Source != SourceSubKeys,
			})
		case *Map:
			keys = append(keys, &payloadKey{
				schema:  t.Schema,
				parents: t.Parents,
				key:     t.Key,
				root:    t.Source != SourceSubKeys,
			})
		}
	}

	// mark every possible key name
	names := make(map[string]int)
	for _, key := range keys {
		fq := key.fullyQualified()
		for start := len(fq) - 1; start >= 0; start-- {
			name := strings.Join(fq[start:], "")
			names[name] += 1
		}
	}

	// find minimal unique names
	count := make(map[string]int)
outer:
	for _, key := range keys {
		fq := key.fullyQualified()
		for start := len(fq) - 1; start >= 0; start-- {
			name := strings.Join(fq[start:], "")
			if names[name] == 1 {
				key.name = name
				namer.keyNames[key.key] = key
				continue outer
			}
		}
		// If we got here, the fully qualified identifier isn't unique
		// as a safeguard, return the name suffixed with an index
		// we shouldn't ever hit this in real life
		prefix := strings.Join(fq, "")
		count[prefix] += 1
		key.name = fmt.Sprintf("%s%d", prefix, count[prefix])
		namer.keyNames[key.key] = key
	}

	return namer
}

// KeyName returns a unique name for the given key, amongst the registered schemas
func (n *GlobalNamer) KeyName(pk *PayloadKey) string {
	key, ok := n.keyNames[pk]
	if !ok {
		panic(fmt.Sprintf("%#v PayloadKey is not registered. (Did you include its schema in the NewGlobalNamer call?",
			pk.Key,
		))
	}
	return key.name
}

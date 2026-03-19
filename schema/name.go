package schema

import (
	"fmt"
	"strings"

	"github.com/korylprince/go-adm/utils/replace"
	"github.com/korylprince/go-adm/utils/text"
)

type payloadKey struct {
	schema  *Schema
	parents []*PayloadKey
	key     *PayloadKey
	root    bool
	name    string
	repTyp  replace.ReplacementType
}

// normalizeANYKeyName returns a normalized name for a PayloadKey,
// using SubKeyType or the descriptor for ANY keys instead of the raw key name.
// Returns "" for plain "ANY" keys with no SubKeyType or descriptor, which
// causes the segment to be skipped in the fully qualified name.
func normalizeANYKeyName(key *PayloadKey) string {
	if !key.IsANY() {
		return text.NormalizeName(key.Key)
	}
	if key.SubKeyType != "" {
		return text.NormalizeName(key.SubKeyType)
	}
	if desc, ok := strings.CutPrefix(key.Key, KeyANY+" "); ok {
		// Replace spaces with hyphens so NormalizeName splits each word
		return text.NormalizeName(strings.ReplaceAll(desc, " ", "-"))
	}
	// Plain "ANY" with no descriptor or SubKeyType — skip this segment
	return ""
}

func (key *payloadKey) fullyQualified() []string {
	var name []string
	if !key.root {
		name = []string{text.NormalizeName(key.schema.Title)}
	}
	for _, parent := range key.parents {
		if n := normalizeANYKeyName(parent); n != "" {
			name = append(name, n)
		}
	}
	if n := normalizeANYKeyName(key.key); n != "" {
		name = append(name, n)
	}
	return name
}

// GlobalNamer registers Schemas to determine globally unique names for all generated types.
// When replacements are provided, collision detection uses post-replacement names
// so that two raw names that produce the same replaced name are correctly disambiguated.
type GlobalNamer struct {
	keyNames map[*PayloadKey]*payloadKey
}

func NewGlobalNamer(file *File, reps replace.Replacements) *GlobalNamer {
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
				repTyp:  replace.Const,
			})
		case *Struct:
			keys = append(keys, &payloadKey{
				schema:  t.Schema,
				parents: t.Parents,
				key:     t.Key,
				root:    t.Source != SourceSubKeys,
				repTyp:  replace.Struct,
			})
		case *Map:
			keys = append(keys, &payloadKey{
				schema:  t.Schema,
				parents: t.Parents,
				key:     t.Key,
				root:    t.Source != SourceSubKeys,
				repTyp:  replace.Struct,
			})
		}
	}

	// resolveName applies replacements to a candidate name for collision detection.
	// This ensures names that differ pre-replacement but collide post-replacement
	// are detected and disambiguated with longer qualified prefixes.
	resolveName := func(rawName string, repTyp replace.ReplacementType) string {
		return reps.Replace(rawName, repTyp)
	}

	// mark every possible key name (using post-replacement names for collision counting)
	names := make(map[string]int)
	for _, key := range keys {
		fq := key.fullyQualified()
		for start := len(fq) - 1; start >= 0; start-- {
			rawName := strings.Join(fq[start:], "")
			resolved := resolveName(rawName, key.repTyp)
			names[resolved] += 1
		}
	}

	// find minimal unique names (checking uniqueness of post-replacement names)
	count := make(map[string]int)
outer:
	for _, key := range keys {
		fq := key.fullyQualified()
		for start := len(fq) - 1; start >= 0; start-- {
			rawName := strings.Join(fq[start:], "")
			resolved := resolveName(rawName, key.repTyp)
			if names[resolved] == 1 {
				key.name = rawName
				namer.keyNames[key.key] = key
				continue outer
			}
		}
		// If we got here, the fully qualified identifier isn't unique
		// as a safeguard, return the name suffixed with an index.
		// Use the resolved prefix for counting so that different raw
		// prefixes that resolve to the same name get distinct indices.
		prefix := strings.Join(fq, "")
		resolvedPrefix := resolveName(prefix, key.repTyp)
		count[resolvedPrefix] += 1
		key.name = fmt.Sprintf("%s%d", prefix, count[resolvedPrefix])
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

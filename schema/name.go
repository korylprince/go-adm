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

func (key *payloadKey) fullyQualified(overrides map[*PayloadKey]string) []string {
	var name []string
	if !key.root {
		name = []string{text.NormalizeName(key.schema.Title)}
	}
	for _, parent := range key.parents {
		// A pinned name is already fully qualified, so it replaces what has
		// been accumulated above it rather than extending it. Without this,
		// a name under a pinned ancestor would repeat the whole path leading
		// to that ancestor as well.
		if pinned, ok := overrides[parent]; ok {
			name = []string{pinned}
			continue
		}
		if n := normalizeANYKeyName(parent); n != "" {
			name = append(name, n)
		}
	}
	if n := normalizeANYKeyName(key.key); n != "" {
		name = append(name, n)
	}
	return name
}

// collapseRepeats drops each segment that simply repeats the one before it.
//
// Apple names an array's element key after the array itself -- `reasons`
// holding `_Reasons` -- so a qualified name picks the same word up twice and
// reads as ...ReasonsReasons. The repeat carries no information.
func collapseRepeats(segments []string) []string {
	collapsed := make([]string, 0, len(segments))
	for i, seg := range segments {
		if i > 0 && seg == segments[i-1] {
			continue
		}
		collapsed = append(collapsed, seg)
	}
	return collapsed
}

// GlobalNamer registers Schemas to determine globally unique names for all generated types.
// When replacements are provided, collision detection uses post-replacement names
// so that two raw names that produce the same replaced name are correctly disambiguated.
type GlobalNamer struct {
	keyNames map[*PayloadKey]*payloadKey
}

type namerOptions struct {
	overrides map[*PayloadKey]string
	qualified map[*Schema]bool
}

type NamerOption func(*namerOptions)

// WithNameOverrides pins the name of specific PayloadKeys instead of deriving it
// from the minimal unique suffix of their qualified path.
//
// The minimal-suffix rule keeps everyday names short, but it makes each name a
// function of every *other* key registered in the same File, so adding one
// schema upstream can silently rename an unrelated type. Callers that synthesize
// a key tree and need its names to stay put across schema updates pin them here
// -- see generator/statustree.
//
// Overridden names are given pre-replacement, the same form the namer derives
// internally, so -repl rules still apply to them. They take part in collision
// detection, so an unpinned key that would otherwise have claimed the same name
// falls back to a longer prefix. Pinned names are assumed to be unique amongst
// themselves; two keys pinned to one name will collide.
func WithNameOverrides(overrides map[*PayloadKey]string) NamerOption {
	return func(opts *namerOptions) {
		opts.overrides = overrides
	}
}

// WithFullyQualifiedNames names every type generated from the given Schemas
// after its whole key path, rather than the shortest suffix of that path no
// other type has claimed.
//
// The shortest-suffix rule keeps names readable when a File holds many small
// schemas, because each schema's title already qualifies what is under it. It
// works badly for one large synthesized schema: the suffixes collapse to bare
// words like Type, Reason, Valid, and Errors, which say nothing at package
// scope and shift whenever a sibling is added. Field names are unaffected --
// those follow their own key, and stay as short as the JSON they map to.
//
// Paths are shortened in two ways. A pinned ancestor (see WithNameOverrides)
// replaces the path above it, and a segment that merely repeats the one before
// it is dropped.
func WithFullyQualifiedNames(schemas ...*Schema) NamerOption {
	return func(opts *namerOptions) {
		if opts.qualified == nil {
			opts.qualified = make(map[*Schema]bool)
		}
		for _, s := range schemas {
			opts.qualified[s] = true
		}
	}
}

func NewGlobalNamer(file *File, reps replace.Replacements, opts ...NamerOption) *GlobalNamer {
	nOpts := new(namerOptions)
	for _, opt := range opts {
		opt(nOpts)
	}

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

	// fixedName returns the name a key is given outright, rather than by
	// searching for the shortest unique suffix of its path.
	fixedName := func(key *payloadKey) (string, bool) {
		if name, ok := nOpts.overrides[key.key]; ok {
			return name, true
		}
		if nOpts.qualified[key.schema] {
			return strings.Join(collapseRepeats(key.fullyQualified(nOpts.overrides)), ""), true
		}
		return "", false
	}

	// Assign fixed names up front. They're counted once so the searches below
	// treat them as taken, but they never participate in a search themselves.
	names := make(map[string]int)
	fixed := make(map[string]int)
	for _, key := range keys {
		name, ok := fixedName(key)
		if !ok {
			continue
		}
		// Distinct keys shouldn't qualify to the same name, but fall back to an
		// index rather than declaring the same Go type twice if they do.
		base := resolveName(name, key.repTyp)
		if n := fixed[base]; n > 0 {
			name = fmt.Sprintf("%s%d", name, n+1)
		}
		fixed[base] += 1

		key.name = name
		namer.keyNames[key.key] = key
		names[resolveName(name, key.repTyp)] += 1
	}

	// mark every possible key name (using post-replacement names for collision counting)
	for _, key := range keys {
		if _, ok := fixedName(key); ok {
			continue
		}
		fq := key.fullyQualified(nOpts.overrides)
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
		if _, ok := fixedName(key); ok {
			continue
		}
		fq := key.fullyQualified(nOpts.overrides)
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

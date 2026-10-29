package replace

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"

	"github.com/korylprince/go-yaml"
)

type ReplacementType string

const (
	Const  ReplacementType = "const"
	Struct ReplacementType = "struct"
	Field  ReplacementType = "field"
)

// Replacement is regexp match and replacement strings
type Replacement struct {
	Match       string
	Replacement string
	Types       []string
	// Once limits this replacement to at most one application per identifier.
	// Other replacements continue to be applied.
	Once bool
}

// replace replaces s with r.Replacement if s matches r.Match and the
// replacement applies to one of typs. It reports whether the replacement
// matched.
func (r *Replacement) replace(s string, typs ...ReplacementType) (string, bool) {
	matchedType := false
	for _, typ := range typs {
		if slices.Contains(r.Types, string(typ)) {
			matchedType = true
			break
		}
	}
	if !matchedType {
		return s, false
	}

	reg := regexp.MustCompile(r.Match)
	if reg.MatchString(s) {
		return reg.ReplaceAllString(s, r.Replacement), true
	}
	return s, false
}

// Replace replaces s with r.Replacement if s matches r.Match
func (r *Replacement) Replace(s string, typ ReplacementType) string {
	s, _ = r.replace(s, typ)
	return s
}

type Replacements []*Replacement

// HasOnce reports whether any replacement is marked Once.
func (r Replacements) HasOnce() bool {
	for _, rep := range r {
		if rep.Once {
			return true
		}
	}
	return false
}

// ReplaceRepeated applies every repeated replacement, in order, until no
// further changes occur. Once replacements are ignored. This is the behavior
// used while assembling parts of an identifier.
func (r Replacements) ReplaceRepeated(s string, typs ...ReplacementType) string {
	for {
		prev := s
		for _, rep := range r {
			if rep.Once {
				continue
			}
			s, _ = rep.replace(s, typs...)
		}
		if s == prev {
			break
		}
	}
	return s
}

// ReplaceOnce applies every Once replacement at most once, in order. It
// repeats until no further Once replacements match so that one Once rule can
// create input for another, but an applied rule is never run again.
func (r Replacements) ReplaceOnce(s string, typs ...ReplacementType) string {
	applied := make([]bool, len(r))
	for {
		prev := s
		for i, rep := range r {
			if !rep.Once || applied[i] {
				continue
			}
			var matched bool
			s, matched = rep.replace(s, typs...)
			if matched {
				applied[i] = true
			}
		}
		if s == prev {
			break
		}
	}
	return s
}

// Replace normalizes s with the repeated replacements, applies the Once
// replacements, then normalizes the result again. The final pass lets ordinary
// replacements process the output of a Once rule.
//
// When no replacement is marked Once, Replace is exactly the previous
// repeat-to-a-fixpoint behavior.
func (r Replacements) Replace(s string, typ ReplacementType) string {
	s = r.ReplaceRepeated(s, typ)
	if !r.HasOnce() {
		return s
	}
	s = r.ReplaceOnce(s, typ)
	return r.ReplaceRepeated(s, typ)
}

// NewReplacements parses Replacements from a simple yaml key: value list
func NewReplacements(data []byte) (Replacements, error) {
	return decodeReplacements(bytes.NewReader(data))
}

// NewReplacementsFromFile parses Replacements from a file containing a simple yaml key: value list
func NewReplacementsFromFile(path string) (Replacements, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open: %w", err)
	}
	defer f.Close()

	return decodeReplacements(f)
}

func decodeReplacements(r io.Reader) (Replacements, error) {
	var ms yaml.MapSlice

	if err := yaml.NewDecoder(r).Decode(&ms); err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	reps := make(Replacements, len(ms))
	for idx, item := range ms {
		match, ok := item.Key.(string)
		if !ok {
			return nil, fmt.Errorf("replacement %d: match is not a string", idx)
		}

		vmap, ok := item.Value.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("replacement %q: expected a mapping", match)
		}

		repl, ok := vmap["repl"].(string)
		if !ok {
			return nil, fmt.Errorf("replacement %q: repl is not a string", match)
		}

		typesRaw, ok := vmap["types"].([]any)
		if !ok {
			return nil, fmt.Errorf("replacement %q: types is not a list", match)
		}
		types := make([]string, len(typesRaw))
		for i, typRaw := range typesRaw {
			types[i], ok = typRaw.(string)
			if !ok {
				return nil, fmt.Errorf("replacement %q: type %d is not a string", match, i)
			}
		}

		once, _ := vmap["once"].(bool)

		reps[idx] = &Replacement{
			Match:       match,
			Replacement: repl,
			Types:       types,
			Once:        once,
		}
	}

	return reps, nil
}

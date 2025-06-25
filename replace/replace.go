package replace

import (
	"bytes"
	"fmt"
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
}

// Replace replaces s with r.Replacement if s matches r.Match
func (r *Replacement) Replace(s string, typ ReplacementType) string {
	if !slices.Contains(r.Types, string(typ)) {
		return s
	}
	reg := regexp.MustCompile(r.Match)
	if reg.MatchString(s) {
		return regexp.MustCompile(r.Match).ReplaceAllString(s, r.Replacement)
	}
	return s
}

type Replacements []*Replacement

// Replace processes all replacements on s in order
func (r Replacements) Replace(s string, typ ReplacementType) string {
	for _, rep := range r {
		s = rep.Replace(s, typ)
	}
	return s
}

// NewReplacements parses Replacements from a simple yaml key: value list
func NewReplacements(data []byte) (Replacements, error) {
	var ms yaml.MapSlice

	if err := yaml.NewDecoder(bytes.NewReader(data)).Decode(&ms); err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	reps := make(Replacements, len(ms))
	for idx, item := range ms {
		vmap := item.Value.(map[string]any)
		var types []string
		for _, typ := range vmap["types"].([]any) {
			types = append(types, typ.(string))
		}
		reps[idx] = &Replacement{Match: item.Key.(string), Replacement: vmap["repl"].(string), Types: types}
	}

	return reps, nil
}

// NewReplacements parses Replacements from a file containing a simple yaml key: value list
func NewReplacementsFromFile(path string) (Replacements, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open: %w", err)
	}
	defer f.Close()

	var ms yaml.MapSlice

	if err = yaml.NewDecoder(f).Decode(&ms); err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	reps := make(Replacements, len(ms))
	for idx, item := range ms {
		vmap := item.Value.(map[string]any)
		var types []string
		for _, typ := range vmap["types"].([]any) {
			types = append(types, typ.(string))
		}
		reps[idx] = &Replacement{Match: item.Key.(string), Replacement: vmap["repl"].(string), Types: types}
	}

	return reps, nil
}

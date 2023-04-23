package replace

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/korylprince/go-yaml"
)

// Replacement is regexp match and replacement strings
type Replacement struct {
	Match       string
	Replacement string
}

// Replace replaces s with r.Replacement if s matches r.Match
func (r *Replacement) Replace(s string) string {
	reg := regexp.MustCompile(r.Match)
	if reg.MatchString(s) {
		return regexp.MustCompile(r.Match).ReplaceAllString(s, r.Replacement)
	}
	return s
}

type Replacements []*Replacement

// Replace processes all replacements on s in order
func (r Replacements) Replace(s string) string {
	for _, rep := range r {
		s = rep.Replace(s)
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
		reps[idx] = &Replacement{Match: item.Key.(string), Replacement: item.Value.(string)}
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
		reps[idx] = &Replacement{Match: item.Key.(string), Replacement: item.Value.(string)}
	}

	return reps, nil
}

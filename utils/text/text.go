package text

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	splitRegexp = regexp.MustCompile(`[_\-:]`)
	cleanRegexp = regexp.MustCompile("[^a-zA-Z0-9]")
)

// NormalizeName normalizes the name to an idomatic Go identifier
func NormalizeName(name string) string {
	parts := splitRegexp.Split(name, -1)
	norm := ""
	for _, p := range parts {
		norm += cases.Title(language.AmericanEnglish, cases.NoLower).String(cleanRegexp.ReplaceAllLiteralString(p, ""))
	}

	return norm
}

// DocComment formats the doc comment with leading //, taking care of multiline strings
func DocComment(doc string) string {
	var comment []string
	for _, line := range strings.Split(strings.TrimSpace(doc), "\n") {
		if c := strings.TrimSpace(line); c != "" {
			comment = append(comment, "// "+c)
		}
	}
	return strings.Join(comment, "\n")
}

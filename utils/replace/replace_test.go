package replace_test

import (
	"testing"

	"github.com/korylprince/go-adm/utils/replace"
)

func TestReplacementsOnce(t *testing.T) {
	reps, err := replace.NewReplacements([]byte(`
"^":
  repl: "Bar"
  types:
    - "struct"
  once: true
"^(.*)Caldav(.*)$":
  repl: "${1}CalDAV${2}"
  types:
    - "struct"
`))
	if err != nil {
		t.Fatal(err)
	}

	if got, want := reps.Replace("CaldavThing", replace.Struct), "BarCalDAVThing"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReplacementsOnceSuffix(t *testing.T) {
	reps := replace.Replacements{
		{Match: "$", Replacement: "Type", Types: []string{"struct"}, Once: true},
		{Match: `^(.*)Url(.*)$`, Replacement: "${1}URL${2}", Types: []string{"struct"}},
	}

	// The ordinary rule normalizes the name, the Once suffix is applied to
	// that result, and the final normalization pass leaves it stable.
	if got, want := reps.Replace("CallbackUrl", replace.Struct), "CallbackURLType"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReplacementsOnceSeesNormalizedResult(t *testing.T) {
	reps := replace.Replacements{
		{Match: `^Caldav$`, Replacement: "CalDAV", Types: []string{"struct"}},
		{Match: "^CalDAV$", Replacement: "BarCalDAV", Types: []string{"struct"}, Once: true},
	}

	if got, want := reps.Replace("Caldav", replace.Struct), "BarCalDAV"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Existing repl files have no Once rules. Ordinary stacks must still apply in
// order and repeat to a fixpoint.
func TestReplacementsWithoutOnce(t *testing.T) {
	reps := replace.Replacements{
		{Match: `^(.*)Caldav(.*)$`, Replacement: "${1}CalDAV${2}", Types: []string{"struct"}},
		{Match: `^(.*)Url(.*)$`, Replacement: "${1}URL${2}", Types: []string{"struct"}},
	}

	for _, tc := range []struct {
		name string
		want string
	}{
		{"CaldavUrl", "CalDAVURL"},
		{"UrlUrlUrl", "URLURLURL"},
	} {
		if got := reps.Replace(tc.name, replace.Struct); got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, got, tc.want)
		}
	}
}

// Non-Once rules must still run to a fixpoint, including over every occurrence
// in a name assembled by GlobalNamer.
func TestReplacementsStillRepeat(t *testing.T) {
	reps := replace.Replacements{
		{Match: `^(.*)Url(.*)$`, Replacement: "${1}URL${2}", Types: []string{"struct"}},
	}

	if got, want := reps.Replace("UrlUrlUrl", replace.Struct), "URLURLURL"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Once rules can feed each other in a single stack, and ordinary rules still
// normalize the result. Without Once the two expansion rules would grow the
// name forever.
func TestReplacementsOnceStack(t *testing.T) {
	reps := replace.Replacements{
		{Match: `^(.*)Cert(.*)$`, Replacement: "${1}Certificate${2}", Types: []string{"struct"}, Once: true},
		{Match: `^(.*)Certificate(.*)$`, Replacement: "${1}CertificateInfo${2}", Types: []string{"struct"}, Once: true},
		{Match: `^(.*)Device(.*)$`, Replacement: "${1}Hardware${2}", Types: []string{"struct"}},
	}

	if got, want := reps.Replace("DeviceCert", replace.Struct), "HardwareCertificateInfo"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNewReplacementsOnceOptional(t *testing.T) {
	reps, err := replace.NewReplacements([]byte(`
"^Foo":
  repl: "Bar"
  types:
    - "struct"
`))
	if err != nil {
		t.Fatal(err)
	}

	if len(reps) != 1 {
		t.Fatalf("got %d replacements, want 1", len(reps))
	}
	if reps[0].Once {
		t.Error("expected Once to default to false")
	}
}

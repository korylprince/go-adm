package text_test

import (
	"testing"

	"github.com/korylprince/go-adm/text"
)

var nameTests = map[string]string{
	"test":                "Test",
	"test_test":           "TestTest",
	"test-test":           "TestTest",
	"test:test":           "TestTest",
	"test_test-test:test": "TestTestTestTest",
	"testTest":            "TestTest",
	"test-testTest":       "TestTestTest",
}

func TestNormalizeName(t *testing.T) {
	for test, want := range nameTests {
		if have := text.NormalizeName(test); have != want {
			t.Errorf("test (%s) failed: have: %s, want: %s", test, have, want)
		}
	}
}

var commentTests = map[string]string{
	"test":             "// test",
	"test\ntest":       "// test\n// test",
	"test\ntest\n":     "// test\n// test",
	"test\ntest\ntest": "// test\n// test\n// test",
	"test\n\ntest":     "// test\n// test",
	"test\n\ntest\n":   "// test\n// test",
}

func TestDocComment(t *testing.T) {
	for test, want := range commentTests {
		if have := text.DocComment(test); have != want {
			t.Errorf("test (%s) failed: have: %s, want: %s", test, have, want)
		}
	}
}

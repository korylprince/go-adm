package declarative_test

import (
	"encoding/json"
	"testing"

	"github.com/korylprince/go-adm/declarative"
)

func TestNewFromType(t *testing.T) {
	for typ := range declarative.DeclarationMap {
		decl, err := declarative.NewFromType(typ, "com.example", "v1.0.0")
		if err != nil {
			t.Errorf("could not create declaration for %s: %v", typ, err)
			continue
		}

		if _, err := json.Marshal(decl); err != nil {
			t.Errorf("could not marshal json for %s: %v", typ, err)
		}
	}
}

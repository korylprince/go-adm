package declarations_test

import (
	"encoding/json"
	"testing"

	"github.com/korylprince/go-adm/declarations"
)

func TestNewFromType(t *testing.T) {
	for typ := range declarations.DeclarationMap {
		decl, err := declarations.NewFromType(typ, "com.example", "v1.0.0")
		if err != nil {
			t.Errorf("could not create declaration for %s: %v", typ, err)
			continue
		}

		if _, err := json.Marshal(decl); err != nil {
			t.Errorf("could not marshal json for %s: %v", typ, err)
		}
	}
}

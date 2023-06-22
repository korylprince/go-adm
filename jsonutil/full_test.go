package jsonutil_test

import (
	"encoding/json"
	"testing"

	"github.com/korylprince/go-adm/declarations"
	"github.com/korylprince/go-adm/jsonutil"
)

func TestFullFields(t *testing.T) {
	for typ := range declarations.DeclarationMap {
		decl, err := declarations.NewFromType(typ, "id", "tok")
		if err != nil {
			t.Errorf("could not generate declaration %s: %v", typ, err)
			continue
		}

		_, err = json.Marshal(jsonutil.FullFields(decl))
		if err != nil {
			t.Errorf("could not marshal declaration %s: %v", typ, err)
		}
	}

}

package tagutil_test

import (
	"encoding/json"
	"testing"

	"github.com/korylprince/go-adm/declarative"
	"github.com/korylprince/go-adm/tagutil"
)

func TestFullFields(t *testing.T) {
	for typ := range declarative.DeclarationMap {
		decl, err := declarative.NewFromType(typ, "id", "tok")
		if err != nil {
			t.Errorf("could not generate declaration %s: %v", typ, err)
			continue
		}

		_, err = json.Marshal(tagutil.FullFields(decl))
		if err != nil {
			t.Errorf("could not marshal declaration %s: %v", typ, err)
		}
	}

}

//go:generate yamlschemagen -repo "https://github.com/apple/device-management.git" -commit "7d4ba1a2bde50a4053fa5a5e0ed6c17388d82ab2" -path "docs/schema.yaml" -pkg schema -repl ./repls.yaml -out schema.gen.go
package schema

import (
	"bytes"
	"fmt"
	"os"
	"slices"

	"github.com/korylprince/go-yaml"
)

// KeyANY is a special *PayloadKey.Key to represent a generic map[string]any instead of a struct
const KeyANY = "ANY"

func New(data []byte) (*Schema, error) {
	cmd := new(Schema)
	if err := yaml.NewDecoder(bytes.NewReader(data)).Decode(cmd); err != nil {
		return nil, fmt.Errorf("could not unmarshal schema: %w", err)
	}

	return cmd, nil
}

func NewFromFile(path string) (*Schema, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open: %w", err)
	}
	defer f.Close()

	cmd := new(Schema)
	if err = yaml.NewDecoder(f).Decode(cmd); err != nil {
		return nil, fmt.Errorf("could not unmarshal schema: %w", err)
	}

	return cmd, nil
}

func (s *Schema) Iter(f func(parents []*PayloadKey, key *PayloadKey)) {
	var dfs func(parents []*PayloadKey, child *PayloadKey)
	dfs = func(parents []*PayloadKey, child *PayloadKey) {
		f(parents, child)
		newParents := append(slices.Clone(parents), child)
		for _, key := range child.SubKeys {
			dfs(newParents, key)
		}
	}
	for _, key := range s.PayloadKeys {
		dfs(nil, key)
	}
	for _, key := range s.ResponseKeys {
		dfs(nil, key)
	}
}

// IsStruct returns if the PayloadKey should be rendered as a struct
func (key *PayloadKey) IsStruct() bool {
	if key.Type != PayloadKeyTypeDictionary {
		return false
	}
	if len(key.SubKeys) > 0 && key.SubKeys[0].Type == PayloadKeyTypeAny {
		return false
	}
	return true
}

// IsMap returns if the PayloadKey should be rendered as a map
func (key *PayloadKey) IsMap() bool {
	if key.Type != PayloadKeyTypeDictionary {
		return false
	}
	if len(key.SubKeys) > 0 && key.SubKeys[0].Type == PayloadKeyTypeAny {
		return true
	}
	return false
}

// IsEnum returns if the PayloadKey should be rendered as an enum
func (key *PayloadKey) IsEnum() bool {
	if len(key.Rangelist) > 0 {
		return true
	}
	if key.Type == PayloadKeyTypeArray && len(key.SubKeys) == 1 && len(key.SubKeys[0].Rangelist) > 0 {
		return true
	}
	return false
}

//go:generate /bin/bash -c "yamlschemagen -repo 'https://github.com/apple/device-management.git' -commit \"$(cat ../GENERATE_SHA)\" -path 'docs/schema.yaml' -pkg schema -repl ./repls.yaml -out schema.gen.go"
package schema

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/korylprince/go-yaml"
)

// KeyANY is a special *PayloadKey.Key to represent a generic map[string]any instead of a struct
const KeyANY = "ANY"

// IsANY returns true if the key name indicates a dynamic/generic key.
// This matches both "ANY" and descriptors like "ANY restriction name".
func (key *PayloadKey) IsANY() bool {
	return key.Key == KeyANY || strings.HasPrefix(key.Key, KeyANY+" ")
}

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
	seen := make(map[*PayloadKey]struct{})
	var dfs func(parents []*PayloadKey, child *PayloadKey)
	dfs = func(parents []*PayloadKey, child *PayloadKey) {
		if _, ok := seen[child]; ok {
			return
		}
		seen[child] = struct{}{}
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
	if len(key.SubKeys) > 0 && key.SubKeys[0].IsANY() {
		return false
	}
	return true
}

// IsMap returns if the PayloadKey should be rendered as a map
func (key *PayloadKey) IsMap() bool {
	if key.Type != PayloadKeyTypeDictionary {
		return false
	}
	if len(key.SubKeys) > 0 && key.SubKeys[0].IsANY() {
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

//go:generate yamlschemagen -repo "https://github.com/apple/device-management.git" -commit "release" -path "docs/schema.yaml" -pkg schema -repl ./repls.yaml -out schema.gen.go
package schema

import (
	"bytes"
	"fmt"
	"os"

	"github.com/korylprince/go-yaml"
)

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

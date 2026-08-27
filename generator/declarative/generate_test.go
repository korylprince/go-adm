package gen_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	gen "github.com/korylprince/go-adm/generator/declarative"
)

// read returns the contents of a generated file, failing the test if it wasn't
// written.
func read(t *testing.T, path string) string {
	t.Helper()

	buf, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("could not read generated %s: %v", path, err)
	}
	return string(buf)
}

func mustContain(t *testing.T, src string, wants ...string) {
	t.Helper()

	for _, want := range wants {
		if !strings.Contains(src, want) {
			t.Errorf("generated code is missing %q:\n%s", want, src)
		}
	}
}

func mustNotContain(t *testing.T, src string, wants ...string) {
	t.Helper()

	for _, want := range wants {
		if strings.Contains(src, want) {
			t.Errorf("generated code unexpectedly contains %q:\n%s", want, src)
		}
	}
}

// TestGenerateFromPaths covers the local tree mode: schemas read off disk with
// no git repo, one package per source directory.
func TestGenerateFromPaths(t *testing.T) {
	out := t.TempDir()

	if err := gen.GenerateFromPaths([]string{"testdata"}, nil, out); err != nil {
		t.Fatal(err)
	}

	t.Run("one package per source directory", func(t *testing.T) {
		status := read(t, filepath.Join(out, "status", "status.go"))
		mustContain(t, status, "package status")

		protocol := read(t, filepath.Join(out, "protocol", "protocol.go"))
		mustContain(t, protocol, "package protocol")
	})

	t.Run("status items are decomposed into a nested tree", func(t *testing.T) {
		src := read(t, filepath.Join(out, "status", "status.go"))

		mustContain(t, src,
			"type StatusItemsDevice struct",
			"type StatusItemsDeviceModel struct",
			"type StatusItemsDeviceIdentifier struct",
			`Family *string `+"`"+`json:"family,omitempty"`+"`",
			`Udid *string `+"`"+`json:"udid,omitempty"`+"`",
			`IsPresent *bool `+"`"+`json:"is-present,omitempty"`+"`",
		)

		// the dotted key must not survive as a literal json tag
		mustNotContain(t, src, `json:"device.model.family"`)

		// the identifier stays discoverable in the docs
		mustContain(t, src, "Status item: `device.model.family`.")
	})

	t.Run("the envelope is grafted around the tree", func(t *testing.T) {
		// The envelope lives in protocol/ while the items live in status/, so
		// this also covers finding it across directories.
		src := read(t, filepath.Join(out, "status", "status.go"))

		mustContain(t, src,
			"type StatusReport struct",
			`StatusItems StatusItems `+"`"+`json:"StatusItems"`+"`",
			`FullReport *bool `+"`"+`json:"FullReport,omitempty"`+"`",
		)
	})

	t.Run("the dotted identifiers are generated as values", func(t *testing.T) {
		// The tree turns each identifier into a path of fields, so nothing else
		// in the output holds the identifier itself -- and that is the form the
		// protocol uses in subscriptions and in Errors[].StatusItem.
		src := read(t, filepath.Join(out, "status", "status.go"))

		mustContain(t, src,
			"var StatusItemTypes = []string{",
			`"device.model.family"`,
			`"device.identifier.udid"`,
			`"passcode.is-present"`,
			"StatusItemTypes lists every DDM status item",
		)
	})

	t.Run("protocol keeps the lossless untyped report", func(t *testing.T) {
		src := read(t, filepath.Join(out, "protocol", "protocol.go"))

		mustContain(t, src, "type StatusReport struct", "StatusItems map[string]any")
		mustNotContain(t, src, "type StatusItems struct{}")
	})
}

// TestGeneratePackageFromPaths covers single package mode, the shape a
// downstream repo uses to drop the types into a package of its own.
func TestGeneratePackageFromPaths(t *testing.T) {
	t.Run("status directory only", func(t *testing.T) {
		buf := new(strings.Builder)
		if err := gen.GeneratePackageFromPaths([]string{"testdata/status"}, "ddmstatus", nil, buf); err != nil {
			t.Fatal(err)
		}

		src := buf.String()
		mustContain(t, src, "package ddmstatus", "type StatusItems struct", "type StatusItemsDeviceModel struct")

		// no envelope was loaded, so there is no report type to wrap it
		mustNotContain(t, src, "type StatusReport struct")
	})

	t.Run("everything merged into one package", func(t *testing.T) {
		buf := new(strings.Builder)
		if err := gen.GeneratePackageFromPaths([]string{"testdata"}, "ddm", nil, buf); err != nil {
			t.Fatal(err)
		}

		src := buf.String()
		mustContain(t, src, "package ddm", "type StatusReport struct", "type StatusItemsDeviceModel struct")

		// The envelope is in the same package as the tree here, so the grafted
		// copy replaces it rather than being generated alongside it.
		if n := strings.Count(src, "type StatusReport struct"); n != 1 {
			t.Errorf("found %d StatusReport types, want 1:\n%s", n, src)
		}
		mustNotContain(t, src, "StatusItems map[string]any")
	})

	t.Run("individually named files", func(t *testing.T) {
		buf := new(strings.Builder)
		err := gen.GeneratePackageFromPaths([]string{
			"testdata/status/device.model.family.yaml",
			"testdata/status/device.identifier.udid.yaml",
		}, "ddm", nil, buf)
		if err != nil {
			t.Fatal(err)
		}

		src := buf.String()
		mustContain(t, src, "type StatusItemsDeviceModel struct", "type StatusItemsDeviceIdentifier struct")
		mustNotContain(t, src, "IsPresent")
	})
}

// Pointing -path at leaf directories puts every schema at the top of its own
// root, where the package name has to come from the root rather than a
// subdirectory -- and each still needs a directory of its own, since two
// package clauses in one directory is not a buildable Go package.
func TestGenerateFromPathsLeafDirectories(t *testing.T) {
	out := t.TempDir()

	err := gen.GenerateFromPaths([]string{"testdata/status", "testdata/protocol"}, nil, out)
	if err != nil {
		t.Fatal(err)
	}

	mustContain(t, read(t, filepath.Join(out, "status", "status.go")), "package status")
	mustContain(t, read(t, filepath.Join(out, "protocol", "protocol.go")), "package protocol")

	for _, stray := range []string{"status.go", "protocol.go"} {
		if _, err := os.Stat(filepath.Join(out, stray)); err == nil {
			t.Errorf("%s was written to the output root, not its own directory", stray)
		}
	}
}

func TestGenerateFromPathsErrors(t *testing.T) {
	if err := gen.GenerateFromPaths([]string{"testdata/does-not-exist"}, nil, t.TempDir()); err == nil {
		t.Error("expected an error for a missing path")
	}
}

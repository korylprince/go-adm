package gen_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/korylprince/go-adm/generated/declarative/status"
)

// statusItemPaths is every status item Apple defines at the schema commit this
// package was generated from, as its dotted identifier.
//
// The list is checked against what StatusItems actually exposes. Its job is to
// catch items going missing: a transform bug, or a rename upstream, otherwise
// shows up only as a field that quietly stopped being generated. Regenerating
// against a newer schema commit is expected to add entries here.
var statusItemPaths = []string{
	"account.list.caldav",
	"account.list.carddav",
	"account.list.exchange",
	"account.list.google",
	"account.list.ldap",
	"account.list.mail.incoming",
	"account.list.mail.outgoing",
	"account.list.subscribed-calendar",
	"app.managed.list",
	"device.identifier.serial-number",
	"device.identifier.udid",
	"device.model.family",
	"device.model.identifier",
	"device.model.marketing-name",
	"device.model.number",
	"device.operating-system.build-version",
	"device.operating-system.family",
	"device.operating-system.marketing-name",
	"device.operating-system.supplemental.build-version",
	"device.operating-system.supplemental.extra-version",
	"device.operating-system.version",
	"device.power.battery-health",
	"diskmanagement.filevault.enabled",
	"management.client-capabilities",
	"management.declarations",
	"mdm.app",
	"package.list",
	"passcode.is-compliant",
	"passcode.is-present",
	"screensharing.connection.group.unresolved-connection",
	"security.certificate.list",
	"services.background-task",
	"softwareupdate.beta-enrollment",
	"softwareupdate.device-id",
	"softwareupdate.failure-reason",
	"softwareupdate.install-reason",
	"softwareupdate.install-state",
	"softwareupdate.pending-version",
	"test.array-value",
	"test.boolean-value",
	"test.dictionary-value",
	"test.error-value",
	"test.integer-value",
	"test.real-value",
	"test.string-value",
}

// reachablePaths walks StatusItems and returns the dotted path of every status
// item leaf, rebuilt from the json tags the way a device nests them.
//
// Interior namespace nodes are told apart by their generated name: declgen pins
// those to a StatusItems-prefixed name derived from the path, precisely so they
// are stable and identifiable. Anything else is a leaf value type -- Apple's own
// declared shape for the item -- and is not descended into.
func reachablePaths(t reflect.Type, prefix []string) []string {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	var paths []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, _, _ := strings.Cut(field.Tag.Get("json"), ",")
		path := append(append([]string{}, prefix...), tag)

		typ := field.Type
		for typ.Kind() == reflect.Pointer {
			typ = typ.Elem()
		}

		if typ.Kind() == reflect.Struct && strings.HasPrefix(typ.Name(), "StatusItems") {
			paths = append(paths, reachablePaths(typ, path)...)
			continue
		}

		paths = append(paths, strings.Join(path, "."))
	}

	return paths
}

func TestStatusItemCoverage(t *testing.T) {
	got := reachablePaths(reflect.TypeOf(status.StatusItems{}), nil)

	want := make(map[string]bool, len(statusItemPaths))
	for _, path := range statusItemPaths {
		want[path] = true
	}

	for _, path := range got {
		if !want[path] {
			t.Errorf("unexpected status item %q in StatusItems", path)
		}
		delete(want, path)
	}

	for path := range want {
		t.Errorf("status item %q is not reachable in StatusItems", path)
	}
}

// TestDecodeReports decodes each report fixture with unknown fields disallowed,
// then re-marshals it and compares.
//
// Strict decoding is the point: it proves the typed tree accounts for every key
// in the document rather than quietly dropping the ones it has no field for,
// which is exactly the failure a status-report parser must not have.
func TestDecodeReports(t *testing.T) {
	paths, err := filepath.Glob("testdata/reports/*.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) == 0 {
		t.Fatal("no report fixtures found")
	}

	for _, path := range paths {
		t.Run(filepath.Base(path), func(t *testing.T) {
			raw, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			var report status.StatusReport
			dec := json.NewDecoder(bytes.NewReader(raw))
			dec.DisallowUnknownFields()
			if err = dec.Decode(&report); err != nil {
				t.Fatalf("strict decode: %v", err)
			}

			out, err := json.Marshal(&report)
			if err != nil {
				t.Fatal(err)
			}

			// Compare as normalized JSON rather than bytes: key order and
			// whitespace in the fixture are not what's under test.
			if a, b := normalize(t, raw), normalize(t, out); a != b {
				t.Errorf("round trip changed the document\n--- fixture ---\n%s\n--- re-marshaled ---\n%s", a, b)
			}
		})
	}
}

func normalize(t *testing.T, raw []byte) string {
	t.Helper()

	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		t.Fatal(err)
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestDecodeValues(t *testing.T) {
	raw, err := os.ReadFile("testdata/reports/report.device.json")
	if err != nil {
		t.Fatal(err)
	}

	var report status.StatusReport
	if err = json.Unmarshal(raw, &report); err != nil {
		t.Fatal(err)
	}

	items := report.StatusItems

	t.Run("dotted items resolve through the tree", func(t *testing.T) {
		if got := deref(t, items.Device.Model.Family); got != "Mac" {
			t.Errorf("device.model.family = %q", got)
		}
		if got := deref(t, items.Device.Identifier.SerialNumber); got != "X0X0X0X0X0X0" {
			t.Errorf("device.identifier.serial-number = %q", got)
		}
		if got := deref(t, items.Device.OperatingSystem.Version); got != "15.3.1" {
			t.Errorf("device.operating-system.version = %q", got)
		}
		// the deepest namespace: supplemental is interior, not a value
		if got := deref(t, items.Device.OperatingSystem.Supplemental.ExtraVersion); got != "(a)" {
			t.Errorf("device.operating-system.supplemental.extra-version = %q", got)
		}
	})

	t.Run("typed leaf values", func(t *testing.T) {
		if got := deref(t, items.Passcode.IsPresent); !got {
			t.Error("passcode.is-present should be true")
		}
		if got := deref(t, items.DiskManagement.FileVault.Enabled); !got {
			t.Error("diskmanagement.filevault.enabled should be true")
		}
		if got := deref(t, items.Device.Power.BatteryHealth); got != status.DevicePowerBatteryHealthNormal {
			t.Errorf("device.power.battery-health = %q", got)
		}
		if got := deref(t, items.SoftwareUpdate.InstallState); got != status.SoftwareUpdateInstallStateNone {
			t.Errorf("softwareupdate.install-state = %q", got)
		}
	})

	t.Run("array and dictionary leaves keep Apple's declared shape", func(t *testing.T) {
		certs := deref(t, items.Security.Certificate.List)
		if len(certs) != 1 {
			t.Fatalf("security.certificate.list = %d entries, want 1", len(certs))
		}
		if certs[0].Identifier != "com.example.cert" {
			t.Errorf("certificate identifier = %q", certs[0].Identifier)
		}
		if !certs[0].IsIdentity {
			t.Error("certificate should be an identity")
		}

		decls := items.Management.Declarations
		if decls == nil || len(decls.Configurations) != 1 {
			t.Fatal("management.declarations.configurations not decoded")
		}
		cfg := decls.Configurations[0]
		if cfg.Identifier != "com.example.config" || cfg.Valid != status.ManagementDeclarationsActivationsValidValid || !cfg.Active {
			t.Errorf("configuration = %+v", cfg)
		}
	})

	t.Run("errors carry the dotted identifier", func(t *testing.T) {
		if len(report.Errors) != 1 {
			t.Fatalf("Errors = %d entries, want 1", len(report.Errors))
		}
		// test.error-value exists to provoke exactly this
		if got := report.Errors[0].StatusItem; got != "test.error-value" {
			t.Errorf("Errors[0].StatusItem = %q", got)
		}
	})

	t.Run("FullReport is typed", func(t *testing.T) {
		if report.FullReport == nil || !*report.FullReport {
			t.Error("FullReport should decode as true")
		}
	})
}

func deref[T any](t *testing.T, v *T) T {
	t.Helper()

	if v == nil {
		var zero T
		t.Error("expected a value, got nil")
		return zero
	}
	return *v
}

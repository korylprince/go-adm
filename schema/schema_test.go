package schema_test

import (
	"bytes"
	"fmt"
	"net"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"github.com/korylprince/go-adm/git"
	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-yaml"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// e2e is a helper function to "clean" a schema file for comparison. It passes the yaml buf through unmarshaling/marshaling to help diff two schemas
// e2e is not used in any tests, but is helpful for exploring the schemas.
func e2e(buf []byte) ([]byte, error) {
	v := new(any)
	if err := yaml.Unmarshal(buf, v); err != nil {
		return nil, fmt.Errorf("could not unmarshal: %w", err)
	}

	buf, err := yaml.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("could not marshal: %w", err)
	}

	// manual fixes because of optional fields.
	buf = regexp.MustCompile(`.*description: "".*\n`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*requiresdep: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*supervised: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*userapprovedmdm: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*allowmanualinstall: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*always-skippable: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*default: "?false"?.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*default: "?true"?.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*userchannel: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*devicechannel: false.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*sharedipad:.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*macOS:.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*tvOS:.*`).ReplaceAll(buf, nil)
	buf = regexp.MustCompile(`\n.*subkeys: null.*`).ReplaceAll(buf, nil)

	return buf, nil
}

func Test(t *testing.T) {
	if _, err := net.LookupAddr("github.com"); err != nil {
		t.Skip("could not resolve github.com, skipping")
	}
	repo, err := git.New("https://github.com/apple/device-management.git", schema.DeviceManagementGenerateHash)
	if err != nil {
		t.Fatalf("could not checkout repo: %v", err)
	}

	t.Run("Schema", func(t *testing.T) {
		t.Run("MDM", func(t *testing.T) {
			testSchema(t, repo, "mdm")
		})

		t.Run("Declarative", func(t *testing.T) {
			testSchema(t, repo, "declarative")
		})

		t.Run("Other", func(t *testing.T) {
			testSchema(t, repo, "other")
		})
	})
	t.Run("Consts", func(t *testing.T) {
		t.Run("MDM", func(t *testing.T) {
			testConsts(t, repo, "mdm")
		})

		t.Run("Declarative", func(t *testing.T) {
			testConsts(t, repo, "declarative")
		})

		t.Run("Other", func(t *testing.T) {
			testConsts(t, repo, "other")
		})
	})
}

func testSchema(t *testing.T, repo billy.Filesystem, path string) {
	info, err := repo.Stat(path)
	if err != nil {
		t.Fatalf("could not read %s: %v", path, err)
	}

	if info.IsDir() {
		entries, err := repo.ReadDir(path)
		if err != nil {
			t.Fatalf("could not read %s directory: %v", path, err)
		}
		for _, e := range entries {
			if e.IsDir() || strings.HasSuffix(e.Name(), ".yaml") {
				name := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
				t.Run(cases.Title(language.AmericanEnglish).String(name), func(t *testing.T) {
					testSchema(t, repo, filepath.Join(path, e.Name()))
				})
			}
		}
		return
	}

	// read original
	orig, err := util.ReadFile(repo, path)
	if err != nil {
		t.Fatalf("could not read %s: %v", path, err)
	}

	// unmarshal to schema
	cmd := new(schema.Schema)
	if err := yaml.Unmarshal(orig, cmd); err != nil {
		t.Fatalf("could not unmarshal %s: %v", path, err)
	}

	// marshal schema
	buf, err := yaml.Marshal(cmd)
	if err != nil {
		t.Errorf("could not marshal %s: %v", path, err)
	}

	// unmarshal marshalled schema
	cmd2 := new(schema.Schema)
	if err := yaml.Unmarshal(orig, cmd2); err != nil {
		t.Fatalf("could not unmarshal marshaled schema %s: %v", path, err)
	}

	// marshal schema
	buf2, err := yaml.Marshal(cmd2)
	if err != nil {
		t.Errorf("could not marshal e2e schema %s: %v", path, err)
	}

	// check e2e marshal/unmarshal is equal
	// ideally, we'd be able to marshal to a schema equal to the original file. However due to the complexity of yaml and the fact that some of Apple's schema files are out of spec with their own schema, this is very hard at least, and possibly impossible.
	if !bytes.Equal(buf, buf2) {
		t.Errorf("e2e schema not equal %s: %v", path, err)
	}
}

func checkiPadMode(name string, mode schema.SharediPadMode, t *testing.T) {
	switch mode {
	case "", schema.SharediPadModeAllowed, schema.SharediPadModeForbidden, schema.SharediPadModeIgnored, schema.SharediPadModeRequired:
	default:
		t.Errorf("unknown SharediPadMode: %s: %s", name, mode)
	}
}

func checkUserMode(name string, mode schema.UserEnrollmentMode, t *testing.T) {
	switch mode {
	case "", schema.UserEnrollmentModeAllowed, schema.UserEnrollmentModeForbidden, schema.UserEnrollmentModeIgnored, schema.UserEnrollmentModeRequired:
	default:
		t.Errorf("unknown UserEnrollmentMode: %s: %s", name, mode)
	}
}

func checkOS(name string, os *schema.OS, t *testing.T) {
	if os.SharediPad != nil {
		checkiPadMode(name, os.SharediPad.Mode, t)
	}
	if os.UserEnrollment != nil {
		checkUserMode(name, os.UserEnrollment.Mode, t)
	}
}

func checkSupportedOS(name string, os *schema.SupportedOS, t *testing.T) {
	if os.MacOS != nil {
		checkOS(name, os.MacOS, t)
	}
	if os.IOS != nil {
		checkOS(name, os.IOS, t)
	}
	if os.TVOS != nil {
		checkOS(name, os.TVOS, t)
	}
	if os.WatchOS != nil {
		checkOS(name, os.WatchOS, t)
	}
}

func checkKey(name string, key *schema.PayloadKey, t *testing.T) {
	if key.SupportedOS != nil {
		checkSupportedOS(name, key.SupportedOS, t)
	}

	switch key.Type {
	case "", schema.PayloadKeyTypeAny, schema.PayloadKeyTypeArray, schema.PayloadKeyTypeBoolean, schema.PayloadKeyTypeData, schema.PayloadKeyTypeDate, schema.PayloadKeyTypeDictionary, schema.PayloadKeyTypeInteger, schema.PayloadKeyTypeReal, schema.PayloadKeyTypeString:
	default:
		t.Errorf("unknown KeyType: %s: %s", name, key.Type)
	}

	switch key.SubType {
	case "", schema.PayloadKeySubTypeEmail, schema.PayloadKeySubTypeHostname, schema.PayloadKeySubTypeURL:
	default:
		t.Errorf("unknown KeySubType: %s: %s", name, key.SubType)
	}

	switch key.Presence {
	case "", schema.PayloadKeyPresenceOptional, schema.PayloadKeyPresenceRequired:
	default:
		t.Errorf("unknown KeyPresence: %s: %s", name, key.Presence)
	}

	for _, k := range key.SubKeys {
		checkKey(name, k, t)
	}
}

func testConsts(t *testing.T, repo billy.Filesystem, path string) {
	info, err := repo.Stat(path)
	if err != nil {
		t.Fatalf("could not read %s: %v", path, err)
	}

	if info.IsDir() {
		entries, err := repo.ReadDir(path)
		if err != nil {
			t.Fatalf("could not read %s directory: %v", path, err)
		}
		for _, e := range entries {
			if e.IsDir() || strings.HasSuffix(e.Name(), ".yaml") {
				name := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
				t.Run(cases.Title(language.AmericanEnglish).String(name), func(t *testing.T) {
					testConsts(t, repo, filepath.Join(path, e.Name()))
				})
			}
		}
		return
	}

	buf, err := util.ReadFile(repo, path)
	if err != nil {
		t.Fatalf("could not read %s: %v", path, err)
	}

	cmd := new(schema.Schema)
	if err := yaml.Unmarshal(buf, cmd); err != nil {
		t.Fatalf("could not parse schema %s: %v", path, err)
	}

	if cmd.Payload.SupportedOS != nil {
		checkSupportedOS(path, cmd.Payload.SupportedOS, t)
	}

	for _, k := range cmd.PayloadKeys {
		checkKey(path, k, t)
	}

	for _, k := range cmd.ResponseKeys {
		checkKey(path, k, t)
	}
}

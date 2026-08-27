package gen

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-git/go-billy/v5/util"
	"github.com/korylprince/go-adm/generator/statustree"
	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/utils/git"
	"github.com/korylprince/go-adm/utils/replace"
)

// sourceSchema is a parsed schema together with where it was found.
//
// Tree mode turns each source directory into its own Go package, so the
// directory has to survive parsing; single package mode ignores it.
type sourceSchema struct {
	// root is the input path -- a git path, local directory, or local file --
	// that this schema was loaded from.
	root string
	// dir is the schema's directory relative to root, empty for a schema
	// sitting directly in the root.
	dir string
	s   *schema.Schema
}

// source returns the directory the schema came from, for header comments.
func (src *sourceSchema) source() string {
	return filepath.Join(src.root, src.dir)
}

func newSourceSchema(root, filePath string, data []byte) (*sourceSchema, error) {
	s, err := schema.New(data)
	if err != nil {
		return nil, fmt.Errorf("could not parse %s: %w", filePath, err)
	}

	rel, err := filepath.Rel(root, filePath)
	if err != nil {
		return nil, fmt.Errorf("could not get relative path for %s: %w", filePath, err)
	}

	dir := filepath.Dir(rel)
	if dir == "." {
		dir = ""
	}

	return &sourceSchema{root: root, dir: dir, s: s}, nil
}

// loadFromGit parses every .yaml under each of paths in the checked out repo.
func loadFromGit(repo *git.Repository, paths []string) ([]*sourceSchema, error) {
	var srcs []*sourceSchema

	for _, path := range paths {
		if err := util.Walk(repo, path, func(filePath string, info fs.FileInfo, _ error) error {
			if !strings.HasSuffix(info.Name(), ".yaml") {
				return nil
			}

			buf, err := util.ReadFile(repo, filePath)
			if err != nil {
				return fmt.Errorf("could not read %s: %w", filePath, err)
			}

			src, err := newSourceSchema(path, filePath, buf)
			if err != nil {
				return err
			}
			srcs = append(srcs, src)

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return srcs, nil
}

// loadFromPaths parses schemas from local files and directories.
//
// Directories are walked recursively for .yaml files. A path naming a file is
// read whether or not it has that extension, so a caller can point at a single
// unconventionally named schema, and is treated as living at the top of its own
// parent directory.
func loadFromPaths(paths []string) ([]*sourceSchema, error) {
	var srcs []*sourceSchema

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("could not stat %s: %w", path, err)
		}

		if !info.IsDir() {
			buf, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("could not read %s: %w", path, err)
			}
			src, err := newSourceSchema(filepath.Dir(path), path, buf)
			if err != nil {
				return nil, err
			}
			srcs = append(srcs, src)
			continue
		}

		if err = filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(d.Name(), ".yaml") {
				return nil
			}

			buf, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("could not read %s: %w", filePath, err)
			}

			src, err := newSourceSchema(path, filePath, buf)
			if err != nil {
				return err
			}
			srcs = append(srcs, src)

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return srcs, nil
}

// pkgGroup is the set of schemas that become one generated Go package.
type pkgGroup struct {
	// dir is the output directory, relative to the output root.
	dir string
	// pkg is the Go package name.
	pkg  string
	srcs []*sourceSchema
}

// sources returns the input directories the group was built from, for header
// comments.
func (g *pkgGroup) sources() []string {
	var sources []string
	for _, src := range g.srcs {
		if s := src.source(); !slices.Contains(sources, s) {
			sources = append(sources, s)
		}
	}
	slices.Sort(sources)
	return sources
}

// groupByDir splits sources into one group per source directory, each of which
// tree mode emits as its own package.
//
// A package takes its name from its directory. Schemas sitting directly in an
// input root have no directory of their own, so they take the root's name and
// are written at the top of the output -- that is what lets `-path
// declarative/status` produce a `status` package instead of one named ".".
func groupByDir(srcs []*sourceSchema) []*pkgGroup {
	var groups []*pkgGroup
	index := make(map[string]*pkgGroup)

	for _, src := range srcs {
		dir, pkg := src.dir, pkgNameFromPath(src.dir)
		if dir == "" {
			// Schemas sitting directly in an input root have no directory of
			// their own to be named after, so they take the root's name -- and
			// are written into a directory of that name rather than at the top
			// of the output. Two roots would otherwise put two different
			// packages in one directory, which is not a buildable Go package.
			pkg = pkgNameFromPath(src.root)
			dir = pkg
		}

		// Two input roots can each contribute schemas at dir "", so the name
		// has to be part of the key or they would merge into one package.
		key := dir + "\x00" + pkg
		g, ok := index[key]
		if !ok {
			g = &pkgGroup{dir: dir, pkg: pkg}
			index[key] = g
			groups = append(groups, g)
		}
		g.srcs = append(g.srcs, src)
	}

	return groups
}

// pkgNameFromPath derives a Go package name from a filesystem path, resolving
// "." and ".." against the working directory so they don't leak into the
// generated `package` clause.
func pkgNameFromPath(path string) string {
	base := filepath.Base(filepath.Clean(path))
	if base == "." || base == ".." || base == string(filepath.Separator) {
		abs, err := filepath.Abs(path)
		if err != nil {
			return base
		}
		base = filepath.Base(abs)
	}
	return base
}

// findStatusReport returns the status report envelope, or nil if it wasn't
// loaded.
//
// It is searched for across every input path rather than within the status
// directory, because Apple keeps the envelope under declarative/protocol/ while
// the items it carries live under declarative/status/.
func findStatusReport(srcs []*sourceSchema) *schema.Schema {
	for _, src := range srcs {
		if statustree.IsStatusReport(src.s) {
			return src.s
		}
	}
	return nil
}

// applyStatusTree replaces a package's flat status item schemas with the single
// nested schema that models a real status report.
//
// The flat schemas are dropped rather than kept alongside. Apple describes each
// status item as though it were a standalone document, but `{"device.model.family":
// "Mac"}` appears nowhere in the protocol: items exist only merged into a
// report's StatusItems dictionary, so a per-item struct can neither parse nor
// produce anything real. Dropping them is also required, not merely tidy --
// see the note on statustree.leafKey about sharing PayloadKeys across schemas.
//
// envelope, when present, is grafted around the tree so the package gets a whole
// typed report. If the envelope is itself one of this package's schemas it is
// dropped, since the grafted copy supersedes it and two StatusReport types in
// one package would collide.
func applyStatusTree(schemas []*schema.Schema, envelope *schema.Schema) ([]*schema.Schema, []EncodeOption, error) {
	var items, rest []*schema.Schema
	for _, s := range schemas {
		switch {
		case statustree.IsStatusItem(s):
			items = append(items, s)
		case s == envelope:
			// superseded by the grafted copy below
		default:
			rest = append(rest, s)
		}
	}

	// No status items here, so nothing is superseded and the envelope -- if this
	// is the package that holds it -- stays as it was.
	if len(items) == 0 {
		return schemas, nil, nil
	}

	tree, err := statustree.Transform(items, envelope)
	if err != nil {
		return nil, nil, fmt.Errorf("could not decompose status items: %w", err)
	}

	// The synthesized schema holds all 45-odd status items at once, so the
	// encoder's usual shortest-unique-suffix naming has nothing to qualify
	// names against and collapses them to bare words. Name its types after
	// their full path instead.
	treeOpts := []EncodeOption{
		WithNamerOptions(
			schema.WithNameOverrides(tree.NameOverrides),
			schema.WithFullyQualifiedNames(tree.Schema),
		),
		WithStatusItemTypes(tree.ItemTypes),
	}

	return append(rest, tree.Schema), treeOpts, nil
}

// renderOptions is everything render needs beyond the schemas themselves.
type renderOptions struct {
	pkg string
	// sourceHeader renders the "generated from" comment for the input
	// directories a package was built from.
	sourceHeader func(sources []string) string
	// sources are the input directories, passed to sourceHeader.
	sources []string
	// hash, when set, is emitted as the DeviceManagementGenerateHash constant.
	hash string
	// envelope is the status report schema, if it was loaded. See applyStatusTree.
	envelope *schema.Schema
	reps     replace.Replacements
	encOpts  []EncodeOption
}

// gitSourceHeader renders the "generated from" comment for schemas read out of a
// git repository.
func gitSourceHeader(repoURL, hash string) func([]string) string {
	return func(sources []string) string {
		return fmt.Sprintf("generated from %s:%s/%s", repoURL, hash, strings.Join(sources, ","))
	}
}

// localSourceHeader renders the "generated from" comment for schemas read off
// the local filesystem, where there is no repository or commit to cite.
func localSourceHeader(sources []string) string {
	return fmt.Sprintf("generated from %s", strings.Join(sources, ","))
}

// render encodes schemas into a single Go source file.
func render(schemas []*schema.Schema, out io.Writer, ropts renderOptions) error {
	schemas, treeOpts, err := applyStatusTree(schemas, ropts.envelope)
	if err != nil {
		return err
	}

	f := jen.NewFile(ropts.pkg)
	f.HeaderComment("DO NOT EDIT")
	f.HeaderComment(ropts.sourceHeader(ropts.sources))

	if ropts.hash != "" {
		f.Const().Id("DeviceManagementGenerateHash").Op("=").Lit(ropts.hash)
	}

	encOpts := append([]EncodeOption{WithReplacements(ropts.reps)}, ropts.encOpts...)
	encOpts = append(encOpts, treeOpts...)

	NewEncoder(f, encOpts...).Encode(schema.NewFile(schemas))

	if err = f.Render(out); err != nil {
		return fmt.Errorf("could not render code: %w", err)
	}

	return nil
}

// generateTree writes one package per source directory under output.
func generateTree(srcs []*sourceSchema, output string, ropts renderOptions) error {
	envelope := findStatusReport(srcs)

	for _, g := range groupByDir(srcs) {
		schemas := make([]*schema.Schema, 0, len(g.srcs))
		for _, src := range g.srcs {
			schemas = append(schemas, src.s)
		}

		gOpts := ropts
		gOpts.pkg = g.pkg
		gOpts.envelope = envelope
		gOpts.sources = g.sources()

		buf := new(bytes.Buffer)
		if err := render(schemas, buf, gOpts); err != nil {
			return fmt.Errorf("could not generate package %s: %w", g.pkg, err)
		}

		dir := filepath.Join(output, g.dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("could not create output directory %s: %w", dir, err)
		}

		path := filepath.Join(dir, g.pkg+".go")
		if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("could not write %s: %w", path, err)
		}
	}

	return nil
}

// generatePackage writes every schema into one package.
func generatePackage(srcs []*sourceSchema, out io.Writer, ropts renderOptions) error {
	schemas := make([]*schema.Schema, 0, len(srcs))
	for _, src := range srcs {
		schemas = append(schemas, src.s)
	}

	ropts.envelope = findStatusReport(srcs)

	return render(schemas, out, ropts)
}

// checkout clones the repo at commit and returns it with its resolved hash.
func checkout(repoURL, commit string) (*git.Repository, string, error) {
	repo, err := git.New(repoURL, commit)
	if err != nil {
		return nil, "", fmt.Errorf("could not check out repository: %w", err)
	}

	hash, err := repo.Hash()
	if err != nil {
		return nil, "", fmt.Errorf("could not get hash: %w", err)
	}

	return repo, hash, nil
}

// GenerateFromGit generates Go types from the declarative schemas at the given
// git repo, commit, and paths, using the optional replacements.
//
// One Go package is written per source directory, named for that directory and
// placed at the matching path under output.
func GenerateFromGit(repoURL, commit string, paths []string, reps replace.Replacements, output string, opts ...EncodeOption) error {
	repo, hash, err := checkout(repoURL, commit)
	if err != nil {
		return err
	}

	srcs, err := loadFromGit(repo, paths)
	if err != nil {
		return err
	}

	return generateTree(srcs, output, renderOptions{
		sourceHeader: gitSourceHeader(repoURL, hash),
		hash:         hash,
		reps:         reps,
		encOpts:      opts,
	})
}

// GenerateFromPaths generates Go types from declarative schemas in the given
// local files and directories, using the optional replacements.
//
// One Go package is written per source directory, as with GenerateFromGit.
func GenerateFromPaths(paths []string, reps replace.Replacements, output string, opts ...EncodeOption) error {
	srcs, err := loadFromPaths(paths)
	if err != nil {
		return err
	}

	return generateTree(srcs, output, renderOptions{
		sourceHeader: localSourceHeader,
		reps:         reps,
		encOpts:      opts,
	})
}

// GeneratePackageFromGit generates a single Go package named pkg from every
// declarative schema at the given git repo, commit, and paths, written to out.
//
// Unlike GenerateFromGit, source directories don't become packages -- everything
// is merged into one file. This is the mode for a downstream repo that wants the
// types dropped into a package of its own choosing.
func GeneratePackageFromGit(repoURL, commit string, paths []string, pkg string, reps replace.Replacements, out io.Writer, opts ...EncodeOption) error {
	repo, hash, err := checkout(repoURL, commit)
	if err != nil {
		return err
	}

	srcs, err := loadFromGit(repo, paths)
	if err != nil {
		return err
	}

	return generatePackage(srcs, out, renderOptions{
		pkg:          pkg,
		sourceHeader: gitSourceHeader(repoURL, hash),
		sources:      paths,
		hash:         hash,
		reps:         reps,
		encOpts:      opts,
	})
}

// GeneratePackageFromPaths generates a single Go package named pkg from the
// declarative schemas in the given local files and directories, written to out.
func GeneratePackageFromPaths(paths []string, pkg string, reps replace.Replacements, out io.Writer, opts ...EncodeOption) error {
	srcs, err := loadFromPaths(paths)
	if err != nil {
		return err
	}

	return generatePackage(srcs, out, renderOptions{
		pkg:          pkg,
		sourceHeader: localSourceHeader,
		sources:      paths,
		reps:         reps,
		encOpts:      opts,
	})
}

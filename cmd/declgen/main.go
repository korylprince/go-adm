package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	declarations "github.com/korylprince/go-adm/generator/declarative"
	"github.com/korylprince/go-adm/utils/replace"
)

// defaultGitPath is where Apple keeps the declarative schemas in their repo.
// It only applies to -repo mode; a local run has no sensible default.
const defaultGitPath = "declarative"

type paths []string

func (p *paths) String() string {
	return strings.Join(*p, ",")
}

func (p *paths) Set(v string) error {
	*p = append(*p, v)
	return nil
}

func printHelp() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] [yamlFile [yamlFile [...]]]\n\n", os.Args[0])
	fmt.Fprint(flag.CommandLine.Output(), `Generate Go types for Declarative Device Management (DDM) from Apple's schemas.

Schemas are read from a git repository with -repo/-commit, or from local files
and directories otherwise.

Without -pkg, one Go package is generated per source directory and -out is the
directory to write them under. With -pkg, every schema is merged into that one
package and -out is a file, or stdout when empty.

Flags:
`)
	flag.PrintDefaults()
	fmt.Fprint(flag.CommandLine.Output(), `
Examples:
  generate a package tree from Apple's repo:
    declgen -repo "https://github.com/apple/device-management.git" -commit "b838baacf2e790db729b6ca3f52724adc8bfb96d" -repl ./repls.yaml

  generate a single package from a local checkout:
    declgen -path ./device-management/declarative/status -pkg status -out status.gen.go

`)
}

func run() error {
	var flPaths paths
	flag.Var(&flPaths, "path", "path to a schema directory or file, repeatable. Rooted in the git repo if -repo is given (default \""+defaultGitPath+"\" with -repo)")
	flRepo := flag.String("repo", "", "git repository URL. Omit to read schemas from local paths")
	flCommit := flag.String("commit", "", "git commit. Required with -repo")
	flPkg := flag.String("pkg", "", "Go package name. Setting this merges all schemas into a single package")
	flOut := flag.String("out", "", "output directory, or output file when -pkg is set. Empty means the working directory, or stdout when -pkg is set")
	flRepl := flag.String("repl", "", "path to replacements file")
	flTags := flag.String("tags", "json", "comma-separated struct tag names to generate")
	flReqDef := flag.Bool("reqdef", false, "generate required and default struct tags")
	flag.Usage = printHelp
	flag.Parse()

	// positional arguments are local schema files, as in structgen
	inputs := append(flPaths, flag.Args()...)

	if *flRepo != "" {
		if *flCommit == "" {
			return errors.New("-commit must be specified with -repo")
		}
		if len(flPaths) == 0 {
			inputs = paths{defaultGitPath}
		}
		if len(flag.Args()) > 0 {
			return errors.New("positional arguments name local files and cannot be combined with -repo")
		}
	} else {
		if *flCommit != "" {
			return errors.New("-commit requires -repo")
		}
		if len(inputs) == 0 {
			return errors.New("-path or a yaml file argument must be specified when -repo is not given")
		}
	}

	var (
		repl replace.Replacements
		err  error
	)
	if *flRepl != "" {
		if repl, err = replace.NewReplacementsFromFile(*flRepl); err != nil {
			return fmt.Errorf("could not open replacements from %s: %w", *flRepl, err)
		}
	}

	var opts []declarations.EncodeOption
	if *flReqDef {
		opts = append(opts, declarations.WithRequiredDefault())
	}
	if tags := splitTags(*flTags); len(tags) > 0 {
		opts = append(opts, declarations.WithTags(tags))
	}

	// no -pkg: a package per source directory, written under -out
	if *flPkg == "" {
		out := *flOut
		if out == "" {
			out = "."
		}
		if *flRepo != "" {
			err = declarations.GenerateFromGit(*flRepo, *flCommit, inputs, repl, out, opts...)
		} else {
			err = declarations.GenerateFromPaths(inputs, repl, out, opts...)
		}
		if err != nil {
			return fmt.Errorf("could not generate code: %w", err)
		}
		return nil
	}

	out := io.Writer(os.Stdout)
	if *flOut != "" {
		f, err := os.OpenFile(*flOut, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("could not open %s: %w", *flOut, err)
		}
		defer f.Close()
		out = f
	}

	if *flRepo != "" {
		err = declarations.GeneratePackageFromGit(*flRepo, *flCommit, inputs, *flPkg, repl, out, opts...)
	} else {
		err = declarations.GeneratePackageFromPaths(inputs, *flPkg, repl, out, opts...)
	}
	if err != nil {
		return fmt.Errorf("could not generate code: %w", err)
	}

	return nil
}

func splitTags(s string) []string {
	var tags []string
	for _, tag := range strings.Split(s, ",") {
		if tag = strings.TrimSpace(tag); tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

func main() {
	if err := run(); err != nil {
		printHelp()
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

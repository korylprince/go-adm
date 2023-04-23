package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/korylprince/go-adm/replace"
	yamlschema "github.com/korylprince/go-adm/yamlschema"
)

func printHelp() {
	fmt.Println("usage:", os.Args[0], "-path <path> -pkg <pkg> [options]")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Print("generate schema from local file:\n\t")
	fmt.Println(os.Args[0], "-path path/to/schema.yaml -pkg schema")
	fmt.Print("generate schema from git repo:\n\t")
	fmt.Println(os.Args[0], `-repo "https://github.com/apple/device-management.git" -commit "eb51fb0cb9626cac4717858556912c257a734ce0" -path "docs/schema.yaml" -pkg schema -repl ./repls.yaml -out ./schema.gen.go`)
	fmt.Println()
}

func checkFlags() error {
	flPath := flag.String("path", "", "path to yaml schema. If -repo is given, path is rooted in git repo")
	flPkg := flag.String("pkg", "schema", "Go package name")
	flRepo := flag.String("repo", "", "git repository URL")
	flCommit := flag.String("commit", "", "git commit")
	flOut := flag.String("out", "", "output path. Leave empty for stdout")
	flRepl := flag.String("repl", "", "path to replacements file")
	flag.Parse()

	if *flPath == "" {
		return errors.New("-path must be specified")
	}

	if *flPkg == "" {
		return errors.New("-pkg must be specified")
	}

	if *flRepo != "" && *flCommit == "" {
		return errors.New("if -repo is specified, -commit must be specified")
	}

	var out io.Writer
	if *flOut == "" || *flOut == "-" {
		out = os.Stdout
	} else {
		f, err := os.Create(*flOut)
		if err != nil {
			return fmt.Errorf("could not open %s: %w", *flOut, err)
		}
		defer f.Close()
		out = f
	}
	var (
		buf []byte
		err error
	)

	var repl replace.Replacements
	if *flRepl != "" {
		repl, err = replace.NewReplacementsFromFile(*flRepl)
		if err != nil {
			return fmt.Errorf("could not open replacements from %s: %w", *flRepl, err)
		}
	}

	if *flRepo != "" {
		buf, err = yamlschema.GenerateFromGit(*flRepo, *flCommit, *flPath, *flPkg, repl)
		if err != nil {
			return fmt.Errorf("could not generate schema from git: %w", err)
		}

	} else {
		buf, err = yamlschema.GenerateFromFile(*flPath, *flPkg, repl)
		if err != nil {
			return fmt.Errorf("could not generate schema from %s: %w", *flPath, err)
		}
	}

	if _, err := out.Write(buf); err != nil {
		return fmt.Errorf("could not write output: %w", err)
	}

	return nil
}

func main() {
	if err := checkFlags(); err != nil {
		printHelp()
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

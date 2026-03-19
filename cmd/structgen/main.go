package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	structgen "github.com/korylprince/go-adm/generator/structgen"
	"github.com/korylprince/go-adm/schema"
	"github.com/korylprince/go-adm/utils/replace"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] [yamlFile [yamlFile [...]]]\n\nGenerate Go structs from YAML schema files.\nEither provide local YAML files as arguments or use -repo/-commit/-path to fetch from a git repository.\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	var (
		flRepl   = flag.String("repl", "", "path to replacements file")
		flOut    = flag.String("out", "", "output file. Leave empty for stdout")
		flPkg    = flag.String("pkg", "yamlschema", "Go package name")
		flTags   = flag.String("tags", "json,plist", "tag names to include, comma separated")
		flReqDef = flag.Bool("reqdef", false, "generate required and default struct tags")
		flRepo   = flag.String("repo", "", "git repository URL")
		flCommit = flag.String("commit", "", "git commit")
		flPath   = flag.String("path", "", "path within git repository to YAML schema directory")
	)
	flag.Parse()

	var err error

	var repl replace.Replacements
	if *flRepl != "" {
		repl, err = replace.NewReplacementsFromFile(*flRepl)
		if err != nil {
			log.Fatalf("could not open replacements from %s: %v", *flRepl, err)
		}
	}

	var f *os.File
	if *flOut == "" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(*flOut, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer f.Close()

	tags := strings.Split(*flTags, ",")

	var encOpts []structgen.EncodeOption
	if *flReqDef {
		encOpts = append(encOpts, structgen.WithSchemaEncoderOption(schema.WithRequiredDefault()))
	}

	if *flRepo != "" {
		if *flCommit == "" {
			log.Fatal("-commit must be specified when using -repo")
		}
		if *flPath == "" {
			log.Fatal("-path must be specified when using -repo")
		}
		err = structgen.GenerateFromGit(*flRepo, *flCommit, *flPath, *flPkg, repl, tags, f, encOpts...)
	} else {
		if len(flag.Args()) < 1 {
			fmt.Fprintf(flag.CommandLine.Output(), "no yamlFile names provided\n")
			flag.Usage()
			os.Exit(1)
		}
		err = structgen.GenerateFromFiles(flag.Args(), *flPkg, repl, tags, f, encOpts...)
	}
	if err != nil {
		log.Fatal(err)
	}
}

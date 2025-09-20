package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/korylprince/go-adm/replace"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <yamlFile> [yamlFile [...]]:\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	var (
		flRepl = flag.String("repl", "", "path to replacements file")
		flOut  = flag.String("out", "", "output file. Leave empty for stdout")
		flPkg  = flag.String("pkg", "yamlschema", "Go package name")
		flTags = flag.String("tags", "json,plist", "tag names to include, comma separated")
	)
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintf(flag.CommandLine.Output(), "no yamlFile names provided\n")
		flag.Usage()
		os.Exit(1)
	}

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

	err = GenerateFromFiles(flag.Args(), *flPkg, repl, strings.Split(*flTags, ","), f)
	if err != nil {
		log.Fatal(err)
	}
}

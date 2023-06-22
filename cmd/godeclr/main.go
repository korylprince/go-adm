package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/google/uuid"
	"github.com/korylprince/go-adm/declarations"
	"github.com/korylprince/go-adm/jsonutil"
)

func main() {
	flIdentifier := flag.String("id", "", "declaration identifier (auto-generated UUID if not specified)")
	flServerToken := flag.String("token", "", "declaration ServerToken")
	flType := flag.String("type", "", "declaration type. Use -types to list all supported types")
	flFull := flag.Bool("full", false, "output all fields in the declaration")
	flTypes := flag.Bool("types", false, "list all supported declaration types")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "usage: godeclr -id IDENTIFIER -token TOKEN -type TYPE")
		fmt.Fprintf(flag.CommandLine.Output(), "\n\tgodeclr generates DDM declarations\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flTypes {
		var typs []string
		for typ := range declarations.DeclarationMap {
			typs = append(typs, typ)
		}
		sort.Strings(typs)
		fmt.Println("supported declaration types:")
		for _, typ := range typs {
			fmt.Println("\t" + typ)
		}
		os.Exit(0)
	}

	if *flIdentifier == "" {
		*flIdentifier = uuid.New().String()
	}

	if *flType == "" {
		flag.Usage()
		fmt.Println("-type must be set")
		os.Exit(1)
	}

	if _, ok := declarations.DeclarationMap[*flType]; !ok {
		fmt.Printf("unsupported declaration type %s. Use -types to see all supported types\n", *flType)
	}

	decl, err := declarations.NewFromType(*flType, *flIdentifier, *flServerToken)
	if err != nil {
		fmt.Println("could not generate declaration:", err)
		os.Exit(1)
	}

	var declobj any = decl
	if *flFull {
		payload := jsonutil.FullFields(decl.Payload)
		if err = jsonutil.SetDefaults(payload); err != nil {
			fmt.Println("could not fill out declaration:", err)
			os.Exit(1)
		}

		m := map[string]any{

			"Type":       decl.Type(),
			"Identifier": decl.Identifier,
			"Payload":    payload,
		}
		if decl.ServerToken != "" {
			m["ServerToken"] = decl.ServerToken
		}
		declobj = m
	}

	buf, err := json.MarshalIndent(declobj, "", "\t")
	if err != nil {
		fmt.Println("could not json marshal declaration:", err)
		os.Exit(1)
	}

	fmt.Println(string(buf))
}

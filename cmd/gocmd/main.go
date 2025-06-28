package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/google/uuid"
	"github.com/korylprince/go-adm/commands"
	gencommands "github.com/korylprince/go-adm/commands/commands"
	"github.com/korylprince/go-adm/tagutil"
	"github.com/micromdm/plist"
)

func main() {
	flUUID := flag.String("uuid", "", "command uuid (auto-generated UUID if not specified)")
	flType := flag.String("type", "", "command type. Use -types to list all supported types")
	flFull := flag.Bool("full", false, "output all fields in the command")
	flTypes := flag.Bool("types", false, "list all supported command types")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "usage: gocmd -uuid CMD_UUID -type TYPE")
		fmt.Fprintf(flag.CommandLine.Output(), "\n\tgocmd generates MDM commands\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flTypes {
		var typs []string
		for typ := range gencommands.CommandMap {
			typs = append(typs, typ)
		}
		sort.Strings(typs)
		fmt.Println("supported command types:")
		for _, typ := range typs {
			fmt.Println("\t" + typ)
		}
		os.Exit(0)
	}

	if *flUUID == "" {
		*flUUID = uuid.New().String()
	}

	if *flType == "" {
		flag.Usage()
		fmt.Println("-type must be set")
		os.Exit(1)
	}

	if _, ok := gencommands.CommandMap[*flType]; !ok {
		fmt.Printf("unsupported command type %s. Use -types to see all supported types\n", *flType)
	}

	cmd, err := commands.NewFromType(*flType, *flUUID)
	if err != nil {
		fmt.Println("could not generate command:", err)
		os.Exit(1)
	}

	v, err := cmd.PlistValue()
	if err != nil {
		fmt.Println("could not convert type:", err)
	}

	if *flFull {
		v = tagutil.FullFields(v)
		if err = tagutil.SetDefaults(v); err != nil {
			fmt.Println("could not fill out command:", err)
			os.Exit(1)
		}
	}

	buf, err := plist.MarshalIndent(v, "\t")
	if err != nil {
		fmt.Println("could not plist marshal command:", err)
		os.Exit(1)
	}

	fmt.Println(string(buf))
}

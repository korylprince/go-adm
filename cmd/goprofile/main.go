package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/google/uuid"
	"github.com/korylprince/go-adm/profiles"
	genprofiles "github.com/korylprince/go-adm/profiles/profiles"
	"github.com/korylprince/go-adm/tagutil"
	"github.com/micromdm/plist"
)

func main() {
	flProfileUUID := flag.String("profile-uuid", "", "profile UUID (auto-generated UUID if not specified)")
	flPayloadUUID := flag.String("payload-uuid", "", "payload UUID (auto-generated UUID if not specified)")
	flType := flag.String("type", "", "payload type. Use -types to list all supported types")
	flFull := flag.Bool("full", false, "output all fields in the profile")
	flTypes := flag.Bool("types", false, "list all supported payload types")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "usage: goprofile -type TYPE")
		fmt.Fprintf(flag.CommandLine.Output(), "\n\tgoprofile generates MDM configuration profiles\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *flTypes {
		var typs []string
		for typ := range genprofiles.PayloadMap {
			typs = append(typs, typ)
		}
		sort.Strings(typs)
		fmt.Println("supported payload types:")
		for _, typ := range typs {
			fmt.Println("\t" + typ)
		}
		os.Exit(0)
	}

	if *flProfileUUID == "" {
		*flProfileUUID = uuid.New().String()
	}

	if *flPayloadUUID == "" {
		*flPayloadUUID = uuid.New().String()
	}

	if *flType == "" {
		flag.Usage()
		fmt.Println("-type must be set")
		os.Exit(1)
	}

	if _, ok := genprofiles.PayloadMap[*flType]; !ok {
		fmt.Printf("unsupported payload type %s. Use -types to see all supported types\n", *flType)
		os.Exit(1)
	}

	profile, err := profiles.NewFromType(*flType, *flProfileUUID, *flPayloadUUID)
	if err != nil {
		fmt.Println("could not generate profile:", err)
		os.Exit(1)
	}

	v, err := profile.PlistValue()
	if err != nil {
		fmt.Println("could not convert type:", err)
		os.Exit(1)
	}

	if *flFull {
		v = tagutil.FullFields(v)
		if err = tagutil.SetDefaults(v); err != nil {
			fmt.Println("could not fill out profile:", err)
			os.Exit(1)
		}
	}

	buf, err := plist.MarshalIndent(v, "\t")
	if err != nil {
		fmt.Println("could not plist marshal profile:", err)
		os.Exit(1)
	}

	fmt.Println(string(buf))
}

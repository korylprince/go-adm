package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	commands "github.com/korylprince/go-adm/generator/commands"
	"github.com/korylprince/go-adm/utils/replace"
)

func printHelp() {
	fmt.Println("usage:", os.Args[0], "-repo <repo url> -commit <commit> [options]")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Print("generate schema from git repo:\n\t")
	fmt.Println(os.Args[0], `-repo "https://github.com/apple/device-management.git" -commit "b838baacf2e790db729b6ca3f52724adc8bfb96d" -repl ./repls.yaml`)
	fmt.Println()
}

func checkFlags() error {
	flPath := flag.String("path", "mdm/commands", "path to commands directory. If -repo is given, path is rooted in git repo")
	flRepo := flag.String("repo", "", "git repository URL")
	flCommit := flag.String("commit", "", "git commit")
	flOut := flag.String("out", ".", "output directory. Leave empty for stdout")
	flRepl := flag.String("repl", "", "path to replacements file")
	flTags := flag.String("tags", "json,plist", "comma-separated struct tag names to generate")
	flReqDef := flag.Bool("reqdef", false, "generate required and default struct tags")
	flag.Parse()

	if *flPath == "" {
		return errors.New("-path must be specified")
	}

	if *flRepo == "" {
		return errors.New("-repo must be specified")
	}

	if *flCommit == "" {
		return errors.New("-commit must be specified")
	}

	var (
		repl replace.Replacements
		err  error
	)
	if *flRepl != "" {
		repl, err = replace.NewReplacementsFromFile(*flRepl)
		if err != nil {
			return fmt.Errorf("could not open replacements from %s: %w", *flRepl, err)
		}
	}

	var opts []commands.EncodeOption
	if *flReqDef {
		opts = append(opts, commands.WithRequiredDefault())
	}

	if *flTags != "" {
		var tags []string
		for _, tag := range strings.Split(*flTags, ",") {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}
			tags = append(tags, tag)
		}
		opts = append(opts, commands.WithTags(tags))
	}

	if err = commands.GenerateFromGit(*flRepo, *flCommit, *flPath, repl, *flOut, opts...); err != nil {
		return fmt.Errorf("could not generate code: %w", err)
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

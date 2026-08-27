# Command Reference

This directory contains the CLI tools for `go-adm`. There are two categories:

- **Code generators** (`profilegen`, `cmdgen`, `declgen`, `structgen`, `yamlschemagen`) — generate Go source code from Apple's device management schemas.
- **Runtime tools** (`goprofile`, `gocmd`, `godeclr`) — generate MDM payloads (plist/JSON) from the generated Go types.

## Code Generators

### profilegen

Generates Go profile payload types from [Apple's device management git repo](https://github.com/apple/device-management).

```
Usage:
  -commit string
    	git commit
  -out string
    	output directory. Leave empty for stdout (default ".")
  -path string
    	path to profiles directory. If -repo is given, path is rooted in git repo (default "mdm/profiles")
  -repl string
    	path to replacements file
  -repo string
    	git repository URL
  -reqdef
    	generate required and default struct tags
  -tags string
    	comma-separated struct tag names to generate (default "json,plist")
```

### cmdgen

Generates Go command request/response types from [Apple's device management git repo](https://github.com/apple/device-management).

```
Usage:
  -commit string
    	git commit
  -out string
    	output directory. Leave empty for stdout (default ".")
  -path string
    	path to commands directory. If -repo is given, path is rooted in git repo (default "mdm/commands")
  -repl string
    	path to replacements file
  -repo string
    	git repository URL
  -reqdef
    	generate required and default struct tags
  -tags string
    	comma-separated struct tag names to generate (default "json,plist")
```

### declgen

Generates Go types for Declarative Device Management (DDM) — declarations, the protocol envelopes, and status reports — from [Apple's device management git repo](https://github.com/apple/device-management) or from a local checkout.

```
Usage: declgen [flags] [yamlFile [yamlFile [...]]]

Generate Go types for Declarative Device Management (DDM) from Apple's schemas.

Schemas are read from a git repository with -repo/-commit, or from local files
and directories otherwise.

Without -pkg, one Go package is generated per source directory and -out is the
directory to write them under. With -pkg, every schema is merged into that one
package and -out is a file, or stdout when empty.

Flags:
  -commit string
    	git commit. Required with -repo
  -out string
    	output directory, or output file when -pkg is set. Empty means the working directory, or stdout when -pkg is set
  -path value
    	path to a schema directory or file, repeatable. Rooted in the git repo if -repo is given (default "declarative" with -repo)
  -pkg string
    	Go package name. Setting this merges all schemas into a single package
  -repl string
    	path to replacements file
  -repo string
    	git repository URL. Omit to read schemas from local paths
  -reqdef
    	generate required and default struct tags
  -tags string
    	comma-separated struct tag names to generate (default "json")
```

#### Status items

Apple keys each DDM status item on a dotted path (`device.model.family`), but a
status report from a device nests those segments instead. `declgen` decomposes
the dotted paths into that nested shape, so the generated `StatusItems` type can
parse a real report — see [Status reports](../README.md#ddm-status-reports).

When the status report envelope (`declarative/protocol/statusreport.yaml`) is
among the schemas read, the tree is grafted into it and the status package gains
a fully typed `StatusReport`. Point `-path` at both `declarative/status` and
`declarative/protocol` to get it:

```bash
declgen \
  -path ./device-management/declarative/status \
  -path ./device-management/declarative/protocol \
  -pkg status -reqdef -out status.gen.go
```

Note that `-pkg` merges the envelope into the same package as the tree, so the
grafted `StatusReport` replaces the one generated from `statusreport.yaml`. Unlike
the package-per-directory mode, there is no second, `map[string]any` report type
to fall back on for status items newer than the schema commit.

### structgen

Generates Go structs from one or more YAML schema files. Supports both local files and fetching from a git repository.

```
Usage: structgen [flags] [yamlFile [yamlFile [...]]]

Generate Go structs from YAML schema files.
Either provide local YAML files as arguments or use -repo/-commit/-path to fetch from a git repository.

Flags:
  -commit string
    	git commit
  -out string
    	output file. Leave empty for stdout
  -path string
    	path within git repository to YAML schema directory
  -pkg string
    	Go package name (default "yamlschema")
  -repl string
    	path to replacements file
  -repo string
    	git repository URL
  -reqdef
    	generate required and default struct tags
  -tags string
    	tag names to include, comma separated (default "json,plist")
```

### yamlschemagen

Generates Go code from a single YAML schema file (the root schema). Supports both local files and fetching from a git repository.

```
Usage:
  -commit string
    	git commit
  -out string
    	output path. Leave empty for stdout
  -path string
    	path to yaml schema. If -repo is given, path is rooted in git repo
  -pkg string
    	Go package name (default "schema")
  -repl string
    	path to replacements file
  -repo string
    	git repository URL
```

## Runtime Tools

### goprofile

Generates MDM configuration profiles as plist output.

```
Usage: goprofile -type TYPE

  -full
    	output all fields in the profile
  -payload-uuid string
    	payload UUID (auto-generated UUID if not specified)
  -profile-uuid string
    	profile UUID (auto-generated UUID if not specified)
  -type string
    	payload type. Use -types to list all supported types
  -types
    	list all supported payload types
```

### gocmd

Generates MDM commands as plist output.

```
Usage: gocmd -uuid CMD_UUID -type TYPE

  -full
    	output all fields in the command
  -type string
    	command type. Use -types to list all supported types
  -types
    	list all supported command types
  -uuid string
    	command uuid (auto-generated UUID if not specified)
```

### godeclr

Generates DDM declarations as JSON output.

```
Usage: godeclr -id IDENTIFIER -token TOKEN -type TYPE

  -full
    	output all fields in the declaration
  -id string
    	declaration identifier (auto-generated UUID if not specified)
  -token string
    	declaration ServerToken
  -type string
    	declaration type. Use -types to list all supported types
  -types
    	list all supported declaration types
```

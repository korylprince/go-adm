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

Generates Go declaration types for Declarative Device Management (DDM) from [Apple's device management git repo](https://github.com/apple/device-management).

```
Usage:
  -commit string
    	git commit
  -out string
    	output directory. Leave empty for stdout (default ".")
  -path string
    	path to declarative directory. If -repo is given, path is rooted in git repo (default "declarative")
  -repl string
    	path to replacements file
  -repo string
    	git repository URL
  -reqdef
    	generate required and default struct tags
  -tags string
    	comma-separated struct tag names to generate (default "json")
```

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

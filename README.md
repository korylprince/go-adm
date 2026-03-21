[![Go Reference](https://pkg.go.dev/badge/github.com/korylprince/go-adm.svg)](https://pkg.go.dev/github.com/korylprince/go-adm)

# About

go-adm is a Go code generator and parser for [Apple's Device Management schemas](https://github.com/apple/device-management). It can generate Go structs for MDM commands, configuration profiles, and Declarative Device Management (DDM) declarations directly from Apple's schema repository.

This repo includes pre-generated Go types for all of Apple's schemas, but the generator tools are also provided so you can generate your own Go code independently of this repo.

## Quick Start

Install the generator you need:

```bash
go install github.com/korylprince/go-adm/cmd/profilegen@latest
go install github.com/korylprince/go-adm/cmd/cmdgen@latest
go install github.com/korylprince/go-adm/cmd/declgen@latest
go install github.com/korylprince/go-adm/cmd/structgen@latest
```

### Generate profile types

```bash
profilegen \
  -repo "https://github.com/apple/device-management.git" \
  -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" \
  -out ./profiles \
  -reqdef
```

### Generate command types

```bash
cmdgen \
  -repo "https://github.com/apple/device-management.git" \
  -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" \
  -out ./commands \
  -reqdef
```

### Generate DDM declaration types

```bash
declgen \
  -repo "https://github.com/apple/device-management.git" \
  -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" \
  -out ./declarations \
  -reqdef
```

### Generate generic structs from specific schema files

```bash
structgen \
  -repo "https://github.com/apple/device-management.git" \
  -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" \
  -path "mdm/checkin" \
  -pkg checkin \
  -reqdef \
  -out checkin.gen.go
```

## Tools

go-adm ships with **code generators** that produce Go source from Apple's schemas, and **runtime tools** that generate payloads (plist/JSON) from the generated types.

See [cmd/README.md](cmd/README.md) for full flag reference for every tool.

### Packages

* [**yamlschema**](https://pkg.go.dev/github.com/korylprince/go-adm/yamlschema) — YAML Schema ([JSON Schema](https://json-schema.org/) in YAML) parser. Generic in theory, tested on Apple's *root schema*.

* [**schema**](https://pkg.go.dev/github.com/korylprince/go-adm/schema) — Parser for Apple's device management schemas (commands, profiles, declarations). Parses every schema in [Apple's repo](https://github.com/apple/device-management) into a [Schema AST](https://pkg.go.dev/github.com/korylprince/go-adm/schema#Schema).

### Code Generators

| Tool | Description |
|------|-------------|
| `profilegen` | Generate Go profile payload structs from Apple's repo |
| `cmdgen` | Generate Go command request/response structs from Apple's repo |
| `declgen` | Generate Go DDM declaration structs from Apple's repo |
| `structgen` | Generate Go structs from arbitrary YAML schema files |
| `yamlschemagen` | Generate Go structs for the root schema |

### Runtime Tools

| Tool | Description |
|------|-------------|
| `goprofile` | Output an MDM configuration profile (plist) |
| `gocmd` | Output an MDM command (plist) |
| `godeclr` | Output a DDM declaration (JSON) |

# Future Work

Eventually I'd like to write schema-based validation tools for profiles, commands, declarations, etc. Pull requests are welcome!

# YAML Fork

This project currently uses [a fork](https://github.com/korylprince/go-yaml) of [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml).
I submitted [a PR to the upstream](https://github.com/goccy/go-yaml/pull/360) to support recursive yaml that Apple uses. Ultimately, the upstream implemented recursive yaml their own way that was incompatible with Apple's usage. Eventually I'd like to switch back to the upstream, but the fork works for now.

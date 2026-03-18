# About

This repo contains a WIP Go parser and generator for [Apple's Device Management schemas](https://github.com/apple/device-management).

All of Apple's device management schemas (e.g. commands, profiles, declarative management, etc) are themselves described by [Apple's Device Management *root schema*](https://github.com/apple/device-management/blob/release/docs/schema.yaml).

## Current Features

### [yamlschema](https://pkg.go.dev/github.com/korylprince/go-adm/yamlschema)

yamlschema is a Go package that contains a YAML Schema (YAML version of [JSON Schema](https://json-schema.org/)) parser. It's a generic parser in theory, but only tested on the *root schema*.

### yamlschemagen

`yamlschemagen` is a tool to generate Go structs for the *root schema*, from a file path or directly from a git repo.

### [schema](https://pkg.go.dev/github.com/korylprince/go-adm/schema)

schema is a Go package that contains a parser (built in part by `yamlschemagen`) for Apple's device management schemas (e.g. commands, profiles, declarative management, etc). It is capable of parsing every schema in [Apple's Device Management repo](https://github.com/apple/device-management) into a [Schema AST](https://pkg.go.dev/github.com/korylprince/go-adm/schema#Schema).

### structgen

`structgen` is a tool to generate Go structs from one or more YAML schema files. It uses the schema parser and supports replacements and required/default struct tags.

### profilegen

`profilegen` generates Go profile payload types from [Apple's device management git repo](https://github.com/apple/device-management). It produces struct and map definitions for all MDM configuration profile payload types.

### cmdgen

`cmdgen` generates Go command types from [Apple's device management git repo](https://github.com/apple/device-management). It produces struct definitions for all MDM command request and response types.

### declgen

`declgen` generates Go declaration types for Declarative Device Management (DDM) from [Apple's device management git repo](https://github.com/apple/device-management). It produces struct definitions for activations, assets, configurations, and management declarations.

### goprofile

`goprofile` is a runtime tool that generates MDM configuration profiles as Go code. It can list all supported payload types or generate a profile with a specific payload type in plist format.

### gocmd

`gocmd` is a runtime tool that generates MDM commands as Go code. It can list all supported command types or generate a command with a specific type in plist format.

### godeclr

`godeclr` is a runtime tool that generates DDM declarations as Go code. It can list all supported declaration types or generate a declaration with a specific type in JSON format.

## TODO

* Write schema validators for commands, profiles, and declarative management schemas

# YAML Fork

This project currently uses [a fork](https://github.com/korylprince/go-yaml) of [github.com/goccy/go-yaml](https://github.com/goccy/go-yaml), since the upstream project [can't currently parse Apple's schema files due to recursive aliases](https://github.com/goccy/go-yaml/pull/360). If/when the PR is merged, this project will revert to the upstream repo.

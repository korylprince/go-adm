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

### Generate DDM declaration and status types

```bash
declgen \
  -repo "https://github.com/apple/device-management.git" \
  -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" \
  -out ./declarations \
  -reqdef
```

`declgen` can also read schemas from a local checkout instead of cloning, and
emit everything into a single package of your choosing rather than a package per
directory:

```bash
declgen \
  -path ./device-management/declarative/status \
  -path ./device-management/declarative/protocol \
  -pkg status \
  -reqdef \
  -out status.gen.go
```

`-path` is repeatable, and positional arguments name individual YAML files.

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

## DDM status reports

Apple defines each DDM status item in its own schema file, keyed by a **dotted**
`statusitemtype` such as `device.model.family`. A status report from a device
does not use those dotted keys — it nests them:

```json
{
  "StatusItems": {
    "device": {
      "model": { "family": "Mac" },
      "operating-system": { "version": "15.3.1" }
    }
  },
  "Errors": [],
  "FullReport": true
}
```

`declgen` splits the dotted paths into that nested shape, so
[`generated/declarative/status`](https://pkg.go.dev/github.com/korylprince/go-adm/generated/declarative/status)
parses a whole report:

```go
var report status.StatusReport
if err := json.Unmarshal(body, &report); err != nil {
    return err
}
family := report.StatusItems.Device.Model.Family // *string, nil if not reported
```

Every status item is a pointer with `omitempty`, because reports are
**incremental** — any subset of items may be present. Each generated field
records its dotted identifier in a doc comment, which is the form that appears
in status subscriptions and in `Errors[].StatusItem`.

Two further things the types deliberately don't express, from Apple's protocol:

* `FullReport: false` means array-valued items carry only *changes*. A server
  must merge them rather than replace.
* `_removed: true` marks a deleted array element, in which case only `_removed`
  and `identifier` are populated.

Status items the generated code doesn't know about — ones Apple added after the
pinned schema commit — are dropped on decode. Regenerate against a newer commit
to pick them up, or use
[`protocol.StatusReport`](https://pkg.go.dev/github.com/korylprince/go-adm/generated/declarative/protocol#StatusReport),
whose `StatusItems` is a `map[string]any` and so is lossless but untyped.

## Replacements (`-repl`)

Apple's schema names do not always map cleanly to idiomatic Go identifiers. In
practice that means generated names can have smashed words or non-Go acronym
casing like `Url`, `Id`, or `Comapple`.

The generator commands accept `-repl <file>` to apply regex-based renames while
generating code. Replacements can target generated `field`, `struct`, and
`const` names.

Example `repls.yaml`:

```yaml
"^(.*)Url(.*)$":
  repl: "${1}URL${2}"
  types:
    - "field"
    - "const"

"^(.*)Id([A-Z].*|)$":
  repl: "${1}ID${2}"
  types:
    - "field"
    - "const"
```

Use it by passing `-repl` to a generator:

```bash
profilegen -repo "https://github.com/apple/device-management.git" -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" -repl ./repls.yaml -out ./profiles
cmdgen     -repo "https://github.com/apple/device-management.git" -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" -repl ./repls.yaml -out ./commands
declgen    -repo "https://github.com/apple/device-management.git" -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" -repl ./repls.yaml -out ./declarations
structgen  -repo "https://github.com/apple/device-management.git" -commit "f878dea98fb88293a3686e44bcfb891f8e78f98f" -path "mdm/checkin" -pkg checkin -repl ./repls.yaml -out checkin.gen.go
```

This repo includes replacement files under `generated/` and `schema/` for code generated in this repo. You can use those as a starting point for your own replacements.

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
| `declgen` | Generate Go DDM declaration and status report structs, from Apple's repo or a local checkout |
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

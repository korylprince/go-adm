# ANY Types in Apple Device Management Schemas

Apple's schemas use `key: ANY` (or `key: ANY <descriptor>`) to represent dictionaries
with dynamic keys. The value type of the ANY subkey determines what Go type is generated.

Apple uses several different patterns for ANY types. The following is an AI-assisted analysis of the patterns.

## Patterns

### 1. `key: ANY` + `type: <any>` → `map[string]any`

```yaml
- key: Configuration
  type: <dictionary>
  subkeys:
  - key: ANY
    type: <any>
```

```go
Configuration map[string]any
```

Apple example: `managed.application.configuration.yaml` — app configuration items.

### 2. `key: ANY` + `type: <string>` → `map[string]string`

```yaml
- key: Headers
  type: <dictionary>
  subkeys:
  - key: ANY
    type: <string>
```

```go
Headers map[string]string
```

Apple example: `com.apple.relay.managed.yaml` — custom HTTP header key/value pairs.

### 3. `key: ANY` + `type: <array>` → `map[string][]T`

```yaml
- key: AllowedExtensions
  type: <dictionary>
  subkeys:
  - key: ANY
    type: <array>
    subkeys:
    - key: Items
      type: <string>
```

```go
AllowedExtensions map[string][]string
```

Apple example: `com.apple.system-extension-policy.yaml` — team ID → bundle IDs.

### 4. `key: ANY` + `type: <dictionary>` → `map[string]Struct`

```yaml
- key: Domains
  type: <dictionary>
  subkeys:
  - key: ANY
    type: <dictionary>
    subkeys:
    - key: Forced
      type: <boolean>
    - key: Value
      type: <string>
```

```go
type DomainEntry struct {
    Forced *bool   `json:"Forced,omitempty"`
    Value  *string `json:"Value,omitempty"`
}

Domains map[string]DomainEntry
```

Apple example: `com.apple.ManagedClient.preferences.yaml` — app domain → preference dict.

### 5. `key: ANY <descriptor>` + `type: <dictionary>` → `map[string]Struct`

Identical to pattern 4 for code generation. The descriptor (e.g. "restriction name")
is documentation only.

```yaml
- key: ANY restriction name
  type: <dictionary>
  subkeys:
  - key: value
    type: <boolean>
```

Apple example: `device.restrictions.list.yaml` — restriction name → parameters.

### 6. Top-level `key: ANY` → `type X map[string]T`

When `key: ANY` is the only entry under `payloadkeys:`, the entire payload is a map.

```yaml
payloadkeys:
- key: ANY
  type: <any>
```

```go
type EthernetConfig map[string]any
```

Apple example: `com.apple.firstethernet.managed.yaml` — 802.1X configuration.

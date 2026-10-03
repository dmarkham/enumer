# Enumer

[![GoDoc](https://godoc.org/github.com/dmarkham/enumer?status.svg)](https://pkg.go.dev/github.com/dmarkham/enumer)
[![GitHub Release](https://img.shields.io/github/release/dmarkham/enumer.svg)](https://github.com/dmarkham/enumer/releases)

Enumer generates Go code that adds useful methods to enums (constants with a specific type).

Given this:

```go
//go:generate go run github.com/dmarkham/enumer@latest -type=Pill -json
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
```

`go generate` writes `pill_enumer.go` and you can do this:

```go
fmt.Println(Aspirin)               // Aspirin
p, err := PillString("ibuprofen")  // p == Ibuprofen (case-insensitive)
PillValues()                       // [Placebo Aspirin Ibuprofen Paracetamol]
PillStrings()                      // [Placebo Aspirin Ibuprofen Paracetamol]
Pill(42).IsAPill()                 // false
json.Marshal(Aspirin)              // "Aspirin"
```

Enumer is a drop-in replacement for [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer).
It generates the same `String()` method plus the extras below, so existing code keeps working.

## Install

Enumer needs Go 1.25 or newer.

The recommended setup pins enumer in `go.mod` as a tool dependency, so `go generate` works on a
fresh clone with no separate install:

```sh
go get -tool github.com/dmarkham/enumer@latest
```

```go
//go:generate go tool enumer -type=Pill
```

[examples/gomods](examples/gomods) is a runnable module set up this way.

Two other options:

```sh
# Run on demand from go:generate without touching go.mod.
//go:generate go run github.com/dmarkham/enumer@latest -type=Pill

# Install a binary on your PATH.
go install github.com/dmarkham/enumer@latest
```

Prebuilt binaries for Linux, macOS, and Windows are on the
[releases page](https://github.com/dmarkham/enumer/releases).

## Usage

```
enumer [flags] -type T [directory]
enumer [flags] -type T files...   # files must be in a single package
```

`-type` is required and takes a comma-separated list of type names. Everything else is optional.
Output goes to `<type>_enumer.go` in the source directory, lowercased, unless you pass `-output`.

### Always generated

| Name | What it does |
|---|---|
| `func (i T) String() string` | Name of the value. Unknown values print as `T(42)`. |
| `func TString(s string) (T, error)` | Value from its name. Matches exact name first, then lowercase. |
| `func TValues() []T` | All values, in declaration order. Aliases are skipped. |
| `func TStrings() []string` | All names, in declaration order. |
| `func (i T) IsAT() bool` | True if the value is one of the declared constants. |

### Encoding flags

Each flag adds the methods for one encoding. Combine as many as you need.

| Flag | Methods added | Interface |
|---|---|---|
| `-json` | `MarshalJSON`, `UnmarshalJSON` | `encoding/json` |
| `-text` | `MarshalText`, `UnmarshalText` | `encoding` |
| `-yaml` | `MarshalYAML`, `UnmarshalYAML` | `gopkg.in/yaml.v2` style, also accepted by yaml.v3 |
| `-sql` | `Value`, `Scan` | `database/sql/driver` |
| `-gqlgen` | `MarshalGQL`, `UnmarshalGQL` | [gqlgen](https://gqlgen.com) |

All of them store the enum as its string name. Unmarshaling an unknown name returns an error.

Use `-text` if the enum is a map key you encode as JSON. Without it, `encoding/json` writes the
numeric value as the key.

`Scan` accepts `string`, `[]byte`, or any `fmt.Stringer`. A `nil` value leaves the receiver unchanged.

### Other flags

| Flag | What it does |
|---|---|
| `-values` | Adds `Values() []string`, which [ent](https://entgo.io/docs/schema-fields/#enum-fields) uses for enum fields. |
| `-validate` | Adds `Validate() error`, which returns an error when the value is not a declared constant. |
| `-flag.value` | Adds `Set(string) error` so the type satisfies `flag.Value`. |
| `-pflag.value` | Adds `Set` and `Type() string` so the type satisfies [pflag.Value](https://pkg.go.dev/github.com/spf13/pflag#Value). `Type` returns all names joined by `\|`. |
| `-typederrors` | Wraps conversion errors with `enumerrs.ErrValueInvalid`. See [Typed errors](#typed-errors). |
| `-linecomment` | Uses the constant's trailing line comment as its name. |
| `-comment` | Adds a comment line to the top of the generated file. Repeatable. |
| `-output` | Output file name. |

## Changing the string names

By default the string name is the Go identifier. Three flags adjust it, applied in this order:

1. `-trimprefix` removes a prefix. Pass a comma-separated list to try several. Names without the prefix are left alone.
2. `-transform` rewrites the case. See the table below.
3. `-addprefix` prepends a string.

The result is what `String()` returns, what `TString()` accepts, and what every encoding uses.

### Transforms

Given a constant named `MyTypeValue`:

| `-transform` | Result |
|---|---|
| `noop` (default) | `MyTypeValue` |
| `snake` | `my_type_value` |
| `snake-upper` | `MY_TYPE_VALUE` |
| `kebab` | `my-type-value` |
| `kebab-upper` | `MY-TYPE-VALUE` |
| `dot` | `my.type.value` |
| `dot-upper` | `MY.TYPE.VALUE` |
| `whitespace` | `my type value` |
| `lower` | `mytypevalue` |
| `upper` | `MYTYPEVALUE` |
| `title` | `MyTypeValue` (first letter uppercased, rest unchanged) |
| `title-lower` | `myTypeValue` (first letter lowercased, rest unchanged) |
| `first` | `M` |
| `first-upper` | `M` |
| `first-lower` | `m` |

Word splitting only works from CamelCase. `snake_upper`, `kebab_upper`, `dot_upper`, `first_upper`,
and `first_lower` are accepted as aliases of the hyphenated names.

### Line comments

With `-linecomment`, a constant's trailing comment replaces its name. Constants without one keep
their identifier.

```go
const (
	Monday Day = iota // lunes
	Tuesday
	Friday // viernes
)
```

`Monday.String()` returns `lunes`, `Tuesday.String()` returns `Tuesday`.

## Typed errors

With `-typederrors`, `TString()`, `Validate()` and the unmarshal methods return an error that matches
`enumerrs.ErrValueInvalid` under `errors.Is`. The message still names the bad input.

```go
import "github.com/dmarkham/enumer/enumerrs"

p, err := PillString("Vitamin")
if errors.Is(err, enumerrs.ErrValueInvalid) {
	// "Vitamin does not belong to Pill values"
}
```

This makes `github.com/dmarkham/enumer` a runtime dependency of your module, since the generated
code imports `enumerrs`.

## Examples

Snake-case JSON for an API:

```sh
enumer -type=Status -json -transform=snake
```

Strip a Go-style prefix and store in a database:

```sh
enumer -type=Color -trimprefix=Color -sql -text
```

Several types at once with a custom file name:

```sh
enumer -type=Pill,Day -output=enums_gen.go
```

## History

Enumer started as a fork of Rob Pike's stringer, was extended by
[Álvaro López Espinosa](https://github.com/alvaroloes/enumer), and continues here.
[jsonenums](https://github.com/campoy/jsonenums) inspired the JSON support.

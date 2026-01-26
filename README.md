# Enumer [![GoDoc](https://godoc.org/github.com/ryanfkeepers/enumer-str?status.svg)](https://godoc.org/github.com/ryanfkeepers/enumer-str) [![Go Report Card](https://goreportcard.com/badge/github.com/ryanfkeepers/enumer-str)](https://goreportcard.com/report/github.com/ryanfkeepers/enumer-str) [![GitHub Release](https://img.shields.io/github/release/ryanfkeepers/enumer-str.svg)](https://github.com/ryanfkeepers/enumer-str/releases)

Enumer is a tool to generate Go code that adds useful methods to Go enums (constants with a specific type).
It started as a fork of [Rob Pike’s Stringer tool](https://godoc.org/golang.org/x/tools/cmd/stringer)
maintained by [Álvaro López Espinosa](https://github.com/alvaroloes/enumer). 
This was again forked here as (https://github.com/dmarkham/enumer) picking up where Álvaro left off.
And the current version was forked here as (https://github.com/ryanfkeepers/enumer-str) to support string consts.

This is an experimental fork.  No support is guaranteed.

```
$ enumer-str --help
Enumer is a tool to generate Go code that adds useful methods to Go enums (constants with a specific type).
Usage of enumer:
        Enumer [flags] -type T [directory]
        Enumer [flags] -type T files... # Must be a single package
For more information, see:
        http://godoc.org/github.com/ryanfkeepers/enumer-str
Flags:
  -type string
        comma-separated list of type names; must be set
  -comment value
        comments to include in generated code, can repeat. Default: ""
  -output string
        output file name; default srcdir/<type>_string.go
```


## Generated functions and methods

When Enumer is applied to a type, it will generate:

- The following basic methods/functions:

  - Function `<Type>Values()`: returns a slice with all the values of the enum
  - Method `IsValid()`: returns true only if the current value is among the values of the enum. Useful for validations.

For example, if we have an enum type called `Pill`,

```go
type Pill string

const (
	Placebo       Pill = "placebo"
	Aspirin       Pill = "aspirin"
	Ibuprofen     Pill = "ibuprofen"
	Paracetamol   Pill = "paracetamol"
	Acetaminophen Pill = "acetaminophen"
)
```

executing `enumer -type=Pill -json` will generate a new file with four basic methods and two extra for JSON:

```go
func PillValues() []Pill {
	//...
}

func (i Pill) IsValid() bool {
	//...
}
```

From now on, we can:

```go
// Convert any Pill value to string
var aspirinString string = Aspirin.String()
// (or use it in any place where a Stringer is accepted)
fmt.Println("I need ", Paracetamol) // Will print "I need Paracetamol"

// Get all the values of the string
allPills := PillValues()
fmt.Println(allPills) // Will print [Placebo Aspirin Ibuprofen Paracetamol]

// Check if a value belongs to the Pill enum values
var notAPill Pill = "Erestor"
if notAPill.IsValid() {
	fmt.Println(notAPill, "is not a value of the Pill enum")
}

alsoNotAPill := "Lacrimosa"
if IsAValidPill(alsoNotAPill) {
	fmt.Println(notAPill, "is not a value of the Pill enum")
}
```

## How to use

For a module-aware repo with `enumer` in the `go.mod` file, generation can be called by adding the following to a `.go` source file:

```golang
//go:generate go run github.com/ryanfkeepers/enumer-str -type=YOURTYPE
```

## Inspiring projects

- [Álvaro López Espinosa](https://github.com/alvaroloes/enumer)
- [Stringer](https://godoc.org/golang.org/x/tools/cmd/stringer)
- [jsonenums](https://github.com/campoy/jsonenums)

# Go Modules Sample

Shows how to run enumer from `go generate` with the tool pinned in `go.mod`.

## Steps

1. `go get -tool github.com/dmarkham/enumer@latest` adds a `tool` directive to `go.mod`.
   This module already has one, so you can skip this step.
2. `go generate` creates `pill_enumer.go`.
3. `go run .` prints the enum values and a couple of method results.

The `tool` directive keeps the dependency through `go mod tidy`, so `go generate`
works on a fresh clone without any extra install. It needs Go 1.25 or newer.

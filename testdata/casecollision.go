// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Names that differ only by case must each be found by their exact spelling.
// Before the exact and lower-cased names were split into two maps, the
// lower-cased entry for one constant overwrote the exact entry for another.

package main

import "fmt"

type Casecollision int

const (
	example Casecollision = iota
	Example
	EXAMPLE
)

func main() {
	ck(example, "example")
	ck(Example, "Example")
	ck(EXAMPLE, "EXAMPLE")
	ck(-127, "Casecollision(-127)")

	// Exact spellings win.
	ckLookup("example", example)
	ckLookup("Example", Example)
	ckLookup("EXAMPLE", EXAMPLE)

	// A spelling that matches none exactly still resolves case-insensitively.
	if v, err := CasecollisionString("eXaMpLe"); err != nil || !v.IsACasecollision() {
		panic(fmt.Sprintf("casecollision.go: eXaMpLe -> %v, %v", v, err))
	}

	if _, err := CasecollisionString("nope"); err == nil {
		panic("casecollision.go: expected error for nope")
	}
}

func ck(v Casecollision, str string) {
	if fmt.Sprint(v) != str {
		panic("casecollision.go: " + str)
	}
}

func ckLookup(s string, want Casecollision) {
	got, err := CasecollisionString(s)
	if err != nil || got != want {
		panic(fmt.Sprintf("casecollision.go: CasecollisionString(%q) = %v, %v; want %v", s, got, err, want))
	}
}

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// With -typederrors, Validate() must wrap enumerrs.ErrValueInvalid so
// errors.Is works, while still naming the value. Priority has gaps so the
// map-based String() path is exercised too.

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dmarkham/enumer/enumerrs"
)

type Priority int

const (
	Low      Priority = 10
	Medium   Priority = 20
	High     Priority = 30
	Critical Priority = 99
)

func main() {
	ck(Low, "Low")
	ck(Critical, "Critical")
	ck(15, "Priority(15)")

	// Positive: declared values validate and do not match the sentinel.
	for _, p := range PriorityValues() {
		err := p.Validate()
		if err != nil {
			panic(fmt.Sprintf("validateTypedErrors.go: %v.Validate() = %v, want nil", p, err))
		}
		if errors.Is(err, enumerrs.ErrValueInvalid) {
			panic("validateTypedErrors.go: nil error matched ErrValueInvalid")
		}
	}

	// Negative: a gap value, zero, and out-of-range all match the sentinel and name the value.
	for _, bad := range []Priority{15, 0, -1, 100} {
		err := bad.Validate()
		if err == nil {
			panic(fmt.Sprintf("validateTypedErrors.go: Priority(%d).Validate() = nil, want error", int(bad)))
		}
		if !errors.Is(err, enumerrs.ErrValueInvalid) {
			panic(fmt.Sprintf("validateTypedErrors.go: Priority(%d) error is not ErrValueInvalid: %v", int(bad), err))
		}
		want := fmt.Sprintf("Priority(%d) does not belong to Priority values", int(bad))
		if !strings.Contains(err.Error(), want) {
			panic(fmt.Sprintf("validateTypedErrors.go: got %q, want it to contain %q", err.Error(), want))
		}
	}

	// A different sentinel must not match.
	if errors.Is(Priority(0).Validate(), errors.New("other")) {
		panic("validateTypedErrors.go: matched an unrelated error")
	}
}

func ck(p Priority, str string) {
	if fmt.Sprint(p) != str {
		panic("validateTypedErrors.go: " + str)
	}
}

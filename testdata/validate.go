// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Validate() must return nil for every declared constant and a non-nil
// error naming the value for anything else.

package main

import (
	"fmt"
	"strings"
)

type Color int

const (
	Red Color = iota
	Green
	Blue
)

// Validate must be usable through an interface, which is the point of it.
type validator interface {
	Validate() error
}

var _ validator = Color(0)

func main() {
	ck(Red, "Red")
	ck(Blue, "Blue")
	ck(-1, "Color(-1)")

	// Positive: every declared value validates.
	for _, c := range ColorValues() {
		if err := c.Validate(); err != nil {
			panic(fmt.Sprintf("validate.go: %v.Validate() = %v, want nil", c, err))
		}
	}

	// Negative: below, above, and far out of range all fail and name the value.
	for _, bad := range []Color{-1, 3, 99} {
		err := bad.Validate()
		if err == nil {
			panic(fmt.Sprintf("validate.go: Color(%d).Validate() = nil, want error", int(bad)))
		}
		want := fmt.Sprintf("Color(%d) does not belong to Color values", int(bad))
		if err.Error() != want {
			panic(fmt.Sprintf("validate.go: got %q, want %q", err.Error(), want))
		}
	}

	// Through the interface, same answers.
	var v validator = Color(3)
	if v.Validate() == nil || !strings.Contains(v.Validate().Error(), "Color(3)") {
		panic("validate.go: interface call did not report Color(3)")
	}
	v = Green
	if v.Validate() != nil {
		panic("validate.go: interface call rejected Green")
	}
}

func ck(c Color, str string) {
	if fmt.Sprint(c) != str {
		panic("validate.go: " + str)
	}
}

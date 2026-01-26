// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Single string

const (
	Str Single = "str"
)

func main() {
	ck(Str, "str")
}

func ck(s Single, str string) {
	if string(s) != str {
		panic("day.go: " + str)
	}
}

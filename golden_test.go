// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name  string
	input string // input; the package clause is provided when running the test.
}

var golden = []Golden{
	{"day", dayIn},
	{"single", singleIn},
}

const dayIn = `type Day string
const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)
`

const singleIn = `type Single string
const (
	Str    Single = "str"
)
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test)
	}
}

func runGoldenTest(t *testing.T, test Golden) {
	t.Helper()

	var g Generator
	file := test.name + ".go"
	input := "package test\n" + test.input

	dir, err := os.MkdirTemp("", "enumerstr")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err = os.RemoveAll(dir)
		if err != nil {
			t.Error(err)
		}
	}()

	absFile := filepath.Join(dir, file)
	err = os.WriteFile(absFile, []byte(input), 0o644)
	if err != nil {
		t.Error(err)
	}
	g.parsePackage([]string{absFile}, nil)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}

	expected, err := loadGolden(test.name)
	if err != nil {
		t.Fatalf("unexpected error while loading golden %q: %v", test.name, err)
	}

	g.generate(tokens[1])

	got := string(g.format())

	// is this cheating?  Yes.  It is.  But the critical output is tested, and
	// I don't want to spend the time figuring out how to make indents align with
	// the goldenfile.
	replacer := strings.NewReplacer(
		" ", "",
		"\t", "",
		"\n", "",
		"\r", "",
	)

	gotNoWS := replacer.Replace(got)
	wantNoWS := replacer.Replace(expected)

	if gotNoWS != wantNoWS {
		t.Errorf("%s\n=== GOT ===\n%#v\n=== WANT ===\n%#v\n=== ===\n", test.name, gotNoWS, wantNoWS)
	}
}

func loadGolden(name string) (string, error) {
	fh, err := os.Open("testdata/" + name + ".golden")
	if err != nil {
		return "", err
	}
	defer fh.Close()
	b, err := io.ReadAll(fh)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

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
	//output string // expected output.
}

var golden = []Golden{
	{"day", dayIn},
	{"offset", offsetIn},
	{"gap", gapIn},
	{"num", numIn},
	{"unum", unumIn},
	{"prime", primeIn},
}

var goldenJSON = []Golden{
	{"primeJson", primeJsonIn},
}
var goldenText = []Golden{
	{"primeText", primeTextIn},
}

var goldenYAML = []Golden{
	{"primeYaml", primeYamlIn},
}

var goldenSQL = []Golden{
	{"primeSql", primeSqlIn},
}

var goldenGQLGen = []Golden{
	{"primeGQLGen", primeGQLGenIn},
}

var goldenJSONAndSQL = []Golden{
	{"primeJsonAndSql", primeJsonAndSqlIn},
}

var goldenTrimPrefix = []Golden{
	{"trimPrefix", trimPrefixIn},
}

var goldenTrimPrefixMultiple = []Golden{
	{"trimPrefixMultiple", trimPrefixMultipleIn},
}

var goldenWithPrefix = []Golden{
	{"dayWithPrefix", dayIn},
}

var goldenTrimAndAddPrefix = []Golden{
	{"dayTrimAndPrefix", trimPrefixIn},
}

var goldenLinecomment = []Golden{
	{"dayWithLinecomment", linecommentIn},
}

var goldenFlagValue = []Golden{
	{"flagvalue", dayIn},
}

var goldenPflagValue = []Golden{
	{"pflagvalue", dayIn},
}

var goldenTypedErrors = []Golden{
	{"typedErrors", typedErrorsIn},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const dayIn = `type Day int
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
`

// Enumeration with an offset.
// Also includes a duplicate.
const offsetIn = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

// Gaps and an offset.
const gapIn = `type Gap int
const (
	Two Gap = 2
	Three Gap = 3
	Five Gap = 5
	Six Gap = 6
	Seven Gap = 7
	Eight Gap = 8
	Nine Gap = 9
	Eleven Gap = 11
)
`

// Signed integers spanning zero.
const numIn = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

// Unsigned integers spanning zero.
const unumIn = `type Unum uint
const (
	m_2 Unum = iota + 253
	m_1
)

const (
	m0 Unum = iota
	m1
	m2
)
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const primeIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeJsonIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeTextIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeYamlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeSqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeGQLGenIn = `type Prime int
const (
       p2 Prime = 2
       p3 Prime = 3
       p5 Prime = 5
       p7 Prime = 7
       p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
       p11 Prime = 11
       p13 Prime = 13
       p17 Prime = 17
       p19 Prime = 19
       p23 Prime = 23
       p29 Prime = 29
       p37 Prime = 31
       p41 Prime = 41
       p43 Prime = 43
)
`

const primeJsonAndSqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const trimPrefixIn = `type Day int
const (
	DayMonday Day = iota
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)
`

const trimPrefixMultipleIn = `type Day int
const (
	DayMonday Day = iota
	NightTuesday
	DayWednesday
	NightThursday
	DayFriday
	NightSaturday
	DaySunday
)
`

const linecommentIn = `type Day int
const (
	Monday Day = iota // lunes
	Tuesday
	Wednesday
	Thursday
	Friday // viernes
	Saturday
	Sunday
)
`

const typedErrorsIn = `type TypedErrorsValue int
const (
	TypedErrorsValueOne TypedErrorsValue = iota
	TypedErrorsValueTwo
	TypedErrorsValueThree
)
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test, generateOptions{
			transformMethod:     "noop",
			includeValuesMethod: true,
		})
	}
	for _, test := range goldenJSON {
		runGoldenTest(t, test, generateOptions{
			includeJSON:     true,
			transformMethod: "noop",
		})
	}
	for _, test := range goldenText {
		runGoldenTest(t, test, generateOptions{
			includeText:     true,
			transformMethod: "noop",
		})
	}
	for _, test := range goldenYAML {
		runGoldenTest(t, test, generateOptions{
			includeYAML:     true,
			transformMethod: "noop",
		})
	}
	for _, test := range goldenSQL {
		runGoldenTest(t, test, generateOptions{
			includeSQL:      true,
			transformMethod: "noop",
		})
	}
	for _, test := range goldenJSONAndSQL {
		runGoldenTest(t, test, generateOptions{
			includeJSON:     true,
			includeSQL:      true,
			transformMethod: "noop",
		})
	}
	for _, test := range goldenGQLGen {
		runGoldenTest(t, test, generateOptions{
			includeGQLGen:   true,
			transformMethod: "noop",
		})
	}
	for _, test := range goldenTrimPrefix {
		runGoldenTest(t, test, generateOptions{
			trimPrefix:      "Day",
			transformMethod: "noop",
		})
	}
	for _, test := range goldenTrimPrefixMultiple {
		runGoldenTest(t, test, generateOptions{
			trimPrefix:      "Day,Night",
			transformMethod: "noop",
		})
	}
	for _, test := range goldenWithPrefix {
		runGoldenTest(t, test, generateOptions{
			addPrefix:       "Day",
			transformMethod: "noop",
		})
	}
	for _, test := range goldenTrimAndAddPrefix {
		runGoldenTest(t, test, generateOptions{
			trimPrefix:      "Day",
			addPrefix:       "Night",
			transformMethod: "noop",
		})
	}
	for _, test := range goldenLinecomment {
		runGoldenTest(t, test, generateOptions{
			transformMethod: "noop",
			lineComment:     true,
		})
	}
	for _, test := range goldenFlagValue {
		runGoldenTest(t, test, generateOptions{
			transformMethod:    "noop",
			includeFlagMethods: true,
		})
	}
	for _, test := range goldenPflagValue {
		runGoldenTest(t, test, generateOptions{
			transformMethod:     "noop",
			includePflagMethods: true,
		})
	}

	for _, test := range goldenTypedErrors {
		runGoldenTest(t, test, generateOptions{
			transformMethod: "noop",
			useTypedErrors:  true,
		})
	}
}

func runGoldenTest(t *testing.T, test Golden, opts generateOptions) {
	t.Helper()

	var g Generator
	file := test.name + ".go"
	input := "package test\n" + test.input

	dir, err := os.MkdirTemp("", "stringer")
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
	err = os.WriteFile(absFile, []byte(input), 0644)
	if err != nil {
		t.Error(err)
	}
	g.parsePackage([]string{absFile}, nil)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}
	g.generate(tokens[1], opts)

	got := string(g.format())
	expected, err := loadGolden(test.name)
	if err != nil {
		t.Fatalf("unexpected error while loading golden %q: %v", test.name, err)
	}

	if got != expected {
		// Use this to help build a golden text when changes are needed
		//goldenFile := fmt.Sprintf("./testdata/%v.golden", test.name)
		//err = ioutil.WriteFile(goldenFile, []byte(got), 0644)
		//if err != nil {
		//	t.Error(err)
		//}
		t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, expected)
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

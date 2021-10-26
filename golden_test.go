// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"io/ioutil"
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

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test, false, false, false, false, false, false, true, "", "")
	}
	for _, test := range goldenJSON {
		runGoldenTest(t, test, true, false, false, false, false, false, false, "", "")
	}
	for _, test := range goldenText {
		runGoldenTest(t, test, false, false, false, true, false, false, false, "", "")
	}
	for _, test := range goldenYAML {
		runGoldenTest(t, test, false, true, false, false, false, false, false, "", "")
	}
	for _, test := range goldenSQL {
		runGoldenTest(t, test, false, false, true, false, false, false, false, "", "")
	}
	for _, test := range goldenJSONAndSQL {
		runGoldenTest(t, test, true, false, true, false, false, false, false, "", "")
	}
	for _, test := range goldenGQLGen {
		runGoldenTest(t, test, false, false, false, false, false, true, false, "", "")
	}
	for _, test := range goldenTrimPrefix {
		runGoldenTest(t, test, false, false, false, false, false, false, false, "Day", "")
	}
	for _, test := range goldenTrimPrefixMultiple {
		runGoldenTest(t, test, false, false, false, false, false, false, false, "Day,Night", "")
	}
	for _, test := range goldenWithPrefix {
		runGoldenTest(t, test, false, false, false, false, false, false, false, "", "Day")
	}
	for _, test := range goldenTrimAndAddPrefix {
		runGoldenTest(t, test, false, false, false, false, false, false, false, "Day", "Night")
	}
	for _, test := range goldenLinecomment {
		runGoldenTest(t, test, false, false, false, false, true, false, false, "", "")
	}
}

func runGoldenTest(t *testing.T, test Golden,
	generateJSON, generateYAML, generateSQL, generateText, linecomment, generateGQLGen, generateValuesMethod bool,
	trimPrefix string, prefix string) {

	var g Generator
	file := test.name + ".go"
	input := "package test\n" + test.input

	dir, err := ioutil.TempDir("", "stringer")
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
	err = ioutil.WriteFile(absFile, []byte(input), 0644)
	if err != nil {
		t.Error(err)
	}
	g.parsePackage([]string{absFile}, nil)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}
	g.generate(tokens[1], generateJSON, generateYAML, generateSQL, generateText, generateGQLGen, "noop", trimPrefix, prefix, linecomment, generateValuesMethod)
	got := string(g.format())
	if got != loadGolden(test.name) {
		// Use this to help build a golden text when changes are needed
		//goldenFile := fmt.Sprintf("./testdata/%v.golden", test.name)
		//err = ioutil.WriteFile(goldenFile, []byte(got), 0644)
		//if err != nil {
		//	t.Error(err)
		//}
		t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, loadGolden(test.name))
	}
}

func loadGolden(name string) string {
	fh, err := os.Open("testdata/" + name + ".golden")
	if err != nil {
		return ""
	}
	defer fh.Close()
	b, err := ioutil.ReadAll(fh)
	if err != nil {
		return ""
	}
	return string(b)

}

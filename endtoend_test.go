// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go command is not available on android

// +build !android

package main

import (
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	// GOEXE defines the executable file name suffix (".exe" on Windows, "" on other systems).
	// Must be defined here, cannot be read from ENVIRONMENT variables
	GOEXE = ""
)

func init() {
	// Set GOEXE for Windows platform
	if runtime.GOOS == "windows" {
		GOEXE = ".exe"
	}
}

// This file contains a test that compiles and runs each program in testdata
// after generating the string method for its type. The rule is that for testdata/x.go
// we run stringer -type X and then compile and run the program. The resulting
// binary panics if the String method for X is not correct, including for error cases.

func TestEndToEnd(t *testing.T) {
	dir, err := ioutil.TempDir("", "stringer")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create stringer in temporary directory.
	stringer := filepath.Join(dir, fmt.Sprintf("stringer%s", GOEXE))
	err = run("go", "build", "-o", stringer)
	if err != nil {
		t.Fatalf("building stringer: %s", err)
	}
	// Read the testdata directory.
	fd, err := os.Open("testdata")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()
	names, err := fd.Readdirnames(-1)
	if err != nil {
		t.Fatalf("Readdirnames: %s", err)
	}
	// Generate, compile, and run the test programs.
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		if name == "cgo.go" && !build.Default.CgoEnabled {
			t.Logf("cgo is no enabled for %s", name)
			continue
		}

		// Names are known to be ASCII and long enough.
		var typeName string
		var transformNameMethod string

		switch name {
		case "transform_snake.go":
			typeName = "SnakeCaseValue"
			transformNameMethod = "snake"
		case "transform_snake_upper.go":
			typeName = "SnakeUpperCaseValue"
			transformNameMethod = "snake-upper"
		case "transform_kebab.go":
			typeName = "KebabCaseValue"
			transformNameMethod = "kebab"
		case "transform_kebab_upper.go":
			typeName = "KebabUpperCaseValue"
			transformNameMethod = "kebab-upper"
		case "transform_upper.go":
			typeName = "UpperCaseValue"
			transformNameMethod = "upper"
		case "transform_lower.go":
			typeName = "LowerCaseValue"
			transformNameMethod = "lower"
		case "transform_title.go":
			typeName = "TitleCaseValue"
			transformNameMethod = "title"
		case "transform_first.go":
			typeName = "FirstCaseValue"
			transformNameMethod = "first"
		case "transform_first_upper.go":
			typeName = "FirstUpperCaseValue"
			transformNameMethod = "first-upper"
		case "transform_first_lower.go":
			typeName = "FirstLowerCaseValue"
			transformNameMethod = "first-lower"
		case "transform_whitespace.go":
			typeName = "WhitespaceSeparatedValue"
			transformNameMethod = "whitespace"
		default:
			typeName = fmt.Sprintf("%c%s", name[0]+'A'-'a', name[1:len(name)-len(".go")])
			transformNameMethod = "noop"
		}

		stringerCompileAndRun(t, dir, stringer, typeName, name, transformNameMethod)
	}
}

// stringerCompileAndRun runs stringer for the named file and compiles and
// runs the target binary in directory dir. That binary will panic if the String method is incorrect.
func stringerCompileAndRun(t *testing.T, dir, stringer, typeName, fileName, transformNameMethod string) {
	t.Logf("run: %s %s\n", fileName, typeName)
	source := filepath.Join(dir, fileName)
	err := copy(source, filepath.Join("testdata", fileName))
	if err != nil {
		t.Fatalf("copying file to temporary directory: %s", err)
	}
	stringSource := filepath.Join(dir, typeName+"_string.go")
	// Run stringer in temporary directory.
	err = run(stringer, "-type", typeName, "-output", stringSource, "-transform", transformNameMethod, source)
	if err != nil {
		t.Fatal(err)
	}
	// Run the binary in the temporary directory.
	err = run("go", "run", stringSource, source)
	if err != nil {
		t.Fatal(err)
	}
}

// copy copies the from file to the to file.
func copy(to, from string) error {
	toFd, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFd.Close()
	fromFd, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFd.Close()
	_, err = io.Copy(toFd, fromFd)
	return err
}

// run runs a single command and returns an error if it does not succeed.
// os/exec should have this function, to be honest.
func run(name string, arg ...string) error {
	return runInDir(".", name, arg...)
}

// runInDir runs a single command in directory dir and returns an error if
// it does not succeed.
func runInDir(dir, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

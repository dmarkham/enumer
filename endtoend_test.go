// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go command is not available on android

//go:build !android
// +build !android

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// GOEXE defines the executable file name suffix (".exe" on Windows, "" on other systems).
// Must be defined here, cannot be read from ENVIRONMENT variables
var GOEXE = ""

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
	t.Skip("jjazil")
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("TestEndToEnd panicked: %v", r)
		}
	}()

	dir, err := ioutil.TempDir("", "enumerstr")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create stringer in temporary directory.
	stringer := filepath.Join(dir, fmt.Sprintf("enumerstr%s", GOEXE))
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

		// Names are known to be ASCII and long enough.
		typeName := fmt.Sprintf("%c%s", name[0]+'A'-'a', name[1:len(name)-len(".go")])

		stringerCompileAndRun(t, dir, stringer, typeName, name)
	}
}

// stringerCompileAndRun runs stringer for the named file and compiles and
// runs the target binary in directory dir. That binary will panic if the String method is incorrect.
func stringerCompileAndRun(t *testing.T, dir, stringer, typeName, fileName string) {
	t.Logf("run: %s %s\n", fileName, typeName)
	source := filepath.Join(dir, fileName)
	err := copy(source, filepath.Join("testdata", fileName))
	if err != nil {
		t.Fatalf("copying file to temporary directory: %s", err)
	}
	stringSource := filepath.Join(dir, typeName+"_string.go")
	// Run stringer in temporary directory.
	args := []string{"-type", typeName, "-output", stringSource}
	args = append(args, source)
	err = run(stringer, args...)
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

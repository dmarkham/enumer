// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Day string

const (
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"
	Sunday    Day = "sunday"
)

func main() {
	ck(Monday, "monday")
	ck(Tuesday, "tuesday")
	ck(Wednesday, "wednesday")
	ck(Thursday, "thursday")
	ck(Friday, "friday")
	ck(Saturday, "saturday")
	ck(Sunday, "sunday")
}

func ck(day Day, str string) {
	if string(day) != str {
		panic("day.go: " + str)
	}
}

package main

import "fmt"

type FirstUpperCaseValue int

const (
	male FirstUpperCaseValue = iota
	female
	unknown
)

func main() {
	ck(male, "M")
	ck(female, "F")
	ck(unknown, "U")
	ck(-127, "")
	ck(127, "")
}

func ck(value FirstUpperCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_first_upper.go: " + str)
	}
}

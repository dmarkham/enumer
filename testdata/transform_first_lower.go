package main

import "fmt"

type FirstLowerCaseValue int

const (
	Male FirstLowerCaseValue = iota
	Female
	Unknown
)

func main() {
	ck(Male, "m")
	ck(Female, "f")
	ck(Unknown, "u")
	ck(-127, "")
	ck(127, "")
}

func ck(value FirstLowerCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_first_lower.go: " + str)
	}
}

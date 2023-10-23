package main

import "fmt"

type FirstCaseValue int

const (
	Male FirstCaseValue = iota
	Female
	unknown
)

func main() {
	ck(Male, "M")
	ck(Female, "F")
	ck(unknown, "u")
	ck(-127, "")
	ck(127, "")
}

func ck(value FirstCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_first.go: " + str)
	}
}

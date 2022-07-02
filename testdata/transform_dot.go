package main

import "fmt"

type DotCaseValue int

const (
	DotCaseValueOne DotCaseValue = iota
	DotCaseValueTwo
	DotCaseValueThree
)

func main() {
	ck(DotCaseValueOne, "dot.case.value.one")
	ck(DotCaseValueTwo, "dot.case.value.two")
	ck(DotCaseValueThree, "dot.case.value.three")
	ck(-127, "DotCaseValue(-127)")
	ck(127, "DotCaseValue(127)")
}

func ck(value DotCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_dot.go: " + str)
	}
}

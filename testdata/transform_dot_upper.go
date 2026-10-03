package main

import "fmt"

type DotUpperCaseValue int

const (
	DotUpperCaseValueOne DotUpperCaseValue = iota
	DotUpperCaseValueTwo
	DotUpperCaseValueThree
)

func main() {
	ck(DotUpperCaseValueOne, "DOT.UPPER.CASE.VALUE.ONE")
	ck(DotUpperCaseValueTwo, "DOT.UPPER.CASE.VALUE.TWO")
	ck(DotUpperCaseValueThree, "DOT.UPPER.CASE.VALUE.THREE")
	ck(-127, "DotUpperCaseValue(-127)")
	ck(127, "DotUpperCaseValue(127)")
}

func ck(value DotUpperCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_dot_upper.go: " + str)
	}
}

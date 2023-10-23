package main

import "fmt"

type KebabUpperCaseValue int

const (
	KebabUpperCaseValueOne KebabUpperCaseValue = iota
	KebabUpperCaseValueTwo
	KebabUpperCaseValueThree
)

func main() {
	ck(KebabUpperCaseValueOne, "KEBAB-UPPER-CASE-VALUE-ONE")
	ck(KebabUpperCaseValueTwo, "KEBAB-UPPER-CASE-VALUE-TWO")
	ck(KebabUpperCaseValueThree, "KEBAB-UPPER-CASE-VALUE-THREE")
	ck(-127, "")
	ck(127, "")
}

func ck(value KebabUpperCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_kebab_upper.go: " + str)
	}
}

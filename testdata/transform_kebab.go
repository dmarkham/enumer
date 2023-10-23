package main

import "fmt"

type KebabCaseValue int

const (
	KebabCaseValueOne KebabCaseValue = iota
	KebabCaseValueTwo
	KebabCaseValueThree
)

func main() {
	ck(KebabCaseValueOne, "kebab-case-value-one")
	ck(KebabCaseValueTwo, "kebab-case-value-two")
	ck(KebabCaseValueThree, "kebab-case-value-three")
	ck(-127, "")
	ck(127, "")
}

func ck(value KebabCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_kebab.go: " + str)
	}
}

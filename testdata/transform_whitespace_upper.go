package main

import "fmt"

type WhitespaceUpperSeparatedValue int

const (
	WhitespaceUpperSeparatedValueOne WhitespaceUpperSeparatedValue = iota
	WhitespaceUpperSeparatedValueTwo
	WhitespaceUpperSeparatedValueThree
)

func main() {
	ck(WhitespaceUpperSeparatedValueOne, "WHITESPACE UPPER SEPARATED VALUE ONE")
	ck(WhitespaceUpperSeparatedValueTwo, "WHITESPACE UPPER SEPARATED VALUE TWO")
	ck(WhitespaceUpperSeparatedValueThree, "WHITESPACE UPPER SEPARATED VALUE THREE")
	ck(-127, "WhitespaceUpperSeparatedValue(-127)")
	ck(127, "WhitespaceUpperSeparatedValue(127)")
}

func ck(value WhitespaceUpperSeparatedValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_whitespace_upper.go: " + str)
	}
}

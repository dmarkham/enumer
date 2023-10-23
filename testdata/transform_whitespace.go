package main

import "fmt"

type WhitespaceSeparatedValue int

const (
	WhitespaceSeparatedValueOne WhitespaceSeparatedValue = iota
	WhitespaceSeparatedValueTwo
	WhitespaceSeparatedValueThree
)

func main() {
	ck(WhitespaceSeparatedValueOne, "whitespace separated value one")
	ck(WhitespaceSeparatedValueTwo, "whitespace separated value two")
	ck(WhitespaceSeparatedValueThree, "whitespace separated value three")
	ck(-127, "")
	ck(127, "")
}

func ck(value WhitespaceSeparatedValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_whitespace.go: " + str)
	}
}

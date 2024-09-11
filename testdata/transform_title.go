package main

import "fmt"

type TitleCaseValue int

const (
	titlecasevalueone TitleCaseValue = iota
	titlecasevaluetwo
	titlecasevaluethree
)

func main() {
	ck(titlecasevalueone, "Titlecasevalueone")
	ck(titlecasevaluetwo, "Titlecasevaluetwo")
	ck(titlecasevaluethree, "Titlecasevaluethree")
	ck(-127, "")
	ck(127, "")
}

func ck(value TitleCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transform_title.go: " + str)
	}
}

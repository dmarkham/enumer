package main

import "fmt"

type Thresholdeq int

const (
	req1  Thresholdeq = 2
	req2  Thresholdeq = 4
	req3  Thresholdeq = 6
	req4  Thresholdeq = 8
	req5  Thresholdeq = 10
	req6  Thresholdeq = 12
	req7  Thresholdeq = 14
	req8  Thresholdeq = 16
	req9  Thresholdeq = 18
	req10 Thresholdeq = 20
)

func main() {
	ck(1, "Thresholdeq(1)")
	ck(req1, "req1")
	ck(3, "Thresholdeq(3)")
	ck(req2, "req2")
	ck(5, "Thresholdeq(5)")
	ck(req3, "req3")
	ck(7, "Thresholdeq(7)")
	ck(req4, "req4")
	ck(9, "Thresholdeq(9)")
	ck(req5, "req5")
	ck(11, "Thresholdeq(11)")
	ck(req6, "req6")
	ck(13, "Thresholdeq(13)")
	ck(req7, "req7")
	ck(15, "Thresholdeq(15)")
	ck(req8, "req8")
	ck(17, "Thresholdeq(17)")
	ck(req9, "req9")
	ck(19, "Thresholdeq(19)")
	ck(req10, "req10")
}

func ck(thresholdeq Thresholdeq, str string) {
	if fmt.Sprint(thresholdeq) != str {
		panic("thresholdeq.go: " + str)
	}
}

package main

import "fmt"

type weighter interface {
	weigh() int
}

type scalesV1 struct {
}

func (s scalesV1) weigh() int {
	return 85
}

type scalesV2 struct {
}

func (s scalesV2) weigh() int {
	return 100
}

func conveer(w weighter) int {
	return w.weigh()
}

func main1() {
	var m scalesV1
	fmt.Println(conveer(m))
	var p scalesV2
	msg := conveer(p)
	fmt.Println(msg)
}

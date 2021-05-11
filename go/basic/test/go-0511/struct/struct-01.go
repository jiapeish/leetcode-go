package main

import "fmt"

type ss struct {
	v int
	subsum int
}

func (r ss) Add(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	var a ss
	fmt.Println("the value: ", a.Add(2, 3))
}

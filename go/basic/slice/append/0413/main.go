package main

import "fmt"

func main() {
	s := []int{}
	fmt.Printf("before append, s %v.\n", s)
	o := appendSlice(s)
	fmt.Printf("after append, s %v, o %v.\n", s, o)
}

func appendSlice(input []int) []int {
	for i := 1; i < 10; i++ {
		input = append(input, i)
	}
	fmt.Printf("input: %v\n", input)
	return input
}

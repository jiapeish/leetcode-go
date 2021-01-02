package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 7, 11, 2, 4, 28, 8, 6, 10}
	sort.Ints(s)
	fmt.Println("after sorted,", s)

	a := []int{5, 7, 11, 2, 4, 28, 8, 6, 10}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("after sorted,", a)

	b := []int{5, 7, 11, 2, 4, 28, 8, 6, 10}
	fmt.Println(sort.SearchInts(b, 7))

}
package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5, 6, 7}
	slice := array[1:4]
	fmt.Printf("array %v slice %v\n", array, slice)
	slice[1] = 88
	fmt.Printf("array %v slice %v\n", array, slice)
	//slice1 := slice
	slice1 := append(slice, 77)
	fmt.Printf("array %v slice %v slice1 %v\n", array, slice, slice1)
	slice2 := slice1
	array[5] = 66
	fmt.Printf("array %v slice %v slice1 %v slice2 %v\n", array, slice, slice1, slice2)

	for _, s := range slice {
		addr := &s
		s++
		*addr = s
		fmt.Printf("s %d\t", s)
	}
	fmt.Printf("\n")

	fmt.Printf("array %v slice %v slice1 %v slice2 %v\n", array, slice, slice1, slice2)


}
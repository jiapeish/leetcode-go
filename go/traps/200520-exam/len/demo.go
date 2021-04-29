package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var postcode int = 200201
	a := struct{}{}
	b := struct {
		age  int
		name string
		address []byte
		home *int
	}{1, "alice", []byte{'c', 'd'}, &postcode}
	c := uintptr(5)

	fmt.Println("a", unsafe.Sizeof(a))
	fmt.Println("b.age", unsafe.Sizeof(b.age))
	fmt.Println("b.name", unsafe.Sizeof(b.name))
	fmt.Println("b.address", unsafe.Sizeof(b.address))
	fmt.Println("b.home", unsafe.Sizeof(b.home))
	fmt.Println("b", unsafe.Sizeof(b))
	fmt.Println("c", unsafe.Sizeof(c))

	str1 := "panda"
	fmt.Println("str1:", unsafe.Sizeof(str1))

	arr1 := [...]int{1, 2}
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr1:", unsafe.Sizeof(arr1))
	fmt.Println("arr2:", unsafe.Sizeof(arr2))

	slice1 := []int{1, 2}
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice1:", unsafe.Sizeof(slice1))
	fmt.Println("slice2:", unsafe.Sizeof(slice2))


}



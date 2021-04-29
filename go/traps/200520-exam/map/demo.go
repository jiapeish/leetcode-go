package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	var mapA map[string]int

	fmt.Println("mapA", mapA)
	fmt.Printf("mapA: %v\n", mapA)
	fmt.Println("mapA size:", unsafe.Sizeof(mapA))

	mapA = make(map[string]int, 1)
	mapA["a"] = 1
	mapA["b"] = 2

	fmt.Println("after initialization...")

	fmt.Println("mapA", mapA)
	fmt.Printf("mapA: %v\n", mapA)
	fmt.Println("mapA size:", unsafe.Sizeof(mapA))

	type User struct {
		name string
		age int
	}

	ma := make(map[int]User)
	andes := User{
		name: "andes",
		age: 18,
	}

	ma[1] = andes
	fmt.Println("ma", ma)
	fmt.Println("ma[1]", ma[1])

	andes.age = 19
	fmt.Println("andes", andes)
	fmt.Println("ma", ma)

	fmt.Println("after replacement...")

	ma[1] = andes
	fmt.Println("andes", andes)
	fmt.Println("ma", ma)

	strings.Join()




}

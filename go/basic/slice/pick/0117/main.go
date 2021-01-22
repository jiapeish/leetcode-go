package main

import (
	"fmt"
)

func main() {
	page := make([]byte, 100)
	for i := 0; i < len(page); i++ {
		page[i] = byte(i)
	}
	fmt.Println("page", page)
	mem0 := page[10 : 20]
	fmt.Printf("mem0 %v len %v cap %v\n", mem0, len(mem0), cap(mem0))
	mem := page[10 : 20 : 20]
	fmt.Printf("mem %v len %v cap %v\n", mem, len(mem), cap(mem))
}

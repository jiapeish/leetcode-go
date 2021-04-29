package main

import "fmt"

func a() {
	i := 0
	defer fmt.Printf("%d ,", i)
	i++
	//defer fmt.Printf("%d ,", i)
	return
}

func b() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ,", i)
	}
}

func c() (i int) {
	defer func() {
		i++
	}()
	return i
}


func main() {
	a()
	b()
	fmt.Printf("c: %d ", c())
}

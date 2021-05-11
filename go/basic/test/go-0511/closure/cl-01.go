package main

import "fmt"

func main() {
	var list func()

	for i := 0; i < 1; i++ {
		i := i
		list = func() {
			fmt.Println(i)
		}
	}
	list()

	for i := 0; i < 1; i++ {
		list = func() {
			fmt.Println(i)
		}
	}
	list()
}

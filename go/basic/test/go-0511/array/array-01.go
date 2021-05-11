package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println("in", a)
		}
		a[i] = v + 100
		//fmt.Println("in, v", v)
	}
	fmt.Println("after", a)
}

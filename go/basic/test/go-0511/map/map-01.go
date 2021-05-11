package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	fmt.Println("m", m)
	//m["x"].name = "two"

	a := make([]int, 5)
	a = append(a, 1)
	fmt.Println("a", a)



}

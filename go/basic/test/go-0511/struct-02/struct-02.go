package main

import "fmt"

type Foo struct {
	bar string
}

func main() {
	list := []Foo{
		{"A"},
		{"B"},
	}

	list2 := make([]*Foo, len(list))
	for i, value := range list {
		fmt.Println("i", i, "value", value, "&value", &value)
		list2[i] = &value
		fmt.Println("list2[i]", list2[i])
	}
	fmt.Println(list[0], list[1])
	fmt.Println(list2[0], list2[1])
}

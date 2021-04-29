package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]*Student{"people": {"zhoujielun"}}
	//m["people"] = Student{"zhangsan"}
	m["people"].name = "lisi"

	fmt.Printf("m: %v, m.people: %v", m, m["people"])
}
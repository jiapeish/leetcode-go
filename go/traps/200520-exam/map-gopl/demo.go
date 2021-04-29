package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)

	//ages := map[string]int{
	//	"alice": 31,
	//	"charlie": 33,
	//}

	ages["alice"] = 31
	ages["charlie"] = 33
	ages["dog"] = 10
	ages["apple"] = 20

	fmt.Println(ages)

	ages["bob"]++
	delete(ages, "charlie")
	fmt.Println("bob", ages["bob"])
	fmt.Println("ages", ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println("before sorting, names are", names)
	sort.Strings(names)
	fmt.Println("after sorting, names are", names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	names2 := make([]string, 0, len(ages))
	names3 := make([]string, 0)

	fmt.Println(names2 == nil)
	fmt.Println(names3 == nil)
	fmt.Println(len(names2) == 0)

	fmt.Println("test map below...")

	var ages2 map[string]int
	fmt.Println(ages2 == nil)
	fmt.Println(len(ages2) == 0)
	fmt.Println("ages2 before delete", ages2)
	delete(ages2, "alice")
	fmt.Println("ages2", ages2)
	fmt.Println("ages2 alice", ages2["alice"])

	if age, ok := ages2["bob"]; !ok {
		fmt.Println("ages2 bob not exist", age)
	}
	for k, v := range ages2 {
		fmt.Println(k, v)
	}


}

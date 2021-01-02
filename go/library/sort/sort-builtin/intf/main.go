package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age }) // 按年龄升序排序
	fmt.Println("Sort by age:", people)

	people2 := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.SliceStable(people2, func(i, j int) bool { return people2[i].Age > people2[j].Age }) // 按年龄降序排序
	fmt.Println("Sort2 by age:", people2)

	people3 := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people3, func(i, j int) bool { return people3[i].Age > people3[j].Age }) // 按年龄降序排序
	fmt.Println("Sort3 by age:", people3)
	fmt.Println("Sorted3:",sort.SliceIsSorted(people3,func(i, j int) bool { return people3[i].Age < people3[j].Age }))


	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })   // 升序排序
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}

}

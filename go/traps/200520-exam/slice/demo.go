package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	slice1 := slice[1:3]
	slice2 := slice
	println("Loc1", slice, slice1, slice2)
	fmt.Println("Loc1", slice, slice1, slice2)

	slice1 = append(slice1, 5)
	println("Loc2", slice, slice1, slice2)
	fmt.Println("Loc2", slice, slice1, slice2)

	slice1 = append(slice1, 6)
	println("Loc2-2", slice, slice1, slice2)
	fmt.Println("Loc2-2", slice, slice1, slice2)

	slice1[1] = 5
	println("Loc3", slice, slice1, slice2)
	fmt.Println("Loc3", slice, slice1, slice2)

	slice[2] = 6
	slice2[0] = 7
	println("Loc4", slice, slice1, slice2)
	fmt.Println("Loc4", slice, slice1, slice2)

	slice = append(slice, 8)
	//slice2 = append(slice2, 9)
	println("Loc5", slice, slice1, slice2)
	fmt.Println("Loc5", slice, slice1, slice2)

	slice[0] = 1
	slice2[2] = 3
	println("Loc6", slice, slice1, slice2)
	fmt.Println("Loc6", slice, slice1, slice2)

	s1 := []int {1,2,3,4,5}
	s2 := s1[0:5]
	println("Loc7", s1, s2)
	fmt.Println("Loc7", s1, s2)

	s2 = append(s2, 6)
	s1[3] = 30
	println("Loc8", s1, s2)
	fmt.Println("Loc8", s1, s2)


	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)


	dict := make(map[string][]int)
	dict["a"] = []int{1, 2}
	dict["b"] = []int{5, 6}
	dict["c"] = []int{9, 10}

	for _, slice := range dict {
		for i := 0; i < len(slice); i++ {
			fmt.Println(slice[i])
		}
	}

}
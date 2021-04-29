package main

import "fmt"

func main() {
	sa := []int{1, 2, 3}
	for v := range sa {
		v = 0
		fmt.Println("va", v)
	}
	fmt.Println("sa", sa)

	sb := []int{1, 2, 3}
	for i, _ := range sb {
		sb[i] = i
	}
	fmt.Println("sb", sb)

	sc := []int{1, 2, 3}
	for _, v := range sc {
		v = 0
		fmt.Println("vc", v)
	}
	fmt.Println("sc", sc)

	sd := []int{1, 2, 3}
	for i, v := range sd {
		v = i
		fmt.Println("vd", v)
	}
	fmt.Println("sd", sd)



}

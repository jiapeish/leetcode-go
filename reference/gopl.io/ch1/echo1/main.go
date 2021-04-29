// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import "fmt"

func main() {
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println(s)
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("internal error: %v", r)
			fmt.Printf("err: %v", err)
		}
	}()

	defer func() {
		fmt.Printf("this is 4th panic\n")
		panic("4th panic")
	}()

	defer func() {
		fmt.Printf("this is 3rd panic\n")
		panic("3rd panic")
	}()

	defer func() {
		fmt.Printf("this is 2nd panic\n")
		panic("2nd panic")
	}()

	defer func() {
		fmt.Printf("this is 1st panic\n")
		panic("1st panic")
	}()

	panic("main panic")


}

//!-

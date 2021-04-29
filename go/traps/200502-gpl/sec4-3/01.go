package main

import (
	"fmt"
)

func main()  {
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}

	ages["alice"] = 32

	ages["bob"] = ages["bob"] + 1

	_ = &ages["bob"]

	//_ = &ages["bob"]

	fmt.Printf("ages are: %v", ages)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Count("helloworldworworwwor", "wor")
	fmt.Println("a", a)
}

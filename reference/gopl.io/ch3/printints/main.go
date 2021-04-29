// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74.

// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

//!+
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	s := []int{1, 2, 3}
	m := make(map[int]*int)

	fmt.Printf("s[0] %d s[1] %d s[2] %d\n", &s[0], &s[1], &s[1])

	for ks, vs := range s {
		m[ks] = &vs
		fmt.Printf("m[%d] = %d, %d\n", ks, &vs, vs)
	}

	fmt.Printf("-------\n")

	for km, vm := range m {
		fmt.Printf("m[%d] = %d, %d\n", km, vm, *vm)
	}


}

//!-

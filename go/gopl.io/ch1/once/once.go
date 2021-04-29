package main

import (
	"fmt"
	"sync"
)

var g sync.Once

func Sync() {
	g.Do(func() {
		fmt.Printf("this is Sync...\n")
		Sync()
	})
}

func main() {
	fmt.Printf("Begin main...\n")
	g = sync.Once{}
	Sync()
	fmt.Printf("End main...\n")

}

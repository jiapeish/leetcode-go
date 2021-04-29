package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i, " ,")
		}(i)
	}

	for i := 0; i < 10; i++{
		go func() {
			fmt.Println(i, " ,")
		}()
	}
	time.Sleep(5 * time.Second)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			a <- i
		}
		//close(a)	// 1
	}()

	//time.Sleep(time.Second)

	close(a)	// 2
	fmt.Println("before print value")
	for i := range a {
		fmt.Println("value i", i)
	}

}

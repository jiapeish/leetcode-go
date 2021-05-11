package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int

func Count(lock sync.Mutex) {
	lock.Lock()
	count++
	fmt.Printf("%d\n", count)
	lock.Unlock()
}

func main() {
	lock := sync.Mutex{}
	N := 5
	for i := 0; i < N; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := count
		lock.Unlock()
		fmt.Printf("%d read\n", c)
		runtime.Gosched()
		if c >= N {
			break
		}
	}
}

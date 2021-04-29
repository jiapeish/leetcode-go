package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initialize() {
	fmt.Println("This function is called to initialize")
}

func main() {
	go once.Do(initialize)
	go initialize()
	time.Sleep(time.Second)
}

//func main() {
//	go func() {
//		once.Do(initialize)
//	}()
//	go initialize()
//	time.Sleep(time.Second)
//}

//func doInitialize() {
//	once.Do(initialize)
//}
//func main() {
//	go doInitialize()
//	doInitialize()
//	time.Sleep(time.Second)
//}

//func main() {
//	go once.Do(initialize)
//	go func() {
//		once.Do(initialize)
//	}()
//	time.Sleep(3*time.Second)
//}




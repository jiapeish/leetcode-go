package main

import "fmt"

func main() {
	ch := make(chan bool)
	ch <- true
	fmt.Println("1")
	go func() {
		<- ch
		fmt.Println("2")
	}()

}
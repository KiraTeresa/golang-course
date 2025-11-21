package main

import "fmt"

func main() {
	twoWayChan()
	sendOnlyChan()
	receiveOnlyChan()
}

func twoWayChan() {
	fmt.Println("---- two way chan ----")
	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}

func sendOnlyChan() {
	fmt.Println("---- send only chan ----")
	c := make(chan<- int, 2)
	c <- 44
	c <- 45
	//fmt.Println(<-c) // doesn't work - Invalid operation: <-c (receive from the send-only type chan<- int)
	//fmt.Println(<-c) // doesn't work -Invalid operation: <-c (receive from the send-only type chan<- int)
	fmt.Printf("%T\n", c)
}

func receiveOnlyChan() {
	fmt.Println("---- receive only chan ----")
	c := make(<-chan int, 2)
	//c <- 46 // doesn't work - Invalid operation: c <- 42 (send to the receive-only type <-chan int)
	//c <- 47 // doesn't work - Invalid operation: c <- 42 (send to the receive-only type <-chan int)
	//fmt.Println(<-c) // doesn't work - no values to be pulled off yet
	//fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}

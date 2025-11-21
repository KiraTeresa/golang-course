package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go send(c)

	// receive
	// no go keyword needed, the receive func already waits until it can do what it is supposed to do
	// -> it runs as soon as there is a value to be pulled off of the chan
	// and therefore waits until the send() goroutine has finished its job
	receive(c)

	fmt.Println("about to exit")
}

// send
func send(c chan<- int) {
	c <- 42
}

// receive
func receive(c <-chan int) {
	fmt.Println(<-c)
}

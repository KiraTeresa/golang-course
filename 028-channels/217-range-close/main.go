package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() { // blocks until its closed
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c { // <- pulls off, until the channel is closed (thus the close() in the send func)
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

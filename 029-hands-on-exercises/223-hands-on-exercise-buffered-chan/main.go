package main

import (
	"fmt"
)

/*
*
get this code working:
- a buffered channel

	func main() {
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	}
*/
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

package main

import "fmt"

func main() {
	doesRun()
	doesAlsoRun()
	doesAlsoRun2()
}

func doesNotRun() {
	c := make(chan int)
	c <- 42 // channels BLOCK!
	fmt.Println(<-c)
}

func doesRun() {
	c := make(chan int)

	go func() {
		c <- 42 // only blocks within this goroutine, but main keeps on going
	}()

	fmt.Println(<-c)
}

// successful buffer
func doesAlsoRun() {
	c := make(chan int, 1) // <- buffer
	c <- 44
	fmt.Println(<-c)
}

// unsuccessful buffer
func doesNotRunEither() {
	c := make(chan int, 1)
	c <- 44
	c <- 50
	fmt.Println(<-c)
}

// successful buffer
func doesAlsoRun2() {
	c := make(chan int, 2) // <- buffer
	c <- 11
	c <- 12
	fmt.Println(<-c)
	fmt.Println(<-c)
}

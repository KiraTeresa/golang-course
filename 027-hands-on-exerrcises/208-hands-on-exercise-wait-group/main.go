package main

import (
	"fmt"
	"sync"
)

/*
● in addition to the main goroutine, launch two additional goroutines
	○ each additional goroutine should print something out
● use wait groups to make sure each goroutine finishes before your program exists
*/

var wg sync.WaitGroup

func main() {
	fmt.Println("Main is running")

	wg.Add(2)

	go greeting()
	go favColor("green")

	fmt.Println("Main is done")
	wg.Wait()
}

func greeting() {
	fmt.Println("Good morning")
	wg.Done()
}

func favColor(c string) {
	fmt.Println("My favourite color is", c)
	wg.Done()
}

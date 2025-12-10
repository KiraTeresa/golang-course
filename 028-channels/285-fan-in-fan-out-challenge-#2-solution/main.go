package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// fan out
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)
	c10 := factorial(in)
	c11 := factorial(in)
	c12 := factorial(in)

	// fan in
	y := 0
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12) {
		fmt.Printf("%v: %v\n", y, n)
		y++
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(c ...<-chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(c))
	for _, v := range c {
		go func() {
			for n := range v {
				out <- n
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

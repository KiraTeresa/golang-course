package main

import "math/rand"

// use the "statement statement" idiom to
// - initialize x with a random int between 0 inclusive and 5 exclusive
// - if x is 3 --> print "x is 3"
// run that code 100 times

func main() {
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			println("x is 3 at iteration", i)
		}
	}
}

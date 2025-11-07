package main

import "fmt"

/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func foo() int {
	return 15
}

func bar() (int, string) {
	return 42, "Hello"
}

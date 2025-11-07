package main

import "fmt"

/*
● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6}

	y := foo(x...)
	fmt.Println(y)

	z := bar(x)
	fmt.Println(z)
}

func foo(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func bar(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

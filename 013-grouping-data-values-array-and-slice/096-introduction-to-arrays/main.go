package main

import "fmt"

func main() {
	// array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "World"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	fmt.Printf("The type of a is %T and has length %v\n", a, len(a))
	fmt.Printf("The type of c is %T and has length %v\n", c, len(c))
}

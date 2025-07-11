package main

import "fmt"

func main() {
	a := 42 // short declaration operator can only be used within a function
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	var g int
	fmt.Println(g)

	var h string
	fmt.Println(h)

	g = 15
	fmt.Println(g)
}

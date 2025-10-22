package main

import "fmt"

func main() {
	a, b, c := "Hello world", 42, 10.234

	fmt.Printf("%v is of type %T \n", a, a)
	fmt.Printf("%v is of type %T \n", b, b)
	fmt.Printf("%v is of type %T \n", c, c)
}

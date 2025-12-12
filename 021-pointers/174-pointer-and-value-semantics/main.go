package main

import "fmt"

/*
	value semantics

- passing in a value to addOne and assigning it to a new variable (v)
- v has its own memory address
*/
func addOne(v int) int {
	return v + 1
}

/*
	pointer semantics

- passing in a pointer to an int
*/
func addOneP(v *int) {
	*v++
}

func main() {
	// value semantics
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(a)) // 2
	fmt.Println(a)         // 1

	fmt.Println("--------------")

	// pointer semantics
	b := 1
	fmt.Println(b) // 1
	addOneP(&b)
	fmt.Println(b) // 2
}

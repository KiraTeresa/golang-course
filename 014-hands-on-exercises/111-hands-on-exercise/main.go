package main

import "fmt"

/*
Using the code from the previous example,
use slicing to create the following new slices which are then printed:
- [42,43,44,45,46]
- [47,48,49,50,51]
- [44,45,46,47,48]
- [43,44,45,46,47]
*/
func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := xs[:5]
	b := xs[5:]
	c := xs[2:7]
	d := xs[1:6]

	fmt.Println("values of a: ", a)
	fmt.Println("values of b: ", b)
	fmt.Println("values of c: ", c)
	fmt.Println("values of d: ", d)
}

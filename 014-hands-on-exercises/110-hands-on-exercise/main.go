package main

import "fmt"

/*
Using a composite literal:
- create a slice of type int
- assign these 10 values: 42,43,44,45,46,47,48,49,50,51
Range over the slice and print the values out.
- print out the value and the type
*/

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range xs {
		fmt.Printf("Index %v: value is %v of type %T\n", i, v, v)
	}
}

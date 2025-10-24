package main

import "fmt"

/*
Using a composite literal:
- create an array which holds 5 values of type int
- assign values to each index position.
Range over the array and print the values out.
- print out the value and the type
*/

func main() {
	//var a [5]int
	a := [5]int{}
	a[0] = 10
	a[1] = 11
	a[2] = 12
	a[3] = 13
	a[4] = 14

	for i, v := range a {
		fmt.Printf("%v: value is %v - type %T\n", i, v, v)
	}
}

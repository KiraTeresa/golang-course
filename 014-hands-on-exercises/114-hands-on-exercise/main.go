package main

import "fmt"

/**
Create a slice to store the names of all states in the USA
- use make and append to do this
- Goal: do not have the array that underlies the slice created more than once
Print out:
- len
- cap
- values, along with their index position, without using the range clause
*/

func main() {
	x := make([]string, 0, 50)
	fmt.Printf("initial len of x is %v\n", len(x))
	fmt.Printf("initial cap of x is %v\n", cap(x))
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Printf("len of x is %v\n", len(x))
	fmt.Printf("cap of x is %v\n", cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Printf("%v: value is %v\n", i, x[i])
	}
}

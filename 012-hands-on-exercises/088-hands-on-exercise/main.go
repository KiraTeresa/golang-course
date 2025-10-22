package main

import "fmt"

// below is the code to create a data structure called a slice of ints
// put this code into a program: xi := []int{42, 43, 44, 45, 46, 47}
// use a for range loop to print each value and the index position of each value

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, x := range xi {
		fmt.Printf("Number %v is at index %v\n", x, i)
	}
}

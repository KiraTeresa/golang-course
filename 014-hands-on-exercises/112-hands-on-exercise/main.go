package main

import "fmt"

/*
*
Start with this slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
append to that slice this value: 52
Print out the slice
In one statement append to that slice these values: 53, 54, 55
Print out the slice
Append to the slice this slice:
y := []int{56, 57, 58, 59, 60}
Print out the slice
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("Starting slice: ", x)

	x = append(x, 52)
	fmt.Println("After adding 52: ", x)

	x = append(x, 53, 54, 55)
	fmt.Println("Added three more: ", x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("After adding y: ", x)
}

package main

import "fmt"

func main() {
	xs := []int{42, 43, 44}
	fmt.Println(xs)
	fmt.Println("----")
	xs = append(xs, 55, 56, 57)
	fmt.Println(xs)
}

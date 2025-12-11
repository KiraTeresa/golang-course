package main

import (
	"fmt"
	"golang-course/033-documentation/245-hands-on-exercise/dog"
)

func main() {
	hy := 13
	dy := dog.Years(hy)
	fmt.Printf("The dog is %v years old, which is %v years in dog years\n", hy, dy)
}

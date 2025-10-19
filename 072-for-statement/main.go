package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("First loop, %v is less than 5\n", i)
	}

	y := 0
	for y < 10 {
		fmt.Printf("Second loop, %v is less than 10\n", y)
		y++
	}

	z := 15
	for {
		fmt.Printf("Third loop, %v is less than 20\n", z)
		if z >= 19 {
			break
		}
		z++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("Fourth loop, %v is an even number\n", i)
	}
}

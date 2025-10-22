package main

import "fmt"

func main() {
	x := 0

	for {
		if x > 15 {
			break
		}
		fmt.Printf("x is %v\n", x)
		x++
	}
}

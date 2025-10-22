package main

import "fmt"

func main() {
	x := 0
	for x > 15 {
		fmt.Printf("Looping till x is  15, now it's %v\n", x)
		x++
	}
}

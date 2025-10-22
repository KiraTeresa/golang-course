package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i < 43; i++ {
		x := rand.Intn(5)

		fmt.Printf("Iteration number %v - random int is: %v\t", i, x)

		switch x {
		case 0:
			fmt.Println("Got a 0")
		case 1:
			fmt.Println("One it is")
		case 2:
			fmt.Println("Here is a two")
		case 3:
			fmt.Println("This time it's a three")
		case 4:
			fmt.Println("Four is the highest possible number")
		}
	}
}

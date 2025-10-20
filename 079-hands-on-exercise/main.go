package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("Random number %v was created\n", x)

	if x > 0 && x < 100 {
		fmt.Println("Is between 0 and 100")
	} else if x > 101 && x < 200 {
		fmt.Println("Is between 101 and 200")
	} else if x > 201 {
		fmt.Println("Is between 201 and 250")
	} else {
		fmt.Println("Must be 0")
	}

	switch {
	case x > 0 && x < 100:
		fmt.Println("Is between 0 and 100")
	case x > 101 && x < 200:
		fmt.Println("Is between 101 and 200")
	case x > 201:
		fmt.Println("Is between 201 and 250")
	default:
		fmt.Println("Must be 0")
	}
}

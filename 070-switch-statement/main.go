package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(40)
	fmt.Printf("+++ value of x is %v +++\n", x)

	// no default fallthrough! no need for a keyword like 'break'
	switch {
	case x < 32:
		fmt.Println("x is less than 42")
	case x == 32:
		fmt.Println("x is equal 42")
	case x > 32:
		fmt.Println("x is greater than 42")
	}

	switch x {
	case 20:
		fmt.Println("x is 20")
	case 35:
		fmt.Println("x is 35")
	default:
		fmt.Println("this is the default case")
	}

	// fallthrough does NOT mean, that the next case is also checked
	// if fallthrough is used, the following case will ALSO be EXECUTED!
	switch {
	case x == 39:
		fmt.Println("x is equal 39")
		fallthrough
	case x < 30:
		fmt.Println("x is less than 30")
		fallthrough
	case x < 20:
		fmt.Println("x is less than 20")
		fallthrough
	case x < 10:
		fmt.Println("x is less than 10")
		fallthrough
	case x < 5:
		fmt.Println("x is less than 5")
	}
}

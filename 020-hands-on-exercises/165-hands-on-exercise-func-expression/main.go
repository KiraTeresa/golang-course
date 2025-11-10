package main

import "fmt"

/*
● Assign a func to a variable, then call that func
*/

func main() {
	sayHello := func(s string) {
		fmt.Println("Hello", s)
	}

	sayHello("Jenny")
}

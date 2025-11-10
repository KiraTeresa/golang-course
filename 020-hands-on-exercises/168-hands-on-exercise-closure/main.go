package main

import "fmt"

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:
*/

func main() {
	outer()
}

func outer() {
	fmt.Println("Outer func is running")

	func() {
		x := 1
		for x <= 10 {
			fmt.Printf("Inner func is running %v time\n", x)
			x++
		}
	}()

	fmt.Println("Outer func is exiting")
}

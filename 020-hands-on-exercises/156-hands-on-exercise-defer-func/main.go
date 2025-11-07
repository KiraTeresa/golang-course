package main

import "fmt"

/*
“defer” multiple functions in main
○ show that a deferred func runs after the func containing it exits.
○ determine the order in which the multiple defer funcs run
*/

func main() {
	fmt.Println("Main started")

	defer fmt.Println("I run last")
	defer fmt.Println("I am the middle")
	defer fmt.Println("Runs first after main exits")

	fmt.Println("Second line of main")
}

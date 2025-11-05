package main

import "fmt"

func main() {
	defer foo() // a deferred func runs, after the function surrounding it exits
	bar()
}

func foo() {
	fmt.Println("I am foo")
}

func bar() {
	fmt.Println("This is bar")
}

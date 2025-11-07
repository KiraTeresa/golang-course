package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Kira")
}

func foo() {
	fmt.Println("Foo ran")
}

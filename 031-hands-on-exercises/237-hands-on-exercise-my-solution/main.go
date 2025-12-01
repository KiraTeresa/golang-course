package main

import "fmt"

/**
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”. If you need a hint, here is one: https://go.dev/play/p/L5QWV8-p11
*/

type customErr struct {
	error
}

func main() {
	e := customErr{
		fmt.Errorf("this is a custom error"),
	}

	foo(e)
}

func foo(err error) {
	fmt.Println("foo prints the error:", err)
}

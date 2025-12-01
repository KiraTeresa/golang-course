package main

import "fmt"

/**
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”. If you need a hint, here is one: https://go.dev/play/p/L5QWV8-p11
*/

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the eror: %v", ce.info)
}

func main() {
	e := customErr{
		"this is a custom error info",
	}

	foo(e)
}

func foo(err error) {
	fmt.Println("foo prints the error:", err)

	// type error does not have the info field, only customErr does
	// need to make an assertion, when we want to access the info data field
	fmt.Println("foo ran –", err, "\n", err.(customErr).info)
}

package main

import "fmt"

func main() {
	// escapes to the heap
	// variable shared between main() and Println()
	//a := 1
	//fmt.Println(a)

	// moved to heap
	// variable moved to the heap
	b := 2
	fmt.Println(&b)
}

/*
Escape Analysis:
To see escape analysis in Go, you use the '-gcflags' flag followed by '-m' when running the `go
build` or `go run` command. The '-m' option instructs the compiler to print escape analysis
information.

go run -gcflags '-m' main.go

You can use `-m=2` for more detailed information:
go run -gcflags '-m=2' main.go

This will print out escape analysis and inlining decisions.
Inlining Decisions
Inlining in programming is an optimization that involves replacing a function call site with the
body of the called function.
*/

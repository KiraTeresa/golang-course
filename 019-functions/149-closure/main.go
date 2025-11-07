package main

import "fmt"

func main() {
	f := incrementer()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := incrementer()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

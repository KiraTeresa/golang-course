package main

import "fmt"

func main() {
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("%#b\n", a)
	fmt.Printf("%#b\n", b)
	fmt.Printf("%#b\n", c)
	fmt.Printf("%#b\n", d)
	fmt.Printf("%#b\n", e)
	fmt.Printf("%#b\n", f)

	fmt.Printf("%#x\n", a)
	fmt.Printf("%#x\n", b)
	fmt.Printf("%#x\n", c)
	fmt.Printf("%#x\n", d)
	fmt.Printf("%#x\n", e)
	fmt.Printf("%#x\n", f)
}

package main

import "fmt"

type myNumbers interface {
	int | float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(3, 4))
	fmt.Println(addT(3.3, 4.4))
}

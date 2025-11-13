package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type myNumbers interface {
	constraints.Integer | constraints.Float
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var n myAlias = 42

	fmt.Println(addT(n, 4))
	fmt.Println(addT(3.3, 4.4))
}

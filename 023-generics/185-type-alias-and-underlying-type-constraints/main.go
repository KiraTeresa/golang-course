package main

import "fmt"

type myNumbers interface {
	~int | ~float64 // adding ~ tells this interface to not only include all values of type int/float64, but also all values whos underlying type is of type int/float64 (like our myAlias type below)
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

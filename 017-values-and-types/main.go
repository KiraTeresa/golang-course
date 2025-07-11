package main

import "fmt"

func main() {
	m := 43.742 // <-- type float64
	fmt.Printf("%v of type %T \n", m, m)

	var n float32 = 43.742 // <-- type float32
	fmt.Printf("%v of type %T \n", n, n)

	//	convert float32 into float64
	var z = float64(n)
	fmt.Printf("%v of type %T \n", z, z)
}

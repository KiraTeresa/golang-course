package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// [inclusive:exclusive]
	fmt.Printf("xs - %#v\n", xs[1:4])
	fmt.Println("----")

	// [:exclusive]
	fmt.Printf("xs - %#v\n", xs[:7])
	fmt.Println("----")

	// [:inclusive]
	fmt.Println(xs)
	fmt.Printf("xs - %#v\n", xs[4:])
	fmt.Println("----")

	// [:]
	fmt.Println(xs)
	fmt.Printf("xs - %#v\n", xs[:])
	fmt.Println("----")
}

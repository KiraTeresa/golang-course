package main

import "fmt"

func main() {
	xs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("xs - %#v\n", xs)
	fmt.Println("----")

	xs = append(xs[:4], xs[5:]...) // unfurling a slice with the ...
	fmt.Printf("new xs - %#v\n", xs)
	fmt.Println("----")
}

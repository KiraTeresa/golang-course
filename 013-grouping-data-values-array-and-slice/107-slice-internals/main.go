package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6) // this creates a new slice with its own underlying array
	copy(b, a)          // this copies the values of a into b, which is why changes on either slice WON'T affect the other (values got copied, and not pointed at!)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-------------")

	a[0] = 7 // WON'T change the value at index position 0 of b

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-------------")

	b[1] = 66 // WON'T change the value at index position 1 of a

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-------------")
}

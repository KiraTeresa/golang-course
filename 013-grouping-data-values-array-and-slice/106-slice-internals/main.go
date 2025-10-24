package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a // this is not copying, b is pointing to a, which is why changes on either slice affects the other as well

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-------------")

	a[0] = 7 // will also change the value at index position 0 of b

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-------------")

	b[1] = 66 // will also change the value at index position 1 of a

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-------------")
}

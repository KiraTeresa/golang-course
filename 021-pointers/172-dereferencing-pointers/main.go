package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	y := &x // is storing the memory address of x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)
	fmt.Println(*&x) // dereferencing the address of x

	*y = 43         // assigns 43 to the address, y is pointing to, which is x, and therefore changes the value for both
	fmt.Println(x)  // prints the value
	fmt.Println(y)  // prints the address
	fmt.Println(*y) // prints the value
}

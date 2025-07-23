package main

import "fmt"

func main() {
	var a int8 = 127
	var b uint8 = 255

	fmt.Printf("%v is of type %T \n", a, a)
	fmt.Printf("%v is of type %T \n", b, b)
}

package main

import "fmt"

var a = "I am a var"

const b = "I am a const"

func main() {
	c := "I am a block level var"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

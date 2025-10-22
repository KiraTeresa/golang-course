package main

import "fmt"

func main() {
	var zeroValue int
	shortDeclaration := "I am a short declared value"
	a, b, c := 1, "surprise", 100
	var specificity string
	d, _, f := 22, 33, "fifty"

	fmt.Println(zeroValue)
	fmt.Println(shortDeclaration)
	fmt.Printf("%v \t %v \t %v \n", a, b, c)
	fmt.Println(specificity)
	fmt.Println(d, f)
}

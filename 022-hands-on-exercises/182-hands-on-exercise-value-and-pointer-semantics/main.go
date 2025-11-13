package main

import "fmt"

/*
Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
	○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS
	○ this function will return nothing
● don't use methods
*/

type person struct {
	first string
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func updateName(p *person, s string) {
	p.first = s
}

func main() {
	p1 := person{first: "James"}
	fmt.Println(p1)
	p1 = changeName(p1, "Bond")
	fmt.Println(p1)

	p2 := person{"Jenny"}
	fmt.Println(p2)
	updateName(&p2, "Moneypenny")
	fmt.Println(p2)
}

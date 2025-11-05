package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	// using an anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Printf("%T \n", p1)

	// using a named struct
	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Printf("%T \n", p2)
}

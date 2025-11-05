package main

import "fmt"

/*
Create your own type "person" which will have an underlying type of "struct" so that it can store the following data:
- first name
- last name
- favorite ice cream flavors
Create two values of type person.
Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

type person struct {
	first string
	last  string
	icf   []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		icf:   []string{"Strawberry", "Chocolate"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		icf:   []string{"Caramel", "Nut", "Coffee"},
	}

	fmt.Println(p1)
	for i, v := range p1.icf {
		fmt.Printf("%v: %v\n", i, v)
	}

	fmt.Println(p2)
	for i, v := range p2.icf {
		fmt.Printf("%v: %v\n", i, v)
	}
}

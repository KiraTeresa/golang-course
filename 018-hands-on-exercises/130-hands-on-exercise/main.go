package main

import "fmt"

/*
Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
Access each value in the map.
Print out the values, ranging over the slice.
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

	m := make(map[string]person)
	m[p1.last] = p1
	m[p2.last] = p2

	for k, v := range m {
		fmt.Printf("Values of key '%v' are: %v\n", k, v)

		for i, val := range v.icf {
			fmt.Printf("%v: %v\n", i, val)
		}
	}
}

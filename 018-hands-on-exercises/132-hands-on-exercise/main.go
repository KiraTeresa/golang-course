package main

import "fmt"

/*
Create and use an anonymous struct with these fields:
- first string
- friends map[string]int
- favDrinks []string
*/

func main() {
	p := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny": 27,
			"Todd":  42,
		},
		favDrinks: []string{
			"Martini", "Coffee",
		},
	}

	fmt.Println(p)
	fmt.Println(p.favDrinks)

	for k, v := range p.friends {
		fmt.Printf("%v friend %v is %v years old\n", p.first, k, v)
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

// values must be uppercase in order to allow json.Marshal to access them
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people) // bs = byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

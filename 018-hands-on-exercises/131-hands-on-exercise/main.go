package main

import "fmt"

/*
Create a type engine struct, and include this field: electric bool
Create a type vehicle struct, and include these fields:
- engine
- make
- model
- doors
- color
Create two values of type vehicle: use a composite literal
Print out each of these values.
Print out a single field from each of these values. */

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{electric: true},
		make:   "VW",
		model:  "Polo",
		doors:  3,
		color:  "black",
	}

	v2 := vehicle{
		engine: engine{electric: false},
		make:   "Opel",
		model:  "Astra",
		doors:  5,
		color:  "Red",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.model)
	fmt.Println(v2.model)
}

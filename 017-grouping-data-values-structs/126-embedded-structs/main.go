package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Jenny",
			last:  "Moneypenny",
			age:   27,
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Printf("%T \t %#v\n", sa2, sa2)

	fmt.Println(sa1.first, sa1.last, sa1.age)
}

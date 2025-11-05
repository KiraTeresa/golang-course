package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

// anything that has the same methods as the human interface is also of type human
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p := person{
		first: "Jenny",
		last:  "Moneypenny",
	}

	sa.speak()
	p.speak()
	sa.person.speak()

	saySomething(sa)
	saySomething(p)
}

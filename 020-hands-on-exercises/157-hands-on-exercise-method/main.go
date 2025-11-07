package main

import "fmt"

/*
● Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first
■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %v and my age is %v", p.first, p.age)
}

func main() {
	p := person{
		first: "James",
		age:   37,
	}

	p.speak()
}

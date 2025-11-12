package main

import "fmt"

/*
Consistency!
If some functions on a type use pointer semantics and
others use value semantics, it can lead to confusion. Typically, once a type has a
method with pointer semantics, all methods on that type should have pointer
semantics.
=> this code would not be in production, it's only for learning purposes
*/

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	//youngRun(d1) <-- does not work! Cannot use 'd1' (type dog) as the type youngin Type does not implement 'youngin' as the 'run' method has a pointer receiver

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2) // <-- works!
}

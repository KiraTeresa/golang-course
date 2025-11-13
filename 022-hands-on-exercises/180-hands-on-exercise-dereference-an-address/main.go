package main

import "fmt"

/*
In the provided code, there are variables that store VALUES of TYPE *int and TYPE *string. The
values stored are memory addresses. Using the variables provided, do the following:
● print the VALUE stored in each variable
	○ these will be memory addresses
● print the TYPE of each variable
● print the data stored at memory locations
	○ dereference the stored memory address
*/

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Println("--------- a ---------")
	fmt.Printf("value: %v\n", *a)
	fmt.Printf("address: %v\n", a)
	fmt.Printf("type: %T\n", a)

	fmt.Println("--------- b ---------")
	fmt.Printf("value: %v\n", *b)
	fmt.Printf("address: %v\n", b)
	fmt.Printf("type: %T\n", b)

	fmt.Println("--------- c ---------")
	fmt.Printf("value: %v\n", *c)
	fmt.Printf("address: %v\n", c)
	fmt.Printf("type: %T\n", c)

	fmt.Println("--------- d ---------")
	fmt.Printf("value: %v\n", *d)
	fmt.Printf("address: %v\n", d)
	fmt.Printf("type: %T\n", d)
}

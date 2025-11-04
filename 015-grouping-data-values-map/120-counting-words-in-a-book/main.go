package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
}

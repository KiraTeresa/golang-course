package main

import "fmt"

func main() {
	g := greeting("Johanna")
	fmt.Println(g)
}

func greeting(s string) string {
	return fmt.Sprint("Hello ", s)
}

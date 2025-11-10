package main

import "fmt"

/*
● Create a func
○ which returns a func
■ which returns 42
● assign the returned func to a variable
● call the returned func
● print
*/

func main() {
	x := getFortyTwo()
	fmt.Println(x())
}

func getFortyTwo() func() int {
	return func() int {
		return 42
	}
}

package main

import "fmt"

// below is the code to create a data structure called a map
// m := map[string]int{
// "James": 42,
// "Moneypenny": 32,
// }
// use a for range loop to print each value and the key associated with each value

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("The value for key %v is %v\n", k, v)
	}
}

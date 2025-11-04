package main

import "fmt"

/*
Create a map with a key of type string which is a person's "last_first" name, and a value of type []string which stores their favorite things.
Store three records in your map.
Print out all the values, along with their index position in the slice.
*/

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	m["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range m {
		fmt.Printf("Key %v has the following values stored:\n", k)
		for i, val := range v {
			fmt.Printf("%v: %v\n", i, val)
		}
	}
}

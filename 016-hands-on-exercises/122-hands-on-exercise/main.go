package main

import "fmt"

/*
Using the code from the previous example, add a record to your map.
Now print out the map using the range loop.
*/
func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	m["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range m {
		fmt.Printf("Key %v has the following values stored:\n", k)
		for i, val := range v {
			fmt.Printf("%v: %v\n", i, val)
		}
	}
}

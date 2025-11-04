package main

import "fmt"

/*
Using the code from the previous example, delete a record from your map.
Now print the map out using the range loop.
*/

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	m["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Printf("Key %v has the following values stored:\n", k)
		for i, val := range v {
			fmt.Printf("%v: %v\n", i, val)
		}
	}
}

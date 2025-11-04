package main

import "fmt"

/**
Create a slice of a slice of string ([][]string).
Store the following data in the multidimensional slice:
- "James", "Bond", "Shaken, not stirred"
- "Miss", "Moneypenny", "I'm 008."
Range over the records, then range over the data in each record.
*/

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "I'm 008"}
	z := [][]string{x, y}

	for _, v := range z {
		fmt.Println(v)

		for i, r := range v {
			fmt.Printf("%v: %v\n", i, r)
		}
	}
}

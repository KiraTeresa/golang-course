package main

import "fmt"

// use the code from the previous exercise
// add this code to print a single value stored in the map
// age := m["James"]
// fmt.Println(age)
// now use similar code to check the map for the value "Q" then print that value
// now use the "comma ok" idiom to test whether "Q" is actually stored in the map,
// then print out a statement if it is not stored in the map

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("The value for key %v is %v\n", k, v)
	}

	age := m["James"]
	fmt.Println(age)

	if v, ok := m["Q"]; !ok {
		fmt.Println("Couldn't find Q in the map ", v)
	}
}

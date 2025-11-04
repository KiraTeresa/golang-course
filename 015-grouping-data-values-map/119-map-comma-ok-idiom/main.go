package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}
	fmt.Println("The age of Henry was: ", am["Henry"])

	am["Lucas"] = 28
	am["Steph"] = 27
	fmt.Println(am)

	delete(am, "Lucas")
	fmt.Println(am)

	vi, ok := am["Step"]
	if ok {
		fmt.Println(vi)
	} else {
		fmt.Println("Key doesn't exist")
	}

	if v, ok := am["Todd"]; !ok {
		fmt.Println("Key doesn't exist")
	} else {
		fmt.Println(v)
	}
}

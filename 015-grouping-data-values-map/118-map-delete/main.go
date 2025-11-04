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

	delete(am, "Step")
	fmt.Println(am)
}

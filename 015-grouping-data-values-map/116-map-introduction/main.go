package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was: ", am["Henry"])

	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 27

	fmt.Println(an)
}

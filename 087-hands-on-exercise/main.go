package main

import "fmt"

// create a loop that runs 5 times
// nest a loop within the first loop that also prints 5 times
// print something in each loop to illustrate what is occurring

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Starting OUTER loop no %v\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("INNER loop - %v\n", j)
		}
	}
}

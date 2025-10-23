package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)"}
	fmt.Printf("Length of the slice is %v and its values are: %v\n", len(xs), xs)

	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])
}

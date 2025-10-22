package main

import "fmt"

// use a slice literal to store these elements in a slice,
// then also print out the slice and the length of the slice
//"Almond Biscotti Café",
//"Banana Pudding ",
//"Balsamic Strawberry (GF)",
//"Bittersweet Chocolate (GF)",
//"Blueberry Pancake (GF)",
//"Bourbon Turtle (GF)",
//"Browned Butter Cookie Dough",
//"Chocolate Covered Black Cherry (GF)",
//"Chocolate Fudge Brownie",
//"Chocolate Peanut Butter (GF)",
//"Coffee (GF)",
//"Cookies & Cream",
//"Fresh Basil (GF)",
//"Garden Mint Chip (GF)",
//"Lavender Lemon Honey (GF)",
//"Lemon Bar",
//"Madagascar Vanilla (GF)",
//"Matcha (GF)",
//"Midnight Chocolate Sorbet (GF, V)",
//"Non-Dairy Chocolate Peanut Butter (GF, V)",
//"Non-Dairy Coconut Matcha (GF, V)",
//"Non-Dairy Sunbutter Cinnamon (GF, V)",
//"Orange Cream (GF) ",
//"Peanut Butter Cookie Dough",
//"Raspberry Sorbet (GF, V)",
//"Salty Caramel (GF)",
//"Slate Slate Different",
//"Strawberry Lemonade Sorbet (GF, V)",
//"Vanilla Caramel Blondie",
//"Vietnamese Cinnamon (GF)",
//"Wolverine Tracks (GF)"
// if you can, try to get a "for range" loop working on the slice

func main() {
	xs := []string{"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}
	fmt.Printf("Length of the slice is %v and its values are: %v\n", len(xs), xs)

	for i, v := range xs {
		fmt.Printf("%v: %v\n", i, v)
	}
}

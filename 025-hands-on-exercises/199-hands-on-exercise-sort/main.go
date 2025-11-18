package main

import (
	"fmt"
	"sort"
)

/**
Starting with this code, sort the []int and []string for each person.
*/

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	//slices.SortFunc(xs, func(a, b string) int {
	//	return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	//})
	sort.Strings(xs)
	fmt.Println(xs)
}

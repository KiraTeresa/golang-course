// Package someMath provides functions to run simple mathematics equations
package someMath

// Sum adds an unlimited number of values of type int
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}

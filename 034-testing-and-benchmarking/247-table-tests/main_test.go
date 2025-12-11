package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{13, 1, 25}, 39},
		test{[]int{1, 1}, 2},
		test{[]int{143, 31, 2}, 176},
		test{[]int{3, 6, 5}, 14},
		test{[]int{4, 7, 9}, 20},
	}

	for _, v := range tests {
		sum := mySum(v.data...)
		if sum != v.answer {
			t.Errorf("Expected %v, but got %v\n", v.answer, sum)
		}
	}

}

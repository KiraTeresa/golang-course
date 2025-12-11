package main

import "testing"

func TestMySum(t *testing.T) {
	sum := mySum(2, 3)
	if sum != 5 {
		t.Error("Expected 5, but got", sum)
	}
}

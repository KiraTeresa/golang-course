package main

import "testing"

func TestSum(t *testing.T) {
	s := Sum(2, 3)
	if s != 5 {
		t.Errorf("Sum(2, 3) = %d; expected 5", s)
	}
}

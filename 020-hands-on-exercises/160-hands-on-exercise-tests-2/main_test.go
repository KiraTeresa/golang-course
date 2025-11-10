package main

import "testing"

func TestDoMath(t *testing.T) {
	total := doMath(42, 16, add)
	if total != 58 {
		t.Errorf("Result was incorrect, got %d, want: %d", total, 58)
	}

	result := doMath(42, 16, subtract)
	if result != 26 {
		t.Errorf("Result was incorrect, got %d, want: %d", total, 26)
	}
}

func TestAdd(t *testing.T) {
	total := add(10, 7)
	if total != 17 {
		t.Errorf("Result was incorrect, got %d, want: %d", total, 17)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(10, 7)
	if total != 3 {
		t.Errorf("Result was incorrect, got %d, want: %d", total, 3)
	}
}

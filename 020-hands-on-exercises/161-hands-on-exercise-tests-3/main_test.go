package main

import "testing"

func TestGreeting(t *testing.T) {
	g := greeting("Hannah")
	if g != "Hello Hannah" {
		t.Errorf("Incorrect greeting, expected %v, got %v", "Hello Hanna", g)
	}
}

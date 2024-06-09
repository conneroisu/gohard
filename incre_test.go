package main

import (
	"testing"
)

// TestIncre tests the incre function
func TestIncre(t *testing.T) {
	num := 1
	expected := 2
	output := incre(num)
	if output != expected {
		t.Errorf("Expected %q got %q", expected, output)
	}
}

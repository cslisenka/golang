package main

import "testing"

// This is unit-test for sum function
// We run it by calling go test inside this package
func TestSum(t *testing.T) {
	result := Sum(1, 2)
	if result != 3 {
		t.Error("Expected 3, got", result)
	}
}

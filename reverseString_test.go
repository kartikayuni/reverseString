package main

import "testing"

func TestReverse(t *testing.T) {
	var1 := []byte{'a', 'b', 'c'}
	reverseString(var1)
	expected := "cba"

	if string(var1) != expected {
		t.Errorf("Test failed, expected: %s, got %s\n ", expected, string(var1))
	}
}

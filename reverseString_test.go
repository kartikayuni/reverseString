package main

import "testing"

type TestCase struct {
	Input    []byte
	Expected string
}

var tc []TestCase = []TestCase{
	{Input: []byte{'a', 'b', 'c'}, Expected: "cba"},
	{Input: []byte{'a', 'b', 'c', 'd'}, Expected: "dcba"},
}

func TestPerverse(t *testing.T) {
	for _, tst := range tc {
		reverseString(tst.Input)
		if string(tst.Input) != tst.Expected {
			t.Errorf("Expected %s, got %s\n", tst.Expected, tst.Input)
		}
	}
}

// Below is more sensible variation:
// prioritize readability of your test cases

type TestCase2 struct {
	Input    string
	Expected string
}

var tc2 []TestCase2 = []TestCase2{
	{Input: "abc", Expected: "cba"},
	{Input: "abcd", Expected: "dcba"},
	{Input: "Hello world", Expected: "dlrow olleH"},
	{Input: "A", Expected: "A"},
}

func TestReverse(t *testing.T) {
	for _, tst := range tc2 {
		testInput := []byte(tst.Input)
		reverseString(testInput)
		if string(testInput) != tst.Expected {
			t.Errorf("Expected %s, got %s\n", tst.Expected, tst.Input)
		}
	}
}

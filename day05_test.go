package main

import (
	"strings"
	"testing"
)

func TestDay5a(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	expect := 5
	result := day5a(strings.NewReader(input))
	if expect != result {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay5a() {
	Day5a()

	// Output: day 5a: 6311
}

func TestDay5b(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	expect := 12
	result := day5b(strings.NewReader(input))
	if expect != result {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay5b() {
	Day5b()

	// Output: day 5b: 19929
}

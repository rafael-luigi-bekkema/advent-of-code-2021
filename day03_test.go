package main

import (
	"strings"
	"testing"
)

func TestDay3a(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	expect := 198
	result := day3a(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay3a() {
	Day3a()
	// Output: day 3a: 3847100
}

func TestDay3b(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	expect := 230
	result := day3b(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay3b() {
	Day3b()
	// Output: day 3b: 4105235
}

package main

import (
	"testing"
)

func TestDay9a(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	expect := 15
	result := day9a([]byte(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay9a() {
	Day9a()
	// Output: day 9a: 489
}

func TestDay9b(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	expect := 1134
	result := day9b([]byte(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay9b() {
	Day9b()
	// Output: day 9b: 1056330
}

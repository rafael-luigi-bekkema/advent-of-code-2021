package main

import "testing"

func TestDay17a(t *testing.T) {
	input := `target area: x=20..30, y=-10..-5`
	expect := 45
	result := day17a(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay17a() {
	Day17a()

	// Output: day 17a: 5995
}

func TestDay17b(t *testing.T) {
	input := `target area: x=20..30, y=-10..-5`
	expect := 112
	result := day17b(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay17b() {
	Day17b()

	// Output: day 17b: 3202
}

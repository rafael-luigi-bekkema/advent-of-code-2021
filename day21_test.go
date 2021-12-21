package main

import "testing"

func TestDay21a(t *testing.T) {
	startPos1 := 4
	startPos2 := 8
	expect := 739785

	result := day21a(startPos1, startPos2)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay21a() {
	Day21a()

	// Output: day 21a: 853776
}

func TestDay21b(t *testing.T) {
	startPos1 := 4
	startPos2 := 8
	expect := 444356092776315

	result := day21b(startPos1, startPos2)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

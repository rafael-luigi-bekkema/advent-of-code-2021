package main

import "testing"

func TestDay7a(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expect := 37
	result := day7a(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay7a() {
	Day7a()
	// Output: day 7a: 356992
}

func TestDay7b(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expect := 168
	result := day7b(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay7b() {
	Day7b()
	// Output: day 7b: 101268110
}

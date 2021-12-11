package main

import (
	"testing"
)

func TestDay11a(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`
	expect := 1656
	result := day11a([]byte(input), 100)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay11a() {
	Day11a()
	// Output: day 11a: 1683
}

func TestDay11b(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`
	expect := 195
	result := day11b([]byte(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay11b() {
	Day11b()
	// Output: day 11b: 788
}

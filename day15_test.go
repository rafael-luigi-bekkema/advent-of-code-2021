package main

import (
	"strings"
	"testing"
)

func TestDay15a(t *testing.T) {
	input := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
	expect := 40
	result := day15a(day15parseInput(strings.NewReader(input)))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay15a() {
	Day15a()
	// Output: day 15a: 540
}

func TestDay15b(t *testing.T) {
	input := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
	expect := 315
	result := day15b(day15parseInput(strings.NewReader(input)))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay15b() {
	Day15b()
	// Output: day 15b: 2879
}

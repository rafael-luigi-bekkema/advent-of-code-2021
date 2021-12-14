package main

import (
	"strings"
	"testing"
)

func TestDay14a(t *testing.T) {
	input := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
	steps := 10
	expect := 1588
	result := day14b(day14parseInput(strings.NewReader(input)), steps)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay14b(t *testing.T) {
	input := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
	steps := 40
	expect := 2_188_189_693_529
	result := day14b(day14parseInput(strings.NewReader(input)), steps)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay14a() {
	Day14a()
	// Output: day 14a: 3143
}

func ExampleDay14b() {
	Day14b()
	// Output: day 14b: 4110215602456
}

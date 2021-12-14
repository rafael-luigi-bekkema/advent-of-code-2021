package main

import (
	"fmt"
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
	tt := []struct {
		steps  int
		expect int
	}{
		{10, 1588},
		// {40, 2_188_189_693_529},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := day14a(day14parseInput(strings.NewReader(input)), tc.steps)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExampleDay14a() {
	Day14a()
	// Output: day 14a: 3143
}

package main

import (
	"fmt"
	"testing"
)

func TestDay6a(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	cases := []struct {
		days   int
		expect int
	}{
		{80, 5934},
		{256, 26_984_457_539},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := day6(input, tc.days)
			if result != tc.expect {
				t.Fatalf("case %d: expected %d, got %d", i+1, tc.expect, result)
			}
		})
	}
}

func ExampleDay6a() {
	Day6a()

	// Output: day 6a: 395627
}

func ExampleDay6b() {
	Day6b()

	// Output: day 6b: 1767323539209
}

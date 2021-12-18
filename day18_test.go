package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay18Sum(t *testing.T) {
	tt := []struct {
		input  string
		expect string
	}{
		{input: "[1,1]\n[2,2]\n[3,3]\n[4,4]", expect: "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{input: "[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]", expect: "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{input: "[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]\n[6,6]", expect: "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
		{input: `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`, expect: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := toStr(day18sum(strings.Split(tc.input, "\n")))
			if result != tc.expect {
				t.Fatalf("expected %q, got %q", tc.expect, result)
			}
		})
	}
}

func TestDay18a(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := day18a(day18sum([]string{tc.input}))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExampleDay18a() {
	Day18a()
	// Output: day 18a: 4433
}

func TestDay18b(t *testing.T) {
	input := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
	expect := 3993
	result := day18b(strings.Split(input, "\n"))
	if result != expect {
		t.Fatalf("expected %d, got %d\n", expect, result)
	}
}

func ExampleDay18b() {
	Day18b()
	// Output: day 18b: 4559
}

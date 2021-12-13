package main

import (
	"strings"
	"testing"
)

func TestDay13a(t *testing.T) {
	input := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`
	expect := 17
	result := day13a(strings.NewReader(input), 1, nil)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay13a() {
	Day13a()
	// Output: day 13a: 729
}

func TestDay13b(t *testing.T) {
	input := `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`
	var out strings.Builder
	day13a(strings.NewReader(input), 10, &out)
	result := out.String()
	expect := "#####\n" +
		"#...#\n" +
		"#...#\n" +
		"#...#\n" +
		"#####\n" +
		".....\n" +
		".....\n"
	if result != expect {
		t.Fatalf("expected\n%s, got\n%s", expect, result)
	}
}

func ExampleDay13b() {
	Day13b()

	// RGZLBHFP

	// Output: day 13b:
	// ###...##..####.#....###..#..#.####.###..
	// #..#.#..#....#.#....#..#.#..#.#....#..#.
	// #..#.#......#..#....###..####.###..#..#.
	// ###..#.##..#...#....#..#.#..#.#....###..
	// #.#..#..#.#....#....#..#.#..#.#....#....
	// #..#..###.####.####.###..#..#.#....#....
}

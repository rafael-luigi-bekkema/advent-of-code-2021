package main

import "testing"

func TestDay20a(t *testing.T) {
	algo := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##" +
		"#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###" +
		".######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#." +
		".#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#....." +
		".#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.." +
		"...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#....." +
		"..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"

	inputImg := "#..#.\n" +
		"#....\n" +
		"##..#\n" +
		"..#..\n" +
		"..###\n"

	expect := 35
	result := day20(algo, inputImg, 2)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay20a() {
	Day20()

	// Output: day20a: 5218
	// day20b: 15527
}

func TestDay20b(t *testing.T) {
	algo := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##" +
		"#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###" +
		".######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#." +
		".#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#....." +
		".#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.." +
		"...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#....." +
		"..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"

	inputImg := "#..#.\n" +
		"#....\n" +
		"##..#\n" +
		"..#..\n" +
		"..###\n"

	expect := 3351
	result := day20(algo, inputImg, 50)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

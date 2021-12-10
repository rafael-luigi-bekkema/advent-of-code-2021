package main

import (
	"strings"
	"testing"
)

func TestDay10a(t *testing.T) {
	input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`
	expect := 26397
	result := day10a(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay10a() {
	Day10a()
	// Output: day 10a: 341823
}

func TestDay10b(t *testing.T) {
	input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`
	expect := 288957
	result := day10b(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay10b() {
	Day10b()
	// Output: day 10b: 2801302861
}

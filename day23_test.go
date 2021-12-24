package main

import "testing"

func TestDay23a(t *testing.T) {
	input := "#############\n" +
		"#...........#\n" +
		"###B#C#B#D###\n" +
		"  #A#D#C#A#\n" +
		"  #########\n"
	expect := 12521
	result := day23a(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

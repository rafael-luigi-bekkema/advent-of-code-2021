package main

import (
	"strings"
	"testing"
)

func TestDay23a(t *testing.T) {
	input := "#############\n" +
		"#...........#\n" +
		"###B#C#B#D###\n" +
		"  #A#D#C#A#\n" +
		"  #########\n"
	expect := 12521
	result := day23a(strings.Split(strings.TrimRight(input, "\n"), "\n"))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay23a2(t *testing.T) {
	input := day23file()
	expect := 10607
	result := day23a(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay23b(t *testing.T) {
	input := "#############\n" +
		"#...........#\n" +
		"###B#C#B#D###\n" +
		"  #D#C#B#A#\n" +
		"  #D#B#A#C#\n" +
		"  #A#D#C#A#\n" +
		"  #########\n"
	expect := 44169
	result := day23a(strings.Split(strings.TrimRight(input, "\n"), "\n"))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay23b2(t *testing.T) {
	input := day23bfile()
	expect := 59071
	result := day23a(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

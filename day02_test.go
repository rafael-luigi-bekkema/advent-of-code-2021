package main

import (
	"strings"
	"testing"
)

func TestDay02a(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`
	expect := 150
	result := day2a(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay2a() {
	Day2a()
	// Output: day 2a: 2272262
}

func TestDay02b(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`
	expect := 900
	result := day2b(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay2b() {
	Day2b()
	// Output: day 2b: 2134882034
}

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
	result, err := day2a(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay2a() {
	if err := Day2a(); err != nil {
		panic(err)
	}
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
	result, err := day2b(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay2b() {
	if err := Day2b(); err != nil {
		panic(err)
	}
	// Output: day 2b: 2134882034
}

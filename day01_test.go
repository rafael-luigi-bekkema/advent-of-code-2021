package main

import (
	"strings"
	"testing"
)

func TestDay1a(t *testing.T) {
	input := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	expect := 7
	result := day1a(strings.NewReader(input))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay1a() {
	Day1a()
	// Output: day 1a: 1462
}

func TestDay1b(t *testing.T) {
	input := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263"
	expect := 5
	count := day1b(strings.NewReader(input))
	if expect != count {
		t.Fatalf("expected %d, got %d", expect, count)
	}
}

func ExampleDay1b() {
	Day1b()
	// Output: day 1b: 1497
}

package main

import (
	"os"
	"testing"
)

func TestDay19a(t *testing.T) {
	f, err := os.Open("input/day19_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expect := 79
	result := day19a(day19input(f))
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

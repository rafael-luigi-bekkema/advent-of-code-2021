package main

import (
	"os"
	"testing"
)

func TestDay19(t *testing.T) {
	f, err := os.Open("input/day19_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expectA := 79
	resultA, resultB := day19a(day19input(f))
	t.Run("a", func(t *testing.T) {
		if resultA != expectA {
			t.Fatalf("expected %d, got %d", expectA, resultA)
		}
	})
	expectB := 3621
	t.Run("b", func(t *testing.T) {
		if resultB != expectB {
			t.Fatalf("expected %d, got %d", expectB, resultB)
		}
	})
}

func TestDay19Final(t *testing.T) {
	f, err := os.Open("input/day19.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expectA := 440
	resultA, resultB := day19a(day19input(f))
	t.Run("a", func(t *testing.T) {
		if resultA != expectA {
			t.Fatalf("expected %d, got %d", expectA, resultA)
		}
	})
	expectB := 13382
	t.Run("b", func(t *testing.T) {
		if resultB != expectB {
			t.Fatalf("expected %d, got %d", expectB, resultB)
		}
	})
}

// func ExampleDay19a() {
// 	Day19a()
// 	// Output: 0
// }

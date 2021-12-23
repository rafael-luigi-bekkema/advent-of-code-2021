package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDay22a(t *testing.T) {
	f, err := os.Open("input/day22_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expect := 590784
	result := day22a(f)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay22a2(t *testing.T) {
	f, err := os.Open("input/day22.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expect := 658691
	result := day22a(f)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func TestDay22b(t *testing.T) {
	f, err := os.Open("input/day22_test2.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expect := 2758514936282235
	result := day22b(f)
	if result != expect {
		t.Fatalf("expected %d (len %d), got %d (len %d)", expect, len(fmt.Sprint(expect)), result, len(fmt.Sprint(result)))
	}
}

func TestDay22b2(t *testing.T) {
	f, err := os.Open("input/day22.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	expect := 1228699515783640
	result := day22b(f)
	if result != expect {
		t.Fatalf("expected %d (len %d), got %d (len %d)", expect, len(fmt.Sprint(expect)), result, len(fmt.Sprint(result)))
	}
}

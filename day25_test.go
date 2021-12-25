package main

import "testing"

func TestDay25a(t *testing.T) {
	input := `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`
	expect := 58
	result := day25a(input)
	if result != expect {
		t.Fatalf("expected %d, got %d", expect, result)
	}
}

func ExampleDay25a() {
	Day25a()
	// Output: day 25a: 516
}

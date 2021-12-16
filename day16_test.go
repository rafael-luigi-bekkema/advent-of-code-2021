package main

import (
	"fmt"
	"testing"
)

func TestDay16a(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			result := day16a(tc.input)
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExampleDay16a() {
	Day16a()
	// Output: day 16a: 984
}

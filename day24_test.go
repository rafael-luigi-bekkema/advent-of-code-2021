package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay24a(t *testing.T) {
	tt := []struct {
		min    bool
		expect int
	}{
		{min: false, expect: 92969593497992},
		{min: true, expect: 81514171161381},
	}

	alu := day24alu()
	for i, tc := range tt {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			result := day24a(alu, tc.min)
			if tc.expect != result {
				t.Fatalf("expected %v, got %v", tc.expect, result)
			}
		})
	}
}

func TestDay24Alu(t *testing.T) {
	tt := []struct {
		program    string
		inputs     []int
		w, x, y, z int
	}{
		{"inp w", []int{5}, 5, 0, 0, 0},
		{"inp w\ninp x\ninp y\ninp z", []int{1, 2, 3, 4}, 1, 2, 3, 4},
		{"inp z\ninp x\nmul z 3\neql z x", []int{2, 6}, 0, 6, 0, 1},
		{"inp w\nadd w 2", []int{5}, 7, 0, 0, 0},
		{"inp w\nmul w 3", []int{2}, 6, 0, 0, 0},
		{"inp w\ndiv w 3", []int{9}, 3, 0, 0, 0},
		{"inp w\nmod w 3", []int{4}, 1, 0, 0, 0},
		{"inp w\nadd x w\neql w x", []int{3}, 1, 3, 0, 0},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			alu := NewALU(strings.NewReader(tc.program))
			alu.run(tc.inputs)
			if tc.w != alu.w {
				t.Fatalf("expected w to be %d, got %d", tc.w, alu.w)
			}
			if tc.x != alu.x {
				t.Fatalf("expected x to be %d, got %d", tc.x, alu.x)
			}
			if tc.y != alu.y {
				t.Fatalf("expected y to be %d, got %d", tc.y, alu.y)
			}
			if tc.z != alu.z {
				t.Fatalf("expected z to be %d, got %d", tc.z, alu.z)
			}
		})
	}
}

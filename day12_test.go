package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay12a(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{{`start-A
start-b
A-c
A-b
b-d
A-end
b-end`, 10},
		{`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, 19},
		{`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`, 226}}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := day12a(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExampleDay12a() {
	Day12a()
	// Output: day 12a: 5333
}

func TestDay12b(t *testing.T) {
	tt := []struct {
		input  string
		expect int
	}{{`start-A
start-b
A-c
A-b
b-d
A-end
b-end`, 36},
		{`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, 103},
		{`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`, 3509}}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := day12b(strings.NewReader(tc.input))
			if result != tc.expect {
				t.Fatalf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func ExampleDay12b() {
	Day12b()
	// Output: day 12b: 146553
}

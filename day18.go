package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tokenizer struct {
	data string
}

type token struct {
	typ   string
	value int
}

func (t *tokenizer) token() token {
	if t.data == "" {
		return token{typ: ""}
	}
	switch c := t.data[0:1]; c {
	case "[", "]", ",":
		t.data = t.data[1:]
		return token{typ: c}
	default:
		var num []rune
		for _, c := range t.data {
			if c < '0' || '9' < c {
				break
			}
			num = append(num, c)
		}
		t.data = t.data[len(num):]
		val, _ := strconv.Atoi(string(num))
		return token{typ: "num", value: val}
	}
}

func toStr(fish *list.List) string {
	var s strings.Builder
	for e := fish.Front(); e != nil; e = e.Next() {
		s.WriteString(fmt.Sprint(e.Value))
	}
	return s.String()
}

func reduce(fish *list.List) {
outer:
	for {
		clb := 0
		for e := fish.Front(); e != nil; e = e.Next() {
			if e.Value == "[" {
				clb++
				if clb == 5 { // explode
					left := e.Next()
					right := left.Next().Next()

					for e := left.Prev(); e != nil; e = e.Prev() {
						if v, ok := e.Value.(int); ok {
							e.Value = v + (left.Value.(int))
							break
						}
					}
					for e := right.Next(); e != nil; e = e.Next() {
						if v, ok := e.Value.(int); ok {
							e.Value = v + (right.Value.(int))
							break
						}
					}

					e.Value = 0
					fish.Remove(e.Next()) // num
					fish.Remove(e.Next()) // ,
					fish.Remove(e.Next()) // num
					fish.Remove(e.Next()) // ]
					continue outer
				}
			}
			if e.Value == "]" {
				clb--
			}
		}
		for e := fish.Front(); e != nil; e = e.Next() {
			if v, ok := e.Value.(int); ok && v >= 10 { // split
				e.Value = "["
				div := v / 2
				e = fish.InsertAfter(div, e)
				e = fish.InsertAfter(",", e)
				if v%2 != 0 {
					div += 1
				}
				e = fish.InsertAfter(div, e)
				e = fish.InsertAfter("]", e)
				continue outer
			}
		}
		break
	}
}

func day18sum(input []string) *list.List {
	fish := list.New()
	for i, line := range input {
		if i > 0 {
			fish.PushFront("[")
			fish.PushBack(",")
		}

		t := tokenizer{line}
		for tok := t.token(); tok.typ != ""; tok = t.token() {
			if tok.typ == "num" {
				fish.PushBack(tok.value)
			} else {
				fish.PushBack(tok.typ)
			}
		}

		if i > 0 {
			fish.PushBack("]")
			reduce(fish)
		}
	}
	return fish
}

// [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
func magnitude(e *list.Element) (int, *list.Element) {
	var left, right int

	e = e.Next()
	if e.Value == "[" {
		left, e = magnitude(e)
	} else {
		left = e.Value.(int)
	}
	e = e.Next().Next() // skip ,

	if e.Value == "[" {
		right, e = magnitude(e)
	} else {
		right = e.Value.(int)
	}

	e = e.Next()
	return 3*left + 2*right, e
}

func day18a(l *list.List) int {
	num, _ := magnitude(l.Front())
	return num
}

func Day18a() {
	data, err := os.ReadFile("input/day18.txt")
	if err != nil {
		panic(err)
	}
	items := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	result := day18a(day18sum(items))
	fmt.Printf("day 18a: %d\n", result)
}

func day18b(input []string) int {
	var poppop int
	for i := range input {
		for j := range input {
			if i == j {
				continue
			}
			num := day18a(day18sum([]string{input[i], input[j]}))
			if num > poppop {
				poppop = num
			}
		}
	}
	return poppop
}

func Day18b() {
	data, err := os.ReadFile("input/day18.txt")
	if err != nil {
		panic(err)
	}
	items := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	result := day18b(items)
	fmt.Printf("day 18b: %d\n", result)
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tokT byte

const (
	tokLeftBr tokT = iota
	tokRightBr
	tokComma
	tokEOL
	tokNum
)

func (t tokT) String() string {
	switch t {
	case tokLeftBr:
		return "["
	case tokRightBr:
		return "]"
	case tokComma:
		return ","
	default:
		return ""
	}
}

type tokenizer struct {
	data string
}

type token struct {
	typ   tokT
	value int
}

func (t *tokenizer) token() token {
	if t.data == "" {
		return token{typ: tokEOL}
	}
	switch c := t.data[0:1]; c {
	case "[":
		t.data = t.data[1:]
		return token{typ: tokLeftBr}
	case "]":
		t.data = t.data[1:]
		return token{typ: tokRightBr}
	case ",":
		t.data = t.data[1:]
		return token{typ: tokComma}
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
		return token{typ: tokNum, value: val}
	}
}

func toStr(fish []interface{}) string {
	var s strings.Builder
	for _, e := range fish {
		s.WriteString(fmt.Sprint(e))
	}
	return s.String()
}

func reduce(fish []interface{}) []interface{} {
outer:
	for {
		clb := 0
		for i, e := range fish {
			if e == tokLeftBr {
				clb++
				if clb == 5 { // explode
					left := fish[i+1]
					right := fish[i+3]

					for j := i; j >= 0; j-- {
						if v, ok := fish[j].(int); ok {
							fish[j] = v + (left.(int))
							break
						}
					}
					for j := i + 4; j <= len(fish)-1; j++ {
						if v, ok := fish[j].(int); ok {
							fish[j] = v + (right.(int))
							break
						}
					}

					fish[i] = 0
					fish = append(fish[:i+1], fish[i+5:]...)
					continue outer
				}
			}
			if e == tokRightBr {
				clb--
			}
		}
		for i, e := range fish {
			if v, ok := e.(int); ok && v >= 10 { // split
				div := v / 2
				div2 := div
				if v%2 != 0 {
					div2 += 1
				}
				rest := make([]interface{}, len(fish[i+1:]))
				copy(rest, fish[i+1:])
				fish = append(fish[:i], tokLeftBr, div, tokComma, div2, tokRightBr)
				fish = append(fish, rest...)
				continue outer
			}
		}
		break
	}
	return fish
}

func day18sum(input []string) []interface{} {
	var fish []interface{}
	for i, line := range input {
		if i > 0 {
			fish = append([]interface{}{tokLeftBr}, fish...)
			fish = append(fish, tokComma)
		}

		t := tokenizer{line}
		for tok := t.token(); tok.typ != tokEOL; tok = t.token() {
			if tok.typ == tokNum {
				fish = append(fish, tok.value)
			} else {
				fish = append(fish, tok.typ)
			}
		}

		if i > 0 {
			fish = append(fish, tokRightBr)
			fish = reduce(fish)
		}
	}
	return fish
}

func magnitude(fish []interface{}) (int, []interface{}) {
	var left, right int

	if fish[1] == tokLeftBr {
		left, fish = magnitude(fish[1:])
	} else {
		left = fish[1].(int)
		fish = fish[2:]
	}

	if fish[1] == tokLeftBr {
		right, fish = magnitude(fish[1:])
		fish = fish[1:]
	} else {
		if tok, ok := fish[1].(tokT); ok {
			fmt.Println(toStr(fish), "=", tok)
		}
		right = fish[1].(int)
		fish = fish[3:]
	}

	return 3*left + 2*right, fish
}

func day18a(input []interface{}) int {
	num, _ := magnitude(input)
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

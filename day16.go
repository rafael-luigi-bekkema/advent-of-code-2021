package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type message string

func (m *message) eat(amount int) string {
	val := (*m)[:amount]
	*m = (*m)[amount:]
	return string(val)
}

type Packet struct {
	ver     int64
	typeID  int64
	literal string
	sub     []Packet
}

func (pak Packet) String() string {
	subs := make([]string, len(pak.sub))
	for i, s := range pak.sub {
		subs[i] = fmt.Sprintf("  %s", s)
	}
	var literal string
	if pak.literal != "" {
		literal = fmt.Sprintf(" literal: %q", pak.literal)
	}
	return fmt.Sprintf("ver: %d, type: %d%s\n%s", pak.ver, pak.typeID, literal, strings.Join(subs, ""))
}

func (pak Packet) value() int64 {
	var total int64
	switch pak.typeID {
	case 0:
		for _, sub := range pak.sub {
			total += sub.value()
		}
	case 1:
		for i, sub := range pak.sub {
			if i == 0 {
				total = sub.value()
				continue
			}
			total *= sub.value()
		}
	case 2: //min
		for i, sub := range pak.sub {
			val := sub.value()
			if i == 0 {
				total = val
				continue
			}
			if val < total {
				total = val
			}
		}
	case 3: //max
		for i, sub := range pak.sub {
			val := sub.value()
			if i == 0 {
				total = val
				continue
			}
			if val > total {
				total = val
			}
		}
	case 4:
		total, _ = strconv.ParseInt(pak.literal, 2, 64)
	case 5:
		if pak.sub[0].value() > pak.sub[1].value() {
			total = 1
		}
	case 6:
		if pak.sub[0].value() < pak.sub[1].value() {
			total = 1
		}
	case 7:
		if pak.sub[0].value() == pak.sub[1].value() {
			total = 1
		}
	default:
		panic(fmt.Sprintf("Unknown typeID: %d", pak.typeID))
	}
	return total
}

func (pak Packet) verSum() int64 {
	total := pak.ver
	for _, sub := range pak.sub {
		total += sub.verSum()
	}
	return total
}

func parsePacketsAll(msg *message) []Packet {
	var paks []Packet
	for len(*msg) > 0 {
		paks = append(paks, parsePacket(msg))
	}
	return paks
}

func parsePackets(msg *message, amount int) []Packet {
	paks := make([]Packet, amount)
	for i := 0; i < amount; i++ {
		paks[i] = parsePacket(msg)
	}
	return paks
}

func parseLiteral(msg *message) string {
	var parts []string
	for {
		prefix := msg.eat(1)
		parts = append(parts, msg.eat(4))

		if prefix == "0" {
			break
		}
	}
	return strings.Join(parts, "")
}

func parsePacket(msg *message) Packet {
	var pak Packet
	pak.ver, _ = strconv.ParseInt(msg.eat(3), 2, 64)
	pak.typeID, _ = strconv.ParseInt(msg.eat(3), 2, 64)

	if pak.typeID == 4 { // literal value
		pak.literal = parseLiteral(msg)
		return pak
	}

	lengthTypeID := msg.eat(1)
	if lengthTypeID == "0" {
		subPacketLen, _ := strconv.ParseInt(msg.eat(15), 2, 64)
		newmsg := message(msg.eat(int(subPacketLen)))
		pak.sub = parsePacketsAll(&newmsg)
		return pak
	}

	// lengthTypeID = 1
	nrSubPackets, _ := strconv.ParseInt(msg.eat(11), 2, 64)
	pak.sub = parsePackets(msg, int(nrSubPackets))
	return pak
}

func day16a(input string) int {
	str := make([]string, len(input))
	for i, c := range input {
		num, _ := strconv.ParseInt(string(c), 16, 64)
		bin := fmt.Sprintf("%04b", num)
		str[i] = bin
	}
	msg := message(strings.TrimRight(strings.Join(str, ""), "0"))
	pak := parsePacket(&msg)

	return int(pak.verSum())
}

func Day16a() {
	data, err := os.ReadFile("input/day16.txt")
	if err != nil {
		panic(err)
	}
	result := day16a(string(data))
	fmt.Printf("day 16a: %d\n", result)
}

func day16b(input string) int {
	str := make([]string, len(input))
	for i, c := range input {
		num, _ := strconv.ParseInt(string(c), 16, 64)
		bin := fmt.Sprintf("%04b", num)
		str[i] = bin
	}
	msg := message(strings.TrimRight(strings.Join(str, ""), "0"))
	pak := parsePacket(&msg)

	return int(pak.value())
}

func Day16b() {
	data, err := os.ReadFile("input/day16.txt")
	if err != nil {
		panic(err)
	}
	result := day16b(string(data))
	fmt.Printf("day 16b: %d\n", result)
}

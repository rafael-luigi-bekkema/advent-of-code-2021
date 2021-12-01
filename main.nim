from std/strutils import Digits, parseInt, splitLines
import std/strformat
import std/sequtils

proc day01a_test() =
    let input = """199
200
208
210
200
207
240
269
260
263"""
    const expect = 7
    var count = 0
    var prev = -1
    for line in splitLines(input):
        let num = parseInt(line)
        if num > prev and prev != -1:
            count += 1
        prev = num
    assert count == expect

proc day01a() =
    let f = open("input/day01.txt")
    defer: f.close()
    var line = ""
    var count = 0
    var prev = -1
    while readLine(f, line):
        let num = parseInt(line)
        if num > prev and prev != -1:
            count += 1
        prev = num

    echo &"total 1a: {count}"

day01a_test()
day01a()

proc day01b_core(lines: seq[string]): int =
    var win = newSeq[int]()
    result = 0
    for line in lines:
        let num = parseInt(line)
        win.add(num)
        if win.len >= 4:
            if win[1..3].foldl(a + b) > win[0..2].foldl(a + b):
                result += 1
            win = win[1..3]

proc day01b_test() =
    let input = """199
200
208
210
200
207
240
269
260
263"""
    const expect = 5
    let count = day01b_core(splitLines(input))
    assert count == expect

proc day01b() =
    let f = toSeq(lines("input/day01.txt"))
    let count = day01b_core(f)
    echo &"total 1b: {count}"


day01b_test()
day01b()


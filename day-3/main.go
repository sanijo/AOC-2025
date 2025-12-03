package main

import (
	"github.com/sanijo/day-3/part1"
	"github.com/sanijo/day-3/part2"
)


func main() {
    part := "part2"
    switch part {
    case "part1":
        part1.Part1()
    case "part2":
        part2.Part2()
    default:
        panic("unknown part")
    }
}


package main

import (
	"fmt"

	"github.com/sanijo/day-4/part1"
	"github.com/sanijo/day-4/part2"
)


func main() {
    part := "part2"
    switch part {
    case "part1":
        fmt.Println("Running Part 1")
        part1.Part1()
    case "part2":
        fmt.Println("Running Part 2")
        part2.Part2()
    default:
        panic("unknown part")
    }
}


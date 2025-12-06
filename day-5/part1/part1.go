package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Bounds represents minimum and maximum bounds read from the input
type Bounds struct {
    Min, Max int
}

func readFromFile(path string) ([]Bounds, []int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var bounds []Bounds
    var ids []int
    firstEmptyLine := false

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text()) // string with the line content
        //fmt.Printf("Read line: %+v\n", line)
        if len(line) == 0 {
            firstEmptyLine = true
            continue
        }
        if firstEmptyLine {
            id, err := strconv.Atoi(strings.TrimSpace(line))
            if err != nil {
                return nil, nil, fmt.Errorf("invalid ID: %s", line)
            }
            ids = append(ids, id)
        } else {
            // split line by - to get Bounds
            parts := strings.Split(line, "-")
            if len(parts) != 2 {
                return nil, nil, fmt.Errorf("invalid bounds format: %s", line)
            }
            lower_bound, err := strconv.Atoi(strings.TrimSpace(parts[0]))
            if err != nil {
                return nil, nil, fmt.Errorf("invalid lower bound: %s", parts[0])
            }
            upper_bound, err := strconv.Atoi(strings.TrimSpace(parts[1]))
            if err != nil {
                return nil, nil, fmt.Errorf("invalid upper bound: %s", parts[1])
            }
            bounds = append(bounds, Bounds{Min: lower_bound, Max: upper_bound}) 
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    return bounds, ids, nil
}

func Part1() {
    bounds, ids, err := readFromFile("input.txt")
    if err != nil {
        log.Fatalf("Error reading from file: %v", err)
    }
    // print bounds and ids
    //fmt.Printf("Bounds: %+v\n", bounds)
    //fmt.Printf("IDs: %+v\n", ids)

    fresh := 0
    for _, id := range ids {
        for _, bound := range bounds {
            if id >= bound.Min && id <= bound.Max {
                fresh++
                break
            }
        }
    }
    fmt.Printf("Fresh IDs count: %d\n", fresh)
}


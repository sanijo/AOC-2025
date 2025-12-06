package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func calculateMergedBounds(bounds []Bounds) []Bounds {
    // sort bounds by Min, then by Max
    sort.Slice(bounds, func(i, j int) bool {
        if bounds[i].Min == bounds[j].Min {
            return bounds[i].Max < bounds[j].Max
        }
        return bounds[i].Min < bounds[j].Min
    })
    
    for i := 0; i < len(bounds); i++ {
        j := i + 1
        for j < len(bounds) {
            L1, U1 := bounds[i].Min, bounds[i].Max
            L2, U2 := bounds[j].Min, bounds[j].Max
    
            // check overlap / touching
            if !(U2 < L1-1 || U1 < L2-1) {
                // they overlap or touch -> merge into bounds[i]
                if L2 < L1 {
                    L1 = L2
                }
                if U2 > U1 {
                    U1 = U2
                }
                bounds[i].Min = L1
                bounds[i].Max = U1
    
                // remove bounds[j] by shrinking the slice
                bounds = append(bounds[:j], bounds[j+1:]...)
    
                // DO NOT j++ here
                // want to check the new element that just moved into index j
                continue
            }
            j++
        }
    }
    return bounds
}

func Part2() {
    bounds, _, err := readFromFile("input.txt")
    if err != nil {
        log.Fatalf("Error reading from file: %v", err)
    }
    // print bounds and ids
    //fmt.Printf("Bounds: %+v\n", bounds)
    //fmt.Printf("IDs: %+v\n", ids)

    // merge bounds
    bounds = calculateMergedBounds(bounds)

    // calculate number of valid IDs
    freshID := 0
    for _, b := range bounds {
        freshID += (b.Max - b.Min + 1)
    }
    fmt.Println()
    fmt.Printf("Number of valid IDs: %d\n", freshID)

}


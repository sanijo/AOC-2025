package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// struct that holds slice of rotation values from loaded file.
// In record slice: first value is direction, second value is distance.
// direction is string, distance is int.
type Rotation struct {
	Direction []string
	Distance  []int
}

// zerosPassed counts how many times the dial points at 0
// during a single rotation of `dist` clicks, starting from `dial`.
func zerosPassed(dial int, dir string, dist int) int {
    // IDEA: find how many clicks to first 0 hit (k0),
    // then count how many full 100-click cycles fit in remaining distance.
    // total hits = 1 (for first hit at k0) + (dist - k0) / 100

    // return immediately if no distance to move
    if dist <= 0 {
        return 0
    }

    // normalize dial position to [0..99] 
    dial_norm := dial % 100
    if dial_norm < 0 {
        dial_norm += 100
    }

    var k0 int // how many clicks to first 0 hit

    switch dir {
    case "R", "r":
        // (dial_norm + k) % 100 == 0  → k ≡ (100 - dial_norm) % 100
        // e.g., if dial_norm == 30, k0 == 70 - need 70 clicks to hit 0
        k0 = (100 - dial_norm) % 100
        if k0 == 0 { // case when dial_norm == 0, takes 100 clicks to hit 0
            k0 = 100
        }

    case "L", "l":
        // (dial_norm - k) % 100 == 0  → k ≡ dial_norm (mod 100)
        // e.g., if dial_norm == 30, k0 == 30 - need 30 clicks to hit 0
        if dial_norm == 0 {
            k0 = 100
        } else {
            k0 = dial_norm
        }

    default:
        log.Fatalf("invalid direction: %s", dir)
    }

    // if first 0 hit is beyond distance, return 0
    if k0 > dist {
        return 0
    }

    return (dist-k0)/100 + 1 // +1 for the first hit at k0
}


func main() {
	// open ./input.csv file
	f, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	rotation := Rotation{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// append direction and distance
		if len(record) == 0 || len(record[0]) == 0 {
			continue
		}

		rawStr := record[0]
		direction := rawStr[:1]       // first character
		distanceStr := rawStr[1:]     // rest of the string
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("record: %+v, direction: %s, distance: %d\n",
			record, direction, distance)

		rotation.Direction = append(rotation.Direction, direction)
		rotation.Distance = append(rotation.Distance, distance)
	}

	fmt.Println()

	// process rotation values to find password
	dial := 50
	password_count := 0

	fmt.Printf("Initial dial position: %d\n", dial)

    for i, dir := range rotation.Direction {
        dist := rotation.Distance[i]
    
        // count all 0-hits during this rotation
        passed_zero := zerosPassed(dial, dir, dist)
        password_count += passed_zero

        //  print how many times 0 was hit during this rotation
        if passed_zero > 0 {
            fmt.Printf("During rotation %s%d, dial passed 0 %d time(s)\n", dir, dist, passed_zero)
        }
    
        // now move dial for end-of-rotation
        switch dir {
        case "R", "r":
            dial = (dial + dist) % 100
        case "L", "l":
            dial = (dial - dist%100 + 100) % 100
        default:
            log.Fatalf("invalid direction: %s", dir)
        }
    
		fmt.Printf("After rotation %s%d, dial is at %d\n", dir, dist, dial)
    }

	fmt.Println()
	fmt.Printf("Password count (times dial points to 0): %d\n", password_count)
}


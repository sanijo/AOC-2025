package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"strconv"
)


func repeatedPatterns(s string) string {
    n := len(s)
    tried := make(map[string]bool)

    // loop through all possible substring lengths
    // for odd lengths, the max repeating substring is n/2 (e.g., "12121" -> "12")
    // because of integer division, n/2 will automatically floor the value
    for i := 1; i <= n/2; i++ {
        // only consider chunks that evenly divide the string length. If it
        // doesn't, skip it. (e.g., "123" can't be made of "12" repeated)
        if n%i == 0 {
            // extract candidate substring
            candidate := s[:i]

            // skip candidates already tested
            if tried[candidate] {
                continue
            }
            tried[candidate] = true

            // build repeated version of candidate
            repeated := ""
            repeatCount := n / i // length of s divided by length of candidate
            for k := 0; k < repeatCount; k++ {
                repeated += candidate
            }

            // compare with original string
            if repeated == s {
                return candidate
            }
        }
    }

    return "" // return empty string if no repeating pattern
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

    var invalidIDsSum int
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) == 0 || len(record[0]) == 0 {
			continue
		}
        // split each field by "-" and create range slices
		for _, field := range record {
			parts := strings.Split(field, "-")
//            fmt.Println("Processing field:", field, "-> parts:", parts)
            if len(parts) != 2 { // check if split resulted in two parts
                log.Fatalf("invalid field format: %s", field)
            }
            start, err := strconv.Atoi(parts[0])
            if err != nil {
                log.Fatal(err)
            }
            end, err := strconv.Atoi(parts[1])
            if err != nil {
                log.Fatal(err)
            }
            // create slice ranging from start to end with start++
            size := (end-start)+1
            if size <= 0 {
                log.Fatalf("invalid range: %s", field)
            }
            rangeSlice := make([]int, size)
            for i := range rangeSlice {
                rangeSlice[i] = start + i
            }
//            fmt.Println(rangeSlice)
            // sum invalid IDs
            for _, id := range rangeSlice {
                strID := strconv.Itoa(id)
                if pattern := repeatedPatterns(strID); pattern != "" {
                    invalidIDsSum += id
                    fmt.Printf("Invalid ID found (repeats pattern '%s'): %d (Sum: %d)\n", pattern, id, invalidIDsSum)
                }
            }
		}
    }

	fmt.Println()
    fmt.Printf("Total Sum of Invalid IDs: %d\n", invalidIDsSum)

}


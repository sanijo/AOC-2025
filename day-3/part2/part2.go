package part2

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	// open ./input.csv file
	f, err := os.Open("input.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)

    outputJoltage := 0
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
        // split record[0] into elements and convert to int slice
        recordStr := strings.Split(record[0], "")
        recordInt := make([]int, len(recordStr))
        for i, string := range recordStr {
            number, err := strconv.Atoi(string)
            if err != nil {
                log.Fatal(err)
            }
            recordInt[i] = number
        }
        fmt.Println(recordInt)

        k := 12
        result := make([]int, 0, k)
        start := 0

        for picks := 0; picks < k; picks++ {
            // remaining digits needed
            need := k - picks
            // last index search can start from
            end := len(recordInt) - need 
            // find max digit in recordInt[start : end+1]
            maxVal := -1
            maxId := start
            for i := start; i <= end; i++ {
                if recordInt[i] > maxVal {
                    maxVal = recordInt[i]
                    maxId = i
                }
            }
            result = append(result, maxVal)
            start = maxId + 1
        }
        // convert result slice to single integer
        joltageStr := ""
        for _, digit := range result {
            joltageStr += strconv.Itoa(digit)
        }
        joltage, err := strconv.Atoi(joltageStr)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Output joltage for record %s: %d\n", record[0], joltage)

        outputJoltage += joltage
    }
    fmt.Println()
    fmt.Printf("Output joltage: %d\n", outputJoltage)
	fmt.Println()

}


package part1

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)


func Part1() {
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
        // find two numbers in slice that combined create largest possible numbers
        // e.g. [2, 3, 5, 8, 1, 9] -> 89
        //largestNum := 0
        var largestNum int
        for i := 0; i<len(recordInt)-1; i++ {
            for j := i+1; j<len(recordInt); j++ {
                combinedNum := strconv.Itoa(recordInt[i]) + strconv.Itoa(recordInt[j])
//                fmt.Printf("Combining %d and %d to form %s\n", recordInt[i], recordInt[j], combinedNum)
                combinedNumInt, err := strconv.Atoi(combinedNum)
                if err != nil {
                    log.Fatal(err)
                }
                if combinedNumInt > largestNum {
                    largestNum = combinedNumInt
//                    fmt.Printf("New largest number found: %d (from %d and %d)\n", largestNum, recordInt[i], recordInt[j])
                }
            }
        }
        outputJoltage += largestNum
        fmt.Printf("Largest number for this record: %d\n", largestNum)
    }
    fmt.Printf("Output joltage: %d\n", outputJoltage)

	fmt.Println()

}


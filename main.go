package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func maxElem(e []int) (bestFound, location int) {
	bestFound = math.MinInt
	location = -1
	for i := 0; i < len(e); i++ {
		if e[i] > bestFound {
			bestFound = e[i]
			location = i
		}
	}
	return bestFound, location
}

func main() {
	f, err := os.Open("./data/day1")
	check(err)
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	elfCalories := make([]int, 0, 100)
	counter := 0
	for scanner.Scan() {
		check(scanner.Err())
		if scanner.Text() == "" {
			// we have added up all the calories an elf has and may record that information
			elfCalories = append(elfCalories, counter)
			fmt.Printf("length %d num cals %d\n", len(elfCalories), counter)
			counter = 0
		} else {
			num, err := strconv.Atoi(scanner.Text())
			check(err)
			counter += num
		}
	}
	best, loc := maxElem(elfCalories)
	fmt.Printf("Elf number %d with %d calories is the winner\n", loc+1, best)

	ff, err := os.Create("test_data/day1")
	check(err)
	// create the test output for comparing the refactor output with
	for i := 0; i < len(elfCalories); i++ {
		ff.WriteString(fmt.Sprintf("%d\n", elfCalories[i]))
	}
	sort.Ints(elfCalories)
	fmt.Printf("calories of top three elves: %d", elfCalories[len(elfCalories)-1]+elfCalories[len(elfCalories)-2]+elfCalories[len(elfCalories)-3])
}

// for scanning input into blocks that are handled individually
// modified from https://go.dev/src/bufio/scan.go?s=9745:9823#L280
func ScanBlock(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexAny(data, "\n\n"); i >= 0 {
		// We have a full double newline-terminated block.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated block. Return it.
	if atEOF {
		data = dropCR(data)
		// for regularity if the very last line is not newline terminated, then add a newline in
		if data[len(data)-1] != '\n' {
			data = append(data, '\n')
		}
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// take whole from https://go.dev/src/bufio/scan.go?s=9745:9823#L337
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

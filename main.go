package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
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
	blockScanner := bufio.NewScanner(f)
	blockScanner.Split(ScanBlock)
	elfCalories := make([]int, 0, 100)
	for blockScanner.Scan() {
		lineScanner := bufio.NewScanner((bytes.NewReader(blockScanner.Bytes())))
		sum := 0
		for lineScanner.Scan() {
			// fmt.Print(lineScanner.Text())
			// fmt.Println(lineScanner.Text())
			n, err := strconv.Atoi(lineScanner.Text())
			check(err)
			sum += n
		}
		fmt.Println(blockScanner.Text())
		fmt.Println("END OF BLOCK")
		elfCalories = append(elfCalories, sum)
	}

	// f, err = os.Create("test_data/day1_refactor")
	// check(err)
	// for _, val := range elfCalories {
	// 	f.WriteString(fmt.Sprintf("%d\n", val))
	// }

}

// for scanning input into blocks that are handled individually
// modified from https://go.dev/src/bufio/scan.go?s=9745:9823#L280
func ScanBlock(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	for i := 0; i < len(data)-1; i++ {
		// We have a full double newline-terminated block.
		if data[i] == '\n' && data[i+1] == '\n' {
			return i + 2, data[0 : i+1], nil
		}
	}
	// If we're at EOF, we have a final, non-terminated block. Return it.
	if atEOF {
		// for regularity if the very last line is not newline terminated, then add a newline in
		// this might be extraneous since the next consumer is another scanner
		if data[len(data)-1] != '\n' {
			data = append(data, '\n')
			return len(data) - 1, data, nil
		}
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

// returns most calories found from day 1 part 1
func day1() int {
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
	best, _ := maxElem(elfCalories)
	return best
}

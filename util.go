package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func throwPoints(throw string) int {
	switch throw {
	case "rock":
		return 1
	case "paper":
		return 2
	case "scissors":
		return 3
	}
	return -1
}

func roundResult(left, right string) (leftResult, rightResult string) {
	if left == right {
		return "draw", "draw"
	}
	if left == "rock" {
		switch right {
		case "scissors":
			return "win", "lose"
		case "paper":
			return "lose", "win"
		}
	}
	if left == "paper" {
		switch right {
		case "scissors":
			return "lose", "win"
		case "rock":
			return "win", "lose"
		}
	}
	if left == "scissors" {
		switch right {
		case "rock":
			return "lose", "win"
		case "paper":
			return "win", "lose"
		}
	}
	return "", ""
}

func pointsFromResult(result string) int {
	switch result {
	case "win":
		return 6
	case "lose":
		return 0
	case "draw":
		return 3
	}
	return -1
}

func throwFromStrategyGuide(stratSymbol string) string {
	switch stratSymbol {
	case "A", "X":
		return "rock"
	case "B", "Y":
		return "paper"
	case "C", "Z":
		return "scissors"
	}
	return "womp womp"
}

func neededResult(stratSymbol string) string {
	switch stratSymbol {
	case "X":
		return "lose"
	case "Y":
		return "draw"
	case "Z":
		return "win"
	}
	return "double womp"
}

func throwForResult(throw, goal string) (neededThrow string) {
	for _, rightThrow := range []string{"rock", "paper", "scissors"} {
		_, rr := roundResult(throw, rightThrow)
		if rr == goal {
			return rightThrow
		}
	}
	return "the rare triple womp!"
}

func Points(left, right string) (leftPoints, rightPoints int) {
	// rock beats scissors, scissors beats paper, paper beats rock
	// winning awards 6 points, drawing awards 3 points, losing awards 0 points
	// throwing rock awards 1 point, paper awards 2 points, scissors awards 3 points
	lr, rr := roundResult(left, right)
	leftPoints = pointsFromResult(lr) + throwPoints(left)
	rightPoints = pointsFromResult(rr) + throwPoints(right)
	return leftPoints, rightPoints
}

type day2Result struct {
	part1 int
	part2 int
}

func day2() day2Result {
	// find total score if you followed strategy guide
	// part two: update meaning of second column
	scanner, fileClose := getScanner("./data/day2")
	defer fileClose()
	runningTotal := 0
	runningTotal2 := 0
	for scanner.Scan() {
		check(scanner.Err())
		round := strings.Split(scanner.Text(), " ")
		left := throwFromStrategyGuide(round[0])
		right := throwFromStrategyGuide(round[1])
		lp, rp := Points(left, right)
		runningTotal += rp

		right = throwForResult(left, neededResult(round[1]))
		lp, rp = Points(left, right)
		runningTotal2 += rp
		lp = lp // just to keep lp around without the linter complaining
	}

	return day2Result{runningTotal, runningTotal2}
}

func getScanner(fn string) (*bufio.Scanner, func() error) {
	f, err := os.Open(fn)
	check(err)
	scanner := bufio.NewScanner(f)
	return scanner, f.Close
}

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

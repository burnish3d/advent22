package main

import (
	"bufio"
	"strings"
)

type StringStack []string

func NewStringStack() StringStack {
	return make(StringStack, 0)
}

func (s *StringStack) Push(in string) {
	*s = append(*s, in)
}

func (s *StringStack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return val
}

func MoveN(stacks []StringStack, from, dest, n int) {
	for i := 0; i < n; i++ {
		val := stacks[from].Pop()
		stacks[dest].Push(val)
	}

}

func MoveN9001(stacks []StringStack, from, dest, n int) {
	for i := 0; i < n; i++ {
		val := stacks[from].Pop()
		defer stacks[dest].Push(val)
	}

}

func ParseStacks(bf *bufio.Scanner) []StringStack {
	// watch out, data.day5 has significant trailing whitespace
	bf.Scan()
	numberOfColumns := 9 // the columns are 4 characters wide except the last, which is 3
	theStacks := make([]StringStack, numberOfColumns)
	for i := 0; i < numberOfColumns; i++ {
		theStacks[i] = NewStringStack()
	}

	for i := 0; i < numberOfColumns; i++ {
		offset := i*4 + 1
		if offset > len(bf.Text()) {
			break // guard against the IDE deleting significant white space :(
		}
		letter := string(bf.Text()[offset])
		if letter != " " {
			defer theStacks[i].Push(letter)
		}
	}
	for {
		bf.Scan()
		if bf.Text()[1] == '1' {
			return theStacks
		}
		for i := 0; i < numberOfColumns; i++ {
			offset := i*4 + 1
			if offset > len(bf.Text()) {
				break // guard against the IDE deleting significant white space :(
			}
			letter := string(bf.Text()[offset])
			if letter != " " {
				defer theStacks[i].Push(letter)
			}
		}
	}
}

func duplicateStacks(stacks []StringStack) []StringStack {
	newStacks := make([]StringStack, len(stacks))
	for i := 0; i < len(stacks); i++ {
		newStacks[i] = make(StringStack, len(stacks[i]))
		copy(newStacks[i], stacks[i])
	}
	return newStacks
}

func ParseCommand(in string) (from, dest, num int) {
	v := strings.Split(in, " ")
	if len(v) < 6 { // represents a line that does not have the format of a command
		return 0, 0, 0
	}
	return toInt(v[3]) - 1, toInt(v[5]) - 1, toInt(v[1]) // stacks are 1 indexed in the data
}

func day5() Result {
	s, c := getScanner("day5")
	defer c()
	var ret Result
	theStacks := ParseStacks(s)
	theStacksForPart2 := duplicateStacks(theStacks)
	for s.Scan() { // we are relying on the blank line between the box stack description and the sequence of move commands
		from, dest, num := ParseCommand(s.Text())
		MoveN(theStacks, from, dest, num)
		MoveN9001(theStacksForPart2, from, dest, num)
	}

	v1 := ""
	for i := 0; i < len(theStacks); i++ {
		v1 += theStacks[i].Pop()
	}
	ret.Part1 = v1

	v2 := ""
	for i := 0; i < len(theStacksForPart2); i++ {
		v2 += theStacksForPart2[i].Pop()
	}
	ret.Part2 = v2

	return ret
}

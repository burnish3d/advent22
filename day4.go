package main

import (
	"strings"
)

type Section struct {
	lowerBound int
	upperBound int
}

func (s Section) Contains(r Section) bool {
	return r.lowerBound >= s.lowerBound && r.upperBound <= s.upperBound
}

// func (s Section) IsContainedBy(r Section) bool {
// 	return s.lowerBound >= r.lowerBound && s.upperBound <= r.upperBound
// }

func NewSection(i string) Section {
	lu := strings.Split(i, "-")
	return Section{
		lowerBound: toInt(lu[0]),
		upperBound: toInt(lu[1]),
	}
}

func day4() Result {
	// part one is all about finding how many overlapping sections exist
	s, c := getScanner("day4")
	defer c()
	overlapCount := 0
	for s.Scan() {
		theSplit := strings.Split(s.Text(), ",")
		ls, rs := NewSection(theSplit[0]), NewSection(theSplit[1])
		if ls.Contains(rs) || rs.Contains(ls) {
			overlapCount += 1
		}
	}

	return Result{Part1: overlapCount}
}

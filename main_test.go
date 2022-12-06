package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringStack(t *testing.T) {
	a := NewStringStack()
	a.Push("I")
	a.Push("B")
	assert.Equal(t, "B", a.Pop())
	assert.Equal(t, "I", a.Pop())
}

func TestParseStacks(t *testing.T) {
	s, c := getScanner("day5")
	defer c()
	// yes this is ugly, BUT, it was also easy and checks if I mess up the parser or if the IDE removes trailing white space from the data file again
	theStacks := ParseStacks(s)
	assert.Equal(t, theStacks[0].Pop(), "T")
	assert.Equal(t, theStacks[1].Pop(), "N")
	assert.Equal(t, theStacks[2].Pop(), "D")
	assert.Equal(t, theStacks[3].Pop(), "L")
	assert.Equal(t, theStacks[4].Pop(), "M")
	assert.Equal(t, theStacks[5].Pop(), "S")
	assert.Equal(t, theStacks[6].Pop(), "W")
	assert.Equal(t, theStacks[7].Pop(), "M")
	assert.Equal(t, theStacks[8].Pop(), "S")
}

func TestDay6(t *testing.T) {
	ret := day6()
	assert.Equal(t, 1640, ret.Part1)
	assert.Equal(t, 3613, ret.Part2)
}

func TestDay6v2(t *testing.T) {
	ret := day6v2()
	assert.Equal(t, 1640, ret.Part1)
	assert.Equal(t, 3613, ret.Part2)
}

var res Result

func BenchmarkDay6(b *testing.B) {
	var r Result
	for i := 0; i < b.N; i++ {
		r = day6()
	}
	res = r
}

func BenchmarkDay6V2(b *testing.B) {
	var r Result
	for i := 0; i < b.N; i++ {
		r = day6v2()
	}
	res = r
}

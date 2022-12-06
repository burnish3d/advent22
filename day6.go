package main

import "bufio"

type RingBuffer struct {
	B     []byte
	Count int
	Size  int
}

func NewRingBuffer(size int) *RingBuffer {
	rb := RingBuffer{}
	rb.B = make([]byte, size)
	rb.Size = size
	return &rb
}

func (r *RingBuffer) Add(in byte) {
	r.B[r.Count%r.Size] = in
	r.Count += 1
}

// returns true if each byte in the buffer is unique
func (r *RingBuffer) Unique() bool {
	// other options include making a set, or making a copy then sorting before checking if any subsequent values match
	// I did it this way because I'd like to compare to a more traditionally better solution
	for i := 0; i < len(r.B); i++ {
		for j := 0; j < len(r.B); j++ {
			if i == j {
				continue
			}
			if r.B[i] == r.B[j] {
				return false
			}
		}
	}
	return true
}

// returning true means 'keep going'
func (r *RingBuffer) FillUntilUnique(in byte) bool {
	if r.Count < r.Size {
		r.Add(in)
		return true
	}
	if r.Count >= r.Size && !r.Unique() {
		r.Add(in)
		return true
	}
	return false
}
func day6() Result {
	// how many characters before the first 'start of packet' marker arrives?
	// the marker is four different characters
	s, c := getScanner("day6")
	defer c()
	s.Split(bufio.ScanBytes)
	rb := NewRingBuffer(14)
	for s.Scan() {
		rb.Add(s.Bytes()[0])
		if rb.Count >= 14 && rb.Unique() {
			break
		}
	}
	ret := Result{Part1: rb.Count}
	return ret
}

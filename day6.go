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

// returns true when unique
func (r *RingBuffer) AddUntilUnique(in byte) bool {
	if r.Count >= r.Size && r.Unique() {
		return true
	}
	r.Add(in)
	return false
}

func (r *RingBuffer) AddUntilUnique2(in byte) {
	if r.Count >= r.Size && r.Unique() {
		return
	}
	r.Add(in)
}

func day6() Result {
	// how many characters before the first 'start of packet' marker arrives?
	// the marker is four different characters
	s, c := getScanner("day6")
	defer c()
	s.Split(bufio.ScanBytes)
	rb := NewRingBuffer(4)
	rb2 := NewRingBuffer(14)
	for s.Scan() {
		theByte := s.Bytes()[0]
		l, r := rb.AddUntilUnique(theByte), rb2.AddUntilUnique(theByte)
		if l && r { // early return
			break
		}

	}
	ret := Result{Part1: rb.Count, Part2: rb2.Count}
	return ret
}

func day6v2() Result {
	// how many characters before the first 'start of packet' marker arrives?
	// the marker is four different characters
	s, c := getScanner("day6")
	defer c()
	s.Split(bufio.ScanBytes)
	rb := NewRingBuffer(4)
	rb2 := NewRingBuffer(14)
	for s.Scan() {
		theByte := s.Bytes()[0]
		rb.AddUntilUnique2(theByte)
		rb2.AddUntilUnique2(theByte)
	}
	ret := Result{Part1: rb.Count, Part2: rb2.Count}
	return ret
}

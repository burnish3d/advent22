package main

type Result struct {
	Part1 int
	Part2 int
}

func toPriority(character int32) int {
	if character >= 'a' {
		return int(character-'a') + 1
	}
	return int(character-'A') + 27
}

func toMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, e := range s {
		m[e] += 1
	}
	return m
}

func findMatches(left, right map[rune]int) []rune {
	r := make([]rune, 0, 10)
	for key, _ := range left {
		if _, ok := right[key]; ok {
			r = append(r, key)
		}
	}
	return r
}

func day3() Result {
	scanner, fileCloser := getScanner("./data/day3")
	defer fileCloser()
	runningTotal := 0
	for scanner.Scan() {
		l := scanner.Text()[0 : len(scanner.Text())/2]
		r := scanner.Text()[len(scanner.Text())/2 : len(scanner.Text())]
		lm, rm := toMap(l), toMap(r)
		matches := findMatches(lm, rm)
		// we only expect one match for this part of the prompt but we will sum any found matches
		// as it is written slightly more general
		for i := 0; i < len(matches); i++ {
			runningTotal += toPriority(matches[i])
		}

	}
	return Result{Part1: runningTotal}
}

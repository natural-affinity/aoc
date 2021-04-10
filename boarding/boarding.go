package boarding

import (
	"bufio"
	"errors"
	"math"
	"os"

	"github.com/natural-affinity/aoc/calendar"
)

// Problem identifier
var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 5: Binary Boarding"}

func region(runes []rune, max int, c1 string, c2 string) int {
	min := 0
	for i := 0; i < len(runes); i++ {
		char := string(runes[i])
		diff := int(math.Ceil(float64(max-min) / 2))

		if char == c1 {
			max = max - diff
		} else if char == c2 {
			min = min + diff
		}
	}

	return min
}

func FindHighest(path string) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, errors.New("invalid boarding pass list")
	}

	highest := -1
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		pass := scanner.Text()

		// TBD: pipeline (r -> c -> id)
		runes := []rune(pass)
		r, c := region(runes[0:7], 127, "F", "B"), region(runes[7:10], 7, "L", "R")

		id := r*8 + c
		if id > highest {
			highest = id
		}
	}

	return highest, nil
}

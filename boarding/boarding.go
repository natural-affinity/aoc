package boarding

import (
	"bufio"
	"errors"
	"math"
	"os"
	"sort"

	"github.com/natural-affinity/aoc/calendar"
)

// Problem identifier
var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 5: Binary Boarding"}

type Plane struct {
	Seats   []int
	Highest int
}

func (p *Plane) FindMySeat() (int, error) {
	for i := 0; i < len(p.Seats); i++ {
		if p.Seats[i+1]-p.Seats[i] == 2 {
			return p.Seats[i] + 1, nil
		}
	}

	return -1, nil
}

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

func FindHighest(path string) (*Plane, error) {
	plane := &Plane{Seats: []int{}}
	fp, err := os.Open(path)
	if err != nil {
		return plane, errors.New("invalid boarding pass list")
	}

	highest := -1
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		pass := scanner.Text()
		runes := []rune(pass)

		r, c := region(runes[:7], 127, "F", "B"), region(runes[7:], 7, "L", "R")

		id := r*8 + c
		plane.Seats = append(plane.Seats, id)

		if id > highest {
			highest = id
		}
	}

	plane.Highest = highest
	sort.Slice(plane.Seats, func(i, j int) bool {
		return plane.Seats[j] > plane.Seats[i]
	})

	return plane, nil
}

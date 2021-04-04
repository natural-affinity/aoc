package trajectory

import (
	"bufio"
	"errors"
	"os"

	"github.com/natural-affinity/aoc/calendar"
)

var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 3: Toboggan Trajectory"}

// Tree (# character)
const Tree rune = 35

type Trail struct {
	Forest [][]rune
}

func (t *Trail) Count(tree rune) (int, error) {
	x, y := 0, 0
	h := len(t.Forest)
	if h == 0 {
		return 0, errors.New("empty map")
	}

	w := len(t.Forest[0])
	count := 0
	for y < h-1 {
		x = (x + 3) % w
		y += 1

		if t.Forest[y][x] == tree {
			count += 1
		}
	}

	return count, nil
}

func Scout(path string) (*Trail, error) {
	fp, err := os.Open(path)
	if err != nil {
		return &Trail{}, errors.New("invalid forest map")
	}

	scanner := bufio.NewScanner(fp)
	trail := &Trail{}
	trail.Forest = [][]rune{}
	for scanner.Scan() {
		columns := []rune(scanner.Text())
		trail.Forest = append(trail.Forest, columns)
	}

	return trail, nil
}

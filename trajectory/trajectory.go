package trajectory

import (
	"bufio"
	"errors"
	"os"

	"github.com/natural-affinity/aoc/calendar"
)

var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 3: Toboggan Trajectory"}
var ErrEmptyMap = errors.New("empty map")
var ErrBadSlope = errors.New("invalid slope")

const Tree rune = 35 // (# character)

type Trail struct {
	forest [][]rune
}

type Slope struct {
	X, Y int
}

// Predict is a good candidate for parallel execution
func (t *Trail) Predict(slopes []Slope) (int, error) {
	p := 1
	for _, s := range slopes {
		count, err := t.Count(Tree, &s)
		if err != nil {
			return 0, err
		}

		p *= count
	}

	return p, nil
}

func (t *Trail) Count(tree rune, s *Slope) (int, error) {
	x, y := 0, 0
	h := len(t.forest)
	if h == 0 {
		return 0, ErrEmptyMap
	}

	if s == nil {
		return 0, ErrBadSlope
	}

	w := len(t.forest[0])
	count := 0
	for y < h-1 {
		x = (x + s.X) % w
		y += s.Y

		if t.forest[y][x] == tree {
			count += 1
		}
	}

	return count, nil
}

func Scout(path string) (*Trail, error) {
	fp, err := os.Open(path)
	if err != nil {
		return &Trail{}, err
	}
	defer fp.Close()

	trail := &Trail{forest: [][]rune{}}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		columns := []rune(scanner.Text())
		trail.forest = append(trail.forest, columns)
	}

	return trail, nil
}

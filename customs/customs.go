package customs

import (
	"bufio"
	"os"

	"github.com/natural-affinity/aoc/calendar"
)

var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 6: Custom Customs"}

func GroupSum(group string) int {
	var yes = make(map[rune]struct{})
	var exists = struct{}{}

	for _, c := range group {
		if _, ok := yes[c]; !ok {
			yes[c] = exists
		}
	}

	delete(yes, rune(32))
	return len(yes)
}

func Declarations(path string) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Split(calendar.SplitMulti)
	sum := 0
	for scanner.Scan() {
		g := GroupSum(scanner.Text())
		sum += g
	}

	return sum, nil
}

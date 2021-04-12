package customs

import (
	"bufio"
	"os"
	"strings"

	"github.com/natural-affinity/aoc/calendar"
)

var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 6: Custom Customs"}

const Space rune = 32

type SumRule func(form string) int

func SumAnyYes(group string) int {
	var exists = struct{}{}
	var yes = make(map[rune]struct{})

	for _, c := range group {
		if _, ok := yes[c]; !ok {
			yes[c] = exists
		}
	}

	delete(yes, Space)
	return len(yes)
}

func SumAllYes(group string) int {
	var yes = make(map[rune]int)
	size := len(strings.Split(group, " "))

	for _, c := range group {
		yes[c] += 1
	}

	delete(yes, Space)
	sum := 0
	for _, v := range yes {
		if v == size {
			sum += 1
		}
	}

	return sum
}

func SumDeclarations(path string, calc SumRule) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer fp.Close()

	sum := 0
	scanner := bufio.NewScanner(fp)
	scanner.Split(calendar.SplitMulti)
	for scanner.Scan() {
		sum += calc(scanner.Text())
	}

	return sum, nil
}

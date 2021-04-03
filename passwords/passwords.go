package passwords

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/natural-affinity/aoc/calendar"
)

// Problem Identifier
var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 2: Password Philosophy"}

func CountValid(path string) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, errors.New("invalid database")
	}

	scanner := bufio.NewScanner(fp)
	valid := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		rng := strings.Split(tokens[0], "-")
		min, _ := strconv.Atoi(rng[0])
		max, _ := strconv.Atoi(rng[1])

		runes := []rune(tokens[1])
		char := string(runes[0])

		count := strings.Count(tokens[2], char)
		if count >= min && count <= max {
			valid += 1
		}
	}

	return valid, nil
}

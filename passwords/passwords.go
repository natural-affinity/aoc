package passwords

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/natural-affinity/aoc/calendar"
)

// Problem Identifier
var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 2: Password Philosophy"}

// Policy for password validity
type Policy interface {
	IsValid(pass string, c *Criteria) (bool, error)
}

// Criteria characteristics
type Criteria struct {
	N1   int
	N2   int
	Char string
}

type OldPolicy struct{}

// IsValid if char is between min (N1) and max (N2)
func (op *OldPolicy) IsValid(pass string, c *Criteria) (bool, error) {
	count := strings.Count(pass, c.Char)
	return (count >= c.N1 && count <= c.N2), nil
}

type NewPolicy struct{}

// IsValid if char is at position (N1) XOR, position (N2)
func (np *NewPolicy) IsValid(pass string, c *Criteria) (bool, error) {
	runes := []rune(pass)
	c1 := string(runes[c.N1-1])
	c2 := string(runes[c.N2-1])

	return (c1 == c.Char) != (c2 == c.Char), nil
}

// Count valid passwords according to policy
func Count(path string, p Policy) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer fp.Close()

	valid := 0
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		rng := strings.Split(tokens[0], "-")
		n1, _ := strconv.Atoi(rng[0])
		n2, _ := strconv.Atoi(rng[1])

		runes := []rune(tokens[1])
		char := string(runes[0])

		c := &Criteria{N1: n1, N2: n2, Char: char}
		if ok, _ := p.IsValid(tokens[2], c); ok {
			valid += 1
		}
	}

	return valid, nil
}

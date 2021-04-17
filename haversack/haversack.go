package haversack

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const Gold = "shiny gold"

var TrimRegex = regexp.MustCompile(` bag(s)?\.?(\s|,\s)?(contain\s?)?`)
var BagRegex = regexp.MustCompile(`(\d+) ([a-z\s-a-z]+){1}`)

type Sack map[string]Bag
type Bag map[string]int
type Set map[string]struct{}

func HasColor(color string, b Bag, sack Sack, found Set, target string) {
	if _, ok := b[target]; ok {
		found[color] = struct{}{}
		return
	}

	for c := range b {
		HasColor(color, sack[c], sack, found, target)
	}
}

func (s Sack) CountColor(target string) int {
	var found = make(Set)
	for color, bag := range s {
		HasColor(color, bag, s, found, target)
	}

	return len(found)
}

func (s Sack) CountNested(b Bag) int {
	sum := 0
	next := []Bag{b}
	for {
		if len(next) == 0 {
			return sum
		}

		current := next[0]
		next = next[1:]
		for color, count := range current {
			sum += count
			for i := 0; i < count; i++ {
				next = append(next, s[color])
			}
		}
	}
}

func Parse(path string) (Sack, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var rules = make(Sack)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		line = TrimRegex.ReplaceAllString(line, ", ")

		colors := strings.Split(line, ",")
		first := strings.TrimSpace(colors[0])
		rules[first] = make(Bag)
		for _, c := range colors[1 : len(colors)-1] {
			var count int
			var color string

			matches := BagRegex.FindAllStringSubmatch(c, -1)
			if len(matches) > 0 {
				count, err = strconv.Atoi(matches[0][1])
				color = matches[0][2]
				if err != nil {
					return nil, err
				}
			} else {
				color = strings.TrimSpace(c)
			}

			rules[first][color] = count
		}
	}

	return rules, err
}

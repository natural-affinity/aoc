package haversack

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var TrimRegex = regexp.MustCompile(` bag(s)?\.?(\s|,\s)?(contain\s?)?(\d+)?`)
var Colors = regexp.MustCompile(`(?P<color>[a-z]+\s[a-z]+){1}(?:\sbag)`)

type Rules map[string]map[string]struct{}
type Cache map[string]struct{}

func HasGold(color string, set Cache, sack Rules, found Cache) {
	if _, ok := set["shiny gold"]; ok {
		found[color] = struct{}{}
		return
	}

	for c := range set {
		HasGold(color, sack[c], sack, found)
	}
}

func Count(sack Rules, target string) int {
	var found = make(Cache)
	for color, set := range sack {
		HasGold(color, set, sack, found)
	}

	return len(found)
}

func Parse(path string) (Rules, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var rules = make(Rules)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		line = TrimRegex.ReplaceAllString(line, ", ")

		colors := strings.Split(line, ",")
		first := strings.TrimSpace(colors[0])

		rules[first] = make(map[string]struct{})
		for _, c := range colors[1:] {
			color := strings.TrimSpace(c)
			rules[first][color] = struct{}{}
		}
	}

	return rules, err
}

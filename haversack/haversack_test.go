package haversack_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/haversack"
)

func TestCountColors(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
	}{
		{"sample", 4},
		{"rules", 151},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		rules, _ := haversack.Parse(p)
		result := haversack.Count(rules, "shiny gold")

		r := !(result == tc.Result)
		if r {
			t.Errorf("Case: %s, Expected: %v, Actual: %v", tc.Name, tc.Result, result)
		}
	}
}

package haversack_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/haversack"
	"github.com/natural-affinity/gotanda"
)

func TestCountBagsWithGold(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"not.found", 0, calendar.ErrFileNotFound},
		{"sample", 4, nil},
		{"rules", 151, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		rules, err := haversack.Parse(p)
		result := haversack.CountColor(rules, haversack.Gold)

		r := !(result == tc.Result)
		e := !(gotanda.CompareError(err, tc.Error))

		if r || e {
			t.Errorf("Case: %s, Expected: %v, %s Actual: %v, %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

func TestBagsInGold(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"not.found", 0, calendar.ErrFileNotFound},
		{"sample", 32, nil},
		{"sample2", 126, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		rules, err := haversack.Parse(p)
		result := haversack.CountNested(rules[haversack.Gold], rules)

		r := !(result == tc.Result)
		e := !(gotanda.CompareError(err, tc.Error))

		if r || e {
			t.Errorf("Case: %s, Expected: %v, %s Actual: %v, %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

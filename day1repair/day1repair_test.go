package day1repair_test

import (
	"errors"
	"path"
	"testing"

	"github.com/natural-affinity/aoc/day1repair"
	"github.com/natural-affinity/gotanda"
)

func TestProductOfTwo(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"one.match", 996075, nil},
		{"no.match", -1, errors.New("No numbers found")},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		report, m, _ := day1repair.ReadReport(p)
		result, err := day1repair.ProductOfTwo(report, m)

		r := !(result == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("\nTest: %s\nExpected:\n %d %s\nActual:\n %d %s",
				tc.Name, tc.Result,
				tc.Error, result,
				err)
		}
	}

}

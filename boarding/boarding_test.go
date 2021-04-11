package boarding_test

import (
	"errors"
	"path"
	"testing"

	"github.com/natural-affinity/aoc/boarding"
	"github.com/natural-affinity/gotanda"
)

func TestFindHighest(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"sample", 820, nil},
		{"passes", 858, nil},
		{"not.found", 0, errors.New("open testdata/not.found.input: The system cannot find the file specified.")},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := boarding.FindHighest(p)

		r := !(result.Highest == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if r || e {
			t.Errorf("Case: %s, Expected: %d %s, Actual: %d %s\n", tc.Name, tc.Result, tc.Error, result.Highest, err)
		}
	}
}

func TestFindMySeat(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
	}{
		{"passes", 557},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		plane, _ := boarding.FindHighest(p)
		result, _ := plane.FindMySeat()

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %d, Actual: %d\n", tc.Name, tc.Result, result)
		}
	}
}

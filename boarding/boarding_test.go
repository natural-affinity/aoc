package boarding_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/boarding"
)

func TestFindHighest(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
	}{
		{"sample", 820},
		{"passes", 858},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, _ := boarding.FindHighest(p)

		if result.Highest != tc.Result {
			t.Errorf("Case: %s, Expected: %d, Actual: %d\n", tc.Name, tc.Result, result.Highest)
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

package halting_test

import (
	"errors"
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/halting"
	"github.com/natural-affinity/gotanda"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		Name  string
		Lines int
		Error error
	}{
		{"not.found", 0, calendar.ErrFileNotFound},
		{"invalid.arg", 0, errors.New(`strconv.Atoi: parsing "+string": invalid syntax`)},
		{"sample", 9, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		result, err := halting.Load(p)

		r := !(result == nil || result.Lines() == tc.Lines)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("\nTest: %s\nExpected:\n %d %s\nActual:\n %v %s",
				tc.Name,
				tc.Lines, tc.Error,
				result, err)
		}
	}
}

func TestRunOnce(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"sample", 5, nil},
		{"boot", 1217, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		boot, _ := halting.Load(p)
		result, err := boot.RunOnce()

		r := !(result == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("Test: %s\nExpected:\n %d %s\nActual:\n %d %s",
				tc.Name,
				tc.Result, tc.Error,
				result, err)
		}
	}
}

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
		_, lines, err := halting.Load(p)

		r := !(lines == tc.Lines)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("\nTest: %s\nExpected:\n %d %s\nActual:\n %v %s",
				tc.Name,
				tc.Lines, tc.Error,
				lines, err)
		}
	}
}

func TestRun(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
	}{
		{"sample", 5},
		{"boot", 1217},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		boot, _, _ := halting.Load(p)
		_, acc := boot.Run()

		if !(acc == tc.Result) {
			t.Errorf("Test: %s\nExpected:\n %d\nActual:\n %d", tc.Name, tc.Result, acc)
		}
	}
}

func TestRepair(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"sample", 8, nil},
		{"boot", 501, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		boot, _, _ := halting.Load(p)
		acc, err := boot.Repair()

		r := !(acc == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if r || e {
			t.Errorf("Test: %s\nExpected:\n %d %s\nActual:\n %d %s", tc.Name, tc.Result, tc.Error, acc, err)
		}
	}
}

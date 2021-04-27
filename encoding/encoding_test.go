package encoding_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/encoding"
	"github.com/natural-affinity/gotanda"
)

func TestRead(t *testing.T) {
	cases := []struct {
		Name  string
		Count int
		Error error
	}{
		{"not.found", 0, calendar.ErrFileNotFound},
		{"sample", 20, nil},
		{"xmas", 1000, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := encoding.Read(p)

		r := !(len(result) == tc.Count)
		e := !(gotanda.CompareError(err, tc.Error))

		if r || e {
			t.Errorf("Case: %s, Want: %v, %s Got: %v, %s", tc.Name, tc.Count, tc.Error, result, err)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		Name   string
		Last   []int
		Num    int
		Result bool
	}{
		{"has.sum.1", []int{20, 5, 10, 4}, 9, true},
		{"has.sum.2", []int{1, 2, 7, 21}, 28, true},
		{"no.sum.1", []int{2, 4, 6, 8}, 16, false},
		{"no.sum.2", []int{20, 35, 27, 68}, 45, false},
		{"no.repeat", []int{10, 11, 12}, 20, false},
	}

	for _, tc := range cases {
		_, result := encoding.HasSum(tc.Num, tc.Last)

		if !(result == tc.Result) {
			t.Errorf("Case: %s, Want: %t Got: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestDecipher(t *testing.T) {
	cases := []struct {
		Name     string
		Preamble int
		Result   int
		Error    error
	}{
		{"sample", -1, -1, encoding.ErrBadPreamble},
		{"sample", 20, -1, encoding.ErrNumNotFound},
		{"sample", 5, 127, nil},
		{"xmas", 25, 14144619, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		xmas, _ := encoding.Read(p)
		result, err := encoding.Decipher(xmas, tc.Preamble)

		r := !(result == tc.Result)
		e := !(gotanda.CompareError(err, tc.Error))

		if r || e {
			t.Errorf("Case: %s, Want: %d, %s Got: %d, %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

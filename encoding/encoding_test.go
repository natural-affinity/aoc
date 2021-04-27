package encoding_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/encoding"
	"github.com/natural-affinity/gotanda"
)

var E = struct{}{}

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

		r := !(result.Len == tc.Count)
		e := !(gotanda.CompareError(err, tc.Error))

		if r || e {
			t.Errorf("Case: %s, Want: %v, %s Got: %v, %s", tc.Name, tc.Count, tc.Error, result, err)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		Name   string
		Last   map[int]struct{}
		Num    int
		Result bool
	}{
		{"has.sum.1", map[int]struct{}{20: E, 5: E, 10: E, 4: E}, 9, true},
		{"has.sum.2", map[int]struct{}{1: E, 2: E, 7: E, 21: E}, 28, true},
		{"no.sum.1", map[int]struct{}{2: E, 4: E, 6: E, 8: E}, 16, false},
		{"no.sum.2", map[int]struct{}{20: E, 35: E, 27: E, 68: E}, 45, false},
		{"no.repeat", map[int]struct{}{10: E, 11: E, 12: E}, 20, false},
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
		result, err := xmas.Decipher(tc.Preamble)

		r := !(result == tc.Result)
		e := !(gotanda.CompareError(err, tc.Error))

		if r || e {
			t.Errorf("Case: %s, Want: %d, %s Got: %d, %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

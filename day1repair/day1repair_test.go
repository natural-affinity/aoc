package day1repair_test

import (
	"errors"
	"path"
	"reflect"
	"testing"

	"github.com/natural-affinity/aoc/day1repair"
	"github.com/natural-affinity/gotanda"
)

func TestReadReport(t *testing.T) {
	cases := []struct {
		Name   string
		Report []int
		Set    map[int]struct{}
		Error  error
	}{
		{"read.empty.report", nil, map[int]struct{}{}, nil},
		{"read.partial.invalid.report", []int{1}, map[int]struct{}{1: {}}, errors.New("strconv.Atoi: parsing \"a\": invalid syntax")},
		{"read.invalid.path", nil, nil, errors.New("open testdata/read.invalid.path.input: The system cannot find the file specified.")},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		report, set, err := day1repair.ReadReport(p)
		r := !reflect.DeepEqual(report, tc.Report)
		s := !reflect.DeepEqual(set, tc.Set)
		e := !gotanda.CompareError(err, tc.Error)

		if r || s || e {
			if r {
				t.Errorf("\nTest: %s\nExpected:\n %v %v %s\nActual:\n %v %v %s",
					tc.Name,
					tc.Report, tc.Set, tc.Error,
					report, set, err)
			}
		}
	}
}

func TestProductOfTwo(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"product.of.two_one.match", 996075, nil},
		{"product.of.two_no.match", -1, errors.New("No numbers found")},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		rpt, set, _ := day1repair.ReadReport(p)
		result, err := day1repair.ProductOfTwo(rpt, set)

		r := !(result == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("\nTest: %s\nExpected:\n %d %s\nActual:\n %d %s",
				tc.Name,
				tc.Result, tc.Error,
				result, err)
		}
	}

}

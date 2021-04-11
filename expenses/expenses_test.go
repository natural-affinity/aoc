package expenses_test

import (
	"errors"
	"path"
	"reflect"
	"testing"

	"github.com/natural-affinity/aoc/expenses"
	"github.com/natural-affinity/gotanda"
)

func TestReadReport(t *testing.T) {
	cases := []struct {
		Name   string
		Report *expenses.Report
		Error  error
	}{
		{"read.empty.report", &expenses.Report{nil, map[int]struct{}{}}, nil},
		{"read.partial.invalid.report", &expenses.Report{[]int{1}, map[int]struct{}{1: {}}}, errors.New("strconv.Atoi: parsing \"a\": invalid syntax")},
		{"read.invalid.path", &expenses.Report{nil, nil}, errors.New("open testdata/read.invalid.path.input: The system cannot find the file specified.")},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		report, err := expenses.ReadReport(p)
		r := !reflect.DeepEqual(report.Itemized, tc.Report.Itemized)
		s := !reflect.DeepEqual(report.Unique, tc.Report.Unique)
		e := !gotanda.CompareError(err, tc.Error)

		if r || s || e {
			if r {
				t.Errorf("\nTest: %s\nExpected:\n %v %v %s\nActual:\n %v %v %s",
					tc.Name,
					tc.Report.Itemized, tc.Report.Unique, tc.Error,
					report.Itemized, report.Unique, err)
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
		{"product.match", 996075, nil},
		{"product.no.match", -1, expenses.ErrNoNum},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		report, _ := expenses.ReadReport(p)
		result, err := report.ProductOfTwo()

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

func TestProductOfThree(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"product.match", 51810360, nil},
		{"product.no.match", -1, expenses.ErrNoNum},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")

		report, _ := expenses.ReadReport(p)
		result, err := report.ProductOfThree()

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

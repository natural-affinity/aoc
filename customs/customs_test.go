package customs_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/customs"
	"github.com/natural-affinity/gotanda"
)

func TestDeclarations(t *testing.T) {
	cases := []struct {
		Name   string
		Rule   customs.SumRule
		Result int
		Error  error
	}{
		{"not.found", nil, -1, calendar.ErrFileNotFound},
		{"sample", customs.SumAnyYes, 11, nil},
		{"forms", customs.SumAnyYes, 6775, nil},
		{"forms", customs.SumAllYes, 3356, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := customs.SumDeclarations(p, tc.Rule)

		r := !(result == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if r || e {
			t.Errorf("Case: %s, Expected: %d, %s, Actual: %d, %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

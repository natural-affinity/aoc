package customs_test

import (
	"errors"
	"path"
	"testing"

	"github.com/natural-affinity/aoc/customs"
	"github.com/natural-affinity/gotanda"
)

func TestDeclarations(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"not.found", -1, errors.New("open testdata/not.found.input: The system cannot find the file specified.")},
		{"sample", 11, nil},
		{"forms", 6775, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := customs.Declarations(p)

		r := !(result == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if r || e {
			t.Errorf("Case: %s, Expected: %d, %s, Actual: %d, %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

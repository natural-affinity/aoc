package passport_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/passport"
	"github.com/natural-affinity/gotanda"
)

func TestCount(t *testing.T) {
	cases := []struct {
		Name      string
		Validator passport.Validator
		Result    int
		Error     error
	}{
		{"sample", passport.HasFields, 2, nil},
		{"batch", passport.HasFields, 235, nil},
		{"sample", passport.HasValidFields, 2, nil},
		{"batch", passport.HasValidFields, 194, nil},
		{"not.found", nil, -1, calendar.ErrFileNotFound},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := passport.Count(p, tc.Validator)

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

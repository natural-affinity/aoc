package passwords_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/calendar"
	"github.com/natural-affinity/aoc/passwords"
	"github.com/natural-affinity/gotanda"
)

func TestPolicy(t *testing.T) {
	cases := []struct {
		Name   string
		Policy passwords.Policy
		Result int
		Error  error
	}{
		{"not.found", &passwords.OldPolicy{}, -1, calendar.ErrFileNotFound},
		{"sample", &passwords.OldPolicy{}, 2, nil},
		{"sample", &passwords.NewPolicy{}, 1, nil},
		{"database", &passwords.OldPolicy{}, 586, nil},
		{"database", &passwords.NewPolicy{}, 352, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := passwords.Count(p, tc.Policy)

		r := !(result == tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("\nTest: %s\nExpected:\n %d %s\nActual:\n %d %s", tc.Name, tc.Result, tc.Error, result, err)
		}
	}
}

package passwords_test

import (
	"path"
	"testing"

	"github.com/natural-affinity/aoc/passwords"
	"github.com/natural-affinity/gotanda"
)

// test old policy
// test new policy

func TestCountValid(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"sample", 2, nil},
		{"database", 586, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := passwords.CountValid(p)

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
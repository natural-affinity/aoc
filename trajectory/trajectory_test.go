package trajectory_test

import (
	"errors"
	"path"
	"reflect"
	"testing"

	"github.com/natural-affinity/aoc/trajectory"
	"github.com/natural-affinity/gotanda"
)

func TestScout(t *testing.T) {
	cases := []struct {
		Name   string
		Result *trajectory.Trail
		Error  error
	}{
		{"invalid", &trajectory.Trail{}, errors.New("invalid forest map")},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		result, err := trajectory.Scout(p)

		r := !reflect.DeepEqual(result, tc.Result)
		e := !gotanda.CompareError(err, tc.Error)

		if e || r {
			t.Errorf("\nTest: %s\nExpected:\n %v %s\nActual:\n %v %s",
				tc.Name,
				tc.Result, tc.Error,
				result, err)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		Name   string
		Result int
		Error  error
	}{
		{"map", 187, nil},
		{"sample", 7, nil},
		{"empty", 0, errors.New("empty map")},
	}

	slope := &trajectory.Slope{X: 3, Y: 1}
	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		trail, _ := trajectory.Scout(p)
		result, err := trail.Count(trajectory.Tree, slope)

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

func TestPredict(t *testing.T) {
	cases := []struct {
		Name   string
		Slopes []trajectory.Slope
		Result int
		Error  error
	}{
		{"sample", []trajectory.Slope{{X: 3, Y: 1}}, 7, nil},
		{"map", []trajectory.Slope{
			{X: 1, Y: 1},
			{X: 3, Y: 1},
			{X: 5, Y: 1},
			{X: 7, Y: 1},
			{X: 1, Y: 2},
		}, 4723283400, nil},
	}

	for _, tc := range cases {
		p := path.Join("testdata", tc.Name+".input")
		trail, _ := trajectory.Scout(p)
		result, err := trail.Predict(tc.Slopes)

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

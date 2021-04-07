package passport_test

import (
	"testing"

	"github.com/natural-affinity/aoc/passport"
)

func TestIsValidBirthYear(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"1919", false},
		{"1920", true},
		{"2000", true},
		{"2002", true},
		{"2003", false},
		{"200", false},
		{"not.a.year", false},
	}

	for _, tc := range cases {
		result := passport.IsValidBirth(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestIsValidIssueYear(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"2009", false},
		{"2010", true},
		{"2011", true},
		{"2020", true},
		{"2021", false},
		{"200", false},
		{"not.a.year", false},
	}

	for _, tc := range cases {
		result := passport.IsValidIssue(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestIsValidExpiryYear(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"2019", false},
		{"2020", true},
		{"2022", true},
		{"2030", true},
		{"2031", false},
		{"200", false},
		{"not.a.year", false},
	}

	for _, tc := range cases {
		result := passport.IsValidExpiry(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestIsValidHeight(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"149cm", false},
		{"150cm", true},
		{"161cm", true},
		{"193cm", true},
		{"194cm", false},
		{"190", false},
		{"58in", false},
		{"59in", true},
		{"65in", true},
		{"76in", true},
		{"77in", false},
		{"not.a.number", false},
	}

	for _, tc := range cases {
		result := passport.IsValidHeight(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestIsValidHairColor(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
		{"#123456", true},
		{"#aaabbf", true},
		{"#12345", false},
	}

	for _, tc := range cases {
		result := passport.IsValidHairColor(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestIsValidEyeColor(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"amb", true},
		{"blu", true},
		{"brn", true},
		{"gry", true},
		{"grn", true},
		{"hzl", true},
		{"oth", true},
		{"1", false},
		{"abc", false},
	}

	for _, tc := range cases {
		result := passport.IsValidEyeColor(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

func TestIsValidPassportId(t *testing.T) {
	cases := []struct {
		Name   string
		Result bool
	}{
		{"000000001", true},
		{"123456789", true},
		{"012345678", true},
		{"a23456789", false},
		{"abcdefghi", false},
		{"1", false},
	}

	for _, tc := range cases {
		result := passport.IsValidPid(tc.Name)

		if result != tc.Result {
			t.Errorf("Case: %s, Expected: %t, Actual: %t", tc.Name, tc.Result, result)
		}
	}
}

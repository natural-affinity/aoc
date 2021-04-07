package passport

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Validator func(passport map[string]string, required map[string]FieldValidator) bool
type FieldValidator func(value string) bool

var required = map[string]FieldValidator{
	"byr": IsValidBirth,
	"iyr": IsValidIssue,
	"eyr": IsValidExpiry,
	"hgt": IsValidHeight,
	"hcl": IsValidHairColor,
	"ecl": IsValidEyeColor,
	"pid": IsValidPid,
}

func HasFields(passport map[string]string, required map[string]FieldValidator) bool {
	for f := range required {
		if _, ok := passport[f]; !ok {
			return false
		}
	}

	return true
}

func HasValidFields(passport map[string]string, required map[string]FieldValidator) bool {
	for f, isValid := range required {
		if v, ok := passport[f]; !ok || !isValid(v) {
			return false
		}
	}

	return true
}

func Count(path string, IsValid Validator) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, errors.New("invalid batch file")
	}

	scanner := bufio.NewScanner(fp)
	scanner.Split(ScanPassport)

	count := 0
	for scanner.Scan() {
		passport := scanner.Text()
		fields := strings.Split(passport, " ")
		mapped := make(map[string]string)

		for _, f := range fields {
			entry := strings.Split(f, ":")
			k, v := entry[0], entry[1]
			mapped[k] = v
		}

		if IsValid(mapped, required) {
			count += 1
		}
	}

	return count, nil
}

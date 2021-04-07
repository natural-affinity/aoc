package passport

import (
	"regexp"
	"strconv"
)

var eyeColors = map[string]struct{}{
	"amb": {},
	"blu": {},
	"brn": {},
	"gry": {},
	"grn": {},
	"hzl": {},
	"oth": {},
}

func IsValidBirth(y string) bool  { return isValidYear(y, 1920, 2002) }
func IsValidIssue(y string) bool  { return isValidYear(y, 2010, 2020) }
func IsValidExpiry(y string) bool { return isValidYear(y, 2020, 2030) }

func isValidYear(y string, min int, max int) bool {
	i, err := strconv.Atoi(y)
	if err != nil {
		return false
	}

	return (i >= min && i <= max)
}

func IsValidHeight(h string) bool {
	re := regexp.MustCompile(`(\d+)(cm|in)`)
	matches := re.FindStringSubmatch(h)

	if len(matches) != 3 {
		return false
	}

	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}

	if matches[2] == "cm" {
		return (i >= 150 && i <= 193)
	}

	return (i >= 59 && i <= 76)
}

func IsValidHairColor(c string) bool {
	re := regexp.MustCompile(`#([0-9a-f]){6}`)
	return re.Match([]byte(c))
}

func IsValidEyeColor(c string) bool {
	_, ok := eyeColors[c]
	return ok
}

func IsValidPid(id string) bool {
	if _, err := strconv.Atoi(id); err != nil {
		return false
	}

	return len(id) == 9
}

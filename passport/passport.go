package passport

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

var required = []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

func Scan(path string) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return -1, errors.New("invalid batch file")
	}

	re := regexp.MustCompile(`[a-z]{3}:`)
	scanner := bufio.NewScanner(fp)
	scanner.Split(ScanPassport)

	valid := 0
	for scanner.Scan() {
		passport := scanner.Text()
		fields := re.FindAllString(passport, -1)
		mapped := make(map[string]struct{})

		// concurrency candidates
		for _, f := range fields {
			mapped[f] = struct{}{}
		}

		for _, r := range required {
			if _, ok := mapped[r]; !ok {
				valid -= 1
				break
			}
		}

		valid += 1

		fmt.Println(passport)
		fmt.Println("----")
	}

	return valid, nil
}

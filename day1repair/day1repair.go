package day1repair

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func ReadReport(path string) ([]int, map[int]struct{}, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	var expenses []int
	m := make(map[int]struct{})
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return expenses, m, err
		}

		m[value] = struct{}{}
		expenses = append(expenses, value)
	}

	return expenses, m, nil
}

func ProductOfTwo(report []int, mapped map[int]struct{}) (int, error) {
	for _, v := range report {
		diff := 2020 - v
		if _, ok := mapped[diff]; ok == true {
			return (diff * v), nil
		}
	}

	return -1, errors.New("No numbers found")
}

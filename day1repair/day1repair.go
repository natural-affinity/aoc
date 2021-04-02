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

	var rpt []int
	set := make(map[int]struct{})

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return rpt, set, err
		}

		set[num] = struct{}{}
		rpt = append(rpt, num)
	}

	return rpt, set, nil
}

func ProductOfTwo(report []int, set map[int]struct{}) (int, error) {
	for _, v := range report {
		diff := 2020 - v
		if _, ok := set[diff]; ok == true {
			return (diff * v), nil
		}
	}

	return -1, errors.New("No numbers found")
}

package expenses

import (
	"bufio"
	"errors"
	"os"
	"strconv"

	"github.com/natural-affinity/aoc/calendar"
)

// Problem identifier
var Problem = &calendar.Puzzle{Event: 2020, Desc: "Day 1: Report Repair"}

type Report struct {
	Itemized []int
	Unique   map[int]struct{}
}

func ReadReport(path string) (Report, error) {
	fp, err := os.Open(path)
	if err != nil {
		return Report{}, err
	}

	var rpt []int
	set := make(map[int]struct{})

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return Report{Itemized: rpt, Unique: set}, err
		}

		set[num] = struct{}{}
		rpt = append(rpt, num)
	}

	return Report{Itemized: rpt, Unique: set}, nil
}

func (r *Report) ProductOfTwo() (int, error) {
	for _, v := range r.Itemized {
		diff := 2020 - v
		if _, ok := r.Unique[diff]; ok == true {
			return (diff * v), nil
		}
	}

	return -1, errors.New("No numbers found")
}
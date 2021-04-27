package encoding

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
)

var ErrBadPreamble = errors.New("invalid preamble length")
var ErrNumNotFound = errors.New("number not found")

func HasSum(n int, last []int) (int, bool) {
	for idx, a := range last {
		diff := int(math.Abs(float64(n - a)))

		for idx2, b := range last {
			if idx != idx2 && (diff-b == 0) {
				return n, true
			}
		}
	}

	return n, false
}

func Decipher(xmas []int, pre int) (int, error) {
	if pre < 0 {
		return -1, ErrBadPreamble
	}

	for i := pre; i < len(xmas); i++ {
		last := xmas[i-pre : i]

		if n, ok := HasSum(xmas[i], last); !ok {
			return n, nil
		}
	}

	return -1, ErrNumNotFound
}

func Read(path string) ([]int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var xmas []int
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return xmas, err
		}

		xmas = append(xmas, n)
	}

	return xmas, nil
}

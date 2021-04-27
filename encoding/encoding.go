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

type LastN map[int]struct{}
type Cipher struct {
	xmas     []int
	preamble LastN
	Len      int
}

func HasSum(n int, m LastN) (int, bool) {
	for a := range m {
		diff := int(math.Abs(float64(n - a)))

		if _, ok := m[diff]; ok && diff != a {
			return n, true
		}
	}

	return n, false
}

func (c *Cipher) Decipher(pre int) (int, error) {
	if pre < 0 {
		return -1, ErrBadPreamble
	}

	for i := 0; i < pre; i++ {
		n := c.xmas[i]
		c.preamble[n] = struct{}{}
	}

	for i := pre; i < len(c.xmas); i++ {
		prev, next := c.xmas[i-pre], c.xmas[i]

		if n, ok := HasSum(next, c.preamble); !ok {
			return n, nil
		}

		delete(c.preamble, prev)
		c.preamble[next] = struct{}{}
	}

	return -1, ErrNumNotFound
}

func Read(path string) (*Cipher, error) {
	fp, err := os.Open(path)
	if err != nil {
		return &Cipher{}, err
	}
	defer fp.Close()

	c := &Cipher{preamble: make(LastN)}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return c, err
		}

		c.xmas = append(c.xmas, n)
	}

	c.Len = len(c.xmas)
	return c, nil
}

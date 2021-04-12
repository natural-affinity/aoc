package calendar

import "errors"

var ErrFileNotFound = errors.New("open testdata/not.found.input: The system cannot find the file specified.")

type Puzzle struct {
	Desc  string
	Event int
}

package calendar

import "bytes"

// SplitMultiline is a modification of bufio.ScanLines (https://golang.org/src/bufio/scan.go?s=11967:12045#L340)
func SplitMulti(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		// We have a full newline-terminated line.
		d := bytes.Trim(data[0:i], "\n")
		d = bytes.ReplaceAll(d, []byte("\n"), []byte(" "))

		return i + 1, d, nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		d := bytes.Trim(data, "\n")
		d = bytes.ReplaceAll(d, []byte("\n"), []byte(" "))

		return len(data), d, nil
	}

	// Request more data.
	return 0, nil, nil
}

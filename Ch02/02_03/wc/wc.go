package wc

import (
	"errors"
	"io"
)

func readByte(r io.Reader) (byte, error) {
	var buf [1]byte
	if _, err := r.Read(buf[:]); err != nil {
		return 0, err
	}

	return buf[0], nil
}

func LineCount(r io.Reader) (int, error) {
	n := 0
	for {
		b, err := readByte(r)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return 0, err
		}

		if b == '\n' {
			n++
		}

	}

	return n, nil
}

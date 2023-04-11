package logs

import (
	"bufio"
	"encoding/json"
	"io"
	"time"
)

type Log struct {
	Time    time.Time `json:"time"`
	Level   string    `json:"level"`
	Message string    `json:"message"`
}

func ParseLogs(r io.Reader, handler func(Log)) (int, error) {
	s := bufio.NewScanner(r)
	count := 0
	for s.Scan() {
		count++

		var lm Log
		if err := json.Unmarshal(s.Bytes(), &lm); err != nil {
			return 0, err
		}

		handler(lm)
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

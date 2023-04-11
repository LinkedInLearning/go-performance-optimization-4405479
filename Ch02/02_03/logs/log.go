package logs

import (
	"sort"
	"time"
)

type Log struct {
	Time string // "Mar 19, 2023 14:32:45"
	// More fields
}

const (
	timeFmt = "Jan 02, 2006 15:04:05"
)

func logTime(log Log) time.Time {
	t, err := time.Parse(timeFmt, log.Time)
	if err != nil {
		panic(err)
	}
	return t
}

func sortByTime(logs []Log) {
	sort.Slice(logs, func(i, j int) bool {
		t1, t2 := logTime(logs[i]), logTime(logs[j])
		return t1.Before(t2)
	})
}

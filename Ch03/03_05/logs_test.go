package logs

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var logData = `2023-04-12T12:12:30Z - ERROR - error message number 1
2023-04-12T12:12:31Z - WARNING - warning message number 2
2023-04-12T12:12:32Z - INFO - info message number 5
`

func asTime(s string) time.Time {
	t, err := parseTime(s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestParseLogs(t *testing.T) {
	expected := []Log{
		{asTime("2023-04-12T12:12:30Z"), Error, "error message number 1"},
		{asTime("2023-04-12T12:12:31Z"), Warning, "warning message number 2"},
		{asTime("2023-04-12T12:12:32Z"), Info, "info message number 5"},
	}

	logs, err := ParseLogs(strings.NewReader(logData))
	require.NoError(t, err)
	require.Equal(t, expected, logs)
}

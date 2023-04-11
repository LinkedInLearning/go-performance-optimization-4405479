package logs

import (
	"os"
	"testing"
)

func handler(l Log) {}

const numLogs = 1024

func TestParseLogs(t *testing.T) {
	file, err := os.Open("testdata/logs.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	n, err := ParseLogs(file, handler)
	if err != nil {
		t.Fatal(err)
	}

	if n != numLogs {
		t.Fatalf("got %d logs, expected %d", n, numLogs)
	}
}

func BenchmarkParseLogs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("testdata/logs.jsonl")
		if err != nil {
			b.Fatal(err)
		}

		ParseLogs(file, handler)
		file.Close()
	}
}

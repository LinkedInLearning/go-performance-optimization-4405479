package logs

import (
	"compress/gzip"
	"os"
	"testing"
)

func BenchmarkParseLogs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		file, err := os.Open("app.log.gz")
		if err != nil {
			b.Fatal(err)
		}
		b.Cleanup(func() { file.Close() })
		r, err := gzip.NewReader(file)
		if err != nil {
			b.Fatal(err)
		}

		b.StartTimer()
		logs, err := ParseLogs(r)
		if err != nil {
			b.Fatal(err)
		}

		if len(logs) != 100_000 {
			b.Fatal(len(logs))
		}

	}
}

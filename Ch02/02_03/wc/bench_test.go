package wc

import (
	"os"
	"testing"
)

func BenchmarkLineCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		file, err := os.Open("sherlock.txt")
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		n, err := LineCount(file)
		file.Close()
		if err != nil {
			b.Fatal(err)
		}
		if n != 12310 {
			b.Fatal(n)
		}
	}
}

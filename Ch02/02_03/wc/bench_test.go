package wc

import (
	"os"
	"testing"
)

var (
	file *os.File
)

func init() {
	var err error
	file, err = os.Open("sherlock.txt")
	if err != nil {
		panic(err)
	}
}

func BenchmarkLineCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file.Seek(0, os.SEEK_SET)

		n, err := LineCount(file)
		if err != nil {
			b.Fatal(err)
		}
		if n != 12310 {
			b.Fatal(n)
		}
	}
}

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CompressFile(src, dest string) (int64, error) {
	in, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	w := gzip.NewWriter(out)
	return io.Copy(w, in)
}

func CompressDir(srcDir, outDir string) (int, int, error) {
	matches, err := filepath.Glob(fmt.Sprintf("%s/*.txt", srcDir))
	if err != nil {
		return 0, 0, err
	}

	var size int64
	for _, src := range matches {
		dest := fmt.Sprintf("%s/%s.gz", outDir, filepath.Base(src))
		n, err := CompressFile(src, dest)
		if err != nil {
			return 0, 0, err
		}
		size += n
	}

	return len(matches), int(size), nil
}

func main() {
	start := time.Now()

	count, size, err := CompressDir("books", "gzipped")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	duration := time.Since(start)
	fmt.Printf("%d files (%d bytes) compressed in %v\n", count, size, duration)
}

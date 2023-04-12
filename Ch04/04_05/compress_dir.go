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

type result struct {
	err  error
	size int64
}

func CompressDir(srcDir, outDir string) (int, int, error) {
	matches, err := filepath.Glob(fmt.Sprintf("%s/*.txt", srcDir))
	if err != nil {
		return 0, 0, err
	}

	ch := make(chan result)
	for _, src := range matches {
		dest := fmt.Sprintf("%s/%s.gz", outDir, filepath.Base(src))
		src := src
		go func() {
			var r result
			r.size, r.err = CompressFile(src, dest)
			ch <- r
		}()
	}

	var size int64
	for range matches {
		r := <-ch
		if r.err != nil {
			return 0, 0, r.err
		}
		size += r.size
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

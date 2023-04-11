package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func URLTime(url string) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%q: bad status - %s", url, resp.Status)
	}

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return 0, err
	}

	return time.Since(start), nil
}

func main() {
	urls := []string{
		"https://go.dev",
		"https://pkg.go.dev",
		"https://www.ardanlabs.com/",
	}

	start := time.Now()
	for _, url := range urls {
		duration, err := URLTime(url)
		status := fmt.Sprintf("%v", duration)
		if err != nil {
			status = fmt.Sprintf("error (%s)", err)
		}
		fmt.Printf("%q: %s\n", url, status)
	}
	duration := time.Since(start)
	fmt.Printf("%d URLs in %v\n", len(urls), duration)
}

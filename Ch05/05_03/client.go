//go:build ignore

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var c http.Client
	url := "http://localhost:8080"

	if os.Getenv("SOCK_TYPE") == "unix" {
		c.Transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/httpd.sock")
			},
		}
		url = "http://x/"
	}

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < 1000; n++ {
				resp, err := c.Get(url)
				if err != nil {
					log.Fatalf("error: %s", err)
				}
				if resp.StatusCode != http.StatusOK {
					log.Fatalf("error: bad status - %s", resp.Status)
				}
				if _, err := io.Copy(io.Discard, resp.Body); err != nil {
					log.Fatalf("error: %s", err)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

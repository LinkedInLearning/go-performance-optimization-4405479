package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Followers(ctx context.Context, login string) (int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%q - bad status: %s", login, resp.Status)
	}

	var reply struct {
		Followers int
	}
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return 0, err
	}

	return reply.Followers, nil
}

type result struct {
	login string
	count int
	err   error
}

func Influence(ctx context.Context, logins []string) (int, error) {
	ch := make(chan result)

	for _, login := range logins {
		login := login
		go func() {
			r := result{login: login}
			r.count, r.err = Followers(ctx, login)
			ch <- r
		}()
	}

	count := 0
	for range logins {
		r := <-ch
		if r.err != nil {
			return 0, r.err
		}
		count += r.count

	}

	return count, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := time.Now()
	logins := []string{"pisush", "matryer", "bradfitz", "robpike", "rakyll", "davecheney"}
	count, err := Influence(ctx, logins)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	duration := time.Since(start)
	fmt.Printf("count: %d (%v)\n", count, duration)
}

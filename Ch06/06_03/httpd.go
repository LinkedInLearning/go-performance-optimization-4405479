package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func utcHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tz, err := time.LoadLocation(query.Get("tz"))
	if err != nil {
		http.Error(w, "bad time zone", http.StatusBadRequest)
		return
	}

	const timeFmt = "2006-01-02T15:04:05"
	local, err := time.ParseInLocation(timeFmt, query.Get("when"), tz)
	if err != nil {
		http.Error(w, "bad time", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, local.UTC().Format(timeFmt))
}

func timingMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			duration := time.Since(start)
			log.Printf("info: timing: %s %s - %v", r.Method, r.URL.Path, duration)
		}()

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func main() {
	http.Handle("/utc", timingMiddleware(http.HandlerFunc(utcHandler)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

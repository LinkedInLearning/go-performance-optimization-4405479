package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	db    *DB
	cache *redis.Client
}

func (s *Server) handleItem(w http.ResponseWriter, r *http.Request) {
	sku := chi.URLParam(r, "sku")
	if sku == "" {
		http.Error(w, "bad sku", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if s.cache != nil {
		data, err := s.cache.Get(r.Context(), sku).Bytes()
		if err == nil {
			w.Write(data)
			return
		}
	}

	item, err := s.db.Get(sku)
	if err != nil {
		log.Printf("warning: sku %q not found", sku)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		log.Printf("error: can't marshal JSON for %q - %s", sku, err)
		http.Error(w, "can't marshal", http.StatusInternalServerError)
		return
	}

	if s.cache != nil {
		if err := s.cache.Set(r.Context(), sku, data, time.Hour).Err(); err != nil {
			log.Printf("warning: can't set cache for %q - %s", sku, err)
		}
	}
	w.Write(data)
}

func main() {
	db, err := NewDB("postgres://postgres@localhost/?sslmode=disable")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer db.Close()

	srv := Server{
		db: db,
	}
	if os.Getenv("USE_CACHE") != "" {
		log.Printf("info: using redis cache")
		srv.cache = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
	}

	mux := chi.NewRouter()
	mux.Get("/item/{sku}", srv.handleItem)

	addr := ":8080"
	log.Printf("info: server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("error: %s", err)
	}
}

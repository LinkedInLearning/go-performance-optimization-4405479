package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	db *DB
}

func (s *Server) handlePrice(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	if symbol == "" {
		http.Error(w, "missing symbol", http.StatusBadRequest)
		return
	}

	price, ok := s.db.Price(symbol)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := map[string]any{
		"symbol": symbol,
		"price":  price,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error: can't encode %#v - %s", resp, err)
	}
}

func (s *Server) handleBatchPrice(w http.ResponseWriter, r *http.Request) {
	symbols := r.URL.Query()["symbol"]
	if len(symbols) == 0 {
		http.Error(w, "missing symbols", http.StatusBadRequest)
		return
	}

	type stock struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price"`
	}
	stocks := make([]stock, 0, len(symbols))
	for _, symbol := range symbols {
		price, ok := s.db.Price(symbol)
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		stocks = append(stocks, stock{symbol, price})
	}

	if err := json.NewEncoder(w).Encode(stocks); err != nil {
		log.Printf("error: can't encode %#v - %s", stocks, err)
	}
}

func main() {
	s := Server{
		db: NewDB(),
	}

	mux := chi.NewRouter()
	mux.Get("/price/{symbol}", s.handlePrice)
	mux.Get("/batch/price", s.handleBatchPrice)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("error: can't start server - %s", err)
	}
}

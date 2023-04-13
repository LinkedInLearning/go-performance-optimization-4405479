package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	c := Client{baseURL}
	return &c
}

func (c *Client) Price(ctx context.Context, symbol string) (float64, error) {
	url := fmt.Sprintf("%s/price/%s", c.baseURL, url.PathEscape(symbol))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status - %s", resp.Status)
	}

	var reply struct {
		Symbol string
		Price  float64
	}

	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return 0, err
	}

	return reply.Price, nil
}

func (c *Client) batchPriceURL(symbols []string) (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/batch/price", c.baseURL))
	if err != nil {
		return "", err
	}
	q := u.Query()
	for _, sym := range symbols {
		q.Add("symbol", sym)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (c *Client) BatchPrice(ctx context.Context, symbols []string) (map[string]float64, error) {
	url, err := c.batchPriceURL(symbols)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status - %s", resp.Status)
	}

	var reply []struct {
		Symbol string
		Price  float64
	}

	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for _, s := range reply {
		prices[s.Symbol] = s.Price
	}
	return prices, nil

}

func main() {
	symbols := []string{
		"AAPL", "ABNB", "ADBE", "ADI", "ADP", "ADSK", "AEP", "ALGN", "AMAT", "AMD", "AMGN", "ANSS",
		"VRSK", "VRTX", "WBA", "WBD", "WDAY", "XEL", "ZM", "ZS",
	}
	c := NewClient("http://localhost:8080")

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	m, err := c.BatchPrice(ctx, symbols)
	cancel()

	if err != nil {
		log.Fatalf("error: %s", err)
	}
	duration := time.Since(start)
	fmt.Println(m)
	fmt.Printf("%d symbols in %v\n", len(symbols), duration)
}

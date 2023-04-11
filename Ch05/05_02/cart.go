package cart

import "time"

type Item struct {
	SKU    string
	Amount int
	Price  float64
}

type Cart struct {
	User    string
	Updated time.Time
	Items   []Item
}

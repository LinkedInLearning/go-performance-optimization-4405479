package cart

import (
	"encoding/json"
	"testing"
	"time"
)

var (
	cart = Cart{
		User:    "joe",
		Updated: time.Date(2023, time.January, 19, 14, 52, 30, 0, time.UTC),
		Items: []Item{
			{SKU: "hammer19", Amount: 1, Price: 3.7},
			{SKU: "nail7", Amount: 100, Price: 0.01},
			{SKU: "glue6", Amount: 2, Price: 2.3},
		},
	}
)

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(cart)
		if err != nil {
			b.Fatal(err)
		}
	}
}

package ads

import (
	"math/rand"
	"testing"
)

var imps []Impression

func init() {
	const size = 1_000_000
	imps = make([]Impression, size)
	for i := 0; i < size; i++ {
		imps[i].Price = rand.Float64()
	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := Total(imps)
		if v <= 0 {
			b.Fatal(v)
		}
	}
}

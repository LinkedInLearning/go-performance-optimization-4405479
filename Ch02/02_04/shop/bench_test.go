package shop

import (
	"fmt"
	"testing"
)

const (
	sku = "sku-93"
)

func tempDB(b *testing.B, useCache bool) *DB {
	tmp := b.TempDir()
	dbFile := fmt.Sprintf("%s/store.db", tmp)
	db, err := NewDB(dbFile, useCache)
	if err != nil {
		b.Fatal(err)
	}
	i := Item{SKU: sku}
	err = db.Set(i)
	if err != nil {
		b.Fatal(err)
	}

	return db
}

func BenchmarkGet(b *testing.B) {
	db := tempDB(b, false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Get(sku)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetCached(b *testing.B) {
	db := tempDB(b, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := db.Get(sku)
		if err != nil {
			b.Fatal(err)
		}
	}
}

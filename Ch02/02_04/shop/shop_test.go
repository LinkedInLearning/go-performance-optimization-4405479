package shop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetGet(t *testing.T) {
	dbFile := fmt.Sprintf("%s/shop.db", t.TempDir())
	t.Logf("db at %q", dbFile)

	db, err := NewDB(dbFile, false)
	require.NoError(t, err)
	defer db.Close()

	for i := 0; i < 10; i++ {
		sku := fmt.Sprintf("sku-%02d", i)
		err := db.Set(Item{SKU: sku})
		require.NoErrorf(t, err, "set: sku = %q", sku)
	}

	sku := "sku-07"
	i, err := db.Get(sku)
	require.NoErrorf(t, err, "get: sku = %q", sku)
	require.Equal(t, sku, i.SKU)
}

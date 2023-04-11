package shop

import (
	"bytes"
	"encoding/gob"
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.etcd.io/bbolt"
)

var (
	bucketName = []byte("items")
)

type Item struct {
	SKU string
	// More fields
}

type DB struct {
	conn  *bbolt.DB
	cache *lru.Cache[string, Item] // SKU -> Item
}

func NewDB(dbFile string, useCache bool) (*DB, error) {
	conn, err := bbolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		return nil, err
	}

	db := DB{
		conn: conn,
	}

	if !useCache {
		return &db, nil
	}

	db.cache, err = lru.New[string, Item](1024)
	if err != nil {
		db.conn.Close()
		return nil, err
	}
	return &db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func marshalItem(i Item) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(i); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func unmarshalItem(data []byte, i *Item) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(i)
}

func (db *DB) Set(i Item) error {
	data, err := marshalItem(i)
	if err != nil {
		return err
	}

	err = db.conn.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Put([]byte(i.SKU), data)
	})
	if err != nil {
		return err
	}

	if db.cache != nil {
		db.cache.Add(i.SKU, i)
	}

	return nil
}

func (db *DB) Get(sku string) (Item, error) {
	if db.cache != nil {
		i, ok := db.cache.Get(sku)
		if ok {
			return i, nil
		}
	}

	var data []byte

	err := db.conn.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		value := b.Get([]byte(sku))
		if value != nil {
			data = make([]byte, len(value))
			copy(data, value)
		}
		return nil
	})

	if err != nil {
		return Item{}, err
	}
	if data == nil {
		return Item{}, fmt.Errorf("item %q not found", sku)
	}

	var i Item
	if err := unmarshalItem(data, &i); err != nil {
		return Item{}, err
	}

	if db.cache != nil {
		db.cache.Add(i.SKU, i)
	}

	return i, nil
}

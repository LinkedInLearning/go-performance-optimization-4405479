package main

import (
	"database/sql"
	_ "embed"

	_ "github.com/lib/pq"
)

var (
	//go:embed sql/get.sql
	getSQL string
)

type Item struct {
	SKU string
	// TODO: More fields
}

type DB struct {
	conn *sql.DB
}

func NewDB(dsn string) (*DB, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db := DB{
		conn: conn,
	}
	return &db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) Get(sku string) (Item, error) {
	row := db.conn.QueryRow(getSQL, sku)
	if err := row.Err(); err != nil {
		return Item{}, err
	}

	var i Item
	if err := row.Scan(&i.SKU); err != nil {
		return Item{}, err
	}
	return i, nil
}

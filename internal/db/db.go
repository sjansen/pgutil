package db

import (
	"github.com/go-pg/pg"
)

type DB struct {
	pool *pg.DB
}

func New(options map[string]string) (*DB, error) {
	merged, err := mergeOptions(options)
	if err != nil {
		return nil, err
	}
	db := &DB{
		pool: pg.Connect(merged),
	}
	return db, nil
}

func (db *DB) Close() error {
	return db.pool.Close()
}

func (db *DB) ServerVersion() (string, error) {
	var version string
	_, err := db.pool.Query(pg.Scan(&version), "SELECT VERSION()")
	if err != nil {
		return "", err
	}
	return version, nil
}

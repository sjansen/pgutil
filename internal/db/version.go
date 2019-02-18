package db

import "github.com/go-pg/pg"

func (db *DB) ServerVersion() (string, error) {
	var version string
	_, err := db.pool.Query(pg.Scan(&version), "SELECT VERSION()")
	if err != nil {
		return "", err
	}
	return version, nil
}

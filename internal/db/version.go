package db

import "github.com/go-pg/pg"

func (c *Connection) ServerVersion() (string, error) {
	var version string
	_, err := c.db.Query(pg.Scan(&version), "SELECT VERSION()")
	if err != nil {
		return "", err
	}
	return version, nil
}

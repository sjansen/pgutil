package db

import (
	"os"
	"os/user"

	"github.com/go-pg/pg"
)

type Connection struct {
	db *pg.DB
}

func New(config map[string]string) (*Connection, error) {
	options := &pg.Options{}
	if addr, ok := config["addr"]; ok {
		options.Addr = addr
	} else if host := os.Getenv("PGHOST"); host != "" {
		if port := os.Getenv("PGPORT"); port != "" {
			options.Addr = host + ":" + port
		} else {
			options.Addr = host + ":5432"
		}
	}

	if database, ok := config["database"]; ok {
		options.Database = database
	} else if database := os.Getenv("PGDATABASE"); database != "" {
		options.Database = database
	}

	if username, ok := config["user"]; ok {
		options.User = username
	} else if username := os.Getenv("PGUSER"); username != "" {
		options.User = username
	} else if u, err := user.Current(); err != nil {
		return nil, err
	} else {
		options.User = u.Username
	}

	if password, ok := config["password"]; ok {
		options.Password = password
	} else if password := os.Getenv("PGPASSWORD"); password != "" {
		options.Password = password
	}

	c := &Connection{
		db: pg.Connect(options),
	}
	return c, nil
}

func (c *Connection) Close() error {
	return c.db.Close()
}

func (c *Connection) ServerVersion() (string, error) {
	var version string
	_, err := c.db.Query(pg.Scan(&version), "SELECT VERSION()")
	if err != nil {
		return "", err
	}
	return version, nil
}

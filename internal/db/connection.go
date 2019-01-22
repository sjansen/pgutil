package db

import (
	"os"
	"os/user"

	"github.com/go-pg/pg"
)

type Connection struct {
	db *pg.DB
}

func New(options map[string]string) (*Connection, error) {
	merged, err := mergeOptions(options)
	if err != nil {
		return nil, err
	}
	c := &Connection{
		db: pg.Connect(merged),
	}
	return c, nil
}

func (c *Connection) Close() error {
	return c.db.Close()
}

func mergeOptions(options map[string]string) (*pg.Options, error) {
	result := &pg.Options{}

	if addr, ok := options["addr"]; ok {
		result.Addr = addr
	} else if host := os.Getenv("PGHOST"); host != "" {
		if port := os.Getenv("PGPORT"); port != "" {
			result.Addr = host + ":" + port
		} else {
			result.Addr = host + ":5432"
		}
	}

	if database, ok := options["database"]; ok {
		result.Database = database
	} else if database := os.Getenv("PGDATABASE"); database != "" {
		result.Database = database
	}

	if username, ok := options["user"]; ok {
		result.User = username
	} else if username := os.Getenv("PGUSER"); username != "" {
		result.User = username
	} else if u, err := user.Current(); err != nil {
		return nil, err
	} else {
		result.User = u.Username
	}

	if password, ok := options["password"]; ok {
		result.Password = password
	} else if password := os.Getenv("PGPASSWORD"); password != "" {
		result.Password = password
	}

	return result, nil
}

package commands

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/db"
)

type PingCmd struct {
	Host     string
	Port     string
	DBName   string
	Username string
}

func (c *PingCmd) Run() error {
	options := map[string]string{}

	if c.Host != "" {
		if c.Port != "" {
			options["addr"] = c.Host + ":" + c.Port
		} else {
			options["addr"] = c.Host + ":5432"
		}
	}
	if c.DBName != "" {
		options["database"] = c.DBName
	}
	if c.Username != "" {
		options["user"] = c.Username
	}

	db, err := db.New(options)
	if err != nil {
		return err
	}

	version, err := db.ServerVersion()
	if err == nil {
		fmt.Println(version)
	}
	return err
}

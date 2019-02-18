package commands

import (
	"fmt"
	"io"
)

type PingCmd struct {
	Host     string
	Port     string
	DBName   string
	Username string
}

func (c *PingCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	db, err := deps.DB(c.dbOptions())
	if err != nil {
		return err
	}
	defer db.Close()

	version, err := db.ServerVersion()
	if err == nil {
		fmt.Fprintln(stdout, version)
	}
	return err
}

func (c *PingCmd) dbOptions() map[string]string {
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

	return options
}

package commands

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/pg"
)

type PingCmd struct {
	Host     string
	Port     uint16
	DBName   string
	Username string

	RequireTLS bool
}

func (c *PingCmd) Run(base *Base) error {
	conn, err := pg.New(&pg.Options{
		Log: base.Log,

		Host:     c.Host,
		Port:     c.Port,
		Username: c.Username,
		Database: c.DBName,

		RequireTLS: c.RequireTLS,
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	version, err := conn.ServerVersion()
	if err == nil {
		fmt.Fprintln(base.Stdout, version)
	}
	return err
}

package commands

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/pg"
)

type PingCmd struct {
	Host     string
	Port     string
	DBName   string
	Username string
}

func (c *PingCmd) Run(base *Base) error {
	opts := c.pgOptions()
	opts.Log = base.Log
	p, err := pg.New(opts)
	if err != nil {
		return err
	}
	defer p.Close()

	version, err := p.ServerVersion()
	if err == nil {
		fmt.Fprintln(base.Stdout, version)
	}
	return err
}

func (c *PingCmd) pgOptions() *pg.Options {
	options := &pg.Options{}
	if c.Host != "" {
		if c.Port != "" {
			options.Address = c.Host + ":" + c.Port
		} else {
			options.Address = c.Host + ":5432"
		}
	}
	if c.DBName != "" {
		options.Database = c.DBName
	}
	if c.Username != "" {
		options.Username = c.Username
	}

	return options
}

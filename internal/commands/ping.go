package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/pg"
)

type PingCmd struct {
	Host     string
	Port     string
	DBName   string
	Username string

	Debug     *os.File
	Verbosity int
}

func (c *PingCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	p, err := pg.New(c.pgOptions())
	if err != nil {
		return err
	}
	defer p.Close()

	version, err := p.ServerVersion()
	if err == nil {
		fmt.Fprintln(stdout, version)
	}
	return err
}

func (c *PingCmd) pgOptions() *pg.Options {
	options := &pg.Options{}
	if c.Debug != nil {
		options.Log = logger.New(c.Verbosity, os.Stderr, c.Debug)
	} else {
		options.Log = logger.New(c.Verbosity, os.Stderr, nil)
	}

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

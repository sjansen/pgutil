package commands

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/pg"
)

// A PingCmd is a request to test database connection options
type PingCmd struct {
	Host     string
	Port     uint16
	SSLMode  string
	DBName   string
	Username string
}

// Run executes the command
func (c *PingCmd) Run(base *Base) error {
	conn, err := pg.New(&pg.Options{
		Log: base.Log,

		Host:     c.Host,
		Port:     c.Port,
		SSLMode:  c.SSLMode,
		Username: c.Username,
		Database: c.DBName,
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

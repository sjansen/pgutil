package commands

import (
	"os"

	"github.com/sjansen/pgutil/internal/ddl"
	"github.com/sjansen/pgutil/internal/pg"
)

// An InspectCmd is a request to describe a database
type InspectCmd struct {
	Host     string
	Port     uint16
	SSLMode  string
	DBName   string
	Username string

	Output string
}

// Run executes the command
func (c *InspectCmd) Run(base *Base) error {
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

	db, err := conn.InspectDatabase()
	if err != nil {
		return err
	}

	w := base.Stdout
	if c.Output != "" && c.Output != "-" {
		w, err = os.OpenFile(c.Output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
	}

	return ddl.Write(w, db)
}

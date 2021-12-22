package commands

import (
	"context"
	"os"

	"github.com/chzyer/readline"

	"github.com/sjansen/pgutil/internal/pg"
)

// An InspectCmd is a request to describe a database
type InspectCmd struct {
	Host     string
	Port     uint16
	SSLMode  string
	DBName   string
	Username string
	Password bool

	Output      string
	SortChecks  bool
	SortColumns bool
	SortIndexes bool
}

// Run executes the command
func (c *InspectCmd) Run(base *Base) error {
	var password string
	if c.Password {
		response, err := readline.Password("Enter password for " + c.Username + ": ")
		if err != nil {
			return err
		}
		password = string(response)
	}

	ctx := context.TODO()
	conn, err := pg.New(ctx, &pg.Options{
		Log: base.Log,

		Host:     c.Host,
		Port:     c.Port,
		SSLMode:  c.SSLMode,
		Username: c.Username,
		Password: password,
		Database: c.DBName,
	})
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	db, err := conn.InspectDatabase(ctx, &pg.InspectOptions{
		SortChecks:  c.SortChecks,
		SortColumns: c.SortColumns,
		SortIndexes: c.SortIndexes,
	})
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

	return db.Write(w)
}

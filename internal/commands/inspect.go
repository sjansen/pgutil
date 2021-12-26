package commands

import (
	"context"
	"os"
	"time"

	"github.com/chzyer/readline"

	"github.com/sjansen/pgutil/internal/ddl"
	"github.com/sjansen/pgutil/internal/pg"
)

// An InspectCmd is a request to describe a database
type InspectCmd struct {
	Host     string
	Port     uint16
	SSLMode  string
	Database string
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
		Database: c.Database,
	})
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	m, err := getMetadata(ctx, conn)
	if err != nil {
		return err
	}

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

	return db.WriteHCL(w, m)
}

func getMetadata(ctx context.Context, conn *pg.Conn) (*ddl.DatabaseMetadata, error) {
	version, err := conn.ServerVersion(ctx)
	if err != nil {
		return nil, err
	}

	m := &ddl.DatabaseMetadata{
		Host:          conn.Host,
		Database:      conn.Database,
		ServerVersion: version,
		Timestamp:     time.Now(),
	}

	return m, nil
}

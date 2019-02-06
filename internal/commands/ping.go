package commands

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/sjansen/pgutil/internal/db"
)

type pingCmd struct {
	host     string
	port     string
	dbname   string
	username string
	// TODO password string
}

func (c *pingCmd) register(app *kingpin.Application) {
	cmd := app.Command("ping", "Create a test connection to a database").Action(c.run)
	cmd.Flag("host", "database server host or socket directory").Short('h').StringVar(&c.host)
	cmd.Flag("port", "database server port").Short('p').StringVar(&c.port)
	cmd.Flag("dbname", "database server port").Short('d').StringVar(&c.dbname)
	cmd.Flag("username", "connect as username").Short('U').StringVar(&c.username)
}

func (c *pingCmd) run(pc *kingpin.ParseContext) error {
	options := map[string]string{}

	if c.host != "" {
		if c.port != "" {
			options["addr"] = c.host + ":" + c.port
		} else {
			options["addr"] = c.host + ":5432"
		}
	}
	if c.dbname != "" {
		options["database"] = c.dbname
	}
	if c.username != "" {
		options["user"] = c.username
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

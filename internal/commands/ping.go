package commands

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/db"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

/* TODO connection args
-h host
--host=host
-p port
--port=port
-d dbname
--dbname=dbname
-U username
--username=username
-P username
--password=password
*/
type pingCmd struct {
}

func (v *pingCmd) register(app *kingpin.Application) {
	app.Command("ping", "Create a test connection to a database").Action(v.run)
}

func (v *pingCmd) run(pc *kingpin.ParseContext) error {
	c, err := db.New(nil)
	if err != nil {
		return err
	}

	version, err := c.ServerVersion()
	if err == nil {
		fmt.Println(version)
	}
	return err
}

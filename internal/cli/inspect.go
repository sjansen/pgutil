package cli

import (
	"github.com/sjansen/pgutil/internal/commands"
)

func registerInspect(p *ArgParser) {
	c := &commands.InspectCmd{}
	cmd := p.addCommand(c, "inspect", "Scan a database and output its configuration")

	cmd.Flag("host", "server hostname or socket directory").Short('h').StringVar(&c.Host)
	cmd.Flag("port", "server port number").Short('p').Uint16Var(&c.Port)
	cmd.Flag("dbname", "database name").Short('d').StringVar(&c.DBName)
	cmd.Flag("username", "connect as username").Short('U').StringVar(&c.Username)
	cmd.Flag("password", "force password prompt").Short('W').BoolVar(&c.Password)
	cmd.Flag("sslmode", "connection security level").EnumVar(&c.SSLMode,
		"disable", "allow", "prefer", "require", "verify-full",
	)

	cmd.Flag("output", "Write to FILENAME instead of stdout").
		Short('o').PlaceHolder("FILENAME").StringVar(&c.Output)
	cmd.Flag("sort-columns", "List columns in alphanumeric order, instead of physical order").
		BoolVar(&c.SortColumns)
	cmd.Flag("sort-check-constraints", "List check constraints order by expression, instead of name").
		BoolVar(&c.SortChecks)
	cmd.Flag("sort-indexes", "List indexes in alphanumeric order, instead of default order").
		BoolVar(&c.SortIndexes)
}

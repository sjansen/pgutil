package cli

import (
	"github.com/sjansen/pgutil/internal/commands"
)

func registerInspect(p *ArgParser) {
	c := &commands.InspectCmd{}
	cmd := p.addCommand(c, "inspect", "Scan a datbase and output its configuration")

	cmd.Flag("output", "Write to FILENAME instead of stdout").
		Short('o').PlaceHolder("FILENAME").StringVar(&c.Output)

	cmd.Flag("host", "server hostname or socket directory").Short('h').StringVar(&c.Host)
	cmd.Flag("port", "server port number").Short('p').Uint16Var(&c.Port)
	cmd.Flag("dbname", "database name").Short('d').StringVar(&c.DBName)
	cmd.Flag("username", "connect as username").Short('U').StringVar(&c.Username)
	// password is deliberately excluded because it is a security risk
	cmd.Flag("sslmode", "connection security level").EnumVar(&c.SSLMode,
		"disable", "allow", "prefer", "require", "verify-full",
	)
}

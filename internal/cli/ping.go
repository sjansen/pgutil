package cli

import (
	"github.com/sjansen/pgutil/internal/commands"
)

func registerPing(p *ArgParser) {
	c := &commands.PingCmd{}
	cmd := p.addCommand(c, "ping", "Create a test connection to a database")
	cmd.Flag("host", "server hostname or socket directory").Short('h').StringVar(&c.Host)
	cmd.Flag("port", "server port number").Short('p').Uint16Var(&c.Port)
	cmd.Flag("dbname", "database name").Short('d').StringVar(&c.DBName)
	cmd.Flag("username", "connect as username").Short('U').StringVar(&c.Username)
	cmd.Flag("require-tls", "require tls").BoolVar(&c.RequireTLS)
}

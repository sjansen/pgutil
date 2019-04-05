package cli

import (
	"os"

	"github.com/sjansen/pgutil/internal/commands"
)

func registerPing(p *ArgParser) {
	c := &commands.PingCmd{}
	cmd := p.addCommand(c, "ping", "Create a test connection to a database")
	cmd.Flag("host", "server hostname or socket directory").Short('h').StringVar(&c.Host)
	cmd.Flag("port", "server port number").Short('p').StringVar(&c.Port)
	cmd.Flag("dbname", "database name").Short('d').StringVar(&c.DBName)
	cmd.Flag("username", "connect as username").Short('U').StringVar(&c.Username)
	cmd.Flag("verbose", "verbose mode, repeat to increase verbosity").Short('v').CounterVar(&c.Verbosity)
	cmd.Flag("debug", "debug log").OpenFileVar(&c.Debug, os.O_CREATE|os.O_WRONLY, 0644)
}

package cli

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/sjansen/pgutil/internal/commands"
)

// RegisterCommands creates the app's standard ArgParser
func RegisterCommands(version string) *ArgParser {
	app := kingpin.
		New("pgutil", "Tools for PostgreSQL").
		UsageTemplate(kingpin.CompactUsageTemplate)

	base := &commands.Base{}
	app.Flag("debug", "debug log").
		OpenFileVar(&base.Debug, os.O_CREATE|os.O_WRONLY, 0644)
	app.Flag("verbose", "verbose mode, repeat to increase verbosity").
		Short('v').CounterVar(&base.Verbosity)

	parser := &ArgParser{
		app:  app,
		base: base,
	}
	registerPing(parser)
	registerRunbook(parser)
	registerVersion(parser, version)
	return parser
}

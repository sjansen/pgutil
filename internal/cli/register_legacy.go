// +build !go1.12

package cli

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func RegisterCommands(version string) *ArgParser {
	app := kingpin.
		New("pgutil", "Tools for PostgreSQL").
		UsageTemplate(kingpin.CompactUsageTemplate)
	parser := &ArgParser{app: app}
	registerPing(parser)
	registerRunbook(parser)
	registerVersion(parser, version)
	return parser
}

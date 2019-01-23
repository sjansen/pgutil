package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/sjansen/pgutil/internal/commands"
)

var build string // set by goreleaser

func main() {
	app := kingpin.
		New("pgutil", "Tools for PostgreSQL").
		UsageTemplate(kingpin.CompactUsageTemplate)
	if build != "" {
		commands.Register(app, build)
	} else {
		commands.Register(app, version)
	}

	if len(os.Args) == 1 {
		app.Usage(os.Args[1:])
		os.Exit(1)
	}

	kingpin.MustParse(app.Parse(os.Args[1:]))
}

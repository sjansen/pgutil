// +build go1.12

package cli

import (
	"runtime/debug"

	"github.com/sjansen/pgutil/internal/commands"
)

func registerVersion(p *ArgParser, build string, info *debug.BuildInfo) {
	c := &commands.VersionCmd{
		App:   "pgutil",
		Build: build,

		BuildInfo: info,
		Verbose:   false,
	}
	cmd := p.addCommand(c, "version", "Print pgutil's version")
	if info != nil {
		cmd.Flag("verbose", "include build details").Short('v').BoolVar(&c.Verbose)
	}
}

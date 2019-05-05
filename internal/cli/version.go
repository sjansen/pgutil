// +build go1.12

package cli

import (
	"runtime/debug"

	"github.com/sjansen/pgutil/internal/commands"
)

func registerVersion(p *ArgParser, build string) {
	info, _ := debug.ReadBuildInfo()
	c := &commands.VersionCmd{
		App:   "pgutil",
		Build: build,

		BuildInfo: info,
		Verbose:   false,
	}
	cmd := p.addCommand(c, "version", "Print pgutil's version")
	cmd.Flag("long", "include build details").Short('l').BoolVar(&c.Verbose)
}

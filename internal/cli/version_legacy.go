// +build !go1.12

package cli

import "github.com/sjansen/pgutil/internal/commands"

func registerVersion(p *ArgParser, build string) {
	c := &commands.VersionCmd{
		App:   "pgutil",
		Build: build,
	}
	p.addCommand(c, "version", "Print pgutil's version")
}

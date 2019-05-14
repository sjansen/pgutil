package cli

import (
	"io"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/logger"
)

// An ArgParser contains the definitions of flags, arguments and commands for an application
type ArgParser struct {
	app     *kingpin.Application
	base    *commands.Base
	cmd     command
	version string
}

// Command is an abstraction for a function that produces human readable output when run
type Command func(stdout, stderr io.Writer) error

type command interface {
	Run(base *commands.Base) error
}

// Parse converts command line arguments into a Command
func (p *ArgParser) Parse(args []string) (Command, error) {
	_, err := p.app.Parse(args)
	if err != nil {
		return nil, err
	}
	fn := func(stdout, stderr io.Writer) error {
		if p.base.Debug == nil {
			p.base.Log = logger.New(p.base.Verbosity, stderr, nil)
		} else {
			p.base.Log = logger.New(p.base.Verbosity, stderr, p.base.Debug)
		}
		p.base.Stdout = stdout
		p.base.Stderr = stderr
		p.base.Log.Infof("pgutil version=%s", p.version)
		return p.cmd.Run(p.base)
	}
	return fn, nil
}

func (p *ArgParser) addCommand(cmd command, name, help string) *kingpin.CmdClause {
	clause := p.app.Command(name, help)
	clause.Action(func(pc *kingpin.ParseContext) error {
		p.cmd = cmd
		return nil
	})
	return clause
}

func (p *ArgParser) addParent(name string) *kingpin.CmdClause {
	return p.app.Command(name, "")
}

func (p *ArgParser) addSubCommand(parent *kingpin.CmdClause, cmd command, name, help string) *kingpin.CmdClause {
	clause := parent.Command(name, help)
	clause.Action(func(pc *kingpin.ParseContext) error {
		p.cmd = cmd
		return nil
	})
	return clause
}

package cli

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type ArgParser struct {
	app *kingpin.Application
	cmd Command
}

type Command interface {
	Run() error
}

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

func (p *ArgParser) Parse(args []string) (Command, error) {
	_, err := p.app.Parse(args)
	if err != nil {
		return nil, err
	}
	return p.cmd, nil
}

func (p *ArgParser) addCommand(cmd Command, name, help string) *kingpin.CmdClause {
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

func (p *ArgParser) addSubCommand(parent *kingpin.CmdClause, cmd Command, name, help string) *kingpin.CmdClause {
	clause := parent.Command(name, help)
	clause.Action(func(pc *kingpin.ParseContext) error {
		p.cmd = cmd
		return nil
	})
	return clause
}

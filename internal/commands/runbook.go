package commands

import (
	"github.com/sjansen/pgutil/internal/runbook"
)

type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

func (c *RunBookRunCmd) Run(base *Base) error {
	return runbook.Run(c.File, base.Stdout, base.Stderr)
}

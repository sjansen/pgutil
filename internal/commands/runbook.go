package commands

import (
	"io"

	"github.com/sjansen/pgutil/internal/runbook"
)

type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

func (c *RunBookRunCmd) Run(stdout, stderr io.Writer, impl *Dependencies) error {
	return runbook.Run(c.File, stdout, stderr)
}

package commands

import (
	"github.com/sjansen/pgutil/internal/runbook"
)

// A RunBookRunCmd is a request to execute the tasks in a runbook
type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

// Run executes the command
func (c *RunBookRunCmd) Run(base *Base) error {
	return runbook.Run(&base.IO, c.File)
}

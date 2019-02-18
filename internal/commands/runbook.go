package commands

import (
	"fmt"
	"io"
)

type RunBookListCmd struct {
	File string
}

type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

func (c *RunBookRunCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	fmt.Fprintln(stdout, "run:", c.File)
	return nil
}

func (c *RunBookListCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	fmt.Fprintln(stdout, "list:", c.File)
	return nil
}

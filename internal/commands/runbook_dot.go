package commands

import (
	"os"
	"path/filepath"

	"github.com/sjansen/pgutil/internal/runbook"
)

// A RunBookDotCmd is a request to describe how a runbook's tasks are related
type RunBookDotCmd struct {
	File    string
	Output  string
	AutoOut bool // base output filename on input filename
	Splines string
}

// Run executes the command
func (c *RunBookDotCmd) Run(base *Base) (err error) {
	output := "-"
	switch {
	case c.Output != "":
		output = c.Output
	case c.AutoOut:
		src := c.File
		ext := filepath.Ext(src)
		output = src[:len(src)-len(ext)] + ".dot"
	}

	w := base.Stdout
	if output != "-" {
		w, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
	}
	return runbook.Dot(&base.IO, c.File, w, c.Splines)
}

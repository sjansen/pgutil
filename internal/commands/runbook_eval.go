package commands

import (
	"os"

	isatty "github.com/mattn/go-isatty"

	"github.com/sjansen/pgutil/internal/runbook"
)

// A RunBookEvalCmd is a request to convert a runbook to pretty-printed JSON
type RunBookEvalCmd struct {
	File   string
	Color  string
	Output string
}

// Run executes the command
func (c *RunBookEvalCmd) Run(base *Base) (err error) {
	output := "-"
	if c.Output != "" {
		output = c.Output
	}

	w := base.Stdout
	if output != "-" {
		w, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
	}

	color := false
	switch c.Color {
	case "auto":
		if x, ok := w.(interface{ Fd() uintptr }); ok {
			color = isatty.IsTerminal(x.Fd())
		}
	case "true":
		color = true
	}
	return runbook.Eval(&base.IO, c.File, w, color)
}

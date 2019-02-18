package commands

import (
	"fmt"
	"io"
	"sort"

	"github.com/sjansen/pgutil/internal/runbook"
)

type RunBookListCmd struct {
	File string
}

func (c *RunBookListCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	cfg, err := runbook.Load(c.File)
	if err != nil {
		return err
	}

	tasks := make([]string, 0, len(cfg.Tasks))
	for t := range cfg.Tasks {
		tasks = append(tasks, t)
	}
	sort.Strings(tasks)

	fmt.Fprintln(stdout, "Tasks:")
	for _, task := range tasks {
		fmt.Fprintf(stdout, "    %s\n", task)
	}
	return nil
}

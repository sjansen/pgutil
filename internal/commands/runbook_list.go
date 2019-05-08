package commands

import (
	"fmt"
	"sort"

	"github.com/sjansen/pgutil/internal/runbook"
)

type RunBookListCmd struct {
	File string
}

func (c *RunBookListCmd) Run(base *Base) error {
	tasks, err := runbook.List(&base.IO, c.File)
	if err != nil {
		return err
	}

	longest := 0
	sorted := make([]string, 0, len(tasks))
	for taskID := range tasks {
		s := string(taskID)
		sorted = append(sorted, s)
		if len(s) > longest {
			longest = len(s)
		}
	}
	sort.Strings(sorted)

	fmt.Fprintln(base.Stdout, "Tasks & Targets:")
	for _, task := range sorted {
		target := tasks[runbook.TaskID(task)]
		fmt.Fprintf(base.Stdout, "  %-*s  %s\n", longest, task, target)
	}
	return nil
}

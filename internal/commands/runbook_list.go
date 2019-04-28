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
	tasks, err := runbook.List(c.File)
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

	fmt.Fprintln(stdout, "Tasks & Targets:")
	for _, task := range sorted {
		target := tasks[runbook.TaskID(task)]
		fmt.Fprintf(stdout, "  %-*s  %s\n", longest, task, target)
	}
	return nil
}

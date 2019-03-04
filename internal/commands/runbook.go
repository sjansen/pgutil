package commands

import (
	"context"
	"io"

	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/tasks"
	"github.com/sjansen/pgutil/internal/tasks/dispatcher"
	"github.com/sjansen/pgutil/internal/tasks/exec"
	"github.com/sjansen/pgutil/internal/tasks/sql"
)

type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

func (c *RunBookRunCmd) Run(stdout, stderr io.Writer, impl *Dependencies) error {
	cfg, err := runbook.Load(c.File)
	if err != nil {
		return err
	}

	db, err := impl.DB(nil)
	if err != nil {
		return err
	}
	defer db.Close()

	deps := make(map[string][]string, len(cfg.Tasks))
	for id, task := range cfg.Tasks {
		deps[id] = task.After
	}

	tasks := map[string]tasks.Task{}
	for id, task := range cfg.Tasks {
		switch {
		case task.Exec != nil:
			tasks[id] = &exec.Task{
				Args:   task.Exec.Args,
				Stdout: stdout,
				Stderr: stderr,
			}
		case task.SQL != nil:
			tasks[id] = &sql.Task{
				C:   db,
				SQL: task.SQL.SQL,
			}
		}
	}

	dispatcher := &dispatcher.Dispatcher{
		Workers: 2,
		Deps:    deps,
		Tasks:   tasks,
	}

	ctx := context.Background()
	_, err = dispatcher.Dispatch(ctx)
	return err
}

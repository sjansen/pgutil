package commands

import (
	"fmt"
	"io"

	"github.com/sjansen/pgutil/internal/graphs"
	"github.com/sjansen/pgutil/internal/runbook"
)

type RunBookRunCmd struct {
	File string

	//DryRun      bool
	//Interactive bool
	//ListTasks   bool
	//Tasks       []string
}

func (c *RunBookRunCmd) Run(stdout, stderr io.Writer, deps *Dependencies) error {
	cfg, taskOrder, err := foo(c.File)
	if err != nil {
		return err
	}

	db, err := deps.DB(nil)
	if err != nil {
		return err
	}
	defer db.Close()

	tasks := cfg.Tasks
	for _, name := range taskOrder {
		fmt.Fprintf(stdout, "begin: %s\n", name)
		task := tasks[name]
		switch {
		case task.Exec != nil:
			p := deps.Process(task.Exec.Args)
			p.Run(stdout, stderr)
		case task.SQL != nil:
			db.Exec(task.SQL.SQL)
		}
		fmt.Fprintf(stdout, "end: %s\n", name)
	}

	return nil
}

func foo(filename string) (cfg *runbook.Config, taskOrder []string, err error) {
	cfg, err = runbook.Load(filename)
	if err != nil {
		return nil, nil, err
	}

	tasks := cfg.Tasks
	nodes := make(map[string][]string, len(tasks))
	for name, node := range tasks {
		nodes[name] = node.After
	}

	g, err := graphs.NewDependencyGraph(
		graphs.NewDirectedGraph(nodes),
	)
	if err != nil {
		return nil, nil, err
	}

	taskOrder, cycle := g.TSort()
	if cycle != nil {
		err = fmt.Errorf("cycle detected: %v", cycle)
		return nil, nil, err
	}

	return cfg, taskOrder, nil
}

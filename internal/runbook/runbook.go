package runbook

import (
	"context"
	"io"

	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/scheduler"
	"github.com/sjansen/pgutil/internal/runbook/strbuf"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

func Run(filename string, stdout, stderr io.Writer) error {
	parser := newParser(stdout, stderr)

	runbook, err := parser.Parse(filename)
	if err != nil {
		return err
	}

	completed := newCompletedChan(runbook.Targets)
	defer close(completed)

	ready := startScheduler(runbook.Targets, runbook.Tasks, completed)

	ctx := context.TODO()
	for taskID := range ready {
		task := runbook.Tasks[string(taskID)]
		target := runbook.Targets[task.Target]
		err = target.Handle(ctx, task.Config)
		if err != nil {
			return err
		}
		completed <- taskID
	}

	return nil
}

func newCompletedChan(targets types.Targets) chan types.TaskID {
	capacity := 0
	for _, t := range targets {
		capacity += t.ConcurrencyLimit()
	}
	capacity *= 2
	return make(chan types.TaskID, capacity)
}

func newParser(stdout, stderr io.Writer) *parser.Parser {
	return &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"strbuf": &strbuf.TargetFactory{
				StdOut: stdout,
			},
		},
	}
}

func startScheduler(targets types.Targets, tasks types.Tasks, completed <-chan types.TaskID) <-chan types.TaskID {
	ch := make(chan types.TaskID)

	go func(ch chan<- types.TaskID) {
		s, ready, err := scheduler.New(targets, tasks)
		for {
			if err != nil {
				break
			}
			for _, taskIDs := range ready {
				for _, taskID := range taskIDs {
					ch <- types.TaskID(taskID)
				}
			}
			if taskID, ok := <-completed; ok {
				ready, err = s.Next(string(taskID))
			} else {
				break
			}
		}
		close(ch)
		for range completed {
			// drain completed
		}
	}(ch)

	return ch
}

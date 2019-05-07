package runbook

import (
	"context"
	"io"
	"sync"

	dot "github.com/awalterschulze/gographviz"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/pg"
	"github.com/sjansen/pgutil/internal/runbook/scheduler"
	"github.com/sjansen/pgutil/internal/runbook/sh"
	"github.com/sjansen/pgutil/internal/runbook/strbuf"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

// TargetID uniquely identifies a target
type TargetID string

// TaskID uniquely identifies a task
type TaskID string

type readyTask struct {
	target TargetID
	taskID TaskID
	task   types.TaskConfig
}

type endedTask struct {
	taskID TaskID
	err    error
}

// Generate a GraphViz compatible description of a runbook's tasks
func Dot(filename string, w io.Writer, splines string) error {
	parser := newParser(nil, nil)
	runbook, err := parser.Parse(filename)
	if err != nil {
		return err
	}

	g := dot.NewEscape()
	g.SetDir(true)
	g.SetName("runbook")
	g.AddAttr("runbook", "newrank", "true")
	if splines != "" {
		g.AddAttr("runbook", "splines", splines)
	}

	for targetID := range runbook.Targets {
		graphID := "cluster_" + targetID
		g.AddSubGraph("runbook", graphID, nil)
		g.AddAttr(graphID, "label", targetID)
	}

	for dstID, task := range runbook.Tasks {
		graphID := "cluster_" + task.Target
		g.AddNode(graphID, dstID, nil)
		for _, srcID := range task.After {
			g.AddEdge(srcID, dstID, true, nil)
		}
	}

	w.Write([]byte(g.String()))
	return nil
}

// List enumerates a runbook's tasks and their targets
func List(filename string) (map[TaskID]TargetID, error) {
	parser := newParser(nil, nil)
	runbook, err := parser.Parse(filename)
	if err != nil {
		return nil, err
	}

	result := map[TaskID]TargetID{}
	for taskID, task := range runbook.Tasks {
		result[TaskID(taskID)] = TargetID(task.Target)
	}

	return result, nil
}

// Run executes the tasks in a runbook
func Run(filename string, stdout, stderr io.Writer) error {
	parser := newParser(stdout, stderr)
	runbook, err := parser.Parse(filename)
	if err != nil {
		return err
	}

	completed := newCompletedChan(runbook.Targets)
	defer close(completed)

	ctx := context.TODO()
	ready := startScheduler(runbook.Targets, runbook.Tasks, completed)
	ended := startTargets(ctx, runbook.Targets, ready)

	for x := range ended {
		if x.err != nil {
			return err
		}
		completed <- x.taskID
	}

	return nil
}

func newCompletedChan(targets types.Targets) chan TaskID {
	capacity := 0
	for _, t := range targets {
		capacity += t.ConcurrencyLimit()
	}
	capacity *= 2
	return make(chan TaskID, capacity)
}

func newParser(stdout, stderr io.Writer) *parser.Parser {
	return &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"pg": &pg.TargetFactory{
				Log: logger.Discard(), // TODO
			},
			"sh": &sh.TargetFactory{
				Stdout: stdout,
				Stderr: stderr,
			},
			"strbuf": &strbuf.TargetFactory{
				Stdout: stdout,
			},
		},
	}
}

func startScheduler(targets types.Targets, tasks types.Tasks, completed <-chan TaskID) <-chan *readyTask {
	capacity := 0
	for _, t := range targets {
		capacity += t.ConcurrencyLimit()
	}
	capacity *= 2
	ch := make(chan *readyTask, capacity)

	go func(ch chan<- *readyTask) {
		defer func() {
			close(ch)
			for range completed {
				// drain completed
			}
		}()

		s, ready, err := scheduler.New(targets, tasks)
		for {
			if err != nil {
				break
			}
			for targetID, taskIDs := range ready {
				for _, taskID := range taskIDs {
					ch <- &readyTask{
						target: TargetID(targetID),
						taskID: TaskID(taskID),
						task:   tasks[taskID].Config,
					}
				}
			}
			if taskID, ok := <-completed; ok {
				ready, err = s.Next(string(taskID))
			} else {
				break
			}
		}
	}(ch)

	return ch
}

func startTarget(
	ctx context.Context, wg *sync.WaitGroup, target types.Target,
	ready <-chan *readyTask, ended chan<- *endedTask,
) {
	go func() {
		target.Start()
		defer target.Stop()
		defer wg.Done()

		for r := range ready {
			err := target.Handle(ctx, r.task)
			if err != nil {
				ended <- &endedTask{taskID: r.taskID, err: err}
				break
			}
			ended <- &endedTask{taskID: r.taskID}
		}

		for range ready {
			// drain ready
		}
	}()
}

func startTargets(ctx context.Context, targets types.Targets, ready <-chan *readyTask) <-chan *endedTask {
	ended := make(chan *endedTask)

	var wg sync.WaitGroup
	channels := make(map[TargetID]chan<- *readyTask)
	for targetID, target := range targets {
		ready := make(chan *readyTask)
		channels[TargetID(targetID)] = ready
		startTarget(ctx, &wg, target, ready, ended)
		wg.Add(1)
	}

	go func() {
		for r := range ready {
			ch := channels[r.target]
			ch <- r
		}
		for _, ch := range channels {
			close(ch)
		}
		wg.Wait()
		close(ended)
	}()

	return ended
}

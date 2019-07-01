package base

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/taskset/types"
)

type Task struct {
	Target  string
	After   []string `hcl:"after,optional"`
	Provide []string `hcl:"provide,optional"`
	Require []string `hcl:"require,optional"`
}

func (t *Task) Dependencies() []string {
	deps := make([]string, 0, len(t.After)+len(t.Require))
	for _, id := range t.After {
		deps = append(deps, "task:"+id)
	}
	return append(deps, t.Require...)
}

func (t *Task) Provides() []string {
	return t.Provide
}

type RunTask func(types.TaskID, types.Task, chan<- types.TaskResults)

func RunTasks(fn RunTask, maxConcurrency int) (chan<- types.TaskBatch, <-chan types.TaskResults) {
	queue := make(chan types.TaskBatch, maxConcurrency)
	results := make(chan types.TaskResults, maxConcurrency)

	go func(queue <-chan types.TaskBatch, results chan<- types.TaskResults) {
		relay := make(chan types.TaskResults)
		running := 0

		for {
			select {
			case tasks, ok := <-queue:
				if ok {
					for id, task := range tasks {
						fmt.Println(id)
						fn(id, task, relay)
						running++
					}
				} else {
					goto shutdown
				}
			case result := <-relay:
				results <- result
				running--
			}
		}
	shutdown:
		for result := range relay {
			results <- result
			running--
			if running < 1 {
				close(relay)
				close(results)
			}
		}
	}(queue, results)

	return queue, results
}

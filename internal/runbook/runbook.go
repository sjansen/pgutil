package runbook

import (
	"github.com/sjansen/pgutil/internal/runbook/scheduler"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

func newCompletedChan(targets types.Targets) chan types.TaskID {
	capacity := 0
	for _, t := range targets {
		capacity += t.ConcurrencyLimit()
	}
	capacity *= 2
	return make(chan types.TaskID, capacity)
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

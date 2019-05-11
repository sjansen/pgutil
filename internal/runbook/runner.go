package runbook

import (
	"context"
	"sync"

	"github.com/sjansen/pgutil/internal/runbook/scheduler"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

type readyMsg struct {
	target TargetID
	taskID TaskID
	task   types.TaskConfig
}

type endedMsg struct {
	taskID TaskID
	err    error
}

type runner struct {
	ctx     context.Context
	targets types.Targets
	tasks   types.Tasks

	done  chan TaskID    // tasks that ended without an error
	ended chan *endedMsg // tasks that ended, possibly with an error
	ready chan *readyMsg // tasks that can be started
}

func newRunner(targets types.Targets, tasks types.Tasks) *runner {
	capacity := 0
	for _, t := range targets {
		capacity += t.ConcurrencyLimit()
	}
	capacity *= 2

	return &runner{
		ctx:     context.TODO(),
		targets: targets,
		tasks:   tasks,

		// producers are responsible for closing
		done:  make(chan TaskID, capacity),    // closed by run()
		ended: make(chan *endedMsg),           // closed by startTargets()
		ready: make(chan *readyMsg, capacity), // closed by startScheduler()
	}
}

func (r *runner) run() error {
	defer close(r.done)

	r.startScheduler()
	err := r.startTargets()
	if err != nil {
		return err
	}

	for x := range r.ended {
		if x.err != nil {
			return err
		}
		r.done <- x.taskID
	}

	return nil
}

func (r *runner) startScheduler() {
	go func(readyChan chan<- *readyMsg) {
		defer close(readyChan)

		s, ready, err := scheduler.New(r.targets, r.tasks)
		for {
			if err != nil {
				break
			}
			for targetID, taskIDs := range ready {
				for _, taskID := range taskIDs {
					readyChan <- &readyMsg{
						target: TargetID(targetID),
						taskID: TaskID(taskID),
						task:   r.tasks[taskID].Config,
					}
				}
			}
			if taskID, ok := <-r.done; ok {
				ready, err = s.Next(string(taskID))
			} else {
				break
			}
		}
	}(r.ready)
}

func (r *runner) startTargets() (err error) {
	wg := &sync.WaitGroup{}
	var channels = make(map[TargetID]chan<- *readyMsg)
	for targetID, target := range r.targets {
		err = target.Start()
		if err != nil {
			break
		}

		wg.Add(1)
		ch := make(chan *readyMsg)
		startTargetGoroutine(r.ctx, wg, target, ch, r.ended)
		channels[TargetID(targetID)] = ch
	}

	// start goroutine, even if there's an error, because it's
	// responsible for closing channels
	go func() {
		for r := range r.ready {
			ch := channels[r.target]
			ch <- r
		}
		for _, ch := range channels {
			close(ch)
		}
		wg.Wait()
		close(r.ended)
	}()

	return err
}

func startTargetGoroutine(
	ctx context.Context,
	wg *sync.WaitGroup,
	target types.Target,
	ready <-chan *readyMsg,
	ended chan<- *endedMsg,
) {
	go func() {
		for x := range ready {
			err := target.Handle(ctx, x.task)
			ended <- &endedMsg{taskID: x.taskID, err: err}
			if err != nil {
				break
			}
		}
		target.Stop()
		wg.Done()
	}()
}

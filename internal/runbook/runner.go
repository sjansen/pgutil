package runbook

import (
	"context"
	"sync"

	"github.com/sjansen/pgutil/internal/runbook/scheduler"
	"github.com/sjansen/pgutil/internal/runbook/types"
	"github.com/sjansen/pgutil/internal/sys"
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
	ctx context.Context
	log sys.Logger

	targets types.Targets
	tasks   types.Tasks

	done  chan TaskID    // tasks that ended without an error
	ended chan *endedMsg // tasks that ended, possibly with an error
	ready chan *readyMsg // tasks that can be started
}

func newRunner(log sys.Logger, targets types.Targets, tasks types.Tasks) *runner {
	capacity := 0
	for _, t := range targets {
		capacity += t.ConcurrencyLimit()
	}
	capacity *= 2

	return &runner{
		ctx:     context.TODO(),
		log:     log,
		targets: targets,
		tasks:   tasks,

		// producers are responsible for closing
		done:  make(chan TaskID, capacity),    // closed by run()
		ended: make(chan *endedMsg, capacity), // closed by startTargets()
		ready: make(chan *readyMsg, capacity), // closed by startScheduler()
	}
}

func (r *runner) closeDone() {
	if r.done != nil {
		close(r.done)
		r.done = nil
	}
}

func (r *runner) run() error {
	defer r.closeDone()

	r.startScheduler()
	err := r.startTargets()
	if err != nil {
		return err
	}

	r.log.Debug("runner: started")
	for x := range r.ended {
		r.log.Debugw("runner: task ended", "task", x.taskID, "err", x.err != nil)
		if x.err != nil {
			err = x.err
			r.closeDone()
			break
		} else {
			r.done <- x.taskID
		}
	}

	r.log.Debug("runner: stopping...")
	for x := range r.ended {
		r.log.Debugw("runner: task ended", "task", x.taskID, "err", x.err != nil)
	}

	r.log.Debug("runner: stopped")
	return err
}

func (r *runner) startScheduler() {
	go func(readyChan chan<- *readyMsg, doneChan <-chan TaskID) {
		defer close(readyChan)

		r.log.Debug("scheduler: started")
		s, ready, err := scheduler.New(r.targets, r.tasks)
		for {
			if err != nil {
				break
			}
			for targetID, taskIDs := range ready {
				for _, taskID := range taskIDs {
					r.log.Debugw("scheduler: task ready", "task", taskID, "target", targetID)
					readyChan <- &readyMsg{
						target: TargetID(targetID),
						taskID: TaskID(taskID),
						task:   r.tasks[taskID].Config,
					}
				}
			}
			r.log.Debug("scheduler: checking for finished tasks")
			if taskID, ok := <-doneChan; ok {
				ready, err = s.Next(string(taskID))
			} else {
				break
			}
		}
		r.log.Debug("scheduler: stopped")
	}(r.ready, r.done)
}

func (r *runner) startTargets() (err error) {
	wg := &sync.WaitGroup{}
	var channels = make(map[TargetID]chan<- *readyMsg)
	for targetID, target := range r.targets {
		r.log.Debugw("target: starting", "task", targetID)
		err = target.Start()
		if err != nil {
			break
		}

		wg.Add(1)
		ch := make(chan *readyMsg)
		startTargetGoroutine(r.ctx, wg, target, ch, r.ended)
		channels[TargetID(targetID)] = ch
		r.log.Debugw("target: started", "task", targetID)
	}

	// start goroutine, even if there's an error, because it's
	// responsible for closing channels
	go func() {
		r.log.Debug("router: started")
		for x := range r.ready {
			r.log.Debugw("router: task ready", "task", x.taskID, "target", x.target)
			ch := channels[x.target]
			ch <- x
		}
		r.log.Debug("router: stopping...")
		for _, ch := range channels {
			close(ch)
		}
		wg.Wait()
		close(r.ended)
		r.log.Debug("router: stopped")
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

package scheduler

import (
	"context"

	"github.com/sjansen/pgutil/internal/graphs"
	"github.com/sjansen/pgutil/internal/tasks"
)

type Scheduler struct {
	Workers int

	Deps  map[string][]string
	Tasks map[string]tasks.Task
}

func (s *Scheduler) Schedule(ctx context.Context) ([]*tasks.Status, error) {
	graph, err := graphs.NewDependencyGraph(s.Deps)
	if err != nil {
		return nil, err
	}

	start := make(chan *readyTask, s.Workers)
	status := make(chan *tasks.Status)
	startWorkers(ctx, start, status, s.Workers)

	tmp := &scheduler{
		ctx:         ctx,
		graph:       graph,
		start:       start,
		status:      status,
		tasks:       s.Tasks,
		terminating: false,
	}

	return tmp.schedule()
}

type readyTask struct {
	id   string
	task tasks.Task
}

func startWorkers(ctx context.Context, start <-chan *readyTask, status chan<- *tasks.Status, workers int) {
	for i := 0; i < workers; i++ {
		go func(start <-chan *readyTask, status chan<- *tasks.Status) {
			for {
				if x, ok := <-start; ok {
					err := x.task.Start(ctx)
					status <- &tasks.Status{
						ID:    x.id,
						Error: err,
					}
				} else {
					break
				}
			}
		}(start, status)
	}
}

type depgraph interface {
	HasPending() bool
	Next(id string) []string
}

type scheduler struct {
	ctx         context.Context
	graph       depgraph
	ready       []*readyTask
	start       chan<- *readyTask
	status      <-chan *tasks.Status
	statuses    []*tasks.Status
	tasks       map[string]tasks.Task
	terminating bool
}

func (s *scheduler) schedule() ([]*tasks.Status, error) {
	s.statuses = make([]*tasks.Status, 0, len(s.tasks))
	for _, id := range s.graph.Next("") {
		s.ready = append(s.ready, &readyTask{id: id, task: s.tasks[id]})
	}
	for {
		tasksReady := len(s.ready) > 0
		tasksPending := s.graph.HasPending()
		statusesPending := len(s.statuses) < cap(s.statuses)
		switch {
		case s.terminating:
			// TODO
		case tasksReady:
			s.maybeStartTask()
		case tasksPending:
			fallthrough
		case statusesPending:
			s.maybeReceiveStatus()
		}
		if !tasksPending && !statusesPending {
			break
		}
	}
	close(s.start)
	return s.statuses, nil
}

func (s *scheduler) maybeReceiveStatus() {
	select {
	case <-s.ctx.Done():
		s.terminating = true
	case tmp := <-s.status:
		s.statuses = append(s.statuses, tmp)
		for _, id := range s.graph.Next(tmp.ID) {
			s.ready = append(s.ready, &readyTask{id: id, task: s.tasks[id]})
		}
	}
}

func (s *scheduler) maybeStartTask() {
	select {
	case <-s.ctx.Done():
		s.terminating = true
	case s.start <- s.ready[0]:
		s.ready = s.ready[1:]
	case tmp := <-s.status:
		s.statuses = append(s.statuses, tmp)
		for _, id := range s.graph.Next(tmp.ID) {
			s.ready = append(s.ready, &readyTask{id: id, task: s.tasks[id]})
		}
	}
}

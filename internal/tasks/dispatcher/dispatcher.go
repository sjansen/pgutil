package dispatcher

import (
	"context"

	"github.com/sjansen/pgutil/internal/graphs"
	t "github.com/sjansen/pgutil/internal/tasks"
)

func Dispatch(ctx context.Context, tasks []t.Task, workers int) ([]*t.TaskStatus, error) {
	graph, taskByID, err := plan(tasks)
	if err != nil {
		return nil, err
	}

	start := make(chan t.Task, workers)
	status := make(chan *t.TaskStatus)
	startWorkers(ctx, start, status, workers)

	d := &dispatcher{
		ctx:         ctx,
		graph:       graph,
		start:       start,
		status:      status,
		taskByID:    taskByID,
		terminating: false,
	}

	return d.dispatch()
}

func plan(tasks []t.Task) (*graphs.DependencyGraph, map[string]t.Task, error) {
	taskByID := map[string]t.Task{}
	nodes := map[string][]string{}
	for _, t := range tasks {
		id := t.ID()
		taskByID[id] = t
		nodes[id] = t.Dependencies()
	}

	g, err := graphs.NewDependencyGraph(nodes)
	if err != nil {
		return nil, nil, err
	}

	return g, taskByID, nil
}

func startWorkers(ctx context.Context, start <-chan t.Task, status chan<- *t.TaskStatus, workers int) {
	for i := 0; i < workers; i++ {
		go func(start <-chan t.Task, status chan<- *t.TaskStatus) {
			for {
				if task, ok := <-start; ok {
					status <- task.Run(ctx)
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

type dispatcher struct {
	ctx         context.Context
	graph       depgraph
	ready       []t.Task
	start       chan<- t.Task
	status      <-chan *t.TaskStatus
	statuses    []*t.TaskStatus
	taskByID    map[string]t.Task
	terminating bool
}

func (d *dispatcher) dispatch() ([]*t.TaskStatus, error) {
	d.statuses = make([]*t.TaskStatus, 0, len(d.taskByID))
	for _, id := range d.graph.Next("") {
		d.ready = append(d.ready, d.taskByID[id])
	}
	for {
		tasksReady := len(d.ready) > 0
		tasksPending := d.graph.HasPending()
		statusesPending := len(d.statuses) < cap(d.statuses)
		switch {
		case d.terminating:
			// TODO
		case tasksReady:
			d.maybeStartTask()
		case tasksPending:
			fallthrough
		case statusesPending:
			d.maybeReceiveStatus()
		}
		if !tasksPending && !statusesPending {
			break
		}
	}
	close(d.start)
	return d.statuses, nil
}

func (d *dispatcher) maybeReceiveStatus() {
	select {
	case <-d.ctx.Done():
		d.terminating = true
	case tmp := <-d.status:
		d.statuses = append(d.statuses, tmp)
		for _, id := range d.graph.Next(tmp.ID) {
			d.ready = append(d.ready, d.taskByID[id])
		}
	}
}

func (d *dispatcher) maybeStartTask() {
	select {
	case <-d.ctx.Done():
		d.terminating = true
	case d.start <- d.ready[0]:
		d.ready = d.ready[1:]
	case tmp := <-d.status:
		d.statuses = append(d.statuses, tmp)
		for _, id := range d.graph.Next(tmp.ID) {
			d.ready = append(d.ready, d.taskByID[id])
		}
	}
}

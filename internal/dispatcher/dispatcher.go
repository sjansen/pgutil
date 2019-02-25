package dispatcher

import (
	"context"

	"github.com/sjansen/pgutil/internal/dtos"
	"github.com/sjansen/pgutil/internal/graphs"
)

func Dispatch(ctx context.Context, tasks []Task, workers int) ([]*dtos.TaskStatus, error) {
	graph, taskByID, err := plan(tasks)
	if err != nil {
		return nil, err
	}

	start := make(chan Task, workers)
	status := make(chan *dtos.TaskStatus)
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

func plan(tasks []Task) (*graphs.DependencyGraph, map[string]Task, error) {
	taskByID := map[string]Task{}
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

func startWorkers(ctx context.Context, start <-chan Task, status chan<- *dtos.TaskStatus, workers int) {
	for i := 0; i < workers; i++ {
		go func(start <-chan Task, status chan<- *dtos.TaskStatus) {
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
	ready       []Task
	start       chan<- Task
	status      <-chan *dtos.TaskStatus
	statuses    []*dtos.TaskStatus
	taskByID    map[string]Task
	terminating bool
}

func (d *dispatcher) dispatch() ([]*dtos.TaskStatus, error) {
	d.statuses = make([]*dtos.TaskStatus, 0, len(d.taskByID))
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

package dispatcher

import (
	"context"

	"github.com/sjansen/pgutil/internal/graphs"
	"github.com/sjansen/pgutil/internal/tasks"
)

func Dispatch(ctx context.Context, tasksByID map[string]tasks.Task, workers int) ([]*tasks.Status, error) {
	graph, err := plan(tasksByID)
	if err != nil {
		return nil, err
	}

	start := make(chan *readyTask, workers)
	status := make(chan *tasks.Status)
	startWorkers(ctx, start, status, workers)

	d := &dispatcher{
		ctx:         ctx,
		graph:       graph,
		start:       start,
		status:      status,
		tasks:       tasksByID,
		terminating: false,
	}

	return d.dispatch()
}

type readyTask struct {
	id   string
	task tasks.Task
}

func plan(tasksByID map[string]tasks.Task) (*graphs.DependencyGraph, error) {
	nodes := map[string][]string{}
	for id, t := range tasksByID {
		nodes[id] = t.Dependencies()
	}

	g, err := graphs.NewDependencyGraph(nodes)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func startWorkers(ctx context.Context, start <-chan *readyTask, status chan<- *tasks.Status, workers int) {
	for i := 0; i < workers; i++ {
		go func(start <-chan *readyTask, status chan<- *tasks.Status) {
			for {
				if x, ok := <-start; ok {
					status <- x.task.Run(ctx, x.id)
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
	ready       []*readyTask
	start       chan<- *readyTask
	status      <-chan *tasks.Status
	statuses    []*tasks.Status
	tasks       map[string]tasks.Task
	terminating bool
}

func (d *dispatcher) dispatch() ([]*tasks.Status, error) {
	d.statuses = make([]*tasks.Status, 0, len(d.tasks))
	for _, id := range d.graph.Next("") {
		d.ready = append(d.ready, &readyTask{id: id, task: d.tasks[id]})
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
			d.ready = append(d.ready, &readyTask{id: id, task: d.tasks[id]})
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
			d.ready = append(d.ready, &readyTask{id: id, task: d.tasks[id]})
		}
	}
}

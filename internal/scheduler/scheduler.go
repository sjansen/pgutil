package scheduler

import (
	"github.com/sjansen/pgutil/internal/graphs"
)

type Task interface {
	Queue() string
	Dependencies() []string
}

type Queue interface {
	Capacity() int
}

type queueID string
type taskID string
type Scheduler struct {
	deps     *graphs.DependencyGraph
	tasks    map[taskID]queueID
	pending  map[queueID][]taskID
	queueCap map[queueID]int
	queueLen map[queueID]int
}

func Start(queues map[string]Queue, tasks map[string]Task) (s *Scheduler, ready map[string][]string, err error) {
	s = &Scheduler{
		tasks:   map[taskID]queueID{},
		pending: map[queueID][]taskID{},

		queueCap: map[queueID]int{},
		queueLen: map[queueID]int{},
	}

	for k, v := range queues {
		qid := queueID(k)
		s.queueCap[qid] = v.Capacity()
		s.queueLen[qid] = 0
	}

	deps := map[string][]string{}
	for k, v := range tasks {
		tid := taskID(k)
		qid := queueID(v.Queue())
		s.tasks[tid] = qid

		deps[k] = v.Dependencies()
	}

	s.deps, err = graphs.NewDependencyGraph(deps)
	if err != nil {
		return nil, nil, err
	}

	s.fillPending(s.deps.Next(""))
	return s, s.buildReady(), nil
}

func (s *Scheduler) Finish(task string) (ready map[string][]string, err error) {
	if !s.deps.HasPending() {
		return nil, ErrNoTasks
	}

	tid := taskID(task)
	qid := s.tasks[tid]
	s.queueLen[qid]--

	s.fillPending(s.deps.Next(task))
	return s.buildReady(), nil
}

func (s *Scheduler) buildReady() map[string][]string {
	ready := map[string][]string{}

	for qid, tasks := range s.pending {
		if len(tasks) < 1 {
			continue
		}

		queueLen := s.queueLen[qid]
		n := s.queueCap[qid] - queueLen
		switch {
		case n < 1:
			continue
		case n > len(tasks):
			n = len(tasks)
		}

		tmp := tasks[:n]
		tids := make([]string, n)
		for i, x := range tmp {
			tids[i] = string(x)
		}
		ready[string(qid)] = tids

		s.pending[qid] = tasks[n:]
		s.queueLen[qid] = queueLen + n
	}

	return ready
}

func (s *Scheduler) fillPending(tasks []string) {
	for _, task := range tasks {
		tid := taskID(task)
		qid := s.tasks[tid]
		pending := s.pending[qid]
		s.pending[qid] = append(pending, tid)
	}
}

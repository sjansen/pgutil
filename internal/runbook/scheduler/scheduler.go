package scheduler

import (
	"github.com/sjansen/pgutil/internal/dag"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

type targetID string
type taskID string

// A Scheduler determines which tasks are ready for execution based on task dependencies
// and target capacity.
type Scheduler struct {
	deps      *dag.DependencyGraph
	tasks     map[taskID]targetID
	pending   map[targetID][]taskID
	targetCap map[targetID]int
	targetLen map[targetID]int
}

// New initializes a new Scheduler, and returns a set of tasks ready for immediate execution
func New(targets types.Targets, tasks types.Tasks) (s *Scheduler, ready map[string][]string, err error) {
	s = &Scheduler{
		tasks:   map[taskID]targetID{},
		pending: map[targetID][]taskID{},

		targetCap: map[targetID]int{},
		targetLen: map[targetID]int{},
	}

	for k, v := range targets {
		qid := targetID(k)
		s.targetCap[qid] = v.ConcurrencyLimit()
		s.targetLen[qid] = 0
	}

	deps := map[string][]string{}
	for k, v := range tasks {
		tid := taskID(k)
		qid := targetID(v.Target)
		s.tasks[tid] = qid

		deps[k] = v.After
	}

	s.deps, err = dag.NewDependencyGraph(deps)
	if err != nil {
		return nil, nil, err
	}

	s.fillPending(s.deps.Next(""))
	return s, s.buildReady(), nil
}

// Next returns a set of tasks ready for execution given a completed task
func (s *Scheduler) Next(completed string) (ready map[string][]string, err error) {
	if s.noPendingTasks() {
		return nil, ErrNoPendingTasks
	}

	tid := taskID(completed)
	qid := s.tasks[tid]
	s.targetLen[qid]--

	s.fillPending(s.deps.Next(completed))
	return s.buildReady(), nil
}

func (s *Scheduler) buildReady() map[string][]string {
	ready := map[string][]string{}

	for qid, tasks := range s.pending {
		if len(tasks) < 1 {
			continue
		}

		targetLen := s.targetLen[qid]
		n := s.targetCap[qid] - targetLen
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
		s.targetLen[qid] = targetLen + n
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

func (s *Scheduler) noPendingTasks() bool {
	if s.deps.HasPending() {
		return false
	}
	for qid, tasks := range s.pending {
		if len(tasks) > 0 {
			return false
		}
		delete(s.pending, qid)
	}
	return true
}

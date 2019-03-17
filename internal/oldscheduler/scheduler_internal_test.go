package oldscheduler

import (
	"context"
	"testing"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/graphs"
	"github.com/sjansen/pgutil/internal/tasks"
	"github.com/sjansen/pgutil/internal/tasks/mocks"
)

type mockGraph struct {
	next map[string][]string
	seen graphs.NodeSet
}

func (m *mockGraph) HasPending() bool {
	pending := m.seen.Size() < len(m.next)-1
	return pending
}

func (m *mockGraph) Next(id string) []string {
	if id != "" {
		m.seen.Add(id)
	}
	return m.next[id]
}

func TestScheduler(t *testing.T) {
	require := require.New(t)
	defer leaktest.Check(t)()

	const workers = 2

	ctx := context.Background()
	config := map[string][]string{
		"":  {"a", "b", "c", "d"},
		"a": nil,
		"b": nil,
		"c": {"e"},
		"d": {"f", "g", "h"},
		"e": nil,
		"f": nil,
		"g": nil,
		"h": nil,
	}
	graph := &mockGraph{
		seen: graphs.NodeSet{},
		next: config,
	}
	tasksByID := map[string]tasks.Task{}
	for id := range config {
		if id != "" {
			tasksByID[id] = &mocks.Task{}
		}
	}

	start := make(chan *readyTask, workers)
	status := make(chan *tasks.Status)
	startWorkers(ctx, start, status, workers)

	s := &scheduler{
		ctx:         ctx,
		graph:       graph,
		start:       start,
		status:      status,
		tasks:       tasksByID,
		terminating: false,
	}

	statuses, err := s.schedule()
	require.NoError(err)
	for id, task := range tasksByID {
		m := task.(*mocks.Task)
		require.Equal(1, m.RunCount)
		require.Contains(statuses, &tasks.Status{ID: id})
	}
}

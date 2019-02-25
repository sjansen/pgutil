package dispatcher

import (
	"context"
	"testing"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/dispatcher/mocks"
	"github.com/sjansen/pgutil/internal/dtos"
	"github.com/sjansen/pgutil/internal/graphs"
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

func TestDispatcher(t *testing.T) {
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
	mockTasks := []*mocks.Task{}
	taskByID := map[string]Task{}
	for id, deps := range config {
		if id != "" {
			m := &mocks.Task{Ident: id, Deps: deps}
			mockTasks = append(mockTasks, m)
			taskByID[id] = m
		}
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

	statuses, err := d.dispatch()
	require.NoError(err)
	for _, m := range mockTasks {
		require.Equal(1, m.RunCount)
		require.Contains(statuses, &dtos.TaskStatus{ID: m.ID()})
	}
}

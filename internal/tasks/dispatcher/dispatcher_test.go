package dispatcher_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/tasks"
	"github.com/sjansen/pgutil/internal/tasks/dispatcher"
	"github.com/sjansen/pgutil/internal/tasks/mocks"
)

func buildDispatcher(config map[string][]string) *dispatcher.Dispatcher {
	tasks := map[string]tasks.Task{}
	for id := range config {
		tasks[id] = &mocks.Task{}
	}
	return &dispatcher.Dispatcher{
		Workers: 2,
		Deps:    config,
		Tasks:   tasks,
	}
}

func TestDispatcher(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	config := map[string][]string{
		"foo": {"bar"},
		"bar": {"foo"},
	}
	d := buildDispatcher(config)
	statuses, err := d.Dispatch(ctx)
	require.Nil(statuses)
	require.Error(err)

	config = map[string][]string{
		"a": {},
		"b": {"c"},
		"c": {},
		"d": {"a", "b"},
		"e": {},
	}
	d = buildDispatcher(config)
	statuses, err = d.Dispatch(ctx)
	require.NoError(err)
	for id, task := range d.Tasks {
		m := task.(*mocks.Task)
		require.Equal(1, m.RunCount)
		require.Contains(
			statuses,
			&tasks.Status{ID: id},
		)
	}
}

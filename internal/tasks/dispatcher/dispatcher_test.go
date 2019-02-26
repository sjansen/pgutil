package dispatcher_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/tasks"
	"github.com/sjansen/pgutil/internal/tasks/dispatcher"
	"github.com/sjansen/pgutil/internal/tasks/mocks"
)

func TestDispatcher(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	tasksByID := map[string]tasks.Task{
		"foo": &mocks.Task{Deps: []string{"bar"}},
		"bar": &mocks.Task{Deps: []string{"foo"}},
	}
	statuses, err := dispatcher.Dispatch(ctx, tasksByID, 2)
	require.Nil(statuses)
	require.Error(err)

	config := map[string][]string{
		"a": {},
		"b": {"c"},
		"c": {},
		"d": {"a", "b"},
		"e": {},
	}
	tasksByID = map[string]tasks.Task{}
	for id, deps := range config {
		m := &mocks.Task{Deps: deps}
		tasksByID[id] = m
	}
	statuses, err = dispatcher.Dispatch(ctx, tasksByID, 2)
	require.NoError(err)
	for id, task := range tasksByID {
		m := task.(*mocks.Task)
		require.Equal(1, m.RunCount)
		require.Contains(
			statuses,
			&tasks.Status{ID: id},
		)
	}
}

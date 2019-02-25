package dispatcher_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/dispatcher"
	"github.com/sjansen/pgutil/internal/dispatcher/mocks"
	"github.com/sjansen/pgutil/internal/dtos"
)

func TestDispatcher(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	tasks := []dispatcher.Task{
		&mocks.Task{Ident: "foo", Deps: []string{"bar"}},
		&mocks.Task{Ident: "bar", Deps: []string{"foo"}},
	}
	statuses, err := dispatcher.Dispatch(ctx, tasks, 2)
	require.Nil(statuses)
	require.Error(err)

	config := map[string][]string{
		"a": {},
		"b": {"c"},
		"c": {},
		"d": {"a", "b"},
		"e": {},
	}
	tasks = []dispatcher.Task{}
	mockTasks := []*mocks.Task{}
	for id, deps := range config {
		m := &mocks.Task{Ident: id, Deps: deps}
		tasks = append(tasks, m)
		mockTasks = append(mockTasks, m)
	}
	statuses, err = dispatcher.Dispatch(ctx, tasks, 2)
	require.NoError(err)
	for _, m := range mockTasks {
		require.Equal(1, m.RunCount)
		require.Contains(
			statuses,
			&dtos.TaskStatus{ID: m.ID()},
		)
	}
}

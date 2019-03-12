package scheduler_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/scheduler"
	"github.com/sjansen/pgutil/internal/tasks"
	"github.com/sjansen/pgutil/internal/tasks/mocks"
)

func buildScheduler(config map[string][]string) *scheduler.Scheduler {
	tasks := map[string]tasks.Task{}
	for id := range config {
		tasks[id] = &mocks.Task{}
	}
	return &scheduler.Scheduler{
		Workers: 2,
		Deps:    config,
		Tasks:   tasks,
	}
}

func TestScheduler(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	config := map[string][]string{
		"foo": {"bar"},
		"bar": {"foo"},
	}
	s := buildScheduler(config)
	statuses, err := s.Schedule(ctx)
	require.Nil(statuses)
	require.Error(err)

	config = map[string][]string{
		"a": {},
		"b": {"c"},
		"c": {},
		"d": {"a", "b"},
		"e": {},
	}
	s = buildScheduler(config)
	statuses, err = s.Schedule(ctx)
	require.NoError(err)
	for id, task := range s.Tasks {
		m := task.(*mocks.Task)
		require.Equal(1, m.RunCount)
		require.Contains(
			statuses,
			&tasks.Status{ID: id},
		)
	}
}

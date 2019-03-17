package scheduler_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/scheduler"
)

type queue struct {
	capacity int
}

func (q *queue) Capacity() int {
	return q.capacity
}

type task struct {
	queue string
	deps  []string
}

func (t *task) Dependencies() []string {
	return t.deps
}

func (t *task) Queue() string {
	return t.queue
}

func TestScheduler(t *testing.T) {
	require := require.New(t)

	queues := map[string]scheduler.Queue{
		"q1": &queue{1},
		"q2": &queue{2},
		"q3": &queue{2},
	}

	tasks := map[string]scheduler.Task{
		"t1": &task{"q1", nil},
		"t2": &task{"q1", nil},
		"t3": &task{"q1", nil},
		"t4": &task{"q2", nil},
		"t5": &task{"q2", nil},
		"t6": &task{"q2", nil},
		"t7": &task{"q3", []string{"t1", "t2"}},
		"t8": &task{"q3", []string{"t3", "t4", "t7"}},
	}

	expected := map[string][]string{
		"q1": {"t1"},
		"q2": {"t4", "t5"},
	}

	s, ready, err := scheduler.Start(queues, tasks)
	require.NoError(err)
	require.NotNil(s)
	require.Equal(expected, ready)

	for _, x := range []struct {
		task     string
		expected map[string][]string
		err      error
	}{{
		"t1", map[string][]string{
			"q1": {"t2"},
		}, nil,
	}, {
		"t2", map[string][]string{
			"q1": {"t3"},
			"q3": {"t7"},
		}, nil,
	}, {
		"t3", map[string][]string{}, nil,
	}, {
		"t4", map[string][]string{
			"q2": {"t6"},
		}, nil,
	}, {
		"t5", map[string][]string{}, nil,
	}, {
		"t6", map[string][]string{}, nil,
	}, {
		"t7", map[string][]string{
			"q3": {"t8"},
		}, nil,
	}, {
		"t8", nil, scheduler.ErrNoTasks,
	}} {
		actual, err := s.Finish(x.task)
		if x.err == nil {
			require.NoError(err, x.task)
			require.Equal(x.expected, actual, x.task)
		} else {
			require.Error(err, x.task)
			require.Nil(actual, x.task)
		}
	}
}

func TestDependencyCycle(t *testing.T) {
	require := require.New(t)

	queues := map[string]scheduler.Queue{
		"q1": &queue{1},
	}

	tasks := map[string]scheduler.Task{
		"t1": &task{"q1", []string{"t1"}},
		"t2": &task{"q1", []string{"t2"}},
	}

	s, ready, err := scheduler.Start(queues, tasks)
	require.Error(err)
	require.Nil(s)
	require.Nil(ready)
}

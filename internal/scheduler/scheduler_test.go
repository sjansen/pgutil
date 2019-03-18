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
		"t01": &task{"q1", nil},
		"t02": &task{"q1", nil},
		"t03": &task{"q1", nil},
		"t04": &task{"q2", nil},
		"t05": &task{"q2", nil},
		"t06": &task{"q2", nil},
		"t07": &task{"q3", []string{"t01", "t02"}},
		"t08": &task{"q3", []string{"t03", "t04", "t07"}},
		"t09": &task{"q1", []string{"t07"}},
		"t10": &task{"q2", []string{"t07"}},
		"t11": &task{"q3", []string{"t05", "t07"}},
		"t12": &task{"q3", []string{"t06", "t07"}},
		"t13": &task{"q2", []string{"t07"}},
	}

	expected := map[string][]string{
		"q1": {"t01"},
		"q2": {"t04", "t05"},
	}

	s, ready, err := scheduler.Start(queues, tasks)
	require.NoError(err, "start")
	require.NotNil(s, "start")
	require.Equal(expected, ready, "start")

	for _, x := range []struct {
		task     string
		expected map[string][]string
		err      error
	}{{
		"t01", map[string][]string{
			"q1": {"t02"},
		}, nil,
	}, {
		"t02", map[string][]string{
			"q1": {"t03"},
			"q3": {"t07"},
		}, nil,
	}, {
		"t03", map[string][]string{}, nil,
	}, {
		"t04", map[string][]string{
			"q2": {"t06"},
		}, nil,
	}, {
		"t05", map[string][]string{}, nil,
	}, {
		"t06", map[string][]string{}, nil,
	}, {
		"t07", map[string][]string{
			"q1": {"t09"},
			"q2": {"t10", "t13"},
			"q3": {"t08", "t11"},
		}, nil,
	}, {
		"t08", map[string][]string{
			"q3": {"t12"},
		}, nil,
	}, {
		"t09", nil, scheduler.ErrNoPendingTasks,
	}, {
		"t10", nil, scheduler.ErrNoPendingTasks,
	}, {
		"t11", nil, scheduler.ErrNoPendingTasks,
	}, {
		"t12", nil, scheduler.ErrNoPendingTasks,
	}, {
		"t13", nil, scheduler.ErrNoPendingTasks,
	}} {
		actual, err := s.Next(x.task)
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
		"t1": &task{"q1", []string{"t2"}},
		"t2": &task{"q1", []string{"t1"}},
	}

	s, ready, err := scheduler.Start(queues, tasks)
	require.Error(err)
	require.Nil(s)
	require.Nil(ready)
}

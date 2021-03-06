package runbook

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

var targets = types.Targets{
	"t1": &fakeTarget{},
	"t2": &fakeTarget{},
}
var tasks = types.Tasks{
	"a": {Target: "t1"},
	"b": {Target: "t2"},
	"c": {Target: "t1", After: []string{"b"}},
	"d": {Target: "t2", After: []string{"a", "c"}},
	"e": {Target: "t1", After: []string{"d"}},
	"f": {Target: "t1", After: []string{"a", "b", "e"}},
	"g": {Target: "t1", After: []string{"b", "e"}},
	"h": {Target: "t2", After: []string{"g"}},
	"i": {Target: "t2", After: []string{"g"}},
	"j": {Target: "t2", After: []string{"g"}},
}

func TestScheduler(t *testing.T) {
	require := require.New(t)

	r := newRunner(logger.Discard(), targets, tasks)
	r.startScheduler()

	seen := map[TaskID]struct{}{}
	for x := range r.ready {
		seen[x.taskID] = struct{}{}
		r.done <- x.taskID
	}

	expected := map[TaskID]struct{}{
		"a": {}, "b": {}, "c": {}, "d": {}, "e": {},
		"f": {}, "g": {}, "h": {}, "i": {}, "j": {},
	}
	require.Equal(expected, seen)
}

func TestSchedulerEarlyTermination(t *testing.T) {
	require := require.New(t)

	r := newRunner(logger.Discard(), targets, tasks)
	r.startScheduler()

	seen := map[TaskID]struct{}{}
	for x := range r.ready {
		seen[x.taskID] = struct{}{}
		if x.taskID == "e" {
			close(r.done)
			break
		} else {
			r.done <- x.taskID
		}
	}

	expected := map[TaskID]struct{}{
		"a": {}, "b": {}, "c": {}, "d": {}, "e": {},
	}
	require.Equal(expected, seen)
}

type fakeTarget struct{}

func (t *fakeTarget) Analyze() error {
	return nil
}
func (t *fakeTarget) ConcurrencyLimit() int {
	return 1
}
func (t *fakeTarget) Handle(ctx context.Context, task types.TaskConfig) error {
	return nil
}
func (t *fakeTarget) NewTaskConfig(class string) (types.TaskConfig, error) {
	return nil, nil
}
func (t *fakeTarget) Start() error {
	return nil
}
func (t *fakeTarget) Stop() error {
	return nil
}

package demo

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/sjansen/pgutil/internal/taskset/base"
	"github.com/sjansen/pgutil/internal/taskset/types"
	"go.uber.org/zap"
)

// TargetFactory instantiates new targets
type TargetFactory struct {
	Log    *zap.SugaredLogger
	Stdout io.Writer
}

// NewTarget create a new target with default settings
func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log:    f.Log,
		stdout: f.Stdout,

		MaxConcurrency: 1,
	}
}

type execer interface {
	exec(*Target) error
}

// Target executes tasks
type Target struct {
	sync.Mutex
	log    *zap.SugaredLogger
	stdout io.Writer

	MaxConcurrency int    `hcl:"max_concurrency,optional"`
	String         string `hcl:"string,attr"`
}

// NewTask creates a new Task of type typ with default settings
func (t *Target) NewTask(typ string) (types.Task, error) {
	switch typ {
	case "echo":
		return &Echo{}, nil
	case "fail":
		return &Fail{}, nil
	case "rev":
		return &Rev{}, nil
	case "rot13":
		return &Rot13{}, nil
	case "sleep":
		return &Sleep{}, nil
	}
	return nil, errors.New("invalid task type") // TODO
}

// Ready verifies the target's settings are valid
func (t *Target) Ready() error {
	return nil
}

// Run executes a single task
func (t *Target) Run(ctx context.Context, task types.Task) error {
	if x, ok := task.(execer); ok {
		return x.exec(t)
	}
	return errors.New("invalid task")
}

// Start should be called before the target starts handling tasks
func (t *Target) Start() (chan<- map[string]types.Task, <-chan map[string]error) {
	fn := func(id string, task types.Task, results chan<- map[string]error) {
		go func() {
			if x, ok := task.(execer); ok {
				results <- map[string]error{
					id: x.exec(t),
				}
			} else {
				results <- map[string]error{
					id: errors.New("invalid task type"),
				}
			}
		}()
	}
	return base.RunTasks(fn, t.MaxConcurrency)
}

// Stop should be called when there are no tasks left for the target to handle
func (t *Target) Stop() error {
	return nil
}

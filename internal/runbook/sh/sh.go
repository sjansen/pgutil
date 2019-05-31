package sh

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/sjansen/pgutil/internal/runbook/types"
	"go.uber.org/zap"
)

// TargetFactory instantiates new targets
type TargetFactory struct {
	Log     *zap.SugaredLogger
	Basedir string
	Stdout  io.Writer
	Stderr  io.Writer
}

// Target executes tasks
type Target struct {
	log     *zap.SugaredLogger
	basedir string
	environ []string
	stdout  io.Writer
	stderr  io.Writer

	Concurrency int
	Env         Env
}

type execer interface {
	exec(context.Context, *Target) error
}

// NewTarget create a new target with default settings
func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log:     f.Log,
		basedir: f.Basedir,
		stdout:  f.Stdout,
		stderr:  f.Stderr,

		Concurrency: 1,
	}
}

// Analyze checks if the target's settings are valid
func (t *Target) Analyze() error {
	if t.Concurrency < 1 {
		return errors.New("invalid concurrency")
	}
	return nil
}

// ConcurrencyLimit reports how many tasks the target can execute concurrently
func (t *Target) ConcurrencyLimit() int {
	return t.Concurrency
}

// Handle executes a task
func (t *Target) Handle(ctx context.Context, task types.TaskConfig) error {
	if x, ok := task.(execer); ok {
		return x.exec(ctx, t)
	}
	return errors.New("invalid task")
}

// NewTaskConfig creates a specific class of TaskConfig with default settings
func (t *Target) NewTaskConfig(class string) (types.TaskConfig, error) {
	switch class {
	case "", "exec":
		return &Exec{}, nil
	}
	return nil, errors.New("invalid task class")
}

// Start should be called before the target starts handling tasks
func (t *Target) Start() error {
	t.environ = t.Env.Apply(os.Environ())
	return nil
}

// Stop should be called when there are no tasks left for the target to handle
func (t *Target) Stop() error {
	return nil
}

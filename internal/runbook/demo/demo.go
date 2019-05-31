package demo

import (
	"context"
	"errors"
	"io"

	"github.com/sjansen/pgutil/internal/runbook/types"
	"go.uber.org/zap"
)

// TargetFactory instantiates new targets
type TargetFactory struct {
	Log    *zap.SugaredLogger
	Stdout io.Writer
}

// Target executes tasks
type Target struct {
	log    *zap.SugaredLogger
	stdout io.Writer
	String string
}

type munger interface {
	munge(*Target) error
}

// NewTarget create a new target with default settings
func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log:    f.Log,
		stdout: f.Stdout,
	}
}

// Analyze verifies the target's settings are valid
func (t *Target) Analyze() error {
	return nil
}

// ConcurrencyLimit reports how many tasks the target can execute concurrently
func (t *Target) ConcurrencyLimit() int {
	return 1
}

// Handle executes a task
func (t *Target) Handle(ctx context.Context, task types.TaskConfig) error {
	if x, ok := task.(munger); ok {
		return x.munge(t)
	}
	return errors.New("invalid task")
}

// NewTaskConfig creates a specific class of TaskConfig with default settings
func (t *Target) NewTaskConfig(class string) (types.TaskConfig, error) {
	switch class {
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
	return nil, errors.New("invalid task class") // TODO
}

// Start should be called before the target starts handling tasks
func (t *Target) Start() error {
	return nil
}

// Stop should be called when there are no tasks left for the target to handle
func (t *Target) Stop() error {
	return nil
}

package sh

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/sjansen/pgutil/internal/runbook/types"
	"go.uber.org/zap"
)

type TargetFactory struct {
	Log    *zap.SugaredLogger
	Stdout io.Writer
	Stderr io.Writer
}

type Target struct {
	log     *zap.SugaredLogger
	stdout  io.Writer
	stderr  io.Writer
	environ []string

	Concurrency int
	Env         Env
}

type execer interface {
	exec(context.Context, *Target) error
}

func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log:    f.Log,
		stdout: f.Stdout,
		stderr: f.Stderr,

		Concurrency: 1,
	}
}

func (t *Target) Analyze() error {
	if t.Concurrency < 1 {
		return errors.New("invalid concurrency")
	}
	return nil
}

func (t *Target) ConcurrencyLimit() int {
	return t.Concurrency
}

func (t *Target) Handle(ctx context.Context, task types.TaskConfig) error {
	if x, ok := task.(execer); ok {
		return x.exec(ctx, t)
	}
	return errors.New("invalid task")
}

func (t *Target) NewTaskConfig(class string) (types.TaskConfig, error) {
	switch class {
	case "", "exec":
		return &Exec{}, nil
	}
	return nil, errors.New("invalid task class")
}

func (t *Target) Start() error {
	t.environ = t.Env.Apply(os.Environ())
	return nil
}

func (t *Target) Stop() error {
	return nil
}

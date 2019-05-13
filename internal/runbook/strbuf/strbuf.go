package strbuf

import (
	"context"
	"errors"
	"io"

	"github.com/sjansen/pgutil/internal/runbook/types"
	"go.uber.org/zap"
)

type TargetFactory struct {
	Log    *zap.SugaredLogger
	Stdout io.Writer
}

type Target struct {
	log    *zap.SugaredLogger
	stdout io.Writer
	Data   string
}

type munger interface {
	munge(*Target) error
}

func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log:    f.Log,
		stdout: f.Stdout,
	}
}

func (t *Target) Analyze() error {
	return nil
}

func (t *Target) ConcurrencyLimit() int {
	return 1
}

func (t *Target) Handle(ctx context.Context, task types.TaskConfig) error {
	if x, ok := task.(munger); ok {
		return x.munge(t)
	}
	return errors.New("invalid task")
}

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

func (t *Target) Start() error {
	return nil
}

func (t *Target) Stop() error {
	return nil
}

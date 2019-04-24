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
	StdOut io.Writer
}

type Target struct {
	log    *zap.SugaredLogger
	stdout io.Writer
	Data   string
}

type Task interface {
	Munge(string) string
}

func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log:    f.Log,
		stdout: f.StdOut,
	}
}

func (t *Target) Analyze() error {
	return nil
}

func (t *Target) ConcurrencyLimit() int {
	return 1
}

func (t *Target) Handle(ctx context.Context, task types.TaskConfig) error {
	if x, ok := task.(Task); ok {
		t.Data = x.Munge(t.Data)
	}
	return errors.New("invalid task")
}

func (t *Target) NewTaskConfig(class string) (types.TaskConfig, error) {
	switch class {
	case "echo":
		return &EchoTask{}, nil
	case "rev":
		return &RevTask{}, nil
	case "rot13":
		return &Rot13Task{}, nil
	}
	return nil, errors.New("invalid task class") // TODO
}

func (t *Target) Start() error {
	return nil
}

func (t *Target) Stop() error {
	return nil
}

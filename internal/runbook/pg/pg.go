package pg

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/sjansen/pgutil/internal/pg"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

type TargetFactory struct {
	Log *zap.SugaredLogger
}

type Target struct {
	log  *zap.SugaredLogger
	conn *pg.Conn

	Concurrency int

	Host     string
	Port     uint16
	Username string
	Password string
	Database string
}

type execer interface {
	exec(context.Context, *Target) error
}

func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log: f.Log,

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
	conn, err := pg.New(&pg.Options{
		Log: t.log,

		Host:     t.Host,
		Port:     t.Port,
		Username: t.Username,
		Password: t.Password,
		Database: t.Database,
	})
	if err == nil {
		t.conn = conn
	}
	return err
}

func (t *Target) Stop() error {
	return t.conn.Close()
}

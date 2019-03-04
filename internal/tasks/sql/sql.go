package sql

import (
	"context"
)

type Task struct {
	C   ConnPool
	SQL string
}

type ConnPool interface {
	Exec(sql string) error
}

func (t *Task) Start(ctx context.Context) error {
	return t.C.Exec(t.SQL)
}

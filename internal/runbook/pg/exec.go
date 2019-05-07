package pg

import (
	"context"
	"errors"
)

var _ execer = &Exec{}

type Exec struct {
	SQL string `json:"sql"`
}

func (x *Exec) exec(ctx context.Context, t *Target) error {
	return t.conn.Exec(x.SQL)
}

func (x *Exec) Check() error {
	if x.SQL == "" {
		return errors.New("missing SQL")
	}
	return nil
}

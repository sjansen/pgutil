package pg

import (
	"context"
	"errors"
)

var _ execer = &Exec{}

// Exec executes SQL
type Exec struct {
	// SQL specifies one or more SQL statments
	SQL string `json:"sql"`
}

func (x *Exec) exec(ctx context.Context, t *Target) error {
	return t.conn.Exec(x.SQL)
}

// Check validates the task's settings
func (x *Exec) Check() error {
	if x.SQL == "" {
		return errors.New("missing SQL")
	}
	return nil
}

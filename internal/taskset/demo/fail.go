package demo

import (
	"errors"

	"github.com/sjansen/pgutil/internal/taskset/base"
)

var _ execer = &Fail{}

// Fail always returns an error when executed
type Fail struct {
	base.Task
}

// Ready validates the task's settings
func (x *Fail) Ready() error {
	return nil
}

var ErrFail = errors.New("fail task executed")

func (x *Fail) exec(t *Target) error {
	return ErrFail
}

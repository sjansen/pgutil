package demo

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/taskset/base"
)

var _ execer = &Echo{}

// Echo prints the target's string when executed
type Echo struct {
	base.Task
}

// Ready validates the task's settings
func (x *Echo) Ready() error {
	return nil
}

func (x *Echo) exec(t *Target) error {
	t.Lock()
	_, err := fmt.Fprintln(t.stdout, t.String)
	t.Unlock()
	return err
}

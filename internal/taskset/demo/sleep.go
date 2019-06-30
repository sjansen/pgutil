package demo

import (
	"time"

	"github.com/sjansen/pgutil/internal/taskset/base"
)

var _ execer = &Sleep{}

// Sleep pauses the current target for Seconds seconds when executed
type Sleep struct {
	base.Task
	Seconds int
}

// Ready validates the task's settings
func (x *Sleep) Ready() error {
	if x.Seconds < 1 {
		x.Seconds = 1
	}
	return nil
}

func (x *Sleep) exec(t *Target) error {
	time.Sleep(
		time.Duration(x.Seconds) * time.Second,
	)
	return nil
}

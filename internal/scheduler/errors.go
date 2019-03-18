package scheduler

import "errors"

var ErrNoPendingTasks = errors.New("all tasks have been scheduled")

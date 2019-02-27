package mocks

import (
	"context"
)

type Task struct {
	RunCount int
}

func (t *Task) Start(ctx context.Context) error {
	t.RunCount++
	return nil
}

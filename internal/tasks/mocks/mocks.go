package mocks

import (
	"context"

	"github.com/sjansen/pgutil/internal/tasks"
)

type Task struct {
	RunCount int
}

func (t *Task) Run(ctx context.Context, id string) *tasks.Status {
	t.RunCount++
	return &tasks.Status{
		ID: id,
	}
}

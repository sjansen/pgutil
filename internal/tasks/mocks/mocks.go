package mocks

import (
	"context"

	"github.com/sjansen/pgutil/internal/tasks"
)

type Task struct {
	Deps     []string
	RunCount int
}

func (t *Task) Dependencies() []string {
	return t.Deps
}

func (t *Task) Run(ctx context.Context, id string) *tasks.Status {
	t.RunCount++
	return &tasks.Status{
		ID: id,
	}
}

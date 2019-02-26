package mocks

import (
	"context"

	"github.com/sjansen/pgutil/internal/tasks"
)

type Task struct {
	Ident    string
	Deps     []string
	RunCount int
}

func (t *Task) ID() string {
	return t.Ident
}

func (t *Task) Dependencies() []string {
	return t.Deps
}

func (t *Task) Run(ctx context.Context) *tasks.TaskStatus {
	t.RunCount++
	return &tasks.TaskStatus{
		ID: t.Ident,
	}
}

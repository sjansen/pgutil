package mocks

import (
	"context"

	"github.com/sjansen/pgutil/internal/dtos"
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

func (t *Task) Run(ctx context.Context) *dtos.TaskStatus {
	t.RunCount++
	return &dtos.TaskStatus{
		ID: t.Ident,
	}
}

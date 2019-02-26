package tasks

import "context"

type Task interface {
	ID() string
	Dependencies() []string
	Run(ctx context.Context) *TaskStatus
}

type TaskStatus struct {
	ID     string
	Status string
	Error  error
}

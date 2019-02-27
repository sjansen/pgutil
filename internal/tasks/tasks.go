package tasks

import "context"

type Status struct {
	ID    string
	Desc  string
	Error error
}

type Task interface {
	Start(ctx context.Context) error
}

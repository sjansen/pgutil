package tasks

import "context"

type Task interface {
	Dependencies() []string
	Run(ctx context.Context, id string) *Status
}

type Status struct {
	ID    string
	Desc  string
	Error error
}

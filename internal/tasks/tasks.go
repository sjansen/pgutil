package tasks

import "context"

type Task interface {
	Run(ctx context.Context, id string) *Status
}

type Status struct {
	ID    string
	Desc  string
	Error error
}

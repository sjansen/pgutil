package types

import (
	"context"
)

// A Runbook describes how to perform a series of tasks
type Runbook struct {
	Targets map[string]Target
	Tasks   map[string]*Task
}

// A TargetFactory instantiates new targets
type TargetFactory interface {
	NewTarget() Target
}

// Targets maps target IDs to targets
type Targets map[string]Target

// A Target executes tasks
type Target interface {
	Analyze() error
	ConcurrencyLimit() int
	Handle(context.Context, TaskConfig) error
	NewTaskConfig(class string) (TaskConfig, error)
	Start() error
	Stop() error
}

// Tasks maps task IDs to tasks
type Tasks map[string]*Task

// Task contains data common to all tasks
type Task struct {
	Target  string
	After   []string
	Message string
	Config  TaskConfig
}

// A TaskConfig contains task-specific data
type TaskConfig interface {
	Check() error
}

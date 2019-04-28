package types

import (
	"context"
)

// Runbook describes how to perform a series of tasks
type Runbook struct {
	Targets map[string]Target
	Tasks   map[string]*Task
}

// TargetFactory instantiates new targets
type TargetFactory interface {
	NewTarget() Target
}

// Targets maps target IDs to targets
type Targets map[string]Target

// Target is concerned with specific classes of tasks
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

// Task contains generic task data
type Task struct {
	After  []string
	Target string
	Config TaskConfig
}

// TaskConfig contains task-specific data
type TaskConfig interface {
	Check() error
}

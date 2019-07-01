package types

// TaskSet describes how to perform a series of tasks
type TaskSet struct {
	Targets map[string]map[TargetID]Target
	Tasks   map[TaskID]Task
}

// TargetID uniquely identifies a target
type TargetID string

// TaskID uniquely identifies a task
type TaskID string

// Target creates and executes tasks
type Target interface {
	NewTask(string) (Task, error)
	Ready() error
	Start() (chan<- TaskBatch, <-chan TaskResults)
}

// Task describes when a task can be executed
type Task interface {
	Dependencies() []string
	Provides() []string
	Ready() error
}

// TargetFactory instantiates new targets
type TargetFactory interface {
	NewTarget() Target
}

// TaskBatch is a set of tasks that should be executed together
type TaskBatch map[TaskID]Task

// TaskResults records how tasks ended
type TaskResults map[TaskID]error

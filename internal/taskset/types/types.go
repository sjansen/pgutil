package types

// TaskSet describes how to perform a series of tasks
type TaskSet struct {
	Targets map[string]map[string]Target
	Tasks   map[string]Task
}

// Target executes tasks
type Target interface {
	NewTask(string) (Task, error)
	Start() (chan<- map[string]Task, <-chan map[string]error)
}

// A TargetFactory instantiates new targets
type TargetFactory interface {
	NewTarget() Target
}

// Task contains data common to all tasks
type Task interface {
	Dependencies() []string
}

package types

type Target interface {
	NewTask(string) (Task, error)
	Start() (chan<- map[string]Task, <-chan map[string]error)
}

type Task interface {
	Dependencies() []string
}

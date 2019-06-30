package base

type Task struct {
	Target  string
	After   []string
	Provide []string
	Require []string
}

func (t *Task) Dependencies() []string {
	deps := make([]string, 0, len(t.After))
	for _, id := range t.After {
		deps = append(deps, "task:"+id)
	}
	return deps
}

package pg

import (
	"errors"
	"fmt"

	"github.com/sjansen/pgutil/internal/taskset/base"
	"github.com/sjansen/pgutil/internal/taskset/types"
)

var _ types.Target = &Target{}

type Target struct{}

func (t *Target) NewTask(class string) (types.Task, error) {
	return &Task{}, nil
}

func (t *Target) Start() (chan<- map[string]types.Task, <-chan map[string]error) {
	fn := func(id string, task types.Task, results chan<- map[string]error) {
		go func() {
			if x, ok := task.(runner); ok {
				results <- map[string]error{
					id: x.Run(),
				}
			} else {
				results <- map[string]error{
					id: errors.New("invalid task type"),
				}
			}
		}()
	}
	return base.RunTasks(fn)
}

type runner interface {
	Run() error
}

type Task struct {
	base.Task
	SQL string `hcl:"sql,attr"`
}

func (t *Task) Run() error {
	fmt.Println("  sql:", t.SQL)
	return nil
}

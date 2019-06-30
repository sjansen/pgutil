package pg

import (
	"errors"
	"fmt"

	"github.com/sjansen/pgutil/internal/taskset/base"
	"github.com/sjansen/pgutil/internal/taskset/types"
	"go.uber.org/zap"
)

var _ types.Target = &Target{}

// TargetFactory instantiates new targets
type TargetFactory struct {
	Log *zap.SugaredLogger
}

// NewTarget create a new target with default settings
func (f *TargetFactory) NewTarget() types.Target {
	return &Target{
		log: f.Log,
	}
}

// Target executes tasks
type Target struct {
	log *zap.SugaredLogger
}

// NewTask creates a new Task of type typ with default settings
func (t *Target) NewTask(typ string) (types.Task, error) {
	return &Task{}, nil
}

// Start should be called before the target starts handling tasks
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

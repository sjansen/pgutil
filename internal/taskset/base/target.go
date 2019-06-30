package base

import (
	"fmt"

	"github.com/sjansen/pgutil/internal/taskset/types"
)

type RunTask func(id string, t types.Task, ch chan<- map[string]error)

func RunTasks(fn RunTask, maxConcurrency int) (chan<- map[string]types.Task, <-chan map[string]error) {
	queue := make(chan map[string]types.Task, maxConcurrency)
	results := make(chan map[string]error, maxConcurrency)

	go func(queue <-chan map[string]types.Task, results chan<- map[string]error) {
		relay := make(chan map[string]error)
		running := 0

		for {
			select {
			case tasks, ok := <-queue:
				if ok {
					for id, task := range tasks {
						fmt.Println(id)
						fn(id, task, relay)
						running++
					}
				} else {
					goto shutdown
				}
			case result := <-relay:
				results <- result
				running--
			}
		}
	shutdown:
		for result := range relay {
			results <- result
			running--
			if running < 1 {
				close(relay)
				close(results)
			}
		}
	}(queue, results)

	return queue, results
}

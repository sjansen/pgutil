package loader

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	jsonnet "github.com/google/go-jsonnet"
)

type Loader struct {
	Queues map[string]func() Queue
	Tasks  map[string]func() TaskConfig
}

type Runbook struct {
	Queues map[string]Queue
	Tasks  map[string]*Task
}

type Task struct {
	Queue  string
	After  []string
	Config TaskConfig
}

type Queue interface {
	ConcurrencyLimit() int
	VerifyConfig() error
	VerifyTask(config interface{}) error
}

type TaskConfig interface {
	VerifyConfig() error
}

type runbook struct {
	Queues map[string]map[string]json.RawMessage
	Tasks  map[string]*task
}

type task struct {
	Queue  string
	After  []string
	Config json.RawMessage
}

func (l *Loader) Load(filename string) (*Runbook, error) {
	directory := filepath.Dir(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vm := jsonnet.MakeVM()
	importer := &jsonnet.FileImporter{
		JPaths: []string{directory},
	}
	vm.Importer(importer)

	evaluated, err := vm.EvaluateSnippet(filename, string(data))
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(
		strings.NewReader(evaluated),
	)
	dec.DisallowUnknownFields()

	tmp := &runbook{}
	err = dec.Decode(tmp)
	if err != nil {
		return nil, err
	}

	book := &Runbook{
		Queues: map[string]Queue{},
		Tasks:  map[string]*Task{},
	}
	for queueType, queues := range tmp.Queues {
		for name, config := range queues {
			dec := json.NewDecoder(
				bytes.NewReader(config),
			)
			dec.DisallowUnknownFields()

			queue := l.Queues[queueType]()
			if err := dec.Decode(queue); err != nil {
				return nil, err
			}
			if err := queue.VerifyConfig(); err != nil {
				return nil, err
			}

			book.Queues[queueType+"/"+name] = queue
		}
	}
	for id, raw := range tmp.Tasks {
		taskType := raw.Queue
		if idx := strings.Index(taskType, "/"); idx >= 0 {
			taskType = taskType[:idx]
		}
		config := l.Tasks[taskType]()

		dec := json.NewDecoder(
			bytes.NewReader(raw.Config),
		)
		dec.DisallowUnknownFields()
		if err := dec.Decode(config); err != nil {
			return nil, err
		}
		if err := config.VerifyConfig(); err != nil {
			return nil, err
		}

		book.Tasks[id] = &Task{
			Queue:  raw.Queue,
			After:  raw.After,
			Config: config,
		}
	}
	return book, nil
}

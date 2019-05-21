package parser

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	jsonnet "github.com/google/go-jsonnet"

	"github.com/sjansen/pgutil/internal/runbook/types"
)

// Parser registers available targets and tasks
type Parser struct {
	Targets map[string]types.TargetFactory
}

type runbook struct {
	Targets map[string]*struct {
		Class  string
		Config json.RawMessage
	}
	Tasks map[string]*struct {
		After  []string
		Target string
		Config json.RawMessage
	}
}

// Load evaluates a runbook file to convert it to JSON
func (p *Parser) Load(filename string) (string, error) {
	directory := filepath.Dir(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	vm := jsonnet.MakeVM()
	vm.Importer(&jsonnet.FileImporter{
		JPaths: []string{directory},
	})

	return vm.EvaluateSnippet(filename, string(data))
}

// Parse loads targets and task from a runbook file
func (p *Parser) Parse(filename string) (*types.Runbook, error) {
	evaluated, err := p.Load(filename)
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

	book := &types.Runbook{
		Targets: make(types.Targets),
		Tasks:   make(types.Tasks),
	}

	if err = p.loadTargets(tmp, book); err != nil {
		return nil, err
	}
	if err = p.loadTasks(tmp, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (p *Parser) loadTargets(tmp *runbook, book *types.Runbook) error {
	for targetID, wrapper := range tmp.Targets {
		factory, ok := p.Targets[wrapper.Class]
		if !ok {
			return errors.New("invalid target class")
		}
		target := factory.NewTarget()

		if len(wrapper.Config) > 0 {
			dec := json.NewDecoder(
				bytes.NewReader(wrapper.Config),
			)
			dec.DisallowUnknownFields()
			if err := dec.Decode(target); err != nil {
				return err
			}
		}
		if err := target.Analyze(); err != nil {
			return err
		}
		book.Targets[targetID] = target
	}
	return nil
}

func (p *Parser) loadTasks(tmp *runbook, book *types.Runbook) error {
	for taskID, task := range tmp.Tasks {
		targetID := task.Target
		taskClass := ""
		if idx := strings.Index(task.Target, "/"); idx != -1 {
			targetID = task.Target[:idx]
			taskClass = task.Target[idx+1:]
		}

		factory, ok := book.Targets[targetID]
		if !ok {
			return errors.New("invalid target")
		}

		taskConfig, err := factory.NewTaskConfig(taskClass)
		if err != nil {
			return err
		}

		if len(task.Config) > 0 {
			dec := json.NewDecoder(
				bytes.NewReader(task.Config),
			)
			dec.DisallowUnknownFields()
			if err := dec.Decode(taskConfig); err != nil {
				return err
			}
		}
		if err := taskConfig.Check(); err != nil {
			return err
		}

		book.Tasks[taskID] = &types.Task{
			After:  task.After,
			Target: targetID,
			Config: taskConfig,
		}
	}
	return nil
}

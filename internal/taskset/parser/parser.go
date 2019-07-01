package parser

import (
	"errors"
	"strings"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"

	"github.com/sjansen/pgutil/internal/taskset/types"
)

type taskSet struct {
	Targets []*target `hcl:"target,block"`
	Tasks   []*task   `hcl:"task,block"`
}

type target struct {
	Type string `hcl:"type,label"`
	ID   string `hcl:"name,label"`

	Config hcl.Body `hcl:",remain"`
}

type task struct {
	Type   string `hcl:"type,label"`
	ID     string `hcl:"name,label"`
	Target string `hcl:"target,optional"`

	After   []string `hcl:"after,optional"`
	Provide []string `hcl:"provides,optional"`
	Require []string `hcl:"requires,optional"`

	Config hcl.Body `hcl:",remain"`
}

// Parser registers available targets and tasks
type Parser struct {
	Targets map[string]types.TargetFactory
}

// Parse loads targets and task from a file
func (p *Parser) Parse(filename string) (*types.TaskSet, error) {
	hp := hclparse.NewParser()
	f, diag := hp.ParseHCLFile(filename)
	if diag != nil && diag.HasErrors() {
		return nil, diag
	}

	raw := &taskSet{}
	diag = gohcl.DecodeBody(f.Body, nil, raw)
	if diag != nil && diag.HasErrors() {
		return nil, diag
	}

	ts := &types.TaskSet{
		Targets: map[string]map[types.TargetID]types.Target{},
		Tasks:   map[types.TaskID]types.Task{},
	}
	err := p.loadDefaultTargets(ts)
	if err != nil {
		return nil, err
	}

	err = p.loadExplicitTargets(raw, ts)
	if err != nil {
		return nil, err
	}

	err = p.loadTasks(raw, ts)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (p *Parser) loadDefaultTargets(ts *types.TaskSet) error {
	for typ, factory := range p.Targets {
		ts.Targets[typ] = map[types.TargetID]types.Target{
			"": factory.NewTarget(),
		}
	}
	return nil
}

func (p *Parser) loadExplicitTargets(raw *taskSet, ts *types.TaskSet) error {
	for _, t := range raw.Targets {
		factory, ok := p.Targets[t.Type]
		if !ok {
			return errors.New("invalid target type")
		}
		target := factory.NewTarget()
		err := gohcl.DecodeBody(t.Config, nil, target)
		if err != nil {
			return err
		}
		ts.Targets[t.Type][types.TargetID(t.ID)] = target
	}
	return nil
}

func (p *Parser) loadTasks(raw *taskSet, ts *types.TaskSet) error {
	for _, t := range raw.Tasks {
		targetType, taskType := t.Type, ""
		if x := strings.SplitN(t.Type, "/", 2); len(x) > 1 {
			targetType = x[0]
			taskType = x[1]
		}

		targetGroup, ok := ts.Targets[targetType]
		if !ok {
			return errors.New("invalid target type")
		}

		target, ok := targetGroup[types.TargetID(t.Target)]
		if !ok {
			return errors.New("invalid target")
		}

		task, err := target.NewTask(taskType)
		if err != nil {
			return err
		}

		diag := gohcl.DecodeBody(t.Config, nil, task)
		if diag != nil && diag.HasErrors() {
			return diag
		}

		ts.Tasks[types.TaskID(t.ID)] = task
	}
	return nil
}

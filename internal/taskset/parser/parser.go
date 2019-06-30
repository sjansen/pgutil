package parser

import (
	"errors"

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
	Name string `hcl:"name,label"`

	Config hcl.Body `hcl:",remain"`
}

type task struct {
	Type    string   `hcl:"type,label"`
	Name    string   `hcl:"name,label"`
	After   []string `hcl:"after,optional"`
	Provide []string `hcl:"after,optional"`
	Require []string `hcl:"after,optional"`
	Target  string   `hcl:"target,optional"`

	Config hcl.Body `hcl:",remain"`
}

// Parser registers available targets and tasks
type Parser struct {
	Targets map[string]types.TargetFactory
}

// ParseFile loads targets and task from a runbook file
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
		Targets: map[string]map[string]types.Target{},
		Tasks:   map[string]types.Task{},
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
		ts.Targets[typ] = map[string]types.Target{
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
		ts.Targets[t.Type][t.Name] = target
	}
	return nil
}

func (p *Parser) loadTasks(raw *taskSet, ts *types.TaskSet) error {
	for _, t := range raw.Tasks {
		task, err := ts.Targets[t.Type][t.Target].NewTask(t.Type)
		if err != nil {
			return err
		}
		diag := gohcl.DecodeBody(t.Config, nil, task)
		if diag != nil && diag.HasErrors() {
			return diag
		}
		ts.Tasks[t.Name] = task
	}
	return nil
}

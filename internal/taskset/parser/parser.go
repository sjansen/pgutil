package parser

import (
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hclparse"

	"github.com/sjansen/pgutil/internal/taskset/pg"
	"github.com/sjansen/pgutil/internal/taskset/sh"
	"github.com/sjansen/pgutil/internal/taskset/types"
)

type TaskSet struct {
	Targets map[string]map[string]types.Target
	Tasks   map[string]types.Task
}

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

func ParseFile(filename string) (*TaskSet, error) {
	p := hclparse.NewParser()
	f, err := p.ParseHCLFile(filename)
	if err != nil {
		return nil, err
	}

	ts := &taskSet{}
	err = gohcl.DecodeBody(f.Body, nil, ts)
	if err != nil {
		return nil, err
	}

	result := &TaskSet{
		Targets: map[string]map[string]types.Target{
			"pg": {"": &pg.Target{}},
			"sh": {"": &sh.Target{}},
		},
		Tasks: map[string]types.Task{},
	}
	for _, t := range ts.Targets {
		var target types.Target
		switch t.Type {
		case "pg":
			target = &pg.Target{}
		case "sh":
			target = &sh.Target{}
		}
		err = gohcl.DecodeBody(t.Config, nil, target)
		if err != nil {
			return nil, err
		}
		result.Targets[t.Type][t.Name] = target
	}

	for _, t := range ts.Tasks {
		task, err := result.Targets[t.Type][t.Target].NewTask(t.Type)
		if err != nil {
			return nil, err
		}
		diag := gohcl.DecodeBody(t.Config, nil, task)
		if diag != nil && diag.HasErrors() {
			return nil, diag
		}
		result.Tasks[t.Name] = task
	}

	return result, nil
}

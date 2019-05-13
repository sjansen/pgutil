package runbook

import (
	"io"

	dot "github.com/awalterschulze/gographviz"

	"github.com/sjansen/pgutil/internal/runbook/demo"
	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/pg"
	"github.com/sjansen/pgutil/internal/runbook/sh"
	"github.com/sjansen/pgutil/internal/runbook/types"
	"github.com/sjansen/pgutil/internal/sys"
)

// TargetID uniquely identifies a target
type TargetID string

// TaskID uniquely identifies a task
type TaskID string

func newParser(sys *sys.IO) *parser.Parser {
	return &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"pg": &pg.TargetFactory{
				Log: sys.Log,
			},
			"sh": &sh.TargetFactory{
				Stdout: sys.Stdout,
				Stderr: sys.Stderr,
			},
			"demo": &demo.TargetFactory{
				Stdout: sys.Stdout,
			},
		},
	}
}

// Dot generates a GraphViz compatible description of a runbook's tasks
func Dot(sys *sys.IO, filename string, w io.Writer, splines string) error {
	p := newParser(sys)
	runbook, err := p.Parse(filename)
	if err != nil {
		return err
	}

	g := dot.NewEscape()
	g.SetDir(true)
	g.SetName("runbook")
	g.AddAttr("runbook", "newrank", "true")
	if splines != "" {
		g.AddAttr("runbook", "splines", splines)
	}

	for targetID := range runbook.Targets {
		graphID := "cluster_" + targetID
		g.AddSubGraph("runbook", graphID, nil)
		g.AddAttr(graphID, "label", targetID)
	}

	for dstID, task := range runbook.Tasks {
		graphID := "cluster_" + task.Target
		g.AddNode(graphID, dstID, nil)
		for _, srcID := range task.After {
			g.AddEdge(srcID, dstID, true, nil)
		}
	}

	w.Write([]byte(g.String()))
	return nil
}

// List enumerates a runbook's tasks and their targets
func List(sys *sys.IO, filename string) (map[TaskID]TargetID, error) {
	p := newParser(sys)
	runbook, err := p.Parse(filename)
	if err != nil {
		return nil, err
	}

	result := map[TaskID]TargetID{}
	for taskID, task := range runbook.Tasks {
		result[TaskID(taskID)] = TargetID(task.Target)
	}

	return result, nil
}

// Run executes the tasks in a runbook
func Run(sys *sys.IO, filename string) error {
	p := newParser(sys)
	runbook, err := p.Parse(filename)
	if err != nil {
		return err
	}

	r := newRunner(sys.Log, runbook.Targets, runbook.Tasks)
	return r.run()
}

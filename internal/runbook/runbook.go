package runbook

import (
	"encoding/json"
	"io"
	"path/filepath"

	dot "github.com/awalterschulze/gographviz"
	"github.com/hokaccha/go-prettyjson"

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

func newParser(sys *sys.IO, basedir string) (*parser.Parser, error) {
	basedir, err := filepath.Abs(basedir)
	if err != nil {
		return nil, err
	}
	p := &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"pg": &pg.TargetFactory{
				Log: sys.Log,
			},
			"sh": &sh.TargetFactory{
				Basedir: basedir,
				Stdout:  sys.Stdout,
				Stderr:  sys.Stderr,
			},
			"demo": &demo.TargetFactory{
				Stdout: sys.Stdout,
			},
		},
	}
	return p, nil
}

// Dot generates a GraphViz compatible description of a runbook's tasks
func Dot(sys *sys.IO, filename string, w io.Writer, splines string) error {
	p, err := newParser(sys, filepath.Dir(filename))
	if err != nil {
		return err
	}
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

// Eval converts a runbook to JSON and pretty prints it
func Eval(sys *sys.IO, filename string, w io.Writer, color bool) error {
	p, err := newParser(sys, filepath.Dir(filename))
	if err != nil {
		return err
	}

	runbook, err := p.Load(filename)
	if err != nil {
		return err
	}

	var unformatted interface{}
	err = json.Unmarshal([]byte(runbook), &unformatted)
	if err != nil {
		return err
	}

	formatter := prettyjson.NewFormatter()
	formatter.DisabledColor = !color
	formatted, err := formatter.Marshal(unformatted)
	if err != nil {
		return err
	}

	_, err = w.Write(formatted)
	return err
}

// List enumerates a runbook's tasks and their targets
func List(sys *sys.IO, filename string) (map[TaskID]TargetID, error) {
	p, err := newParser(sys, filepath.Dir(filename))
	if err != nil {
		return nil, err
	}
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
	p, err := newParser(sys, filepath.Dir(filename))
	if err != nil {
		return err
	}
	runbook, err := p.Parse(filename)
	if err != nil {
		return err
	}

	r := newRunner(sys.Log, runbook.Targets, runbook.Tasks)
	return r.run()
}

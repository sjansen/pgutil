package parser_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
	"github.com/sjansen/pgutil/internal/runbook/types"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	l := parser.Parser{
		Targets: map[string]types.TargetFactory{
			"pg": &pgFactory{},
			"sh": &shFactory{},
		},
	}

	expected := &types.Runbook{
		Targets: types.Targets{
			"sh": &shTarget{},
			"src": &pgTarget{
				Concurrency: 3,
				Host:        "localhost",
				DBName:      "tmp-pgutil-src",
				Username:    "AzureDiamond",
				Password:    "hunter2",
			},
			"dst": &pgTarget{
				Concurrency: 3,
				Host:        "localhost",
				DBName:      "tmp-pgutil-dst",
				Username:    "AzureDiamond",
				Password:    "hunter2",
			},
		},
		Tasks: types.Tasks{
			"create-dir": {
				Target: "sh",
				Config: &shTask{
					Args: []string{"mkdir", "-p", "/tmp/pgutil-simple-example"},
				},
			},
			"remove-dir": {
				After:  []string{"delete-old-measurements", "insert-new-measurements"},
				Target: "sh",
				Config: &shTask{
					Args: []string{"rmdir", "/tmp/pgutil-simple-example"},
				},
			},
			"create-table": {
				Target: "src",
				Config: &pgTask{
					SQL: readfile("testdata/scripts/create.sql"),
				},
			},
			"delete-old-measurements": {
				After:  []string{"insert-new-measurements"},
				Target: "src",
				Config: &pgTask{
					SQL: readfile("testdata/scripts/delete.sql"),
				},
			},
			"insert-new-measurements": {
				After:  []string{"create-table"},
				Target: "src",
				Config: &pgTask{
					SQL: readfile("testdata/scripts/insert.sql"),
				},
			},
		},
	}

	actual, err := l.Parse("testdata/simple.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)

	actual, err = l.Parse("testdata/invalid-filename")
	require.Nil(actual)
	require.Error(err)

	actual, err = l.Parse("testdata/invalid-import.jsonnet")
	require.Nil(actual)
	require.Error(err)
}

func readfile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

type pgFactory struct{}

func (f *pgFactory) NewTarget() types.Target {
	return &pgTarget{}
}

type pgTarget struct {
	Concurrency int
	Host        string
	DBName      string
	Username    string
	Password    string
}

func (t *pgTarget) Analyze() error {
	return nil
}
func (t *pgTarget) ConcurrencyLimit() int {
	if t.Concurrency < 1 {
		return 1
	}
	return t.Concurrency
}
func (t *pgTarget) Handle(ctx context.Context, task types.TaskConfig) error {
	return nil
}
func (t *pgTarget) NewTaskConfig(class string) (types.TaskConfig, error) {
	if class != "exec" {
		return nil, fmt.Errorf("invalid task class: %q", class)
	}
	return &pgTask{}, nil
}
func (t *pgTarget) Start() error {
	return nil
}
func (t *pgTarget) Stop() error {
	return nil
}

type shFactory struct{}

func (f *shFactory) NewTarget() types.Target {
	return &shTarget{}
}

type shTarget struct {
	Concurrency int
}

func (t *shTarget) Analyze() error {
	return nil
}
func (t *shTarget) ConcurrencyLimit() int {
	if t.Concurrency < 1 {
		return 1
	}
	return t.Concurrency
}
func (t *shTarget) Handle(ctx context.Context, task types.TaskConfig) error {
	return nil
}
func (t *shTarget) NewTaskConfig(class string) (types.TaskConfig, error) {
	if class != "" {
		return nil, errors.New("invalid task class")
	}
	return &shTask{}, nil
}
func (t *shTarget) Start() error {
	return nil
}
func (t *shTarget) Stop() error {
	return nil
}

type pgTask struct {
	SQL string
}

func (t *pgTask) Check() error {
	return nil
}

type shTask struct {
	Args []string
}

func (t *shTask) Check() error {
	return nil
}

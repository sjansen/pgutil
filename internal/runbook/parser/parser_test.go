package parser_test

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/parser"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	l := parser.Parser{
		Queues: map[string]func() parser.Queue{
			"pg": func() parser.Queue { return &pgQueue{} },
			"sh": func() parser.Queue { return &shQueue{} },
		},
		Tasks: map[string]func() parser.Task{
			"pg/exec": func() parser.Task { return &pgTask{} },
			"sh":      func() parser.Task { return &shTask{} },
		},
	}

	expected := &parser.Runbook{
		Queues: map[string]parser.Queue{
			"sh": &shQueue{},
			"pg": &pgQueue{},
			"pg/dst": &pgQueue{
				Concurrency: 3,
				Host:        "localhost",
				DBName:      "tmp-pgutil-dst",
				Username:    "AzureDiamond",
				Password:    "hunter2",
			},
			"pg/src": &pgQueue{
				Concurrency: 3,
				Host:        "localhost",
				DBName:      "tmp-pgutil-src",
				Username:    "AzureDiamond",
				Password:    "hunter2",
			},
		},
		Steps: map[string]*parser.Step{
			"create-dir": {
				Queue: "sh",
				Task: &shTask{
					Args: []string{"mkdir", "/tmp/pgutil-simple-example"},
				},
			},
			"remove-dir": {
				Queue: "sh",
				After: []string{"delete-old-measurements", "insert-new-measurements"},
				Task: &shTask{
					Args: []string{"rmdir", "/tmp/pgutil-simple-example"},
				},
			},
			"create-table": {
				Queue: "pg/src",
				Task: &pgTask{
					SQL: readfile("testdata/scripts/create.sql"),
				},
			},
			"delete-old-measurements": {
				Queue: "pg/src",
				After: []string{"insert-new-measurements"},
				Task: &pgTask{
					SQL: readfile("testdata/scripts/delete.sql"),
				},
			},
			"insert-new-measurements": {
				Queue: "pg/src",
				After: []string{"create-table"},
				Task: &pgTask{
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

type pgQueue struct {
	Concurrency int
	Host        string
	DBName      string
	Username    string
	Password    string
}

func (q *pgQueue) ConcurrencyLimit() int {
	if q.Concurrency < 1 {
		return 1
	}
	return q.Concurrency
}
func (q *pgQueue) VerifyConfig() error {
	return nil
}
func (q *pgQueue) VerifyTask(config interface{}) error {
	if _, ok := config.(*pgTask); !ok {
		return errors.New("invalid pg task")
	}
	return nil
}

type shQueue struct {
	Concurrency int
}

func (q *shQueue) ConcurrencyLimit() int {
	if q.Concurrency < 1 {
		return 1
	}
	return q.Concurrency
}
func (q *shQueue) VerifyConfig() error {
	return nil
}
func (q *shQueue) VerifyTask(config interface{}) error {
	if _, ok := config.(*shTask); !ok {
		return errors.New("invalid sh task")
	}
	return nil
}

type pgTask struct {
	SQL string
}

func (t *pgTask) VerifyConfig() error {
	return nil
}

type shTask struct {
	Args []string
}

func (t *shTask) VerifyConfig() error {
	return nil
}

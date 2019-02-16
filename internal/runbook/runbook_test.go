package runbook_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook"
)

func readfile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func TestLoad(t *testing.T) {
	require := require.New(t)

	expected := &runbook.Config{
		Databases: map[string]*runbook.Database{
			"default": {
				Host:     "localhost",
				DBName:   "tmp",
				Username: "AzureDiamond",
				Password: "hunter2",
			},
		},
		Tasks: map[string]*runbook.Task{
			"create-dir": {
				TaskExec: &runbook.TaskExec{
					Args: []string{"mkdir", "/tmp/pgutil-simple-example"},
				},
			},
			"remove-dir": {
				After: []string{"delete-old-measurements", "insert-new-measurements"},
				TaskExec: &runbook.TaskExec{
					Args: []string{"rmdir", "/tmp/pgutil-simple-example"},
				},
			},
			"create-table": {
				After: []string{"create-dir"},
				TaskSQL: &runbook.TaskSQL{
					SQL: readfile("testdata/create.sql"),
				},
			},
			"delete-old-measurements": {
				After: []string{"create-table"},
				TaskSQL: &runbook.TaskSQL{
					SQL: readfile("testdata/delete.sql"),
				},
			},
			"insert-new-measurements": {
				After: []string{"create-table"},
				TaskSQL: &runbook.TaskSQL{
					SQL: readfile("testdata/insert.sql"),
				},
			},
		},
	}

	actual, err := runbook.Load("testdata", "simple.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)

	actual, err = runbook.Load("testdata", "invalid-filename")
	require.Nil(actual)
	require.Error(err)

	actual, err = runbook.Load("testdata", "invalid-import.jsonnet")
	require.Nil(actual)
	require.Error(err)

	actual, err = runbook.Load("testdata", "invalid-task-field.jsonnet")
	require.Nil(actual)
	require.Error(err)
}

package oldrunbook_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/oldrunbook"
	"github.com/sjansen/pgutil/internal/oldrunbook/tasks"
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

	expected := &oldrunbook.Config{
		Databases: map[string]*oldrunbook.Database{
			"default": {
				Host:     "localhost",
				DBName:   "tmp",
				Username: "AzureDiamond",
				Password: "hunter2",
			},
		},
		Tasks: map[string]*oldrunbook.Task{
			"create-dir": {
				Exec: &tasks.Exec{
					Args: []string{"mkdir", "/tmp/pgutil-simple-example"},
				},
			},
			"remove-dir": {
				After: []string{"delete-old-measurements", "insert-new-measurements"},
				Exec: &tasks.Exec{
					Args: []string{"rmdir", "/tmp/pgutil-simple-example"},
				},
			},
			"create-table": {
				After: []string{"create-dir"},
				SQL: &tasks.SQL{
					SQL: readfile("testdata/create.sql"),
				},
			},
			"delete-old-measurements": {
				After: []string{"create-table"},
				SQL: &tasks.SQL{
					SQL: readfile("testdata/delete.sql"),
				},
			},
			"insert-new-measurements": {
				After: []string{"create-table"},
				SQL: &tasks.SQL{
					SQL: readfile("testdata/insert.sql"),
				},
			},
		},
	}

	actual, err := oldrunbook.Load("testdata/simple.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)

	actual, err = oldrunbook.Load("testdata/invalid-filename")
	require.Nil(actual)
	require.Error(err)

	actual, err = oldrunbook.Load("testdata/invalid-import.jsonnet")
	require.Nil(actual)
	require.Error(err)

	actual, err = oldrunbook.Load("testdata/invalid-task-field.jsonnet")
	require.Nil(actual)
	require.Error(err)
}

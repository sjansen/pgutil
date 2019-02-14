package taskset_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/taskset"
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

	expected := &taskset.Config{
		Databases: map[string]map[string]string{
			"default": {
				"host":     "localhost",
				"dbname":   "tmp",
				"username": "AzureDiamond",
				"password": "hunter2",
			},
		},
		Tasks: map[string]map[string]string{
			"create-table": {
				"sql": readfile("testdata/create.sql"),
			},
			"delete-old-measurements": {
				"after": "create-table",
				"sql":   readfile("testdata/delete.sql"),
			},
			"insert-new-measurements": {
				"after": "create-table",
				"sql":   readfile("testdata/insert.sql"),
			},
		},
	}

	actual, err := taskset.Load("testdata", "simple.jsonnet")
	require.NoError(err)
	require.Equal(expected, actual)

	actual, err = taskset.Load("testdata", "invalid-filename")
	require.Nil(actual)
	require.Error(err)

	actual, err = taskset.Load("testdata", "invalid-import.jsonnet")
	require.Nil(actual)
	require.Error(err)
}

package taskset_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/taskset/parser"
	"github.com/sjansen/pgutil/internal/taskset/types"
)

func TestParseFile(t *testing.T) {
	require := require.New(t)

	ts, err := parser.ParseFile("testdata/simple.hcl")
	require.NoError(err)

	pg := ts.Targets["pg"]["src"]
	sh := ts.Targets["sh"][""]

	pgQueue, pgResults := pg.Start()
	shQueue, shResults := sh.Start()

	shQueue <- map[string]types.Task{
		"create-dir": ts.Tasks["create-dir"],
	}
	result := <-shResults
	require.NoError(result["create-dir"])

	pgQueue <- map[string]types.Task{
		"create-table": ts.Tasks["create-table"],
	}
	result = <-pgResults
	require.NoError(result["create-table"])

	pgQueue <- map[string]types.Task{
		"insert-new-measurements": ts.Tasks["insert-new-measurements"],
	}
	result = <-pgResults
	require.NoError(result["insert-new-measurements"])

	pgQueue <- map[string]types.Task{
		"delete-old-measurements": ts.Tasks["delete-old-measurements"],
	}
	result = <-pgResults
	require.NoError(result["delete-old-measurements"])

	shQueue <- map[string]types.Task{
		"remove-dir": ts.Tasks["remove-dir"],
	}
	result = <-shResults
	require.NoError(result["remove-dir"])

	close(pgQueue)
	close(shQueue)
}

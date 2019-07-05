package taskset_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/taskset/demo"
	"github.com/sjansen/pgutil/internal/taskset/parser"
	"github.com/sjansen/pgutil/internal/taskset/pg"
	"github.com/sjansen/pgutil/internal/taskset/sh"
	"github.com/sjansen/pgutil/internal/taskset/types"
)

func newParser() *parser.Parser {
	log := logger.Discard()
	return &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"pg": &pg.TargetFactory{
				Log: log,
			},
			"sh": &sh.TargetFactory{
				Log: log,
			},
		},
	}
}

func TestTaskExecution(t *testing.T) {
	require := require.New(t)

	p := newParser()
	ts, err := p.Parse("testdata/measurements.hcl")
	require.NoError(err)

	pg := ts.Targets["pg"]["src"]
	sh := ts.Targets["sh"][""]

	pgQueue, pgResults := pg.Start()
	shQueue, shResults := sh.Start()

	shQueue <- map[types.TaskID]types.Task{
		"create-dir": ts.Tasks["create-dir"],
	}
	result := <-shResults
	require.NoError(result["create-dir"])

	pgQueue <- map[types.TaskID]types.Task{
		"create-table": ts.Tasks["create-table"],
	}
	result = <-pgResults
	require.NoError(result["create-table"])

	pgQueue <- map[types.TaskID]types.Task{
		"insert-new-measurements": ts.Tasks["insert-new-measurements"],
	}
	result = <-pgResults
	require.NoError(result["insert-new-measurements"])

	pgQueue <- map[types.TaskID]types.Task{
		"delete-old-measurements": ts.Tasks["delete-old-measurements"],
	}
	result = <-pgResults
	require.NoError(result["delete-old-measurements"])

	shQueue <- map[types.TaskID]types.Task{
		"remove-dir": ts.Tasks["remove-dir"],
	}
	result = <-shResults
	require.NoError(result["remove-dir"])

	close(pgQueue)
	close(shQueue)
}

func TestDemoTaskExecution(t *testing.T) {
	require := require.New(t)

	buf := &bytes.Buffer{}
	log := logger.Discard()
	p := &parser.Parser{
		Targets: map[string]types.TargetFactory{
			"demo": &demo.TargetFactory{
				Log:    log,
				Stdout: buf,
			},
		},
	}
	ts, err := p.Parse("testdata/demo.hcl")
	require.NoError(err)

	for _, target := range []types.TargetID{"msg1", "msg2", "msg3"} {
		msg := ts.Targets["demo"][target]
		queue, results := msg.Start()

		for _, name := range []string{"/reverse", "/rotate"} {
			taskID := types.TaskID(string(target) + name)
			queue <- map[types.TaskID]types.Task{
				taskID: ts.Tasks[taskID],
			}
		}

		for i := 0; i < 2; i++ {
			result := <-results
			for _, err := range result {
				require.NoError(err)
			}
		}

		taskID := types.TaskID(string(target) + "/decrypted")
		queue <- map[types.TaskID]types.Task{
			taskID: ts.Tasks[taskID],
		}

		result := <-results
		require.NoError(result[taskID])
	}

	expected := `Spoon!
The cake is a lie.
The world is a vampire.
`
	require.Equal(expected, buf.String())
}

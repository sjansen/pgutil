package sh_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook/sh"
)

func TestTasks(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	factory := &sh.TargetFactory{
		Stdout: &stdout,
		Stderr: &stderr,
	}
	target := factory.NewTarget()

	ctx := context.TODO()
	require.NoError(target.Analyze())
	require.Equal("", stdout.String())
	require.Equal("", stderr.String())

	task := &sh.Exec{}
	require.Error(task.Check())

	for _, tc := range []struct {
		args   []string
		err    bool
		stdout string
		stderr string
	}{{
		args:   []string{"false"},
		err:    true,
		stdout: "",
		stderr: "",
	}, {
		args:   []string{"true"},
		err:    false,
		stdout: "",
		stderr: "",
	}, {
		args:   []string{"echo", "Spoon!"},
		err:    false,
		stdout: "Spoon!\n",
		stderr: "",
	}} {
		task := &sh.Exec{Args: tc.args}
		require.NoError(task.Check())

		stdout.Reset()
		stderr.Reset()
		err := target.Handle(ctx, task)
		if tc.err {
			require.Error(err, tc.args[0])
		} else {
			require.NoError(err, tc.args[0])
		}
		require.Equal(tc.stdout, stdout.String())
		require.Equal(tc.stderr, stderr.String())
	}
}

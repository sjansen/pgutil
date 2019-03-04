package exec_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/tasks/exec"
)

func TestRun(t *testing.T) {
	require := require.New(t)

	ctx := context.Background()
	var stdout, stderr bytes.Buffer
	task := &exec.Task{
		Args:   []string{"echo", "Kilroy", "was", "here."},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	err := task.Start(ctx)
	require.NoError(err)
	expected := "Kilroy was here.\n"
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())

	stdout.Reset()
	stderr.Reset()
	task.Args = []string{"non-existent-command"}
	err = task.Start(ctx)
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())

	stdout.Reset()
	stderr.Reset()
	task.Args = []string{"scripts/non-existent-command"}
	err = task.Start(ctx)
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())
}

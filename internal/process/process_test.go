package process_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/process"
)

func TestRun(t *testing.T) {
	require := require.New(t)

	expected := "Kilroy was here.\n"
	var stdout, stderr bytes.Buffer
	c := process.Command{
		Args: []string{"echo", "Kilroy", "was", "here."},

		Stdout: &stdout,
		Stderr: &stderr,
	}
	err := c.Run()
	require.NoError(err)
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())

	c.Args = []string{"non-existent-command"}
	stdout.Reset()
	stderr.Reset()
	err = c.Run()
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())

	c.Args = []string{"scripts/non-existent-command"}
	stdout.Reset()
	stderr.Reset()
	err = c.Run()
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())
}

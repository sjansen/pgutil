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
	args := []string{"echo", "Kilroy", "was", "here."}
	err := process.Create(args).
		Run(&stdout, &stderr)
	require.NoError(err)
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())

	stdout.Reset()
	stderr.Reset()
	args = []string{"non-existent-command"}
	err = process.Create(args).
		Run(&stdout, &stderr)
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())

	stdout.Reset()
	stderr.Reset()
	args = []string{"scripts/non-existent-command"}
	err = process.Create(args).
		Run(&stdout, &stderr)
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())
}

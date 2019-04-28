package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
)

func TestRunBookRun(t *testing.T) {
	require := require.New(t)

	expected := `.ravgyniB ehbl xaveq bg rehf rO
Be sure to drink your Ovaltine.
`
	cmd := &commands.RunBookRunCmd{
		File: "testdata/message.jsonnet",
	}
	var stdout, stderr bytes.Buffer
	err := cmd.Run(&stdout, &stderr, nil)
	require.NoError(err)
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())

	cmd.File = "non-existent-runbook-file"
	stdout.Reset()
	stderr.Reset()
	err = cmd.Run(&stdout, &stderr, nil)
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())
}

package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/sys"
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
	base := &commands.Base{
		IO: sys.IO{
			Log:    logger.Discard(),
			Stdout: &stdout,
			Stderr: &stderr,
		},
	}
	err := cmd.Run(base)
	require.NoError(err)
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())

	cmd.File = "non-existent-runbook-file"
	stdout.Reset()
	stderr.Reset()
	err = cmd.Run(base)
	require.Error(err)
	require.Empty(stdout.String())
	require.Empty(stderr.String())
}

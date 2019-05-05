package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/logger"
)

func TestPing(t *testing.T) {
	require := require.New(t)

	cmd := &commands.PingCmd{}

	var stdout, stderr bytes.Buffer
	base := &commands.Base{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}

	err := cmd.Run(base)
	require.NoError(err)
	require.NotEmpty(stdout.String())
	require.Empty(stderr.String())
}

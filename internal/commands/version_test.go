package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/sys"
)

func TestVersion(t *testing.T) {
	require := require.New(t)

	expected := "pgmagic 1.0-test\n"
	cmd := &commands.VersionCmd{
		App:   "pgmagic",
		Build: "1.0-test",
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
}

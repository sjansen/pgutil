// +build !go1.12

package cli_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/cli"
	"github.com/sjansen/pgutil/internal/commands"
)

func TestLegacyVersion(t *testing.T) {
	require := require.New(t)

	parser := cli.RegisterCommands("test-version")
	for _, tc := range []struct {
		args        []string
		expected    interface{}
		expectError bool
	}{{
		args: []string{
			"version",
		},
		expected: &commands.VersionCmd{
			App:   "pgutil",
			Build: "test",
		},
	}, {
		args: []string{
			"version", "--long",
		},
		expectError: true,
	}} {
		cmd, err := parser.Parse(tc.args)
		if tc.expectError {
			require.Error(err)
		} else {
			require.NoError(err)

			var stdout, stderr bytes.Buffer
			err = cmd(&stdout, &stderr)
			require.NoError(err)

			require.Contains(stdout.String(), "test-version")
			require.Empty(stderr)
		}
	}
}

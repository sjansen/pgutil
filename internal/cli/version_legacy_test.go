// +build !go1.12

package cli_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/cli"
	"github.com/sjansen/pgutil/internal/commands"
)

func TestVersion(t *testing.T) {
	require := require.New(t)

	parser := cli.RegisterCommands("test")
	for _, tc := range []struct {
		args        []string
		expected    cli.Command
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
			"version", "--verbose",
		},
		expectError: true,
	}} {
		actual, err := parser.Parse(tc.args)
		if tc.expectError {
			require.Error(err)
		} else {
			require.NoError(err)
			require.Equal(tc.expected, actual)
		}
	}
}

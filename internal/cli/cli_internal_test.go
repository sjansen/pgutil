package cli

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
)

func TestArgParser(t *testing.T) {
	require := require.New(t)

	parser := RegisterCommands("test")
	for _, tc := range []struct {
		args        []string
		expected    interface{}
		expectError bool
	}{{
		args: []string{
			"ping",
		},
		expected: &commands.PingCmd{},
	}, {
		args: []string{
			"ping", "-h", "ssh.example.com", "-p", "2222", "-d", "template1", "-U", "postgres",
		},
		expected: &commands.PingCmd{
			Host:     "ssh.example.com",
			Port:     2222,
			DBName:   "template1",
			Username: "postgres",
		},
	}, {
		args: []string{
			"ping", "--host", "ssh.example.com", "--port", "2222",
			"--dbname", "template1", "--username", "postgres",
		},
		expected: &commands.PingCmd{
			Host:     "ssh.example.com",
			Port:     2222,
			DBName:   "template1",
			Username: "postgres",
		},
	}, {
		args: []string{
			"runbook",
		},
		expectError: true,
	}, {
		args: []string{
			"runbook", "testdata/foo",
		},
		expected: &commands.RunBookRunCmd{
			File: "testdata/foo",
		},
	}, {
		args: []string{
			"runbook", "list", "testdata/foo",
		},
		expected: &commands.RunBookListCmd{
			File: "testdata/foo",
		},
	}, {
		args: []string{
			"runbook", "ls", "testdata/foo",
		},
		expected: &commands.RunBookListCmd{
			File: "testdata/foo",
		},
	}, {
		args: []string{
			"runbook", "run", "testdata/foo",
		},
		expected: &commands.RunBookRunCmd{
			File: "testdata/foo",
		},
	}, {
		args: []string{
			"runbook", "run", "testdata/nonexistent",
		},
		expectError: true,
	}} {
		_, err := parser.Parse(tc.args)
		if tc.expectError {
			require.Error(err)
		} else {
			require.NoError(err)
			require.Equal(tc.expected, parser.cmd)
		}
	}
}

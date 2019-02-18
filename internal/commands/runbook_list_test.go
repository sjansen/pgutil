package commands_test

import (
	"bytes"
	"testing"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/stretchr/testify/require"
)

func TestRunBookList(t *testing.T) {
	require := require.New(t)

	expected := `Tasks:
    begin
    end
    hostname
    whoami
`
	cmd := &commands.RunBookListCmd{
		File: "testdata/list.jsonnet",
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

package runbook_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/sys"
)

func TestPg(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	sys := &sys.IO{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}

	err := runbook.Run(sys, "testdata/pg.jsonnet")
	require.NoError(err)
	require.Equal("", stdout.String())
	require.Equal("", stderr.String())
}

func TestPgErrors(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	sys := &sys.IO{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}
	for _, filename := range []string{
		"testdata/pg-invalid-task-class.jsonnet",
		"testdata/pg-invalid-task-field.jsonnet",
	} {
		err := runbook.Run(sys, filename)
		require.Error(err)
	}
}

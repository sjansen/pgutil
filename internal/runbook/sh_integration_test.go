package runbook_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/sys"
)

func TestSh(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	sys := &sys.IO{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}
	err := runbook.Run(sys, "testdata/sh.jsonnet")
	require.NoError(err)

	expected := `The world is a vampire, sent to drain
Secret destroyers, hold you up to the flames
And what do I get, for my pain?
Betrayed desires, and a piece of the game
`
	require.Equal(expected, stdout.String())
	require.Equal("", stderr.String())
}

func TestShErrors(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	sys := &sys.IO{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}
	for _, filename := range []string{
		"testdata/sh-invalid-task-class.jsonnet",
		"testdata/sh-invalid-task-field.jsonnet",
	} {
		err := runbook.Run(sys, filename)
		require.Error(err)
	}
}

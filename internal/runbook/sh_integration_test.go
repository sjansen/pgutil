package runbook_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/runbook"
)

func TestSh(t *testing.T) {
	require := require.New(t)

	var stdout, stderr bytes.Buffer
	err := runbook.Run("testdata/sh.jsonnet", &stdout, &stderr)
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

	for _, filename := range []string{
		"testdata/sh-invalid-task-class.jsonnet",
		"testdata/sh-invalid-task-field.jsonnet",
	} {
		err := runbook.Run(filename, nil, nil)
		require.Error(err)
	}
}

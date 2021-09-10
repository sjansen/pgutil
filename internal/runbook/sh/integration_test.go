//go:build integration
// +build integration

package sh_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/sys"
)

var chorus = `
Despite all my rage I am still just a rat in a cage
Despite all my rage I am still just a rat in a cage
Someone will say what is lost can never be saved
Despite all my rage I am still just a rat in a cage`

var expected = `The world is a vampire, sent to drain
Secret destroyers, hold you up to the flames
And what do I get, for my pain?
Betrayed desires, and a piece of the game

Even though I know, I suppose I'll show
All my cool and cold--like old Job

Despite all my rage I am still just a rat in a cage
Despite all my rage I am still just a rat in a cage
Someone will say what is lost can never be saved
Despite all my rage I am still just a rat in a cage
`

func TestSh(t *testing.T) {
	require := require.New(t)

	err := os.Setenv("PGUTIL_CHORUS", chorus)
	require.NoError(err)

	var stdout, stderr bytes.Buffer
	sys := &sys.IO{
		Log:    logger.Discard(),
		Stdout: &stdout,
		Stderr: &stderr,
	}

	err = runbook.Run(sys, "testdata/sh.jsonnet")
	require.NoError(err)
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

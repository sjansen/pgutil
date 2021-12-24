//go:build integration
// +build integration

package pg_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/runbook"
	"github.com/sjansen/pgutil/internal/sys"
	"github.com/sjansen/pgutil/internal/testutil"
)

func TestPg(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			var stdout, stderr bytes.Buffer
			sys := &sys.IO{
				Log:    logger.Discard(),
				Stdout: &stdout,
				Stderr: &stderr,
			}

			filename := "testdata/pg.jsonnet"
			err := runbook.Run(sys, filename)
			require.NoError(err, filename)
		})
	}
}

func TestPgErrors(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			var stdout, stderr bytes.Buffer
			sys := &sys.IO{
				Log:    logger.Discard(),
				Stdout: &stdout,
				Stderr: &stderr,
			}
			for _, filename := range []string{
				"testdata/pg-bad-target-config.jsonnet",
				"testdata/pg-invalid-task-class.jsonnet",
				"testdata/pg-invalid-task-field.jsonnet",
			} {
				t.Run(filename, func(t *testing.T) {
					require := require.New(t)

					err := runbook.Run(sys, filename)
					require.Error(err, filename)
				})
			}
		})
	}
}

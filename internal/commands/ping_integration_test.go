package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/logger"
	"github.com/sjansen/pgutil/internal/sys"
	"github.com/sjansen/pgutil/internal/testutil"
)

func TestPing(t *testing.T) {
	for _, pghost := range testutil.PGHosts() {
		pghost := pghost
		t.Run(pghost, func(t *testing.T) {
			t.Setenv("PGHOST", pghost)
			require := require.New(t)

			cmd := &commands.PingCmd{}

			var stdout, stderr bytes.Buffer
			base := &commands.Base{
				IO: sys.IO{
					Log:    logger.Discard(),
					Stdout: &stdout,
					Stderr: &stderr,
				},
			}

			t.Setenv("PGHOST", "pg9.6")
			err := cmd.Run(base)
			require.NoError(err)
			require.NotEmpty(stdout.String())
			require.Empty(stderr.String())
		})
	}
}

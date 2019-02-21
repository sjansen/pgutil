package commands_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sjansen/pgutil/internal/commands"
	"github.com/sjansen/pgutil/internal/mocks"
)

func TestPing(t *testing.T) {
	require := require.New(t)

	version := "PostgreSQL 9.6 compatible, mock"
	db := new(mocks.DB)
	db.On("Close").Return(nil)
	db.On("ServerVersion").Return(version, nil)
	deps := &commands.Dependencies{
		DB: func(opts map[string]string) (commands.DB, error) {
			return db, nil
		},
	}

	expected := version + "\n"
	var stdout, stderr bytes.Buffer
	cmd := &commands.PingCmd{}
	err := cmd.Run(&stdout, &stderr, deps)
	require.NoError(err)
	require.Equal(expected, stdout.String())
	require.Empty(stderr.String())
}

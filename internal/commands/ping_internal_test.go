package commands

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPingOptions(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		cmd      *PingCmd
		expected map[string]string
	}{{
		cmd:      &PingCmd{},
		expected: map[string]string{},
	}, {
		cmd: &PingCmd{
			Host: "db.example.com",
		},
		expected: map[string]string{
			"addr": "db.example.com:5432",
		},
	}, {
		cmd: &PingCmd{
			Host:     "db.example.com",
			Port:     "5439",
			DBName:   "foo",
			Username: "bob",
		},
		expected: map[string]string{
			"addr":     "db.example.com:5439",
			"database": "foo",
			"user":     "bob",
		},
	}} {
		actual := tc.cmd.dbOptions()
		require.Equal(tc.expected, actual)
	}
}

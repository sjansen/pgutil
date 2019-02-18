package db

import (
	"os"
	"os/user"
	"testing"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/require"
)

var envVariables = []string{
	"PGHOST", "PGPORT", "PGDATABASE", "PGUSER", "PGPASSWORD",
}

func TestMergeOptions(t *testing.T) {
	require := require.New(t)

	user, err := user.Current()
	require.NoError(err)
	require.NotEmpty(user.Username)

	// Reset environment variables once test is finished
	for _, key := range envVariables {
		value, set := os.LookupEnv(key)
		defer func(key, value string, set bool) {
			if set {
				err := os.Setenv(key, value)
				require.NoError(err)
			} else {
				err := os.Unsetenv(key)
				require.NoError(err)
			}
		}(key, value, set)
		err := os.Unsetenv(key)
		require.NoError(err)
	}

	for _, tc := range []struct {
		environ  map[string]string
		options  map[string]string
		expected *pg.Options
	}{{
		expected: &pg.Options{
			User: user.Username,
		},
	}, {
		environ: map[string]string{
			"PGHOST":     "foo",
			"PGDATABASE": "baz",
			"PGUSER":     "qux",
			"PGPASSWORD": "quux",
		},
		expected: &pg.Options{
			Addr:     "foo:5432",
			Database: "baz",
			User:     "qux",
			Password: "quux",
		},
	}, {
		environ: map[string]string{
			"PGPORT": "bar",
		},
		expected: &pg.Options{
			Addr:     "foo:bar",
			Database: "baz",
			User:     "qux",
			Password: "quux",
		},
	}, {
		options: map[string]string{
			"addr":     "FOO",
			"database": "BAR",
			"user":     "BAZ",
			"password": "QUX",
		},
		expected: &pg.Options{
			Addr:     "FOO",
			Database: "BAR",
			User:     "BAZ",
			Password: "QUX",
		},
	}} {
		for k, v := range tc.environ {
			err := os.Setenv(k, v)
			require.NoError(err)
		}

		actual, err := mergeOptions(tc.options)
		require.NoError(err)
		require.Equal(tc.expected, actual)
	}
}

package pg

import (
	"os"
	"os/user"
	"testing"

	gopg "github.com/go-pg/pg"

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

	// Clear environment variables until test is finished
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
		options  *Options
		expected *gopg.Options
	}{{
		expected: &gopg.Options{
			ApplicationName: "pgutil",
			User:            user.Username,
		},
	}, {
		environ: map[string]string{
			"PGHOST":     "foo",
			"PGDATABASE": "baz",
			"PGUSER":     "qux",
			"PGPASSWORD": "quux",
		},
		expected: &gopg.Options{
			ApplicationName: "pgutil",
			Addr:            "foo:5432",
			Database:        "baz",
			User:            "qux",
			Password:        "quux",
		},
	}, {
		environ: map[string]string{
			"PGPORT": "bar",
		},
		expected: &gopg.Options{
			ApplicationName: "pgutil",
			Addr:            "foo:bar",
			Database:        "baz",
			User:            "qux",
			Password:        "quux",
		},
	}, {
		options: &Options{
			Address:  "FOO",
			Database: "BAR",
			Username: "BAZ",
			Password: "QUX",
		},
		expected: &gopg.Options{
			ApplicationName: "pgutil",
			Addr:            "FOO",
			Database:        "BAR",
			User:            "BAZ",
			Password:        "QUX",
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

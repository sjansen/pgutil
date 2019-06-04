package ddl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseIndex(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Stmt     string
		Expected *Index
	}{{Stmt: `
		CREATE UNIQUE INDEX foo__key
		ON foo USING btree (key)
	`,
		// WHERE value IS NULL
		Expected: &Index{
			Table:  "foo",
			Name:   "foo__key",
			Unique: true,
			Using:  "btree",
			Keys: []*IndexKey{
				{Column: "key"},
			},
		},
	}, {Stmt: `
		CREATE INDEX foo__key__value
		ON foo USING btree (key, value)
	`,
		// WHERE value IS NULL
		Expected: &Index{
			Table: "foo",
			Name:  "foo__key__value",
			Using: "btree",
			Keys: []*IndexKey{
				{Column: "key"},
				{Column: "value"},
			},
		},
	}} {
		actual, err := ParseIndex(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

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
		ON foo USING btree ( key, value )
	`,
		Expected: &Index{
			Table: "foo",
			Name:  "foo__key__value",
			Using: "btree",
			Keys: []*IndexKey{
				{Column: "key"},
				{Column: "value"},
			},
		},
	}, {Stmt: `
		CREATE INDEX foo__uniq__key__null_value
		ON foo ( key )
		WHERE value IS NULL
	`,
		Expected: &Index{
			Table: "foo",
			Name:  "foo__uniq__key__null_value",
			Where: "value IS NULL",
			Keys: []*IndexKey{
				{Column: "key"},
			},
		},
	}, {Stmt: `
		CREATE INDEX foo__key__like
		ON foo (key varchar_pattern_ops)
	`,
		Expected: &Index{
			Table: "foo",
			Name:  "foo__key__like",
			Keys: []*IndexKey{
				{Column: "key", OpClass: "varchar_pattern_ops"},
			},
		},
	}, {Stmt: `
		CREATE INDEX foo__key__value__like
		ON foo ( key varchar_pattern_ops , value varchar_pattern_ops )
	`,
		Expected: &Index{
			Table: "foo",
			Name:  "foo__key__value__like",
			Keys: []*IndexKey{
				{Column: "key", OpClass: "varchar_pattern_ops"},
				{Column: "value", OpClass: "varchar_pattern_ops"},
			},
		},
	}, {Stmt: `
		CREATE INDEX foo__key__upper
		ON foo ( upper(key::text) )
	`,
		Expected: &Index{
			Table: "foo",
			Name:  "foo__key__upper",
			Keys: []*IndexKey{
				{Expression: "upper(key::text)"},
			},
		},
	}, {Stmt: `
		CREATE INDEX foo__value__even_odd
		ON foo ( ( value % 2 ) )
	`,
		Expected: &Index{
			Table: "foo",
			Name:  "foo__value__even_odd",
			Keys: []*IndexKey{
				{Expression: "( value % 2 )"},
			},
		},
	}} {
		actual, err := ParseIndex(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

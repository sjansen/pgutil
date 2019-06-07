package ddl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseForeignKey(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Stmt     string
		Expected *ForeignKey
	}{{Stmt: `
		FOREIGN KEY (foo_id) REFERENCES foo(id)
	`,
		Expected: &ForeignKey{
			Table:      "foo",
			Columns:    []string{"foo_id"},
			Referenced: []string{"id"},
		},
	}, {Stmt: `
		FOREIGN KEY (schema, table)
		REFERENCES tables(schemaName, tableName)
		MATCH FULL
		ON DELETE CASCADE
		ON UPDATE SET DEFAULT
		DEFERRABLE INITIALLY DEFERRED
	`,
		Expected: &ForeignKey{
			Table:             "tables",
			Columns:           []string{"schema", "table"},
			Referenced:        []string{"schemaName", "tableName"},
			Match:             "FULL",
			OnDelete:          "CASCADE",
			OnUpdate:          "SET DEFAULT",
			Deferrable:        true,
			InitiallyDeferred: true,
		},
	}, {Stmt: `
		Foreign Key (a, b, c)
		References T(d, e, f)
		Match   Simple
		On   Delete   Restrict
		On   Update   Set   Null
		Deferrable
		Initially   Deferred
	`,
		Expected: &ForeignKey{
			Table:             "T",
			Columns:           []string{"a", "b", "c"},
			Referenced:        []string{"d", "e", "f"},
			Match:             "SIMPLE",
			OnDelete:          "RESTRICT",
			OnUpdate:          "SET NULL",
			Deferrable:        true,
			InitiallyDeferred: true,
		},
	}} {
		actual, err := ParseForeignKey(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

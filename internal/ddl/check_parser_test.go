package ddl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCheck(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Stmt     string
		Expected *Check
	}{{Stmt: `
		CHECK (0 <= value)
	`,
		Expected: &Check{
			Expression: "0 <= value",
		},
	}, {Stmt: `
		CHECK (0 <= value)
		DEFERRABLE INITIALLY DEFERRED
	`,
		Expected: &Check{
			Expression:        "0 <= value",
			Deferrable:        true,
			InitiallyDeferred: true,
		},
	}, {Stmt: `
		CHECK (jsonb_typeof(data) = 'object'::text)
	`,
		Expected: &Check{
			Expression: "jsonb_typeof(data) = 'object'::text",
		},
	}, {Stmt: `
CHECK (
CASE
    WHEN key::text = 'begin'::text THEN 0 <= value
    ELSE 0 < value
END)`,
		Expected: &Check{
			Expression: `
CASE
    WHEN key::text = 'begin'::text THEN 0 <= value
    ELSE 0 < value
END`,
		},
	}} {
		actual, err := ParseCheck(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

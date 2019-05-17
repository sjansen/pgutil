package ddl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTrigger(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Stmt     string
		Expected *Trigger
	}{{Stmt: `
		  CREATE TRIGGER update_foo_modified
		`,
		Expected: &Trigger{
			Name: "update_foo_modified",
		},
	}, {Stmt: `
		  CREATE CONSTRAINT TRIGGER trigger1
		`,
		Expected: &Trigger{
			Name:       "trigger1",
			Constraint: true,
		},
	}, {Stmt: `
		  CREATE TRIGGER trigger2
		`,
		Expected: &Trigger{
			Name: "trigger2",
		},
	}, {Stmt: `
		  create  constraint  trigger  TRIGGER3
		`,
		Expected: &Trigger{
			Name:       "TRIGGER3",
			Constraint: true,
		},
	}} {
		actual, err := ParseTrigger(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

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
		  BEFORE UPDATE ON foo
		  FOR EACH ROW
		  EXECUTE PROCEDURE update_modified_column()
		`,
		Expected: &Trigger{
			Table:      "foo",
			Name:       "update_foo_modified",
			Function:   "update_modified_column",
			When:       "BEFORE",
			ForEachRow: true,
			Update:     true,
		},
	}, {Stmt: `
		  CREATE CONSTRAINT TRIGGER trigger1
		  AFTER INSERT OR DELETE ON table1
		  DEFERRABLE INITIALLY DEFERRED
		  FOR EACH ROW
		  EXECUTE FUNCTION fn1()
		`,
		Expected: &Trigger{
			Table:             "table1",
			Name:              "trigger1",
			Function:          "fn1",
			When:              "AFTER",
			Constraint:        true,
			Deferrable:        true,
			InitiallyDeferred: true,
			ForEachRow:        true,
			Insert:            true,
			Delete:            true,
		},
	}, {Stmt: `
		  CREATE TRIGGER trigger2
		  BEFORE TRUNCATE ON table2
		  FOR EACH STATEMENT
		  EXECUTE PROCEDURE fn2()
		`,
		Expected: &Trigger{
			Table:    "table2",
			Name:     "trigger2",
			Function: "fn2",
			When:     "BEFORE",
			Truncate: true,
		},
	}, {Stmt: `
		  create  trigger  TRIGGER3
		  instead  of  update  on  VIEW3
		  execute  procedure  FN3()
		`,
		Expected: &Trigger{
			Table:    "VIEW3",
			Name:     "TRIGGER3",
			Function: "FN3",
			When:     "INSTEAD OF",
			Update:   true,
		},
	}, {Stmt: `
		  Create Constraint Trigger
		  Trigger4
		  After Update Of col1, col2 On
		  Table4
		  NOT	Deferrable
		  For	Each	Statement
		  Execute   Procedure
		  Func4()
		`,
		Expected: &Trigger{
			Table:      "Table4",
			Name:       "Trigger4",
			Function:   "Func4",
			When:       "AFTER",
			Constraint: true,
			Update:     true,
			Columns:    []string{"col1", "col2"},
		},
	}, {Stmt: `
		  CREATE CONSTRAINT TRIGGER trigger5
		  AFTER INSERT OR UPDATE OR DELETE ON table5
		  INITIALLY IMMEDIATE
		  FOR ROW
		  EXECUTE PROCEDURE proc5()
		`,
		Expected: &Trigger{
			Table:      "table5",
			Name:       "trigger5",
			Function:   "proc5",
			When:       "AFTER",
			Constraint: true,
			ForEachRow: true,
			Insert:     true,
			Update:     true,
			Delete:     true,
		},
	}} {
		actual, err := ParseTrigger(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

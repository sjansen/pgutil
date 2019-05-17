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
		`,
		Expected: &Trigger{
			Table:  "foo",
			Name:   "update_foo_modified",
			Called: "BEFORE",
			Events: []*TriggerEvent{
				{Event: "UPDATE"},
			},
			ForEach: "ROW",
		},
	}, {Stmt: `
		  CREATE CONSTRAINT TRIGGER trigger1
		  AFTER INSERT OR DELETE ON table1
		  DEFERRABLE INITIALLY DEFERRED
		  FOR EACH ROW
		`,
		Expected: &Trigger{
			Table:      "table1",
			Name:       "trigger1",
			Constraint: true,
			Timing:     "INITIALLY DEFERRED",
			Called:     "AFTER",
			Events: []*TriggerEvent{
				{Event: "INSERT"},
				{Event: "DELETE"},
			},
			ForEach: "ROW",
		},
	}, {Stmt: `
		  CREATE TRIGGER trigger2
		  BEFORE TRUNCATE ON table2
		  FOR EACH STATEMENT
		`,
		Expected: &Trigger{
			Table:  "table2",
			Name:   "trigger2",
			Called: "BEFORE",
			Events: []*TriggerEvent{
				{Event: "TRUNCATE"},
			},
			ForEach: "STATEMENT",
		},
	}, {Stmt: `
		  create  trigger  TRIGGER3
		  instead  of  update  on  VIEW3
		`,
		Expected: &Trigger{
			Table:  "VIEW3",
			Name:   "TRIGGER3",
			Called: "INSTEAD OF",
			Events: []*TriggerEvent{
				{Event: "UPDATE"},
			},
		},
	}, {Stmt: `
		  Create Constraint Trigger
		  Trigger4
		  After Update On
		  Table4
		  NOT	Deferrable
		  For	Each	Statement
		`,
		Expected: &Trigger{
			Table:      "Table4",
			Name:       "Trigger4",
			Constraint: true,
			Timing:     "NOT DEFERRABLE",
			Called:     "AFTER",
			Events: []*TriggerEvent{
				{Event: "UPDATE"},
			},
			ForEach: "STATEMENT",
		},
	}, {Stmt: `
		  CREATE CONSTRAINT TRIGGER trigger5
		  AFTER INSERT OR UPDATE OR DELETE ON table5
		  INITIALLY IMMEDIATE
		  FOR ROW
		`,
		Expected: &Trigger{
			Table:      "table5",
			Name:       "trigger5",
			Constraint: true,
			Timing:     "INITIALLY IMMEDIATE",
			Called:     "AFTER",
			Events: []*TriggerEvent{
				{Event: "INSERT"},
				{Event: "UPDATE"},
				{Event: "DELETE"},
			},
			ForEach: "ROW",
		},
	}} {
		actual, err := ParseTrigger(tc.Stmt)
		require.NoError(err)
		require.Equal(tc.Expected, actual)
	}
}

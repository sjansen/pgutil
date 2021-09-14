package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v2"
	"github.com/sjansen/pgutil/internal/ddl"
)

var statements = []string{`
CREATE TRIGGER update_foo_modified
  BEFORE UPDATE ON bar
  FOR EACH ROW
  EXECUTE PROCEDURE update_modified_column()
`, `
CREATE TRIGGER trigger2
  BEFORE TRUNCATE ON table2
  FOR EACH STATEMENT
  EXECUTE PROCEDURE fn2()
`, `
CREATE TRIGGER log_update
  AFTER UPDATE ON accounts
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE FUNCTION log_account_update()
`, `
CREATE CONSTRAINT TRIGGER trigger1
  AFTER INSERT OR DELETE ON table1
  DEFERRABLE INITIALLY DEFERRED
  FOR EACH ROW
  EXECUTE FUNCTION fn1()
`,
}

func main() {
	for _, stmt := range statements {
		results, err := pg_query.Parse(stmt)
		if err != nil {
			panic(err)
		}

		fmt.Println(stmt)
		for _, result := range results.Stmts {
			// fmt.Println(result.String())
			stmt := result.Stmt.Node.(*pg_query.Node_CreateTrigStmt).CreateTrigStmt
			// fmt.Printf("%#v\n", stmt)

			trigger := &ddl.Trigger{
				Schema: stmt.Relation.Schemaname,
				Table:  stmt.Relation.Relname,
				Name:   stmt.Trigname,

				Function: stmt.Funcname[0].Node.(*pg_query.Node_String_).String_.Str,

				Constraint:        stmt.Isconstraint,
				Deferrable:        stmt.Deferrable,
				InitiallyDeferred: stmt.Initdeferred,
				ForEachRow:        stmt.Row,

				Delete:   stmt.Events&8 != 0,
				Insert:   stmt.Events&4 != 0,
				Truncate: stmt.Events&32 != 0,
				Update:   stmt.Events&16 != 0,
			}
			if stmt.Timing == 0 {
				trigger.Timing = "AFTER"
			} else {
				trigger.Timing = "BEFORE"
			}
			formatted, err := trigger.ToSQL()
			if err != nil {
				panic(err)
			}
			fmt.Println(formatted)
			// fmt.Printf("Trigger: %#v\n", trigger)
			fmt.Println("----")
		}
	}
}

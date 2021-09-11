package main

import (
	"fmt"

	pg_query "github.com/pganalyze/pg_query_go/v2"
)

var statements = []string{`
    CREATE TRIGGER update_foo_modified
    BEFORE UPDATE ON foo
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
			fmt.Println(result.String())
			fmt.Println("----")
			stmt := result.Stmt.Node.(*pg_query.Node_CreateTrigStmt)
			fmt.Printf("%#v\n", stmt.CreateTrigStmt)
		}

		fmt.Println("----")

		stmt, err := pg_query.Deparse(results)
		if err != nil {
			panic(err)
		}
		fmt.Println(stmt)
	}
}

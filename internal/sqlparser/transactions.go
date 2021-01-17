package sqlparser

import "github.com/sjansen/pgutil/internal/sql"

func newBeginStmt(opts []*option) *sql.BeginStmt {
	stmt := &sql.BeginStmt{}
	for _, opt := range opts {
		switch opt.Name {
		case "deferrable":
			stmt.Deferrable = opt.Value.(bool)
		case "isolation_level":
			stmt.IsolationLevel = opt.Value.(string)
		case "read_only":
			stmt.ReadOnly = opt.Value.(bool)
		}
	}
	return stmt
}

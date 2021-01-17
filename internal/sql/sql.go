package sql

// BeginStmt describes a BEGIN statement.
type BeginStmt struct {
	Deferrable     bool
	ReadOnly       bool
	IsolationLevel string
}

// CommitStmt describes a COMMIT statement.
type CommitStmt struct{}

// RollbackStmt describes a ROLLBACK statement.
type RollbackStmt struct{}

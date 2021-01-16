//go:generate goyacc -o parser.go parser.go.y

package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		lexer := &Lexer{
			buf: []byte(arg),
		}
		// yyDebug = 10
		yyErrorVerbose = true
		result := yyParse(lexer)
		if result == 0 {
			for _, line := range lexer.Statement.Lines() {
				fmt.Println(" ", line)
			}
		}
	}
}

type Statement interface {
	Lines() []string
}

type Column struct {
	Alias string
	Name  string
}

type SelectStmt struct {
	Table   string
	Columns []*Column
}

func (stmt *SelectStmt) Lines() []string {
	lines := make([]string, 0, len(stmt.Columns)+3)
	lines = append(lines, "SELECT")
	prefix := "  "
	for _, col := range stmt.Columns {
		if col.Alias == "" {
			lines = append(lines, prefix+col.Name)
		} else {
			lines = append(lines, prefix+col.Name+" AS "+col.Alias)
		}
		prefix = ", "
	}
	lines = append(lines, "FROM "+stmt.Table, ";")
	return lines
}

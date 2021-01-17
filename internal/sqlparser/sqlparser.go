//go:generate goyacc -o generated.go grammar.y

package sqlparser

import "errors"

func init() {
	yyErrorVerbose = true
}

// Statement is a SQL command or fragment.
type Statement interface{}

// Parse parses a SQL statement.
func Parse(buf []byte) (Statement, error) {
	lexer := &Lexer{
		buf: buf,
	}
	if rc := yyParse(lexer); rc == 0 {
		return lexer.Statement, nil
	}
	return nil, errors.New(lexer.err)
}

// EnableDebugLogging enables debug logging to stderr.
func EnableDebugLogging() {
	yyDebug = 10
}

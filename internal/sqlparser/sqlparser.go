//go:generate goyacc -o generated.go grammar.y

package sqlparser

import (
	"errors"

	"github.com/sjansen/pgutil/internal/schema"
)

func init() {
	yyErrorVerbose = true
}

// Statement is a SQL command.
type Statement interface{}

// EnableDebugLogging enables debug logging to stderr.
func EnableDebugLogging() {
	yyDebug = 10
}

// Parse parses a SQL statement.
func Parse(buf []byte) (Statement, error) {
	return parse(buf, 0)
}

// ParseForeignKey parses a foreign key declaration.
func ParseForeignKey(buf []byte) (*schema.ForeignKey, error) {
	tmp, err := parse(buf, MODE_FOREIGN_KEY)
	if err != nil {
		return nil, err
	}
	return tmp.(*schema.ForeignKey), err
}

// ParseTrigger parses a trigger declaration.
func ParseTrigger(buf []byte) (*schema.Trigger, error) {
	tmp, err := parse(buf, MODE_TRIGGER)
	if err != nil {
		return nil, err
	}
	return tmp.(*schema.Trigger), err
}

func parse(buf []byte, mode int) (interface{}, error) {
	lexer := &Lexer{
		buf:  buf,
		mode: mode,
	}
	rc := yyParse(lexer)
	if rc != 0 {
		return nil, errors.New(lexer.err)
	}
	return lexer.result, nil
}

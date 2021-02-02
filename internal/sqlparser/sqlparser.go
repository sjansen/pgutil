//go:generate goyacc -o generated.go grammar.y

package sqlparser

import (
	"errors"

	"github.com/sjansen/pgutil/internal/ddl"
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
func Parse(str string) (Statement, error) {
	return parse(str, 0)
}

// ParseCheck parses a CHECK declaration.
func ParseCheck(str string) (*ddl.Check, error) {
	tmp, err := parse(str, MODE_CHECK)
	if err != nil {
		return nil, err
	}
	return tmp.(*ddl.Check), err
}

// ParseCreateIndex parses a CREATE INDEX statement.
func ParseCreateIndex(str string) (*ddl.Index, error) {
	tmp, err := parse(str, MODE_CREATE_INDEX)
	if err != nil {
		return nil, err
	}
	return tmp.(*ddl.Index), err
}

// ParseCreateTrigger parses a CREATE TRIGGER statement.
func ParseCreateTrigger(str string) (*ddl.Trigger, error) {
	tmp, err := parse(str, MODE_CREATE_TRIGGER)
	if err != nil {
		return nil, err
	}
	return tmp.(*ddl.Trigger), err
}

// ParseForeignKey parses a FOREIGN KEY declaration.
func ParseForeignKey(str string) (*ddl.ForeignKey, error) {
	tmp, err := parse(str, MODE_FOREIGN_KEY)
	if err != nil {
		return nil, err
	}
	return tmp.(*ddl.ForeignKey), err
}

func parse(str string, mode int) (interface{}, error) {
	lexer := &lexer{
		str:  str,
		mode: mode,
	}
	rc := yyParse(lexer)
	if rc != 0 {
		return nil, errors.New(lexer.err)
	}
	return lexer.result, nil
}

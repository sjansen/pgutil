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
func Parse(str string) (Statement, error) {
	return parse(str, 0)
}

// ParseCheck parses a CHECK declaration.
func ParseCheck(str string) (*schema.Check, error) {
	tmp, err := parse(str, MODE_CHECK)
	if err != nil {
		return nil, err
	}
	return tmp.(*schema.Check), err
}

// ParseCreateTrigger parses a CREATE TRIGGER statement.
func ParseCreateTrigger(str string) (*schema.Trigger, error) {
	tmp, err := parse(str, MODE_CREATE_TRIGGER)
	if err != nil {
		return nil, err
	}
	return tmp.(*schema.Trigger), err
}

// ParseForeignKey parses a FOREIGN KEY declaration.
func ParseForeignKey(str string) (*schema.ForeignKey, error) {
	tmp, err := parse(str, MODE_FOREIGN_KEY)
	if err != nil {
		return nil, err
	}
	return tmp.(*schema.ForeignKey), err
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

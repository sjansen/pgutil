// +build integration

// nolint: govet
package ddl

import (
	"testing"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/stretchr/testify/require"
)

var (
	triggerLexer = lexer.Must(lexer.Regexp(`(\s+)` +
		`|(?P<Keyword>(?i)` +
		`AFTER|BEFORE|CONSTRAINT|CREATE|DELETE|EACH|EXECUTE` +
		`|FOR|INSERT|INSTEAD|OF|ON|OR|PROCEDURE|ROW|STATEMENT` +
		`|TRIGGER|TRUNCATE|UPDATE` +
		`)\b` +
		`|(?P<Ident>[a-zA-Z_][a-zA-Z0-9_]*)` +
		`|(?P<Number>[-+]?\d*\.?\d+([eE][-+]?\d+)?)` +
		`|(?P<String>'[^']*')` +
		`|(?P<Operators><>|!=|<=|>=|[-+*/%,.()=<>])`,
	))
	triggerParser = participle.MustBuild(
		&createTrigger{},
		participle.Lexer(triggerLexer),
		participle.CaseInsensitive("Keyword"),
		participle.Unquote("String"),
		participle.Upper("Keyword"),
	)
)

type createTrigger struct {
	Constraint bool            `"CREATE" @"CONSTRAINT"? "TRIGGER"`
	Name       string          `@Ident`
	When       string          `@("BEFORE"|"AFTER"|"INSTEAD" "OF")`
	Events     []*triggerEvent `@@ ("OR" @@)*`
	Table      string          `"ON" @Ident`
	ForEach    string          `("FOR" "EACH"? @("ROW"|"STATEMENT"))?`
	Function   string          `"EXECUTE" "PROCEDURE" @Ident "(" ")"`
}

type triggerEvent struct {
	Event   string   `(@"DELETE"|@"INSERT"|@"TRUNCATE"|@"UPDATE"`
	Columns []string ` ("OF" @Ident ("," @Ident)*)*)`
}

func TestTriggerParser(t *testing.T) {
	require := require.New(t)

	actual := &createTrigger{}
	err := triggerParser.ParseString(
		`CREATE TRIGGER update_foo_modified
		  BEFORE UPDATE ON foo
		  FOR EACH ROW
		  EXECUTE PROCEDURE update_modified_column()
		`,
		actual,
	)
	require.NoError(err)

	expected := &createTrigger{
		Name: "update_foo_modified",
		When: "BEFORE",
		Events: []*triggerEvent{
			{Event: "UPDATE"},
		},
		Table:    "foo",
		ForEach:  "ROW",
		Function: "update_modified_column",
	}
	require.Equal(expected, actual)
}

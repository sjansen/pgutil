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
		`AFTER|BEFORE|CONSTRAINT|CREATE` +
		`|DEFERRABLE|DEFERRED|DELETE` +
		`|EACH|EXECUTE\s+PROCEDURE|FOR|FROM` +
		`|IMMEDIATE|INITIALLY|INSERT|INSTEAD\s+OF` +
		`|NOT|ON|OR|ROW|STATEMENT` +
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
		participle.Map(CollapseWhitespace, "Keyword"),
		participle.Unquote("String"),
		participle.Upper("Keyword"),
	)
)

func CollapseWhitespace(token lexer.Token) (lexer.Token, error) {
	token.Value = whitespace.ReplaceAllString(token.Value, " ")
	return token, nil
}

type createTrigger struct {
	Constraint bool            `"CREATE" @"CONSTRAINT"? "TRIGGER"`
	Name       string          `@Ident`
	When       string          `@("BEFORE"|"AFTER"|"INSTEAD OF")`
	Events     []*triggerEvent `@@ ("OR" @@)*`
	Table      string          `"ON" @Ident`
	From       string          `("FROM" @Ident)?`
	Deferrable bool            `("NOT" "DEFERRABLE"|@"DEFERRABLE"`
	Deferred   bool            ` ("INITIALLY" "IMMEDIATE"|"INITIALLY" @"DEFERRED")?)?`
	ForEach    string          `("FOR" "EACH"? @("ROW"|"STATEMENT"))?`
	Function   string          `"EXECUTE PROCEDURE" @Ident "(" ")"`
}

type triggerEvent struct {
	Event   string   `(@"DELETE"|@"INSERT"|@"TRUNCATE"|@"UPDATE"`
	Columns []string ` ("OF" @Ident ("," @Ident)*)*)`
}

func TestTriggerParser(t *testing.T) {
	require := require.New(t)

	for _, tc := range []struct {
		Stmt     string
		Expected *createTrigger
	}{{Stmt: `
		  CREATE TRIGGER update_foo_modified
		  BEFORE UPDATE ON foo
		  FOR EACH ROW
		  EXECUTE PROCEDURE update_modified_column()
		`,
		Expected: &createTrigger{
			Name: "update_foo_modified",
			When: "BEFORE",
			Events: []*triggerEvent{
				{Event: "UPDATE"},
			},
			Table:    "foo",
			ForEach:  "ROW",
			Function: "update_modified_column",
		},
	}, {Stmt: `
		  CREATE TRIGGER trigger1
		  AFTER DELETE ON table1
		  EXECUTE PROCEDURE fn1()
		`,
		Expected: &createTrigger{
			Name: "trigger1",
			When: "AFTER",
			Events: []*triggerEvent{
				{Event: "DELETE"},
			},
			Table:    "table1",
			Function: "fn1",
		},
	}, {Stmt: `
		  CREATE CONSTRAINT TRIGGER trigger2
		  INSTEAD OF INSERT OR TRUNCATE ON table2
		  DEFERRABLE INITIALLY DEFERRED
		  FOR EACH STATEMENT
		  EXECUTE PROCEDURE fn2()
		`,
		Expected: &createTrigger{
			Constraint: true,
			Name:       "trigger2",
			When:       "INSTEAD OF",
			Events: []*triggerEvent{
				{Event: "INSERT"},
				{Event: "TRUNCATE"},
			},
			Table:      "table2",
			Deferrable: true,
			Deferred:   true,
			ForEach:    "STATEMENT",
			Function:   "fn2",
		},
	}, {Stmt: `
		  create  constraint  trigger  TRIGGER3
		  instead  of  insert  or  truncate  on  TABLE3
		  deferrable  initially  deferred
		  for  each  statement
		  execute  procedure  FN3()
		`,
		Expected: &createTrigger{
			Constraint: true,
			Name:       "TRIGGER3",
			When:       "INSTEAD OF",
			Events: []*triggerEvent{
				{Event: "INSERT"},
				{Event: "TRUNCATE"},
			},
			Table:      "TABLE3",
			Deferrable: true,
			Deferred:   true,
			ForEach:    "STATEMENT",
			Function:   "FN3",
		},
	},
	} {
		actual := &createTrigger{}
		err := triggerParser.ParseString(
			tc.Stmt,
			actual,
		)
		require.NoError(err, tc.Stmt)
		require.Equal(tc.Expected, actual, tc.Stmt)
	}
}

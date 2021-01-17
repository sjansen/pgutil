%{
package main
%}

%union {
  empty struct{}
  str   string
  node  interface{}
}

%token LEX_ERROR
%token<empty> ',' ';'
%token<empty> AS FROM SELECT
%token<str> Identifier

%type<node> select_stmt columns column

%start statement

%%

statement:
  select_stmt semicolon_opt

semicolon_opt:
/*empty*/
| ';'

select_stmt:
  SELECT columns FROM Identifier
  {
    stmt := &SelectStmt{
      Table: $4,
      Columns: $2.([]*Column),
    }
    yylex.(*Lexer).Statement = stmt
  }

column:
  Identifier {
    $$ = &Column{
      Name: $1,
    }
  }
| Identifier AS Identifier {
    $$ = &Column{
      Alias: $3,
      Name: $1,
    }
  }

columns:
  column {
    slice := make([]*Column, 0, 4)
    $$ = append(slice, $1.(*Column))
  }
| columns ',' column {
    slice := $1.([]*Column)
    $$ = append(slice, $3.(*Column))
  }

%{
package sqlparser

import (
  "github.com/sjansen/pgutil/internal/sql"
)

type option struct {
  Name  string
  Value interface{}
}

func newOption(name string, value interface{}) *option {
  return &option{
    Name:  name,
    Value: value,
  }
}

func newOptionList(opts ...*option) []*option {
  slice := make([]*option, 0, 4)
  slice = append(slice, opts...)
  return slice
}
%}

/*
 * NOTES
 *	  CAPITALS are used to represent terminal symbols.
 *	  non-capitals are used to represent non-terminals.
 */

%union {
  ast   interface{}
  bool  bool
  opt   *option
  opts  []*option
  str   string
  strs  []string
}

/* keywords */
%token ABORT ABSOLUTE ACCESS ACTION ADD ADMIN AFTER AGGREGATE ALL
  ALSO ALTER ALWAYS ANALYSE ANALYZE AND ANY ARRAY AS ASC
	ASSERTION ASSIGNMENT ASYMMETRIC AT ATTACH ATTRIBUTE
	AUTHORIZATION

	BACKWARD BEFORE BEGIN BETWEEN BIGINT BINARY BIT BOOLEAN BOTH BY

	CACHE CALL CALLED CASCADE CASCADED CASE CAST CATALOG CHAIN
	CHAR CHARACTER CHARACTERISTICS CHECK CHECKPOINT CLASS CLOSE
	CLUSTER COALESCE COLLATE COLLATION COLUMN COLUMNS COMMENT
	COMMENTS COMMIT COMMITTED CONCURRENTLY CONFIGURATION CONFLICT
	CONNECTION CONSTRAINT CONSTRAINTS CONTENT CONTINUE CONVERSION
	COPY COST CREATE CROSS CSV CUBE CURRENT CURRENT_CATALOG
	CURRENT_DATE CURRENT_ROLE CURRENT_SCHEMA CURRENT_TIME
	CURRENT_TIMESTAMP CURRENT_USER CURSOR CYCLE

	DATA DATABASE DAY DEALLOCATE DEC DECIMAL DECLARE DEFAULT
	DEFAULTS DEFERRABLE DEFERRED DEFINER DELETE DELIMITER
	DELIMITERS DEPENDS DESC DETACH DICTIONARY DISABLE DISCARD
	DISTINCT DO DOCUMENT DOMAIN DOUBLE DROP

	EACH ELSE ENABLE ENCODING ENCRYPTED END ENUM ESCAPE EVENT
	EXCEPT EXCLUDE EXCLUDING EXCLUSIVE EXECUTE EXISTS EXPLAIN
	EXPRESSION EXTENSION EXTERNAL EXTRACT

	FALSE FAMILY FETCH FILTER FIRST FLOAT FOLLOWING FOR FORCE
	FOREIGN FORWARD FREEZE FROM FULL FUNCTION FUNCTIONS

	GENERATED GLOBAL GRANT GRANTED GREATEST GROUP GROUPING GROUPS

	HANDLER HAVING HEADER HOLD HOUR

	IDENTITY IF ILIKE IMMEDIATE IMMUTABLE IMPLICIT IMPORT IN
	INCLUDE INCLUDING INCREMENT INDEX INDEXES INHERIT INHERITS
	INITIALLY INLINE INNER INOUT INPUT INSENSITIVE INSERT INSTEAD
	INT INTEGER INTERSECT INTERVAL INTO INVOKER IS ISNULL
	ISOLATION

	JOIN

	KEY

	LABEL LANGUAGE LARGE LAST LATERAL LEADING LEAKPROOF LEAST
	LEFT LEVEL LIKE LIMIT LISTEN LOAD LOCAL LOCALTIME LOCALTIMESTAMP
	LOCATION LOCK LOCKED LOGGED

	MAPPING MATCH MATERIALIZED MAXVALUE METHOD MINUTE MINVALUE
	MODE MONTH MOVE

	NAME NAMES NATIONAL NATURAL NCHAR NEW NEXT NFC NFD NFKC
	NFKD NO NONE NORMALIZE NORMALIZED NOT NOTHING NOTIFY NOTNULL
	NOWAIT NULL NULLIF NULLS NUMERIC

	OBJECT OF OFF OFFSET OIDS OLD ON ONLY OPERATOR OPTION OPTIONS
	OR ORDER ORDINALITY OTHERS OUT OUTER OVER OVERLAPS OVERLAY
	OVERRIDING OWNED OWNER

	PARALLEL PARSER PARTIAL PARTITION PASSING PASSWORD PLACING
	PLANS POLICY POSITION PRECEDING PRECISION PRESERVE PREPARE
	PREPARED PRIMARY PRIOR PRIVILEGES PROCEDURAL PROCEDURE
	PROCEDURES PROGRAM PUBLICATION

	QUOTE

	RANGE READ REAL REASSIGN RECHECK RECURSIVE REF REFERENCES
	REFERENCING REFRESH REINDEX RELATIVE RELEASE RENAME REPEATABLE
	REPLACE REPLICA RESET RESTART RESTRICT RETURNING RETURNS
	REVOKE RIGHT ROLE ROLLBACK ROLLUP ROUTINE ROUTINES ROW ROWS
	RULE

	SAVEPOINT SCHEMA SCHEMAS SCROLL SEARCH SECOND SECURITY
	SELECT SEQUENCE SEQUENCES SERIALIZABLE SERVER SESSION
	SESSION_USER SET SETS SETOF SHARE SHOW SIMILAR SIMPLE SKIP
	SMALLINT SNAPSHOT SOME SQL STABLE STANDALONE START STATEMENT
	STATISTICS STDIN STDOUT STORAGE STORED STRICT STRIP
	SUBSCRIPTION SUBSTRING SUPPORT SYMMETRIC SYSID SYSTEM

	TABLE TABLES TABLESAMPLE TABLESPACE TEMP TEMPLATE TEMPORARY
	TEXT THEN TIES TIME TIMESTAMP TO TRAILING TRANSACTION
	TRANSFORM TREAT TRIGGER TRIM TRUE TRUNCATE TRUSTED TYPE
	TYPES

	UESCAPE UNBOUNDED UNCOMMITTED UNENCRYPTED UNION UNIQUE
	UNKNOWN UNLISTEN UNLOGGED UNTIL UPDATE USER USING

	VACUUM VALID VALIDATE VALIDATOR VALUE VALUES VARCHAR VARIADIC
	VARYING VERBOSE VERSION VIEW VIEWS VOLATILE

	WHEN WHERE WHITESPACE WINDOW WITH WITHIN WITHOUT WORK WRAPPER
	WRITE

	XML XMLATTRIBUTES XMLCONCAT XMLELEMENT XMLEXISTS XMLFOREST
	XMLNAMESPACES XMLPARSE XMLPI XMLROOT XMLSERIALIZE XMLTABLE

	YEAR YES

	ZONE

%token MODE_FOREIGN_KEY

%token UNEXPECTED_SYMBOL
%token<str> Identifier Name

%type<ast>  ast
%type<strs> column_list column_list_or_empty
%type<bool> deferrable
%type<ast>  foreign_key_decl
%type<str>  foreign_key_action foreign_key_match
%type<opts> foreign_key_actions
%type<opt>  foreign_key_delete foreign_key_update
%type<bool> initially_deferred
%type<bool> transaction_chain
%type<str>  transaction_isolation_level
%type<opt>  transaction_mode_item
%type<opts> transaction_mode_list transaction_mode_list_or_empty
%type<ast>  transaction_stmt

%start ast

%%

ast:
  MODE_FOREIGN_KEY foreign_key_decl { yylex.(*Lexer).result = $2 }
| transaction_stmt semicolon_opt    { yylex.(*Lexer).result = $1 }

semicolon_opt:
/* empty */
| ';'

column_list:
  Identifier {
    slice := make([]string, 0, 4)
    $$ = append(slice, $1)
  }
| column_list ',' Identifier {
  slice := $1
  $$ = append(slice, $3)
}

column_list_or_empty:
/* empty */           { $$ = nil }
| '(' column_list ')' { $$ = $2 }

deferrable:
/* empty */      { $$ = false }
| DEFERRABLE     { $$ = true }
| NOT DEFERRABLE { $$ = false }

initially_deferred:
/* empty */           { $$ = false }
| INITIALLY DEFERRED  { $$ = true }
| INITIALLY IMMEDIATE { $$ = false }

/*****************************************************************************
 *
 *	foreign key declaration
 *
 *****************************************************************************/

/* TODO: Identifier should be qualified_name */
foreign_key_decl:
  FOREIGN KEY '(' column_list ')' REFERENCES Identifier column_list_or_empty
  foreign_key_match foreign_key_actions deferrable initially_deferred {
    fk := newForeignKey($10...)
    fk.Table = $7
    fk.Columns = $4
    fk.Referenced = $8
    fk.Match = $9
    fk.Deferrable = $11
    fk.InitiallyDeferred = $12
    $$ = fk
  }

foreign_key_match:
/* empty */     { $$ = "" }
| MATCH FULL    { $$ = "FULL" }
| MATCH PARTIAL { $$ = "PARTIAL" }
| MATCH SIMPLE  { $$ = "SIMPLE" }

foreign_key_action:
  NO ACTION   { $$ = "NO ACTION" }
| RESTRICT    { $$ = "RESTRICT" }
| CASCADE     { $$ = "CASCADE" }
| SET NULL    { $$ = "SET NULL" }
| SET DEFAULT { $$ = "SET DEFAULT" }

foreign_key_actions:
/* empty */                             { $$ = nil }
| foreign_key_delete                    { $$ = []*option{$1} }
| foreign_key_update                    { $$ = []*option{$1} }
| foreign_key_delete foreign_key_update { $$ = []*option{$1, $2} }
| foreign_key_update foreign_key_delete { $$ = []*option{$1, $2} }

foreign_key_delete:
  ON DELETE foreign_key_action { $$ = newOption("on_delete", $3) }

foreign_key_update:
  ON UPDATE foreign_key_action { $$ = newOption("on_update", $3) }

/*****************************************************************************
 *
 *	transaction management
 *
 *****************************************************************************/

transaction_chain:
/* empty */    { $$ = false }
| AND CHAIN    { $$ = true }
| AND NO CHAIN { $$ = false }

transaction_isolation_level:
  READ COMMITTED		{ $$ = "read committed" }
| READ UNCOMMITTED	{ $$ = "read uncommitted" }
| REPEATABLE READ		{ $$ = "repeatable read" }
| SERIALIZABLE			{ $$ = "serializable" }

transaction_keywords:
/* empty */
| WORK
| TRANSACTION

transaction_mode_item:
  ISOLATION LEVEL transaction_isolation_level {
		$$ = newOption("isolation_level", $3)
  }
| DEFERRABLE     { $$ = newOption("deferrable", true) }
| NOT DEFERRABLE { $$ = newOption("deferrable", false) }
| READ ONLY      { $$ = newOption("read_only", true) }
| READ WRITE     { $$ = newOption("read_only", false) }

transaction_mode_list:
  transaction_mode_item { $$ = newOptionList($1) }
| transaction_mode_list transaction_mode_item     { $$ = append($1, $2) }
| transaction_mode_list ',' transaction_mode_item { $$ = append($1, $3) }

transaction_mode_list_or_empty:
/* empty */             { $$ = nil }
| transaction_mode_list { $$ = $1 }

transaction_stmt:
  ABORT transaction_keywords transaction_chain {
    $$ = &sql.RollbackStmt{Chain: $3}
    if yyDebug > 6 {
      __yyfmt__.Printf("stmt = %#v\n", $$)
    }
  }
| BEGIN transaction_keywords transaction_mode_list_or_empty {
    $$ = newBeginStmt($3)
    if yyDebug > 6 {
      __yyfmt__.Printf("stmt = %#v\n", $$)
    }
  }
| COMMIT transaction_keywords transaction_chain {
    $$ = &sql.CommitStmt{Chain: $3}
    if yyDebug > 6 {
      __yyfmt__.Printf("stmt = %#v\n", $$)
    }
  }
| END transaction_keywords transaction_chain {
    $$ = &sql.CommitStmt{Chain: $3}
    if yyDebug > 6 {
      __yyfmt__.Printf("stmt = %#v\n", $$)
    }
  }
| ROLLBACK transaction_keywords transaction_chain {
    $$ = &sql.RollbackStmt{Chain: $3}
    if yyDebug > 6 {
      __yyfmt__.Printf("stmt = %#v\n", $$)
    }
  }
| START TRANSACTION transaction_mode_list_or_empty {
    $$ = newBeginStmt($3)
    if yyDebug > 6 {
      __yyfmt__.Printf("stmt = %#v\n", $$)
    }
  }

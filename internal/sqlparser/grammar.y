/*-------------------------------------------------------------------------
 * Portions Copyright (c) 1996-2021, PostgreSQL Global Development Group
 * Portions Copyright (c) 1994, Regents of the University of California
 *-------------------------------------------------------------------------
 */
%{
package sqlparser

import (
  "strings"

  "github.com/sjansen/pgutil/internal/ddl"
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
  asts  interface{}
  bool  bool
  opt   *option
  opts  []*option
  str   string
  strs  []string
}

/* non-keyword tokens */
%token BCONST FCONST ICONST SCONST XCONST
%token PARAM
%token COLON_EQUALS EQUALS_GREATER

/* keywords */
%token<str>
  ABORT ABSOLUTE ACCESS ACTION ADD ADMIN AFTER AGGREGATE ALL
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

/*
 * NOT_LA exists so that productions such as NOT LIKE can be given the same
 * precedence as LIKE; otherwise they'd effectively have the same precedence
 * as NOT, at least with respect to their left-hand subexpression.
 * NULLS_LA and WITH_LA are needed to make the grammar LALR(1).
 */
%token NOT_LA NULLS_LA WITH_LA

%token MODE_CHECK
%token MODE_CREATE_INDEX
%token MODE_CREATE_TRIGGER
%token MODE_FOREIGN_KEY
%token MODE_NOT_IMPLEMENTED

%token UNEXPECTED_SYMBOL
%token<str> Identifier Name

/* Precedence: lowest to highest */
%nonassoc	SET				/* see relation_expr_opt_alias */
%left		UNION EXCEPT
%left		INTERSECT
%left		OR
%left		AND
%right		NOT
%nonassoc	IS ISNULL NOTNULL	/* IS sets precedence for IS NULL, etc */
%nonassoc	'<' '>' '=' LESS_EQUALS GREATER_EQUALS NOT_EQUALS
%nonassoc	BETWEEN IN LIKE ILIKE SIMILAR NOT_LA
%nonassoc	ESCAPE			/* ESCAPE must be just above LIKE/ILIKE/SIMILAR */
/*
 * To support target_el without AS, it used to be necessary to assign Identifier an
 * explicit precedence just less than Op.  While that's not really necessary
 * since we removed postfix operators, it's still helpful to do so because
 * there are some other unreserved keywords that need precedence assignments.
 * If those keywords have the same precedence as Identifier then they clearly act
 * the same as non-keywords, reducing the risk of unwanted precedence effects.
 *
 * We need to do this for PARTITION, RANGE, ROWS, and GROUPS to support
 * opt_existing_window_name (see comment there).
 *
 * The frame_bound productions UNBOUNDED PRECEDING and UNBOUNDED FOLLOWING
 * are even messier: since UNBOUNDED is an unreserved keyword (per spec!),
 * there is no principled way to distinguish these from the productions
 * a_expr PRECEDING/FOLLOWING.  We hack this up by giving UNBOUNDED slightly
 * lower precedence than PRECEDING and FOLLOWING.  At present this doesn't
 * appear to cause UNBOUNDED to be treated differently from other unreserved
 * keywords anywhere else in the grammar, but it's definitely risky.  We can
 * blame any funny behavior of UNBOUNDED on the SQL standard, though.
 *
 * To support CUBE and ROLLUP in GROUP BY without reserving them, we give them
 * an explicit priority lower than '(', so that a rule with CUBE '(' will shift
 * rather than reducing a conflicting rule that takes CUBE as a function name.
 * Using the same precedence as Identifier seems right for the reasons given above.
 */
%nonassoc	UNBOUNDED		/* ideally would have same precedence as Identifier */
%nonassoc	Identifier PARTITION RANGE ROWS GROUPS PRECEDING FOLLOWING CUBE ROLLUP
%left		Op OPERATOR		/* multi-character ops and user-defined operators */
%left		'+' '-'
%left		'*' '/' '%'
%left		'^'
/* Unary Operators */
%left		AT				/* sets precedence for AT TIME ZONE */
%left		COLLATE
%right		UMINUS
%left		'[' ']'
%left		'(' ')'
%left		TYPECAST
%left		'.'
/*
 * These might seem to be low-precedence, but actually they are not part
 * of the arithmetic hierarchy at all in their use as JOIN operators.
 * We make them high-precedence to support their use as function names.
 * They wouldn't be given a precedence at all, were it not that we need
 * left-associativity among the JOIN rules themselves.
 */
%left		JOIN CROSS LEFT FULL RIGHT INNER NATURAL

%type<ast>  ast
%type<str>  a_expr_str
%type<strs> column_list column_list_or_empty
%type<bool> deferrable
%type<ast>  check_decl
%type<bool> concurrently_opt
%type<ast>  create_index_stmt create_trigger_stmt
%type<bool> desc_opt
%type<ast>  foreign_key_decl
%type<str>  foreign_key_action foreign_key_match
%type<opts> foreign_key_actions
%type<opt>  foreign_key_delete foreign_key_update
%type<bool> initially_deferred
%type<ast>  index_key
%type<asts> index_keys
%type<str>  index_opclass_opt index_using
%type<str>  name name_opt
%type<bool> transaction_chain
%type<str>  transaction_isolation_level
%type<opt>  transaction_mode_item
%type<opts> transaction_mode_list transaction_mode_list_or_empty
%type<ast>  transaction_stmt
%type<opt>  trigger_event_item
%type<opts> trigger_event_list
%type<bool> trigger_for
%type<str>  trigger_from trigger_timing trigger_when
%type<bool> unique_opt
%type<str>  where_expr_opt

%type<str> col_name_keyword
%type<str> reserved_keyword
%type<str> unreserved_keyword

%start ast

%%

ast:
  MODE_CHECK check_decl                    { yylex.(*lexer).result = $2 }
| MODE_CREATE_INDEX create_index_stmt      { yylex.(*lexer).result = $2 }
| MODE_CREATE_TRIGGER create_trigger_stmt  { yylex.(*lexer).result = $2 }
| MODE_FOREIGN_KEY foreign_key_decl        { yylex.(*lexer).result = $2 }
| MODE_NOT_IMPLEMENTED select_stmt         { /* not implemented */ }
| create_index_stmt semicolon_opt          { yylex.(*lexer).result = $1 }
| create_trigger_stmt semicolon_opt        { yylex.(*lexer).result = $1 }
| transaction_stmt semicolon_opt           { yylex.(*lexer).result = $1 }

/* TODO: require end of input */
semicolon_opt:
/* empty */
| ';'

a_expr_begin:
/* empty */ {
  yylex.(*lexer).setMark()
}

a_expr_str:
  a_expr_begin a_expr {
    $$ = strings.TrimSpace(
      yylex.(*lexer).sinceMark(),
    )
  }

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

concurrently_opt:
/* empty */    { $$ = false }
| CONCURRENTLY { $$ = true }

/* TODO: combine with initially_deferred */
deferrable:
/* empty */      { $$ = false }
| DEFERRABLE     { $$ = true }
| NOT DEFERRABLE { $$ = false }

desc_opt:
/* empty */ { $$ = false }
| ASC       { $$ = false }
| DESC      { $$ = true }

function_or_procedure:
  FUNCTION
|	PROCEDURE

name_opt:
/* empty */ { $$ = "" }
| name      { $$ = $1 }

initially_deferred:
/* empty */           { $$ = false }
| INITIALLY DEFERRED  { $$ = true }
| INITIALLY IMMEDIATE { $$ = false }

name:
  Identifier         { $$ = $1 }
| unreserved_keyword { $$ = $1 }
| col_name_keyword   { $$ = $1 }

unique_opt:
/* empty */ { $$ = false }
| UNIQUE    { $$ = true }

where_expr_opt:
/* empty */        { $$ = "" }
| WHERE a_expr_str { $$ = $2 }

/*****************************************************************************
 *
 *	CHECK declaration
 *
 *****************************************************************************/

check_decl:
  CHECK '(' a_expr_str ')' deferrable initially_deferred
  {
    chk := &ddl.Check{}
    chk.Expression = $3
    chk.Deferrable = $5
    chk.InitiallyDeferred = $6
    $$ = chk
  }

/*****************************************************************************
 *
 *	CREATE INDEX
 *
 *****************************************************************************/

/* TODO
 * - table name -> qualified_name
 * - full grammar
 */
create_index_stmt:
  CREATE unique_opt INDEX concurrently_opt name_opt
  ON name index_using '(' index_keys ')' where_expr_opt {
    idx := &ddl.Index{}
    idx.Unique = $2
    idx.Name = $5
    idx.Table = $7
    idx.Using = $8
    idx.Keys = $10.([]*ddl.IndexKey)
    idx.Where = $12
    $$ = idx
  }
| CREATE unique_opt INDEX concurrently_opt IF NOT EXISTS name
  ON name index_using '(' index_keys ')' where_expr_opt {
    idx := &ddl.Index{}
    idx.Unique = $2
    idx.Name = $8
    idx.Table = $10
    idx.Using = $11
    idx.Keys = $13.([]*ddl.IndexKey)
    idx.Where = $15
    $$ = idx
  }

/* TODO: allow bare function calls (test case #6) */
index_key:
  name index_opclass_opt desc_opt {
    k := &ddl.IndexKey{}
    k.Column = $1
    k.OpClass = $2
    k.Descending = $3
    $$ = k
  }
| '(' a_expr_str ')' index_opclass_opt desc_opt {
    k := &ddl.IndexKey{}
    k.Expression = $2
    k.OpClass = $4
    k.Descending = $5
    $$ = k
  }

index_keys:
  index_key {
    slice := make([]*ddl.IndexKey, 0, 4)
    $$ = append(slice, $1.(*ddl.IndexKey))
  }
| index_keys ',' index_key {
    slice := $1.([]*ddl.IndexKey)
    $$ = append(slice, $3.(*ddl.IndexKey))
  }

index_opclass_opt:
/* empty */ { $$ = "" }
| name      { $$ = $1 }

/* TODO: s/Identifier/name/ */
index_using:
/* empty */        { $$ = "" }
| USING Identifier { $$ = $2 }

/*****************************************************************************
 *
 *	CREATE TRIGGER
 *
 *****************************************************************************/

/* TODO
 * - trigger Identifier should be name
 * - table Identifier qualified_name
 * - function Identifier should be func_name
 * - OR REPLACE
 * - REFERENCING ...
 * - WHEN ( condition )
 * - function arguments
 */
create_trigger_stmt:
  CREATE TRIGGER Identifier trigger_timing trigger_event_list ON Identifier
  trigger_for trigger_when EXECUTE function_or_procedure Identifier '(' ')' {
    t := newTrigger($5...)
    t.Name = $3
    t.Timing = $4
    t.Table = $7
    t.ForEachRow = $8
    t.When = $9
    t.Function = $12
    $$ = t
  }
| CREATE CONSTRAINT TRIGGER Identifier AFTER trigger_event_list ON Identifier
  trigger_from deferrable initially_deferred FOR EACH ROW trigger_when
  EXECUTE function_or_procedure Identifier '(' ')' {
    t := newTrigger($6...)
    t.Constraint = true
    t.Name = $4
    t.Timing = "AFTER"
    t.Table = $8
    t.From = $9
    t.Deferrable = $10
    t.InitiallyDeferred = $11
    t.ForEachRow = true
    t.Function = $18
    $$ = t
  }

trigger_event_item:
  DELETE                { $$ = newOption("delete", true) }
| INSERT                { $$ = newOption("insert", true) }
| TRUNCATE              { $$ = newOption("truncate", true) }
| UPDATE                { $$ = newOption("update", true) }
| UPDATE OF column_list { $$ = newOption("columns", $3) }

trigger_event_list:
  trigger_event_item                       { $$ = newOptionList($1) }
| trigger_event_list OR trigger_event_item { $$ = append($1, $3) }

trigger_for:
/* empty */          { $$ = false }
| FOR EACH ROW       { $$ = true }
| FOR EACH STATEMENT { $$ = false }
| FOR ROW            { $$ = true }
| FOR STATEMENT      { $$ = false }

/* TODO: Identifier should be qualified_name */
trigger_from:
/* empty */       { $$ = "" }
| FROM Identifier { $$ = $2 }

trigger_timing:
  AFTER      { $$ = "AFTER" }
| BEFORE     { $$ = "BEFORE" }
| INSTEAD OF { $$ = "INSTEAD OF"}

trigger_when:
/* empty */               { $$ ="" }
| WHEN '(' a_expr_str ')' { $$ = $3 }

/*****************************************************************************
 *
 *	FOREIGN KEY declaration
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
 *	SELECT
 *
 *****************************************************************************/

select_stmt:
  select_no_parens			%prec UMINUS
| select_with_parens		%prec UMINUS

select_with_parens:
  '(' select_no_parens ')'
| '(' select_with_parens ')'

select_no_parens:
/* TODO: implement */

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

/* --------------------------------------------------------------------------*/
/* --------------------------------------------------------------------------*/
/*  copied from postgres/src/backend/parser/gram.y                           */
/*  commit c9d5298485b78a37923a23f9af9aa0ade06762db                          */
/* --------------------------------------------------------------------------*/
/* --------------------------------------------------------------------------*/

/*****************************************************************************
 *
 *	keyword category lists
 *
 *****************************************************************************/

col_name_keyword:
  BETWEEN
| BIGINT
| BIT
| BOOLEAN
| CHAR
| CHARACTER
| COALESCE
| DEC
| DECIMAL
| EXISTS
| EXTRACT
| FLOAT
| GREATEST
| GROUPING
| INOUT
| INT
| INTEGER
| INTERVAL
| LEAST
| NATIONAL
| NCHAR
| NONE
| NORMALIZE
| NULLIF
| NUMERIC
| OUT
| OVERLAY
| POSITION
| PRECISION
| REAL
| ROW
| SETOF
| SMALLINT
| SUBSTRING
| TIME
| TIMESTAMP
| TREAT
| TRIM
| VALUES
| VARCHAR
| XMLATTRIBUTES
| XMLCONCAT
| XMLELEMENT
| XMLEXISTS
| XMLFOREST
| XMLNAMESPACES
| XMLPARSE
| XMLPI
| XMLROOT
| XMLSERIALIZE
| XMLTABLE

reserved_keyword:
  ALL
| ANALYSE
| ANALYZE
| AND
| ANY
| ARRAY
| AS
| ASC
| ASYMMETRIC
| BOTH
| CASE
| CAST
| CHECK
| COLLATE
| COLUMN
| CONSTRAINT
| CREATE
| CURRENT_CATALOG
| CURRENT_DATE
| CURRENT_ROLE
| CURRENT_TIME
| CURRENT_TIMESTAMP
| CURRENT_USER
| DEFAULT
| DEFERRABLE
| DESC
| DISTINCT
| DO
| ELSE
| END
| EXCEPT
| FALSE
| FETCH
| FOR
| FOREIGN
| FROM
| GRANT
| GROUP
| HAVING
| IN
| INITIALLY
| INTERSECT
| INTO
| LATERAL
| LEADING
| LIMIT
| LOCALTIME
| LOCALTIMESTAMP
| NOT
| NULL
| OFFSET
| ON
| ONLY
| OR
| ORDER
| PLACING
| PRIMARY
| REFERENCES
| RETURNING
| SELECT
| SESSION_USER
| SOME
| SYMMETRIC
| TABLE
| THEN
| TO
| TRAILING
| TRUE
| UNION
| UNIQUE
| USER
| USING
| VARIADIC
| WHEN
| WHERE
| WINDOW
| WITH

type_func_name_keyword:
  AUTHORIZATION
| BINARY
| COLLATION
| CONCURRENTLY
| CROSS
| CURRENT_SCHEMA
| FREEZE
| FULL
| ILIKE
| INNER
| IS
| ISNULL
| JOIN
| LEFT
| LIKE
| NATURAL
| NOTNULL
| OUTER
| OVERLAPS
| RIGHT
| SIMILAR
| TABLESAMPLE
| VERBOSE

unreserved_keyword:
  ABORT
| ABSOLUTE
| ACCESS
| ACTION
| ADD
| ADMIN
| AFTER
| AGGREGATE
| ALSO
| ALTER
| ALWAYS
| ASSERTION
| ASSIGNMENT
| AT
| ATTACH
| ATTRIBUTE
| BACKWARD
| BEFORE
| BEGIN
| BY
| CACHE
| CALL
| CALLED
| CASCADE
| CASCADED
| CATALOG
| CHAIN
| CHARACTERISTICS
| CHECKPOINT
| CLASS
| CLOSE
| CLUSTER
| COLUMNS
| COMMENT
| COMMENTS
| COMMIT
| COMMITTED
| CONFIGURATION
| CONFLICT
| CONNECTION
| CONSTRAINTS
| CONTENT
| CONTINUE
| CONVERSION
| COPY
| COST
| CSV
| CUBE
| CURRENT
| CURSOR
| CYCLE
| DATA
| DATABASE
| DAY
| DEALLOCATE
| DECLARE
| DEFAULTS
| DEFERRED
| DEFINER
| DELETE
| DELIMITER
| DELIMITERS
| DEPENDS
| DETACH
| DICTIONARY
| DISABLE
| DISCARD
| DOCUMENT
| DOMAIN
| DOUBLE
| DROP
| EACH
| ENABLE
| ENCODING
| ENCRYPTED
| ENUM
| ESCAPE
| EVENT
| EXCLUDE
| EXCLUDING
| EXCLUSIVE
| EXECUTE
| EXPLAIN
| EXPRESSION
| EXTENSION
| EXTERNAL
| FAMILY
| FILTER
| FIRST
| FOLLOWING
| FORCE
| FORWARD
| FUNCTION
| FUNCTIONS
| GENERATED
| GLOBAL
| GRANTED
| GROUPS
| HANDLER
| HEADER
| HOLD
| HOUR
| IDENTITY
| IF
| IMMEDIATE
| IMMUTABLE
| IMPLICIT
| IMPORT
| INCLUDE
| INCLUDING
| INCREMENT
| INDEX
| INDEXES
| INHERIT
| INHERITS
| INLINE
| INPUT
| INSENSITIVE
| INSERT
| INSTEAD
| INVOKER
| ISOLATION
| KEY
| LABEL
| LANGUAGE
| LARGE
| LAST
| LEAKPROOF
| LEVEL
| LISTEN
| LOAD
| LOCAL
| LOCATION
| LOCK
| LOCKED
| LOGGED
| MAPPING
| MATCH
| MATERIALIZED
| MAXVALUE
| METHOD
| MINUTE
| MINVALUE
| MODE
| MONTH
| MOVE
| NAME
| NAMES
| NEW
| NEXT
| NFC
| NFD
| NFKC
| NFKD
| NO
| NORMALIZED
| NOTHING
| NOTIFY
| NOWAIT
| NULLS
| OBJECT
| OF
| OFF
| OIDS
| OLD
| OPERATOR
| OPTION
| OPTIONS
| ORDINALITY
| OTHERS
| OVER
| OVERRIDING
| OWNED
| OWNER
| PARALLEL
| PARSER
| PARTIAL
| PARTITION
| PASSING
| PASSWORD
| PLANS
| POLICY
| PRECEDING
| PREPARE
| PREPARED
| PRESERVE
| PRIOR
| PRIVILEGES
| PROCEDURAL
| PROCEDURE
| PROCEDURES
| PROGRAM
| PUBLICATION
| QUOTE
| RANGE
| READ
| REASSIGN
| RECHECK
| RECURSIVE
| REF
| REFERENCING
| REFRESH
| REINDEX
| RELATIVE
| RELEASE
| RENAME
| REPEATABLE
| REPLACE
| REPLICA
| RESET
| RESTART
| RESTRICT
| RETURNS
| REVOKE
| ROLE
| ROLLBACK
| ROLLUP
| ROUTINE
| ROUTINES
| ROWS
| RULE
| SAVEPOINT
| SCHEMA
| SCHEMAS
| SCROLL
| SEARCH
| SECOND
| SECURITY
| SEQUENCE
| SEQUENCES
| SERIALIZABLE
| SERVER
| SESSION
| SET
| SETS
| SHARE
| SHOW
| SIMPLE
| SKIP
| SNAPSHOT
| SQL
| STABLE
| STANDALONE
| START
| STATEMENT
| STATISTICS
| STDIN
| STDOUT
| STORAGE
| STORED
| STRICT
| STRIP
| SUBSCRIPTION
| SUPPORT
| SYSID
| SYSTEM
| TABLES
| TABLESPACE
| TEMP
| TEMPLATE
| TEMPORARY
| TEXT
| TIES
| TRANSACTION
| TRANSFORM
| TRIGGER
| TRUNCATE
| TRUSTED
| TYPE
| TYPES
| UESCAPE
| UNBOUNDED
| UNCOMMITTED
| UNENCRYPTED
| UNKNOWN
| UNLISTEN
| UNLOGGED
| UNTIL
| UPDATE
| VACUUM
| VALID
| VALIDATE
| VALIDATOR
| VALUE
| VARYING
| VERSION
| VIEW
| VIEWS
| VOLATILE
| WHITESPACE
| WITHIN
| WITHOUT
| WORK
| WRAPPER
| WRITE
| XML
| YEAR
| YES
| ZONE

/*****************************************************************************
 *
 *	misc.
 *
 *****************************************************************************/

all_Op:
  Op
| MathOp

any_operator:
  all_Op
| ColId '.' any_operator

attrs:
  '.' attr_name
| attrs '.' attr_name

case_arg:
/* empty */
| a_expr

case_default:
/* empty */
| ELSE a_expr

case_expr:
  CASE case_arg when_clause_list case_default END

columnref:
  ColId
| ColId indirection

explicit_row:
  ROW '(' expr_list ')'
| ROW '(' ')'

extract_arg:
  Identifier
| YEAR
| MONTH
| DAY
| HOUR
| MINUTE
| SECOND
| SCONST

extract_list:
  extract_arg FROM a_expr

filter_clause:
/* empty */
| FILTER '(' WHERE a_expr ')'

frame_bound:
  UNBOUNDED PRECEDING
| UNBOUNDED FOLLOWING
| CURRENT ROW
| a_expr PRECEDING
| a_expr FOLLOWING

frame_extent:
  frame_bound
| BETWEEN frame_bound AND frame_bound

func_arg_expr:
  a_expr
| param_name COLON_EQUALS a_expr
| param_name EQUALS_GREATER a_expr

func_arg_list:
  func_arg_expr
| func_arg_list ',' func_arg_expr

func_arg_list_opt:
/* empty */
| func_arg_list

implicit_row:
  '(' expr_list ',' a_expr ')'

in_expr:
  select_with_parens
| '(' expr_list ')'

indirection:
  indirection_el
| indirection indirection_el

indirection_el:
  '.' attr_name
| '.' '*'
| '[' a_expr ']'
| '[' opt_slice_bound ':' opt_slice_bound ']'

MathOp:
  '+'
| '-'
| '*'
| '/'
| '%'
| '^'
| '<'
| '>'
| '='
| LESS_EQUALS
| GREATER_EQUALS
| NOT_EQUALS

opt_asc_desc:
/* empty */
| ASC
| DESC

opt_asymmetric:
/* empty */
| ASYMMETRIC

/*
 * If we see PARTITION, RANGE, ROWS or GROUPS as the first token after the '('
 * of a window_specification, we want the assumption to be that there is
 * no existing_window_name; but those keywords are unreserved and so could
 * be ColIds.  We fix this by making them have the same precedence as Identifier
 * and giving the empty production here a slightly higher precedence, so
 * that the shift/reduce conflict is resolved in favor of reducing the rule.
 * These keywords are thus precluded from being an existing_window_name but
 * are not reserved for any other purpose.
 */
opt_existing_window_name:
/* empty */ %prec Op
| ColId

/*
 * For frame clauses, we return a WindowDef, but only some fields are used:
 * frameOptions, startOffset, and endOffset.
 */
opt_frame_clause:
/* empty */
| RANGE frame_extent opt_window_exclusion_clause
| ROWS frame_extent opt_window_exclusion_clause
| GROUPS frame_extent opt_window_exclusion_clause

opt_indirection:
/* empty */
| opt_indirection indirection_el

opt_nulls_order:
/* empty */
| NULLS_LA FIRST
| NULLS_LA LAST

opt_partition_clause:
/* empty */
| PARTITION BY expr_list

opt_slice_bound:
/* empty */
| a_expr

opt_sort_clause:
/* empty */
| sort_clause

opt_window_exclusion_clause:
/* empty */
| EXCLUDE CURRENT ROW
| EXCLUDE GROUP
| EXCLUDE TIES
| EXCLUDE NO OTHERS

over_clause:
/* empty */
| OVER window_specification
| OVER ColId

overlay_list:
  a_expr PLACING a_expr FROM a_expr FOR a_expr
| a_expr PLACING a_expr FROM a_expr

param_name:
  type_function_name

/* position_list uses b_expr not a_expr to avoid conflict with general IN */
position_list:
  b_expr IN b_expr

qual_all_Op:
  all_Op
| OPERATOR '(' any_operator ')'

row:
  ROW '(' expr_list ')'
| ROW '(' ')'
| '(' expr_list ',' a_expr ')'

sort_clause:
  ORDER BY sortby_list

sortby:
  a_expr USING qual_all_Op opt_nulls_order
| a_expr opt_asc_desc opt_nulls_order

sortby_list:
  sortby
| sortby_list ',' sortby

sub_type:
  ANY
| SOME
| ALL

subquery_Op:
  all_Op
| OPERATOR '(' any_operator ')'
| LIKE
| NOT_LA LIKE
| ILIKE
| NOT_LA ILIKE

/*
 * SUBSTRING() arguments
 *
 * Note that SQL:1999 has both
 *     text FROM int FOR int
 * and
 *     text FROM pattern FOR escape
 *
 * In the parser we map them both to a call to the substring() function and
 * rely on type resolution to pick the right one.
 *
 * In SQL:2003, the second variant was changed to
 *     text SIMILAR pattern ESCAPE escape
 * We could in theory map that to a different function internally, but
 * since we still support the SQL:1999 version, we don't.  However,
 * ruleutils.c will reverse-list the call in the newer style.
 */
substr_list:
  a_expr FROM a_expr FOR a_expr
| a_expr FOR a_expr FROM a_expr
| a_expr FROM a_expr
| a_expr FOR a_expr
| a_expr SIMILAR a_expr ESCAPE a_expr

qual_Op:
  Op
| OPERATOR '(' any_operator ')'

trim_list:
  a_expr FROM expr_list
| FROM expr_list
| expr_list

unicode_normal_form:
  NFC
| NFD
| NFKC
| NFKD

window_specification:
  '(' opt_existing_window_name opt_partition_clause
      opt_sort_clause opt_frame_clause ')'

within_group_clause:
/* empty */
| WITHIN GROUP '(' sort_clause ')'

when_clause:
  WHEN a_expr THEN a_expr

when_clause_list:
  when_clause
| when_clause_list when_clause

/*****************************************************************************
 *
 *	names and constants
 *
 *****************************************************************************/

AexprConst:
  ICONST
| FCONST
| SCONST
| BCONST
| XCONST
| func_name SCONST
| func_name '(' func_arg_list opt_sort_clause ')' SCONST
| ConstTypename SCONST
| ConstInterval SCONST opt_interval
| ConstInterval '(' ICONST ')' SCONST
| TRUE
| FALSE
| NULL

any_name:
  ColId
| ColId attrs

attr_name:
  ColLabel

ColId:
  Identifier
| unreserved_keyword
| col_name_keyword

ColLabel:
  Identifier
| unreserved_keyword
| col_name_keyword
| type_func_name_keyword
| reserved_keyword

/*
 * The production for a qualified func_name has to exactly match the
 * production for a qualified columnref, because we cannot tell which we
 * are parsing until we see what comes after it ('(' or SCONST for a func_name,
 * anything else for a columnref).  Therefore we allow 'indirection' which
 * may contain subscripts, and reject that case in the C code.  (If we
 * ever implement SQL99-like methods, such syntax may actually become legal!)
 */
func_name:
  type_function_name
| ColId indirection

type_function_name:
  Identifier
| unreserved_keyword
| type_func_name_keyword

/*****************************************************************************
 *
 *	xml
 *
 *****************************************************************************/

document_or_content:
  DOCUMENT
| CONTENT

opt_xml_root_standalone:
/* empty */
| ',' STANDALONE YES
| ',' STANDALONE NO
| ',' STANDALONE NO VALUE

xml_attributes:
  XMLATTRIBUTES '(' xml_attribute_list ')'

xml_attribute_el:
  a_expr AS ColLabel
| a_expr

xml_attribute_list:
  xml_attribute_el
| xml_attribute_list ',' xml_attribute_el

xmlexists_argument:
  PASSING c_expr
| PASSING c_expr xml_passing_mech
| PASSING xml_passing_mech c_expr
| PASSING xml_passing_mech c_expr xml_passing_mech

xml_passing_mech:
  BY REF
| BY VALUE

xml_root_version:
  VERSION a_expr
| VERSION NO VALUE

xml_whitespace_option:
/* empty */
| PRESERVE WHITESPACE
| STRIP WHITESPACE

/*****************************************************************************
 *
 *	data types
 *
 *****************************************************************************/

Typename:
  SimpleTypename opt_array_bounds
| SETOF SimpleTypename opt_array_bounds
/* SQL standard syntax, currently only one-dimensional */
| SimpleTypename ARRAY '[' ICONST ']'
| SETOF SimpleTypename ARRAY '[' ICONST ']'
| SimpleTypename ARRAY
| SETOF SimpleTypename ARRAY

opt_array_bounds:
  /* empty */
| opt_array_bounds '[' ']'
| opt_array_bounds '[' ICONST ']'

SimpleTypename:
  GenericType
| Numeric
| Bit
| Character
| ConstDatetime
| ConstInterval opt_interval
| ConstInterval '(' ICONST ')'

/* We have a separate ConstTypename to allow defaulting fixed-length
 * types such as CHAR() and BIT() to an unspecified length.
 * SQL9x requires that these default to a length of one, but this
 * makes no sense for constructs like CHAR 'hi' and BIT '0101',
 * where there is an obvious better choice to make.
 * Note that ConstInterval is not included here since it must
 * be pushed up higher in the rules to accommodate the postfix
 * options (e.g. INTERVAL '1' YEAR). Likewise, we have to handle
 * the generic-type-name case in AexprConst to avoid premature
 * reduce/reduce conflicts against function names.
 */
ConstTypename:
  Numeric
| ConstBit
| ConstCharacter
| ConstDatetime

/*
 * GenericType covers all type names that don't have special syntax mandated
 * by the standard, including qualified names.  We also allow type modifiers.
 * To avoid parsing conflicts against function invocations, the modifiers
 * have to be shown as expr_list here, but parse analysis will only accept
 * constants for them.
 */
GenericType:
  type_function_name opt_type_modifiers
| type_function_name attrs opt_type_modifiers

opt_type_modifiers:
/* empty */
| '(' expr_list ')'

/*
 * SQL numeric data types
 */
Numeric:
  INT
| INTEGER
| SMALLINT
| BIGINT
| REAL
| FLOAT opt_float
| DOUBLE PRECISION
| DECIMAL opt_type_modifiers
| DEC opt_type_modifiers
| NUMERIC opt_type_modifiers
| BOOLEAN

opt_float:
/* empty */
| '(' ICONST ')'

/*
 * SQL bit-field data types
 * The following implements BIT() and BIT VARYING().
 */
Bit:
  BitWithLength
| BitWithoutLength

/* ConstBit is like Bit except "BIT" defaults to unspecified length */
/* See notes for ConstCharacter, which addresses same issue for "CHAR" */
ConstBit:
  BitWithLength
| BitWithoutLength

BitWithLength:
  BIT opt_varying '(' expr_list ')'

BitWithoutLength:
  BIT opt_varying


/*
 * SQL character data types
 * The following implements CHAR() and VARCHAR().
 */
Character:
  CharacterWithLength
| CharacterWithoutLength

ConstCharacter:
  CharacterWithLength
| CharacterWithoutLength

CharacterWithLength:
  character '(' ICONST ')'

CharacterWithoutLength:
  character

character:
  CHARACTER opt_varying
| CHAR opt_varying
| VARCHAR
| NATIONAL CHARACTER opt_varying
| NATIONAL CHAR opt_varying
| NCHAR opt_varying

opt_varying:
/* empty */
| VARYING

/*
 * SQL date/time types
 */
ConstDatetime:
  TIMESTAMP '(' ICONST ')' opt_timezone
| TIMESTAMP opt_timezone
| TIME '(' ICONST ')' opt_timezone
| TIME opt_timezone

ConstInterval:
  INTERVAL

opt_timezone:
/* empty */
| WITH_LA TIME ZONE
| WITHOUT TIME ZONE

opt_interval:
/* empty */
| YEAR
| MONTH
| DAY
| HOUR
| MINUTE
| interval_second
| YEAR TO MONTH
| DAY TO HOUR
| DAY TO MINUTE
| DAY TO interval_second
| HOUR TO MINUTE
| HOUR TO interval_second
| MINUTE TO interval_second

interval_second:
  SECOND
| SECOND '(' ICONST ')'

/*****************************************************************************
 *
 *	expressions
 *
 *****************************************************************************/

/*
 * We have two expression types: a_expr is the unrestricted kind, and
 * b_expr is a subset that must be used in some places to avoid shift/reduce
 * conflicts.  For example, we can't do BETWEEN as "BETWEEN a_expr AND a_expr"
 * because that use of AND conflicts with AND as a boolean operator.  So,
 * b_expr is used in BETWEEN and we remove boolean keywords from b_expr.
 *
 * Note that '(' a_expr ')' is a b_expr, so an unrestricted expression can
 * always be used by surrounding it with parens.
 *
 * c_expr is all the productions that are common to a_expr and b_expr;
 * it's factored out just to eliminate redundant coding.
 *
 * Be careful of productions involving more than one terminal token.
 * By default, bison will assign such productions the precedence of their
 * last terminal, but in nearly all cases you want it to be the precedence
 * of the first terminal instead; otherwise you will not get the behavior
 * you expect!  So we use %prec annotations freely to set precedences.
 */
a_expr:
  c_expr
| a_expr TYPECAST Typename
| a_expr COLLATE any_name
| a_expr AT TIME ZONE a_expr  %prec AT
 /*
  * These operators must be called out explicitly in order to make use
  * of bison's automatic operator-precedence handling.  All other
  * operator names are handled by the generic productions using "Op",
  * below; and all those operators will have the same precedence.
  *
  * If you add more explicitly-known operators, be sure to add them
  * also to b_expr and to the MathOp list below.
  */
| '+' a_expr  %prec UMINUS
| '-' a_expr  %prec UMINUS
| a_expr '+' a_expr
| a_expr '-' a_expr
| a_expr '*' a_expr
| a_expr '/' a_expr
| a_expr '%' a_expr
| a_expr '^' a_expr
| a_expr '<' a_expr
| a_expr '>' a_expr
| a_expr '=' a_expr
| a_expr LESS_EQUALS a_expr
| a_expr GREATER_EQUALS a_expr
| a_expr NOT_EQUALS a_expr

| a_expr qual_Op a_expr  %prec Op
| qual_Op a_expr         %prec Op

| a_expr AND a_expr
| a_expr OR a_expr
| NOT a_expr
| NOT_LA a_expr          %prec NOT

| a_expr LIKE a_expr
| a_expr LIKE a_expr ESCAPE a_expr               %prec LIKE
| a_expr NOT_LA LIKE a_expr                      %prec NOT_LA
| a_expr NOT_LA LIKE a_expr ESCAPE a_expr        %prec NOT_LA
| a_expr ILIKE a_expr
| a_expr ILIKE a_expr ESCAPE a_expr              %prec ILIKE
| a_expr NOT_LA ILIKE a_expr                     %prec NOT_LA
| a_expr NOT_LA ILIKE a_expr ESCAPE a_expr       %prec NOT_LA

| a_expr SIMILAR TO a_expr                       %prec SIMILAR
| a_expr SIMILAR TO a_expr ESCAPE a_expr         %prec SIMILAR
| a_expr NOT_LA SIMILAR TO a_expr                %prec NOT_LA
| a_expr NOT_LA SIMILAR TO a_expr ESCAPE a_expr  %prec NOT_LA

 /* NullTest clause
  * Define SQL-style Null test clause.
  * Allow two forms described in the standard:
  *     a IS NULL
  *     a IS NOT NULL
  * Allow two SQL extensions
  *     a ISNULL
  *     a NOTNULL
  */
| a_expr IS NULL                                          %prec IS
| a_expr ISNULL
| a_expr IS NOT NULL                                      %prec IS
| a_expr NOTNULL
| row OVERLAPS row
| a_expr IS TRUE                                          %prec IS
| a_expr IS NOT TRUE                                      %prec IS
| a_expr IS FALSE                                         %prec IS
| a_expr IS NOT FALSE                                     %prec IS
| a_expr IS UNKNOWN                                       %prec IS
| a_expr IS NOT UNKNOWN                                   %prec IS
| a_expr IS DISTINCT FROM a_expr                          %prec IS
| a_expr IS NOT DISTINCT FROM a_expr                      %prec IS
| a_expr BETWEEN opt_asymmetric b_expr AND a_expr         %prec BETWEEN
| a_expr NOT_LA BETWEEN opt_asymmetric b_expr AND a_expr  %prec NOT_LA
| a_expr BETWEEN SYMMETRIC b_expr AND a_expr              %prec BETWEEN
| a_expr NOT_LA BETWEEN SYMMETRIC b_expr AND a_expr       %prec NOT_LA
| a_expr IN in_expr
| a_expr NOT_LA IN in_expr                                %prec NOT_LA
| a_expr subquery_Op sub_type select_with_parens          %prec Op
| a_expr subquery_Op sub_type '(' a_expr ')'              %prec Op
| UNIQUE select_with_parens
| a_expr IS DOCUMENT                                      %prec IS
| a_expr IS NOT DOCUMENT                                  %prec IS
| a_expr IS NORMALIZED                                    %prec IS
| a_expr IS unicode_normal_form NORMALIZED                %prec IS
| a_expr IS NOT NORMALIZED                                %prec IS
| a_expr IS NOT unicode_normal_form NORMALIZED            %prec IS
| DEFAULT

/*
 * Restricted expressions
 *
 * b_expr is a subset of the complete expression syntax defined by a_expr.
 *
 * Presently, AND, NOT, IS, and IN are the a_expr keywords that would
 * cause trouble in the places where b_expr is used.  For simplicity, we
 * just eliminate all the boolean-keyword-operator productions from b_expr.
 */
b_expr:
  c_expr
| b_expr TYPECAST Typename
| '+' b_expr  %prec UMINUS
| '-' b_expr  %prec UMINUS
| b_expr '+' b_expr
| b_expr '-' b_expr
| b_expr '*' b_expr
| b_expr '/' b_expr
| b_expr '%' b_expr
| b_expr '^' b_expr
| b_expr '<' b_expr
| b_expr '>' b_expr
| b_expr '=' b_expr
| b_expr LESS_EQUALS b_expr
| b_expr GREATER_EQUALS b_expr
| b_expr NOT_EQUALS b_expr
| b_expr qual_Op b_expr               %prec Op
| qual_Op b_expr                      %prec Op
| b_expr IS DISTINCT FROM b_expr      %prec IS
| b_expr IS NOT DISTINCT FROM b_expr  %prec IS
| b_expr IS DOCUMENT                  %prec IS
| b_expr IS NOT DOCUMENT              %prec IS

/*
 * Productions that can be used in both a_expr and b_expr.
 *
 * Note: productions that refer recursively to a_expr or b_expr mostly
 * cannot appear here.  However, it's OK to refer to a_exprs that occur
 * inside parentheses, such as function arguments; that cannot introduce
 * ambiguity to the b_expr syntax.
 */
c_expr:
  columnref
| AexprConst
| PARAM opt_indirection
| '(' a_expr ')' opt_indirection
| case_expr
| func_expr
| select_with_parens              %prec UMINUS
| select_with_parens indirection
| EXISTS select_with_parens
| ARRAY select_with_parens
| ARRAY array_expr
| explicit_row
| implicit_row
| GROUPING '(' expr_list ')'

array_expr:
  '[' expr_list ']'
| '[' array_expr_list ']'
| '[' ']'

array_expr_list:
  array_expr
| array_expr_list ',' array_expr

expr_list:
  a_expr
| expr_list ',' a_expr

func_application:
  func_name '(' ')'
| func_name '(' func_arg_list opt_sort_clause ')'
| func_name '(' VARIADIC func_arg_expr opt_sort_clause ')'
| func_name '(' func_arg_list ',' VARIADIC func_arg_expr opt_sort_clause ')'
| func_name '(' ALL func_arg_list opt_sort_clause ')'
| func_name '(' DISTINCT func_arg_list opt_sort_clause ')'
| func_name '(' '*' ')'

/*
 * func_expr and its cousin func_expr_windowless are split out from c_expr just
 * so that we have classifications for "everything that is a function call or
 * looks like one".  This isn't very important, but it saves us having to
 * document which variants are legal in places like "FROM function()" or the
 * backwards-compatible functional-index syntax for CREATE INDEX.
 * (Note that many of the special SQL functions wouldn't actually make any
 * sense as functional index entries, but we ignore that consideration here.)
 */
func_expr:
  func_application within_group_clause filter_clause over_clause
| func_expr_common_subexpr

/*
 * Special expressions that are considered to be functions.
 */
func_expr_common_subexpr:
  COLLATION FOR '(' a_expr ')'
| CURRENT_DATE
| CURRENT_TIME
| CURRENT_TIME '(' ICONST ')'
| CURRENT_TIMESTAMP
| CURRENT_TIMESTAMP '(' ICONST ')'
| LOCALTIME
| LOCALTIME '(' ICONST ')'
| LOCALTIMESTAMP
| LOCALTIMESTAMP '(' ICONST ')'
| CURRENT_ROLE
| CURRENT_USER
| SESSION_USER
| USER
| CURRENT_CATALOG
| CURRENT_SCHEMA
| CAST '(' a_expr AS Typename ')'
| EXTRACT '(' extract_list ')'
| NORMALIZE '(' a_expr ')'
| NORMALIZE '(' a_expr ',' unicode_normal_form ')'
| OVERLAY '(' overlay_list ')'
| OVERLAY '(' func_arg_list_opt ')'
| POSITION '(' position_list ')'
| SUBSTRING '(' substr_list ')'
| SUBSTRING '(' func_arg_list_opt ')'
| TREAT '(' a_expr AS Typename ')'
| TRIM '(' BOTH trim_list ')'
| TRIM '(' LEADING trim_list ')'
| TRIM '(' TRAILING trim_list ')'
| TRIM '(' trim_list ')'
| NULLIF '(' a_expr ',' a_expr ')'
| COALESCE '(' expr_list ')'
| GREATEST '(' expr_list ')'
| LEAST '(' expr_list ')'
| XMLCONCAT '(' expr_list ')'
| XMLELEMENT '(' NAME ColLabel ')'
| XMLELEMENT '(' NAME ColLabel ',' xml_attributes ')'
| XMLELEMENT '(' NAME ColLabel ',' expr_list ')'
| XMLELEMENT '(' NAME ColLabel ',' xml_attributes ',' expr_list ')'
| XMLEXISTS '(' c_expr xmlexists_argument ')'
| XMLFOREST '(' xml_attribute_list ')'
| XMLPARSE '(' document_or_content a_expr xml_whitespace_option ')'
| XMLPI '(' NAME ColLabel ')'
| XMLPI '(' NAME ColLabel ',' a_expr ')'
| XMLROOT '(' a_expr ',' xml_root_version opt_xml_root_standalone ')'
| XMLSERIALIZE '(' document_or_content a_expr AS SimpleTypename ')'

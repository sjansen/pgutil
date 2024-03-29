// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package pg10

import (
	"context"

	"github.com/jackc/pgtype"
)

const describeFunction = `-- name: describeFunction :one
SELECT
  pg_get_userbyid(p.proowner) AS "Owner"
, obj_description(p.oid, 'pg_proc') AS "Comment"
, pg_get_function_result(p.oid) AS "Returns"
, l.lanname AS "Language"
, p.prosrc AS "Definition"
FROM pg_proc p
LEFT JOIN pg_namespace n
  ON n.oid = p.pronamespace
LEFT JOIN pg_language l
  ON l.oid = p.prolang
WHERE n.nspname = $1
  AND p.proname = $2
`

type describeFunctionParams struct {
	Schema       string
	FunctionName string
}

type describeFunctionRow struct {
	Owner      string
	Comment    pgtype.Text
	Returns    pgtype.Text
	Language   pgtype.Name
	Definition pgtype.Text
}

func (q *Queries) describeFunction(ctx context.Context, arg describeFunctionParams) (describeFunctionRow, error) {
	row := q.db.QueryRow(ctx, describeFunction, arg.Schema, arg.FunctionName)
	var i describeFunctionRow
	err := row.Scan(
		&i.Owner,
		&i.Comment,
		&i.Returns,
		&i.Language,
		&i.Definition,
	)
	return i, err
}

const describeTrigger = `-- name: describeTrigger :one
SELECT
  pg_get_triggerdef(t.oid, true)::text AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_trigger t
  ON t.tgrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND t.tgname = $3
`

type describeTriggerParams struct {
	Schema    string
	TableName string
	Trigger   string
}

func (q *Queries) describeTrigger(ctx context.Context, arg describeTriggerParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, describeTrigger, arg.Schema, arg.TableName, arg.Trigger)
	var Definition pgtype.Text
	err := row.Scan(&Definition)
	return Definition, err
}

const listChecks = `-- name: listChecks :many
SELECT
  r.conname AS "Name"
, pg_get_constraintdef(r.oid, true) AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_constraint r
  ON r.conrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND r.contype = 'c'
ORDER BY 1
`

type listChecksParams struct {
	Schema    string
	TableName string
}

type listChecksRow struct {
	Name       string
	Definition pgtype.Text
}

func (q *Queries) listChecks(ctx context.Context, arg listChecksParams) ([]listChecksRow, error) {
	rows, err := q.db.Query(ctx, listChecks, arg.Schema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listChecksRow
	for rows.Next() {
		var i listChecksRow
		if err := rows.Scan(&i.Name, &i.Definition); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listColumns = `-- name: listColumns :many
SELECT
  a.attname AS "Name"
, format_type(a.atttypid, a.atttypmod) AS "Type"
, a.attnotnull AS "NotNull"
, (SELECT pg_get_expr(d.adbin, d.adrelid)
   FROM pg_attrdef d
   WHERE d.adrelid = a.attrelid
     AND d.adnum = a.attnum
     AND a.atthasdef) AS "Default"
, col_description(a.attrelid, a.attnum) AS "Comment"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_attribute a
  ON a.attrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND a.attnum > 0
  AND NOT a.attisdropped
ORDER BY a.attnum
`

type listColumnsParams struct {
	Schema    string
	TableName string
}

type listColumnsRow struct {
	Name    string
	Type    pgtype.Text
	NotNull bool
	Default pgtype.Text
	Comment pgtype.Text
}

func (q *Queries) listColumns(ctx context.Context, arg listColumnsParams) ([]listColumnsRow, error) {
	rows, err := q.db.Query(ctx, listColumns, arg.Schema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listColumnsRow
	for rows.Next() {
		var i listColumnsRow
		if err := rows.Scan(
			&i.Name,
			&i.Type,
			&i.NotNull,
			&i.Default,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listForeignKeys = `-- name: listForeignKeys :many
SELECT
  t.conname AS "Name"
, pg_get_constraintdef(t.oid) AS "Definition"
FROM pg_namespace n
JOIN pg_class c1
  ON c1.relnamespace = n.oid
JOIN pg_constraint t
  ON t.conrelid = c1.oid
JOIN pg_class c2
  ON c2.oid = t.confrelid
WHERE n.nspname = $1
  AND c1.relname = $2
  AND t.contype = 'f'
ORDER BY c2.relname, t.conkey
`

type listForeignKeysParams struct {
	Schema    string
	TableName string
}

type listForeignKeysRow struct {
	Name       string
	Definition pgtype.Text
}

func (q *Queries) listForeignKeys(ctx context.Context, arg listForeignKeysParams) ([]listForeignKeysRow, error) {
	rows, err := q.db.Query(ctx, listForeignKeys, arg.Schema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listForeignKeysRow
	for rows.Next() {
		var i listForeignKeysRow
		if err := rows.Scan(&i.Name, &i.Definition); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFunctions = `-- name: listFunctions :many
SELECT
  n.nspname AS "Schema"
, p.proname AS "Name"
, pg_get_userbyid(p.proowner) AS "Owner"
, obj_description(p.oid, 'pg_proc') AS "Comment"
, pg_get_function_result(p.oid) AS "Returns"
, l.lanname AS "Language"
, p.prosrc AS "Definition"
FROM pg_proc p
LEFT JOIN pg_namespace n
  ON n.oid = p.pronamespace
LEFT JOIN pg_language l
  ON l.oid = p.prolang
WHERE pg_function_is_visible(p.oid)
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
ORDER BY "Schema", "Name"
`

type listFunctionsRow struct {
	Schema     pgtype.Name
	Name       string
	Owner      string
	Comment    pgtype.Text
	Returns    pgtype.Text
	Language   pgtype.Name
	Definition pgtype.Text
}

func (q *Queries) listFunctions(ctx context.Context) ([]listFunctionsRow, error) {
	rows, err := q.db.Query(ctx, listFunctions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listFunctionsRow
	for rows.Next() {
		var i listFunctionsRow
		if err := rows.Scan(
			&i.Schema,
			&i.Name,
			&i.Owner,
			&i.Comment,
			&i.Returns,
			&i.Language,
			&i.Definition,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listIndexes = `-- name: listIndexes :many
SELECT
  c2.relname AS "Name"
, i.indisprimary AS "IsPrimary"
, pg_get_indexdef(i.indexrelid, 0, true) AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_index i
  ON i.indrelid = c.oid
JOIN pg_class c2
  ON c2.oid = i.indexrelid
WHERE n.nspname = $1
  AND c.relname = $2
ORDER BY
  i.indisprimary DESC
, i.indisunique DESC
, c2.relname
`

type listIndexesParams struct {
	Schema    string
	TableName string
}

type listIndexesRow struct {
	Name       string
	IsPrimary  bool
	Definition pgtype.Text
}

func (q *Queries) listIndexes(ctx context.Context, arg listIndexesParams) ([]listIndexesRow, error) {
	rows, err := q.db.Query(ctx, listIndexes, arg.Schema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listIndexesRow
	for rows.Next() {
		var i listIndexesRow
		if err := rows.Scan(&i.Name, &i.IsPrimary, &i.Definition); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSchemas = `-- name: listSchemas :many
SELECT
  n.nspname AS "Name"
, pg_get_userbyid(n.nspowner) AS "Owner"
, obj_description(n.oid, 'pg_namespace') AS "Comment"
FROM
  pg_namespace n
WHERE n.nspname !~ '^pg_'
  AND n.nspname <> 'information_schema'
ORDER BY "Name"
`

type listSchemasRow struct {
	Name    string
	Owner   string
	Comment pgtype.Text
}

func (q *Queries) listSchemas(ctx context.Context) ([]listSchemasRow, error) {
	rows, err := q.db.Query(ctx, listSchemas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listSchemasRow
	for rows.Next() {
		var i listSchemasRow
		if err := rows.Scan(&i.Name, &i.Owner, &i.Comment); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSequences = `-- name: listSequences :many
SELECT
  n.nspname AS "Schema"
, c.relname AS "Name"
, pg_get_userbyid(c.relowner) AS "Owner"
, obj_description(c.oid, 'pg_class')::text AS "Comment"
, format_type(s.seqtypid, NULL) AS "Type"
, s.seqstart AS "Start"
, s.seqmin AS "Minimum"
, s.seqmax AS "Maximum"
, s.seqincrement AS "Increment"
, s.seqcache AS "Cache"
, s.seqcycle AS "Cycle"
, n2.nspname AS "OwnedBySchema"
, c2.relname AS "OwnedByTable"
, a2.attname AS "OwnedByColumn"
FROM pg_class c
JOIN pg_namespace n
  ON n.oid = c.relnamespace
JOIN pg_sequence s
  ON s.seqrelid = c.oid
LEFT JOIN pg_depend d
  ON d.objid = c.oid
JOIN pg_class c2
  ON c2.oid = d.refobjid
JOIN pg_namespace n2
  ON n2.oid = c2.relnamespace
JOIN pg_attribute a2
  ON (a2.attrelid = c2.oid AND a2.attnum = d.refobjsubid)
WHERE c.relkind = 'S'
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
  AND n.nspname !~ '^pg_toast'
  AND pg_table_is_visible(c.oid)
  AND d.classid='pg_class'::regclass
  AND d.refclassid='pg_class'::regclass
  AND d.deptype IN ('a', 'i')
ORDER BY "Schema", "Name"
`

type listSequencesRow struct {
	Schema        string
	Name          string
	Owner         string
	Comment       pgtype.Text
	Type          pgtype.Text
	Start         int64
	Minimum       int64
	Maximum       int64
	Increment     int64
	Cache         int64
	Cycle         bool
	OwnedBySchema string
	OwnedByTable  string
	OwnedByColumn string
}

func (q *Queries) listSequences(ctx context.Context) ([]listSequencesRow, error) {
	rows, err := q.db.Query(ctx, listSequences)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listSequencesRow
	for rows.Next() {
		var i listSequencesRow
		if err := rows.Scan(
			&i.Schema,
			&i.Name,
			&i.Owner,
			&i.Comment,
			&i.Type,
			&i.Start,
			&i.Minimum,
			&i.Maximum,
			&i.Increment,
			&i.Cache,
			&i.Cycle,
			&i.OwnedBySchema,
			&i.OwnedByTable,
			&i.OwnedByColumn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTableStorageParameters = `-- name: listTableStorageParameters :many
SELECT UNNEST(c.reloptions)::text AS "Parameter"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
WHERE n.nspname = $1
  AND c.relname = $2
ORDER BY 1
`

type listTableStorageParametersParams struct {
	Schema    string
	TableName string
}

func (q *Queries) listTableStorageParameters(ctx context.Context, arg listTableStorageParametersParams) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, listTableStorageParameters, arg.Schema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var Parameter pgtype.Text
		if err := rows.Scan(&Parameter); err != nil {
			return nil, err
		}
		items = append(items, Parameter)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTables = `-- name: listTables :many
SELECT
  n.nspname as "Schema"
, c.relname as "Name"
, pg_get_userbyid(c.relowner) as "Owner"
, obj_description(c.oid, 'pg_class') as "Comment"
FROM pg_class c
LEFT JOIN pg_namespace n
  ON n.oid = c.relnamespace
WHERE c.relkind IN ('r','p','')
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
  AND n.nspname !~ '^pg_toast'
  AND pg_table_is_visible(c.oid)
ORDER BY "Schema", "Name"
`

type listTablesRow struct {
	Schema  pgtype.Name
	Name    string
	Owner   string
	Comment pgtype.Text
}

func (q *Queries) listTables(ctx context.Context) ([]listTablesRow, error) {
	rows, err := q.db.Query(ctx, listTables)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listTablesRow
	for rows.Next() {
		var i listTablesRow
		if err := rows.Scan(
			&i.Schema,
			&i.Name,
			&i.Owner,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTriggers = `-- name: listTriggers :many
SELECT
  t.tgname AS "Name"
, pg_get_triggerdef(t.oid, true)::text AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_trigger t
  ON t.tgrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND NOT t.tgisinternal
ORDER BY 1
`

type listTriggersParams struct {
	Schema    string
	TableName string
}

type listTriggersRow struct {
	Name       string
	Definition pgtype.Text
}

func (q *Queries) listTriggers(ctx context.Context, arg listTriggersParams) ([]listTriggersRow, error) {
	rows, err := q.db.Query(ctx, listTriggers, arg.Schema, arg.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []listTriggersRow
	for rows.Next() {
		var i listTriggersRow
		if err := rows.Scan(&i.Name, &i.Definition); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

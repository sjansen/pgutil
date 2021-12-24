-- name: describeFunction :one
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
WHERE n.nspname = @schema
  AND p.proname = @function_name
;

-- name: describeTrigger :one
SELECT
  pg_get_triggerdef(t.oid, true)::text AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_trigger t
  ON t.tgrelid = c.oid
WHERE n.nspname = @schema
  AND c.relname = @table_name
  AND t.tgname = @trigger
;

-- name: listChecks :many
SELECT
  r.conname AS "Name"
, pg_get_constraintdef(r.oid, true) AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_constraint r
  ON r.conrelid = c.oid
WHERE n.nspname = @schema
  AND c.relname = @table_name
  AND r.contype = 'c'
ORDER BY 1
;

-- name: listColumns :many
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
WHERE n.nspname = @schema
  AND c.relname = @table_name
  AND a.attnum > 0
  AND NOT a.attisdropped
ORDER BY a.attnum
;

-- name: listForeignKeys :many
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
WHERE n.nspname = @schema
  AND c1.relname = @table_name
  AND t.contype = 'f'
ORDER BY c2.relname, t.conkey
;

-- name: listFunctions :many
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
;

-- name: listIndexes :many
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
WHERE n.nspname = @schema
  AND c.relname = @table_name
ORDER BY
  i.indisprimary DESC
, i.indisunique DESC
, c2.relname
;

-- name: listSchemas :many
SELECT
  n.nspname AS "Name"
, pg_get_userbyid(n.nspowner) AS "Owner"
, obj_description(n.oid, 'pg_namespace') AS "Comment"
FROM
  pg_namespace n
WHERE n.nspname !~ '^pg_'
  AND n.nspname <> 'information_schema'
ORDER BY "Name"
;

-- name: listSequences :many
SELECT
  n.nspname as "Schema"
, c.relname as "Name"
, pg_get_userbyid(c.relowner) as "Owner"
, obj_description(c.oid, 'pg_class') as "Comment"
, n2.nspname AS "OwnedBySchema"
, c2.relname AS "OwnedByTable"
, a2.attname AS "OwnedByColumn"
FROM pg_class c
JOIN pg_namespace n
  ON n.oid = c.relnamespace
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
  AND d.classid='pg_catalog.pg_class'::pg_catalog.regclass
  AND d.refclassid='pg_catalog.pg_class'::pg_catalog.regclass
  AND d.deptype IN ('a', 'i')
ORDER BY "Schema", "Name"
;

-- name: listTableStorageParameters :many
SELECT UNNEST(c.reloptions)::text AS "Parameter"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
WHERE n.nspname = @schema
  AND c.relname = @table_name
ORDER BY 1
;

-- name: listTables :many
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
;

-- name: listTriggers :many
SELECT
  t.tgname AS "Name"
, pg_get_triggerdef(t.oid, true)::text AS "Definition"
FROM pg_namespace n
JOIN pg_class c
  ON c.relnamespace = n.oid
JOIN pg_trigger t
  ON t.tgrelid = c.oid
WHERE n.nspname = @schema
  AND c.relname = @table_name
  AND NOT t.tgisinternal
ORDER BY 1
;

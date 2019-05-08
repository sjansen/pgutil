package pg

import "github.com/sjansen/pgutil/internal/ddl"

var listTables = `
SELECT
  n.nspname as "Schema"
, c.relname as "Name"
, pg_catalog.pg_get_userbyid(c.relowner) as "Owner"
, pg_catalog.obj_description(c.oid, 'pg_class') as "Comment"
FROM pg_catalog.pg_class c
LEFT JOIN pg_catalog.pg_namespace n
  ON n.oid = c.relnamespace
WHERE c.relkind IN ('r','p','')
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
  AND n.nspname !~ '^pg_toast'
  AND pg_catalog.pg_table_is_visible(c.oid)
ORDER BY "Schema", "Name"
`

func (c *Conn) ListTables() ([]*ddl.Table, error) {
	c.log.Infow("listing tables")

	c.log.Debugw("executing query", "query", listTables)
	rows, err := c.conn.Query(listTables)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var tables []*ddl.Table
	for rows.Next() {
		var schema, name, owner, comment *string
		err = rows.Scan(&schema, &name, &owner, &comment)
		if err != nil {
			break
		}
		table := &ddl.Table{
			Schema:  String(schema),
			Name:    String(name),
			Owner:   String(owner),
			Comment: String(comment),
		}
		c.log.Debugw("row scanned", "table", table)
		tables = append(tables, table)
	}
	if err != nil {
		return nil, err
	}

	return tables, nil
}

var listColumns = `
SELECT a.attname
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c
  ON c.relnamespace = n.oid
JOIN pg_catalog.pg_attribute a
  ON a.attrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND a.attnum > 0
  AND NOT a.attisdropped
ORDER BY a.attnum
`

func (c *Conn) ListColumns(schema, table string) ([]string, error) {
	c.log.Infow("listing columns", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listColumns)
	rows, err := c.conn.Query(listColumns, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var columns []string
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			break
		}
		c.log.Debugw("row scanned", "column", name)
		columns = append(columns, name)
	}
	if err != nil {
		return nil, err
	}

	return columns, nil
}

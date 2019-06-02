package pg

import "github.com/sjansen/pgutil/internal/ddl"

var listColumns = `
SELECT
  a.attname
, pg_catalog.format_type(a.atttypid, a.atttypmod)
, a.attnotnull
, (SELECT pg_catalog.pg_get_expr(d.adbin, d.adrelid)
   FROM pg_catalog.pg_attrdef d
   WHERE d.adrelid = a.attrelid
     AND d.adnum = a.attnum
     AND a.atthasdef)
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

// ListColumns describes the columns of a database table
func (c *Conn) ListColumns(schema, table string) ([]*ddl.Column, error) {
	c.log.Infow("listing columns", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listColumns)
	rows, err := c.conn.Query(listColumns, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var columns []*ddl.Column
	for rows.Next() {
		var defaultValue *string
		col := &ddl.Column{}
		err = rows.Scan(
			&col.Name, &col.Type, &col.NotNull, &defaultValue,
		)
		if err != nil {
			break
		}
		col.Default = String(defaultValue)
		c.log.Debugw("row scanned", "column", col.Name, "type", col.Type)
		columns = append(columns, col)
	}
	if err != nil {
		return nil, err
	}

	return columns, nil
}

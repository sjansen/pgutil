package pg

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

// ListColumns describes the columns of a database table
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

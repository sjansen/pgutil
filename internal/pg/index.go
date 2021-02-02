package pg

import (
	"github.com/sjansen/pgutil/internal/ddl"
	"github.com/sjansen/pgutil/internal/sqlparser"
)

var listIndexes = `
SELECT
  c2.relname
, i.indisprimary
, pg_catalog.pg_get_indexdef(i.indexrelid, 0, true)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c
  ON c.relnamespace = n.oid
JOIN pg_catalog.pg_index i
  ON i.indrelid = c.oid
JOIN pg_catalog.pg_class c2
  ON c2.oid = i.indexrelid
WHERE n.nspname = $1
  AND c.relname = $2
ORDER BY
  i.indisprimary DESC
, i.indisunique DESC
, c2.relname
`

// ListIndexes describes the indexes of a database table
func (c *Conn) ListIndexes(schema, table string) ([]*ddl.Index, error) {
	c.log.Infow("listing indexes", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listIndexes)
	rows, err := c.conn.Query(listIndexes, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var indexes []*ddl.Index
	for rows.Next() {
		var name, indexdef string
		var isPrimary bool
		err = rows.Scan(&name, &isPrimary, &indexdef)
		if err != nil {
			break
		}
		c.log.Debugw("row scanned", "index", name, "indexdef", indexdef)

		var index *ddl.Index
		index, err = sqlparser.ParseCreateIndex(indexdef)
		if err != nil {
			break
		}
		index.Schema = schema
		index.Primary = isPrimary
		indexes = append(indexes, index)
	}
	if err != nil {
		return nil, err
	}

	return indexes, nil
}

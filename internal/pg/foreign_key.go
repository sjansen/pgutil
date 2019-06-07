package pg

import "github.com/sjansen/pgutil/internal/ddl"

var listForeignKeys = `
SELECT
  t.conname
, pg_catalog.pg_get_constraintdef(t.oid) as condef
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c1
  ON c1.relnamespace = n.oid
JOIN pg_catalog.pg_constraint t
  ON t.conrelid = c1.oid
JOIN pg_catalog.pg_class c2
  ON c2.oid = t.confrelid
WHERE n.nspname = $1
  AND c1.relname = $2
  AND t.contype = 'f'
ORDER BY c2.relname, t.conkey;
`

// ListForeignKeys describes a database table's dependencies on other tables
func (c *Conn) ListForeignKeys(schema, table string) ([]*ddl.ForeignKey, error) {
	c.log.Infow("listing foreign keys", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listForeignKeys)
	rows, err := c.conn.Query(listForeignKeys, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var fks []*ddl.ForeignKey
	for rows.Next() {
		var name, fkdef string
		err = rows.Scan(&name, &fkdef)
		if err != nil {
			break
		}
		c.log.Debugw("row scanned", "name", name, "fkdef", fkdef)

		var fk *ddl.ForeignKey
		fk, err = ddl.ParseForeignKey(fkdef)
		if err != nil {
			break
		}
		fk.Name = name
		fks = append(fks, fk)
	}
	if err != nil {
		return nil, err
	}

	return fks, nil
}

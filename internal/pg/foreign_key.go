package pg

import (
	"github.com/sjansen/pgutil/internal/schema"
	"github.com/sjansen/pgutil/internal/sqlparser"
)

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
func (c *Conn) ListForeignKeys(namespace, table string) ([]*schema.ForeignKey, error) {
	c.log.Infow("listing foreign keys", "schema", namespace, "table", table)

	c.log.Debugw("executing query", "query", listForeignKeys)
	rows, err := c.conn.Query(listForeignKeys, namespace, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var fks []*schema.ForeignKey
	for rows.Next() {
		var name, fkdef string
		err = rows.Scan(&name, &fkdef)
		if err != nil {
			break
		}
		c.log.Debugw("row scanned", "name", name, "fkdef", fkdef)

		var fk *schema.ForeignKey
		fk, err = sqlparser.ParseForeignKey([]byte(fkdef))
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

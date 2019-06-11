package pg

import "github.com/sjansen/pgutil/internal/ddl"

var listChecks = `
SELECT
  r.conname
, pg_catalog.pg_get_constraintdef(r.oid, true)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c
  ON c.relnamespace = n.oid
JOIN pg_catalog.pg_constraint r
  ON r.conrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND r.contype = 'c'
ORDER BY 1;
`

// ListChecks describes the check constraints of a database table
func (c *Conn) ListChecks(schema, table string) ([]*ddl.Check, error) {
	c.log.Infow("listing checks", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listChecks)
	rows, err := c.conn.Query(listChecks, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var checks []*ddl.Check
	for rows.Next() {
		var name, checkdef string
		err = rows.Scan(&name, &checkdef)
		if err != nil {
			break
		}
		c.log.Debugw("row scanned", "check", name, "checkdef", checkdef)

		var check *ddl.Check
		check, err = ddl.ParseCheck(checkdef)
		if err != nil {
			break
		}
		check.Name = name
		checks = append(checks, check)
	}
	if err != nil {
		return nil, err
	}

	return checks, nil
}

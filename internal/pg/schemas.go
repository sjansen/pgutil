package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/ddl"
)

var listSchemas = `
SELECT
  n.nspname AS "Name"
, pg_catalog.pg_get_userbyid(n.nspowner) AS "Owner"
, pg_catalog.obj_description(n.oid, 'pg_namespace') AS "Comment"
FROM
  pg_catalog.pg_namespace n
WHERE n.nspname !~ '^pg_'
  AND n.nspname <> 'information_schema'
ORDER BY "Name"
`

// ListSchemas describes the schemas in the database
func (c *Conn) ListSchemas(ctx context.Context) ([]*ddl.Schema, error) {
	c.log.Infow("listing schemas")

	c.log.Debugw("executing query", "query", listSchemas)
	rows, err := c.conn.Query(ctx, listSchemas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var schemas []*ddl.Schema
	for rows.Next() {
		var name, owner, comment *string
		err = rows.Scan(&name, &owner, &comment)
		if err != nil {
			break
		}
		schema := &ddl.Schema{
			Name:    String(name),
			Owner:   String(owner),
			Comment: String(comment),
		}
		c.log.Debugw("row scanned", "schema", schema)
		schemas = append(schemas, schema)
	}
	if err != nil {
		return nil, err
	}

	return schemas, nil
}

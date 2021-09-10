package pg

import (
	"context"
	"strings"

	"github.com/sjansen/pgutil/internal/ddl"
)

var listTableStorageParameters = `
SELECT UNNEST(c.reloptions)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c
  ON c.relnamespace = n.oid
WHERE n.nspname = $1
  AND c.relname = $2
ORDER BY 1;
`

// ListTableStorageParameters describes table storage parameters
func (c *Conn) ListTableStorageParameters(
	ctx context.Context, schema, table string,
) (*ddl.TableStorageParameters, error) {
	c.log.Infow("listing table storage parameters", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listTableStorageParameters)
	rows, err := c.conn.Query(ctx, listTableStorageParameters, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	params := &ddl.TableStorageParameters{}
	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			break
		}
		c.log.Debugw("row scanned", "param", tmp)

		parts := strings.SplitN(tmp, "=", 2)
		if len(parts) < 2 {
			err = params.Set(parts[0], "")
		} else {
			err = params.Set(parts[0], parts[1])
		}
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return params, nil
}

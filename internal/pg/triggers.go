package pg

import (
	"github.com/jackc/pgx"
	"github.com/sjansen/pgutil/internal/ddl"
)

var describeTrigger = `
SELECT
  pg_catalog.pg_get_triggerdef(t.oid, true)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c
  ON c.relnamespace = n.oid
JOIN pg_catalog.pg_trigger t
  ON t.tgrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND t.tgname = $3
`

var listTriggers = `
SELECT
  t.tgname
, pg_catalog.pg_get_triggerdef(t.oid, true)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_class c
  ON c.relnamespace = n.oid
JOIN pg_catalog.pg_trigger t
  ON t.tgrelid = c.oid
WHERE n.nspname = $1
  AND c.relname = $2
  AND NOT t.tgisinternal
ORDER BY 1;
`

// DescribeTrigger describes a specific database table trigger
func (c *Conn) DescribeTrigger(schema, table, name string) (*ddl.Trigger, error) {
	c.log.Infow("DescribeTrigger", "schema", schema, "table", table, "name", name)

	var triggerdef string

	c.log.Debugw("executing query", "query", listTriggers)
	err := c.conn.QueryRow(describeTrigger, schema, table, name).Scan(&triggerdef)
	switch {
	case err == pgx.ErrNoRows:
		return nil, ErrNotFound
	case err != nil:
		return nil, err
	}
	c.log.Debugw("scanned", "triggerdef", triggerdef)

	var trigger *ddl.Trigger
	trigger, err = ddl.ParseTrigger(triggerdef)
	if err != nil {
		return nil, err
	}
	trigger.Schema = schema
	c.log.Debugw("parsed", "trigger", trigger)

	return trigger, nil
}

// ListTriggers describes the triggers of a database table
func (c *Conn) ListTriggers(schema, table string) ([]*ddl.Trigger, error) {
	c.log.Infow("ListTriggers", "schema", schema, "table", table)

	c.log.Debugw("executing query", "query", listTriggers)
	rows, err := c.conn.Query(listTriggers, schema, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var triggers []*ddl.Trigger
	for rows.Next() {
		var name, triggerdef string
		err = rows.Scan(&name, &triggerdef)
		if err != nil {
			break
		}
		c.log.Debugw("scanned", "trigger", name, "triggerdef", triggerdef)

		var trigger *ddl.Trigger
		trigger, err = ddl.ParseTrigger(triggerdef)
		if err != nil {
			break
		}
		trigger.Schema = schema
		triggers = append(triggers, trigger)
	}
	if err != nil {
		return nil, err
	}

	return triggers, nil
}

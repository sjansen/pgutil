package pg

import "github.com/sjansen/pgutil/internal/ddl"

var listFunctions = `
SELECT
  n.nspname as "Schema"
, p.proname as "Name"
, pg_catalog.pg_get_userbyid(p.proowner) as "Owner"
, pg_catalog.obj_description(p.oid, 'pg_proc') as "Comment"
, pg_catalog.pg_get_function_result(p.oid) as "Returns"
, l.lanname as "Language"
, p.prosrc as "Definition"
FROM pg_catalog.pg_proc p
LEFT JOIN pg_catalog.pg_namespace n
  ON n.oid = p.pronamespace
LEFT JOIN pg_catalog.pg_language l
  ON l.oid = p.prolang
WHERE pg_catalog.pg_function_is_visible(p.oid)
  AND n.nspname <> 'pg_catalog'
  AND n.nspname <> 'information_schema'
ORDER BY "Schema", "Name"
`

// ListFunctions describes the custom functions in the database
func (c *Conn) ListFunctions() ([]*ddl.Function, error) {
	c.log.Infow("listing functions")

	c.log.Debugw("executing query", "query", listFunctions)
	rows, err := c.conn.Query(listFunctions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	c.log.Debugw("scanning rows")
	var functions []*ddl.Function
	for rows.Next() {
		var schema, name, owner, comment *string
		var returns, language, definition *string
		err = rows.Scan(
			&schema, &name, &owner, &comment,
			&returns, &language, &definition,
		)
		if err != nil {
			break
		}
		function := &ddl.Function{
			Schema:  String(schema),
			Name:    String(name),
			Owner:   String(owner),
			Comment: String(comment),

			Returns:    String(returns),
			Language:   String(language),
			Definition: String(definition),
		}
		c.log.Debugw("row scanned", "function", function)
		functions = append(functions, function)
	}
	if err != nil {
		return nil, err
	}

	return functions, nil
}

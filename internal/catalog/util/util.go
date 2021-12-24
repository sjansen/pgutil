package util

import "github.com/jackc/pgx/v4"

// Identifier returns a SQL safe identifier string.
// For example: ["pg_catalog", "pg_class"] becomes "pg_catalog"."pg_class".
func Identifier(parts ...string) string {
	id := pgx.Identifier(parts)
	return id.Sanitize()
}

package pg

import "github.com/jackc/pgx/v4"

// String converts nullable strings to standard strings
func String(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Identifier returns a SQL safe identifier string.
// For example: ["pg_catalog", "pg_class"] becomes "pg_catalog"."pg_class".
func Identifier(parts ...string) string {
	id := pgx.Identifier(parts)
	return id.Sanitize()
}

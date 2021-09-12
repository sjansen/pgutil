package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListChecks describes the check constraints of a database table
func (c *Conn) ListChecks(ctx context.Context, schema, table string) ([]*ddl.Check, error) {
	db := pg10.New(c.conn)
	return db.ListChecks(ctx, schema, table)
}

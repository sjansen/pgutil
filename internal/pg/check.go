package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListChecks describes the check constraints of a database table
func (c *Conn) ListChecks(ctx context.Context, schema, table string) ([]*ddl.Check, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListChecks(ctx, schema, table)
}

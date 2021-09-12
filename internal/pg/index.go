package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListIndexes describes the indexes of a database table
func (c *Conn) ListIndexes(ctx context.Context, schema, table string) ([]*ddl.Index, error) {
	db := pg10.New(c.conn)
	return db.ListIndexes(ctx, schema, table)
}

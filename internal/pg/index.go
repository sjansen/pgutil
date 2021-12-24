package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListIndexes describes the indexes of a database table
func (c *Conn) ListIndexes(ctx context.Context, schema, table string) ([]*ddl.Index, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListIndexes(ctx, schema, table)
}

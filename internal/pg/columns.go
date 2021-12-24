package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListColumns describes the columns of a database table
func (c *Conn) ListColumns(ctx context.Context, schema, table string) ([]*ddl.Column, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListColumns(ctx, schema, table)
}

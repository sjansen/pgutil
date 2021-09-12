package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListColumns describes the columns of a database table
func (c *Conn) ListColumns(ctx context.Context, schema, table string) ([]*ddl.Column, error) {
	db := pg10.New(c.conn)
	return db.ListColumns(ctx, schema, table)
}

package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListForeignKeys describes a database table's dependencies on other tables
func (c *Conn) ListForeignKeys(ctx context.Context, schema, table string) ([]*ddl.ForeignKey, error) {
	db := pg10.New(c.conn)
	return db.ListForeignKeys(ctx, schema, table)
}

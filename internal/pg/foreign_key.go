package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListForeignKeys describes a database table's dependencies on other tables
func (c *Conn) ListForeignKeys(ctx context.Context, schema, table string) ([]*ddl.ForeignKey, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListForeignKeys(ctx, schema, table)
}

package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListTables describes the tables in the database
func (c *Conn) ListTables(ctx context.Context) ([]*ddl.Table, error) {
	db, err := catalog.New(ctx, c.conn)
	if err != nil {
		return nil, err
	}
	return db.ListTables(ctx)
}

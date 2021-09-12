package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListTables describes the tables in the database
func (c *Conn) ListTables(ctx context.Context) ([]*ddl.Table, error) {
	db := pg10.New(c.conn)
	return db.ListTables(ctx)
}

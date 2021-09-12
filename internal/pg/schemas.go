package pg

import (
	"context"

	"github.com/sjansen/pgutil/internal/catalog/pg10"
	"github.com/sjansen/pgutil/internal/ddl"
)

// ListSchemas describes the schemas in the database
func (c *Conn) ListSchemas(ctx context.Context) ([]*ddl.Schema, error) {
	db := pg10.New(c.conn)
	return db.ListSchemas(ctx)
}
